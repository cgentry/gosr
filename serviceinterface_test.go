package gosr

import (
	"testing"
)
/*
 * Simple test to make sure each of the high level types conforms to a signature
 */

func ensure_type( sr ServiceInterface  ){
	return
}


func TestWRequest( t * testing.T ){

	w := NewWRequest()
	ensure_type( w )
}

