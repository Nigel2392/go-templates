package extensions

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

// View returns a view that renders the extension into the base template.
func View(options *Options, ext Extension) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var buf bytes.Buffer
		var tdata = ext.View(w, r)
		var tmpl *template.Template

		switch ext := ext.(type) {
		case ExtensionWithTemplate:
			tmpl = ext.Template(w, r)
			options.render(&buf, ext, tmpl.Tree.Root.String())
		case ExtensionWithStrings:
			options.render(&buf, ext, ext.String(w, r))
		case ExtensionWithFilename:
			tmpl, err = template.ParseFS(options.ExtensionManager.TplFileSystem, ext.Filename())
			if err != nil {
				defaultErr(options, w, r, err)
				return
			}
			options.render(&buf, ext, tmpl.Tree.Root.String())
		default:
			panic(fmt.Sprintf("Extension %s does not implement any of the extension interfaces", ext.Name()))
		}

		t, err := options.BaseManager.GetFromString(buf.String(), "ext")
		if err != nil {
			defaultErr(options, w, r, err)
			return
		}

		base, err := options.BaseManager.GetBases(nil)
		if err != nil {
			defaultErr(options, w, r, err)
			return
		}
		for _, b := range base.Templates() {
			t.AddParseTree(b.Name(), b.Tree)
		}

		t.Funcs(options.BaseManager.DefaultTplFuncs)
		t.Funcs(options.ExtensionManager.DefaultTplFuncs)

		if options.BeforeRender != nil {
			options.BeforeRender(w, r, t)
		}

		err = t.Execute(w, tdata)
		if err != nil {
			defaultErr(options, w, r, err)
			return
		}
	}
}

func defaultErr(o *Options, w http.ResponseWriter, r *http.Request, err error) {
	if o.OnError != nil {
		o.OnError(w, r, err)
	} else {
		http.Error(w, err.Error(), 500)
	}
}
