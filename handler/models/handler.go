package models

import "net/http"

type MyHttpHandler struct{}

func (h *MyHttpHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	
}

type T struct{
	Path string `yaml:"path"`
	Url string `yaml:"url"`
}