package http

import (
	"net/http"
	"fmt"
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
	r.WResponse.Initialise()
	return r
}

func ( r * Response ) Encode( w http.ResponseWriter ) {

	// We need to copy all the header values over
	for key,values := range *r.WResponse.Parameters {
		for _,value := range values {
			w.Header().Add( key , value )
		}
	}
	w.Header().Set( HEADER_MD5  , r.Content.Signature )
	w.Header().Set( HEADER_TYPE , r.Content.ContentType )
	w.Header().Set( HEADER_TOKEN, r.GetSignature()      )
	w.Header().Set( HEADER_STATUS , r.StatusText )

	w.WriteHeader( r.Status )

	// LAST STEP...
	fmt.Fprintf( w , "%s" , r.Content.Content )
	return
}

func ( r * Response ) SetError( err * gosr.Error ) * Response{
	r.Status = err.Status()
	r.StatusText = err.Error()
	return r
}

