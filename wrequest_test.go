package gosr


import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateRequest( t * testing.T ){
	c := NewWRequest()
	Convey( "Simple create" , t , func(){
		So( c.Signature , ShouldEqual, "" )

		c.Timestamp.SetNow()
		So( c.Timestamp.Verify( 2 ), ShouldBeNil)

		c.Content.Set("abc" , "test/json")
		So( c.Content.Content, ShouldEqual, "abc")
		So( c.Content.ContentType , ShouldEqual, "test/json")
	})
}
