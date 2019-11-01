package view

import (
	"encoding/json"
	"net/http"
)

type JsonResponseView struct {
	writer http.ResponseWriter
}

func InitJsonResponseView(w http.ResponseWriter) *JsonResponseView {
	return &JsonResponseView{
		writer: w,
	}
}

func (v *JsonResponseView) SendResponse(statusCode int, output interface{}) {
	v.writer.Header().Set("Content-Type", "application/json")
	v.writer.WriteHeader(statusCode)
	if output != nil {
		json.NewEncoder(v.writer).Encode(output)
	}
}
