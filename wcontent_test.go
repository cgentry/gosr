package gosr


import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"

	"crypto/md5"
	"encoding/base64"
)


func TestSetContent( t * testing.T ){
	c := NewContent()

	Convey( "Make sure setting values" , t, func(){
		c.Set( "ABC" , "text/json")
		So( c.Content , ShouldEqual, "ABC")
		So( c.ContentType , ShouldEqual , "text/json")

		c.Set( "BCD" , "text/json")
		So( c.Content, ShouldEqual , "BCD")
	})
}

func TestSignature_OK( t * testing.T ){
	c := NewContent()
	Convey( "Calculate signature" , t , func(){
		c.Set( "ABC" , "text/json")
		sum1 := c.CalculateSignature()
		So( sum1 , ShouldEqual , c.Signature )

		c.Set( "CBA" , "" )
		So( sum1, ShouldNotEqual, c.CalculateSignature() )
		So( c.CalculateSignature(), ShouldEqual, c.Signature )
	})

	Convey( "Validate" , t , func(){
		c.Set( "1" , "2")
		var sum string = ""
			d := md5.New()
			d.Write([]byte(c.Content))
			m5 := d.Sum(nil)
			sum = base64.StdEncoding.EncodeToString(m5)
		So( sum , ShouldEqual , c.Signature )

		So( c.Verify(), ShouldBeTrue )
	})
}


