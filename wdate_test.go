package gosr


import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"time"
	"net/http"
)


func setDate( minutes int ) string {
	offset := time.Duration(minutes) * time.Minute
	return time.Now().UTC().Add( offset ).Format( http.TimeFormat)
}

func TestNewDate( t * testing.T ){
	s := NewWDate()
	duration, err := time.ParseDuration( "1 sec")

	Convey( "Date should be good" , t , func(){
		So( err, ShouldBeNil )
		err:=s.Verify( duration )
		So( err, ShouldBeNil )
	})
}

func TestDate_OutsideRange( t * testing.T ){
	var s WDate

	Convey( "Date Future Date" , t , func(){
		s.Parse( setDate( 20 ))

		err:=s.Verify( s.Verify( &time.Duration(15 )) )
		So( err, ShouldNotBeNil )

	})
}

