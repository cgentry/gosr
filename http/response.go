package http

import (
	"net/http"
	"github.com/cgentry/gosr"
)
type Response struct {
	gosr.WResponse
}

func NewResponse() * Response {
	r := new( Response )
	return r.Initialise()
}

func ( r * Response ) Initialise () * Response {
	return r
}

func ( r * Response ) Encode( rhttp * http.Response ) {
	return
}

