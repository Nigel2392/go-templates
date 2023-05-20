package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ResponseStatus string

const (
	ResponseStatusOK       ResponseStatus = "ok"
	ResponseStatusError    ResponseStatus = "error"
	ResponseStatusRedirect ResponseStatus = "redirect"
)

type JSONResponse struct {
	Detail string         `json:"detail,omitempty"`
	Status ResponseStatus `json:"status"`
	Data   interface{}    `json:"data"`
}

// Encode json to a request.
func Json(w http.ResponseWriter, jsonResponse *JSONResponse) error {
	var jsonData, err = json.Marshal(jsonResponse)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	return nil
}

// Render a JSONError to a request.
type JSONError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"` // HTTP status code

	// The error that caused this error
	Err error `json:"-"`
}

func (e *JSONError) Error() string {
	return e.Message
}

func (e *JSONError) Unwrap() error {
	return e.Err
}

// Create a new JSONError
func NewJsonError(message string, statusCode int, err error) *JSONError {
	return &JSONError{
		Message:    message,
		StatusCode: statusCode,
		Err:        err,
	}
}

func (e *JSONError) Write(w http.ResponseWriter) error {
	jsonData, err := json.Marshal(e)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	return nil
}

func JsonError(w http.ResponseWriter, message string, statusCode int, err error) error {
	var response = JSONError{
		Message:    message,
		StatusCode: statusCode,
		Err:        err,
	}

	return response.Write(w)
}

// Render json to a request.
// Response will be in the form of:
//
//	{
//		"status": "ok",
//		"data": {
//			"key": "value"
//		}
//	}
func JsonEncode(w http.ResponseWriter, data interface{}, status ...ResponseStatus) error {
	var response = JSONResponse{
		Data: data,
	}
	if len(status) > 0 {
		response.Status = status[0]
	} else {
		response.Status = ResponseStatusOK
	}
	var jsonData, err = json.Marshal(response)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	return nil
}

// Decoode json from a request, into any.
func JsonDecode(r *http.Request, data interface{}) error {
	// Check header
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("Content-Type is not application/json")
	}
	return json.NewDecoder(r.Body).Decode(data)
}
