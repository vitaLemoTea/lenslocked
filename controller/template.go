package controller

import "net/http"

type Template interface {
	Execute(w http.ResponseWriter,r *http.Request, data interface{})error
}