package http

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"strings"
	"time"
)

func TestGetUser_WithValue(t * testing.T) {

	testData := ""
	r, err := http.NewRequest("POST", "http://example.com/test", strings.NewReader(testData))
	r.Header.Set( HEADER_TIMESTAMP , time.Now().UTC().Format( http.TimeFormat ))

	s := NewRequest()

	Convey("No blanks, Should be 12345", t, func() {
		So(err, ShouldBeNil)
		r.Header.Set("Authorization", "12345:abcde")
		err := s.Decode( r , "" )				// Copy the info over
		So(err, ShouldBeNil)
		val := s.GetUser()

		So(val, ShouldEqual, "12345")
	})
}

func TestUserParams( t * testing.T ){
	testData := ""
	r, _ := http.NewRequest("POST", "http://example.com/test/rqst?a=b&c=d", strings.NewReader(testData))
	r.Header.Set( HEADER_TIMESTAMP , time.Now().UTC().Format( http.TimeFormat ))
	r.Header.Set( "Authorization" , "12345:abcde")

	s := NewRequest()
	err := s.Decode( r , "Extra")

	Convey( "Make sure parms were copied" , t , func() {
		So( err , ShouldBeNil )
		So( s.Action , ShouldEqual , "POST")
		So( s.Operation, ShouldEqual , "/test/rqst")
		So( s.RawQuery , ShouldEqual , "a=b&c=d")
	})
}
