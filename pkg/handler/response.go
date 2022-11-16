package handler

import (
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
)

var _ render.Renderer = (*ErrResponse)(nil)

type ErrResponse struct {
	HTTPStatusCode int    `json:"status"`
	ErrorText      string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func NewErrResponse(HTTPStatusCode int, errorText string) *ErrResponse {
	return &ErrResponse{HTTPStatusCode: HTTPStatusCode, ErrorText: errorText}
}

func RenderErrResponse(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	errResponse := NewErrResponse(statusCode, err.Error())
	if err := render.Render(w, r, errResponse); err != nil {
		logrus.Fatalf("error while rendering: %s", err.Error())
	}
}
