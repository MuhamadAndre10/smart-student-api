package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

type WebResponse[T any] struct {
	Data    T             `json:"data"`
	Message string        `json:"message,omitempty"`
	Token   string        `json:"token,omitempty"`
	Paging  *PageMetadata `json:"paging,omitempty"`
	Errors  string        `json:"errors,omitempty"`
}

type PageResponse[T any] struct {
	Data         []T          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"paging,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}

func DecodeJson[T any](w http.ResponseWriter, r *http.Request, data T) error {

	maxBytes := 10 << 20

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return err
	}

	return nil
}

func WriteJsonBody(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func WriteErrorJson(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusInternalServerError

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload WebResponse[any]
	payload.Errors = err.Error()
	payload.Data = nil

	return WriteJsonBody(w, statusCode, payload)
}
