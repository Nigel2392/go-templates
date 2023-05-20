### Rendering templates

Firstly, we need to define a variable variables in the `github.com/Nigel2392/templates/response` package:

```go
response.TEMPLATE_MANAGER = &templates.Manager{
	// Configure default template settings.
	TplFileSystem:   os.DirFS("templates/"),
	BaseTplSuffixes: []string{".tmpl"},
	BaseTplDirs:     []string{"base"},
	TplDirs:         []string{"templates"},
	UseTplCache:     false,
	DefaultTplFuncs: template.FuncMap{
		"helloworld": func() string {
			return "hello world!"
		},
	},
}
  
```

As you might see from the above code, this follows your file structure.
We do not have to define the regular template directories, but we do have to define the base template directories.
We define the regular directories when rendering them.

```bash
    # The base directory is where the base templates are stored.
    templates/
    ├── base
    │   └── base.tmpl
    └── app
        ├── index.tmpl
        └── user.tmpl
```

Then, we can render templates like so:

```go
func indexFunc(w *http.ResponseWriter, r *http.Request) {
    // Render the template with the given data.
    var err = response.Render(w, "app/index.tmpl", map[string]any{"this":"is_custom_tpl_data"})
    if err != nil {
	req.WriteString(err.Error())
    }
}
```
