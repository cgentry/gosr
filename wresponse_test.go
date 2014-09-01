package gosr


import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreatResponse_goodComplex( t * testing.T ){
	var err error
	c := NewWResponse()
	secret := []byte( "abceasy")
	Convey( "Check with good date" , t , func(){
		c.Timestamp.SetNow()
		c.Content.Set("abc"       , "test/json")
		So( c.Content.Content     , ShouldEqual, "abc")
		So( c.Content.ContentType , ShouldEqual, "test/json")

		c.Timestamp.SetNow()
		So( c.IsVerified, ShouldBeFalse )
		c.Sign( secret )

		err = c.Verify( secret , 1 )
		So( c.IsVerified , ShouldBeTrue )
		So( err , ShouldBeNil )

	})
}

func TestCreatResponse( t * testing.T ){
	c := NewWResponse()
	secret := []byte( "abceasy")
	Convey( "Simple create" , t , func(){
		So( c.Signature , ShouldEqual, "" )

		c.Timestamp.SetNow()
		So( c.Timestamp.Verify( 1 ), ShouldBeNil)

		c.Content.Set("abc" , "test/json")
		So( c.Content.Content, ShouldEqual, "abc")
		So( c.Content.ContentType , ShouldEqual, "test/json")

		So( c.IsVerified, ShouldBeFalse )
		c.Sign( secret )

		err := c.Verify( secret , 1 )
		So( err , ShouldBeNil )
		So( c.IsVerified , ShouldBeTrue )
	})
}

func TestCreatResponse_Fail( t * testing.T ){
	var err error
	c := NewWResponse()
	Convey( "Failure create" , t , func(){
		So( c.Signature , ShouldEqual, "" )

		err = c.Timestamp.Parse( "Mon Jan  2 15:04:05 2006")
		So( err, ShouldBeNil )
		So( c.Timestamp.Verify( 1 ), ShouldNotBeNil)
	})
}





