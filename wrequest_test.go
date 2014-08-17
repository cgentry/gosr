package gosr


import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateRequest( t * testing.T ){
	c := NewWRequest()
	Convey( "Simple create" , t , func(){
		So( c.Signature , ShouldEqual, "" )

		c.Timestamp.Set()
		So( c.Timestamp.Verify( 1 ), ShouldBeNil)

		c.Content.Set("abc" , "test/json")
		So( c.Content.Content, ShouldEqual, "abc")
		So( c.Content.ContentType , ShouldEqual, "test/json")
	})
}
