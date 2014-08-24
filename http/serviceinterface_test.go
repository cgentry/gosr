package http

import (
	"testing"
	"github.com/cgentry/gosr"
)
/*
 * Simple test to make sure each of the high level types conforms to a signature
 */

func ensure_type( sr gosr.ServiceInterface  ){
	return
}


func TestHttp( t * testing.T ){
	w := NewRequest()
	ensure_type( w )
}
