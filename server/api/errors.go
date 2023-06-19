package api

import (
	"net/http"

	"github.com/go-chi/render"
)

type Data map[string]any

type Error struct {
	HTTPStatusCode int    `json:"-"`
	ErrText        string `json:"error,omitempty"`
	Data           Data   `json:"data,omitempty"`
}

func (err *Error) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, err.HTTPStatusCode)
	return nil
}

func NewDerivedErrorRendererWithData(err error, statusCode int, data Data) render.Renderer {
	return &Error{
		HTTPStatusCode: statusCode,
		ErrText:        err.Error(),
		Data:           data,
	}
}

func NewDerivedErrorRenderer(err error, statusCode int) render.Renderer {
	return NewDerivedErrorRendererWithData(err, statusCode, nil)
}

func NewTextErrorRendererWithData(errText string, statusCode int, data Data) render.Renderer {
	return &Error{
		HTTPStatusCode: statusCode,
		ErrText:        errText,
		Data:           data,
	}
}

func NewTextErrorRenderer(errText string, statusCode int) render.Renderer {
	return NewTextErrorRendererWithData(errText, statusCode, nil)
}
