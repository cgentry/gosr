package http

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	//"net/http"
	"net/http/httptest"
	"fmt"
	"encoding/json"

)

func TestSimpleResponse_WithValue(t * testing.T) {

	w := httptest.NewRecorder()
	s := NewResponse()

	s.Content.Set( `TestBody` , "text/html")

	s.Encode( w )

	Convey("No blanks, Should be 12345", t, func() {
		So( s.Content.Signature , ShouldNotEqual, "")
		So( w.Header().Get( HEADER_TYPE ) , ShouldEqual , `text/html`)
		So( w.Header().Get( HEADER_MD5  ) , ShouldEqual ,  s.Content.Signature )
		So( w.Body.String(), ShouldEqual, `TestBody`)
		fmt.Println( "ok")
	})
}

func TestSimpleResponse_LongContent(t * testing.T) {

	w := httptest.NewRecorder()
	s := NewResponse()

	testData := "Hello\nThis is a long test of data\nAnd we should be able to handle it\n"
	s.Content.Set( testData , "text/text")

	s.Encode( w )

	Convey("No blanks, Should be 12345", t, func() {
		So( s.Content.Signature , ShouldNotEqual, "")
		So( w.Header().Get( HEADER_TYPE ) , ShouldEqual , `text/text`)
		So( w.Header().Get( HEADER_MD5  ) , ShouldEqual ,  s.Content.Signature )
		So( w.Body.String(), ShouldEqual, testData )
		fmt.Println( "ok")
	})
}

func TestSimpleResponse_JsonContent(t * testing.T) {

	w := httptest.NewRecorder()
	s := NewResponse()

	jdata := map[string]int{
		"a" : 1 , "b" : 2 ,
		"c" : 3 , "d" : 4  }

	testData, err  := json.Marshal( jdata )
	s.Content.Set(string(testData) , "text/text")

	s.Encode( w )

	Convey("No blanks, Should be 12345", t, func() {
		So( err , ShouldBeNil )
		So( s.Content.Signature , ShouldNotEqual, "")
		So( w.Header().Get( HEADER_TYPE ) , ShouldEqual , `text/text`)
		So( w.Header().Get( HEADER_MD5  ) , ShouldEqual ,  s.Content.Signature )
		So( w.Body.String(), ShouldEqual, string(testData) )
		fmt.Println( "ok")
	})
}
