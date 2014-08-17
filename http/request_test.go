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
		err := s.Decode( r )				// Copy the info over
		So(err, ShouldBeNil)
		val := s.GetUser()

		So(val, ShouldEqual, "12345")
	})
}
