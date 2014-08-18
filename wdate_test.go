package gosr


import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"time"
	"net/http"
)


func setDate( minutes int ) string {
	offset := time.Duration(minutes ) * time.Minute
	now := time.Now().UTC();
	return now.Add( offset ).Format( http.TimeFormat)
}

func TestNewDate( t * testing.T ){
	s := NewWDate()
	s.SetNow()

	Convey( "Date should be good" , t , func(){
		err:=s.Verify( 1 )
		So( err, ShouldBeNil )
	})
}


func TestDate_OutsideRange( t * testing.T ){
	var s WDate

	Convey( "Date Future." , t , func(){
		s.Parse( setDate( 20 ) )
		err:=s.Verify( 15 )

		So( err, ShouldNotBeNil )
	})
}

func TestDate_WithinRange( t * testing.T ){
	var s WDate

	Convey( "Date Future Date" , t , func(){
		s.SetNow()
		err:=s.Verify( 1 )
		So( err, ShouldBeNil )
	})
}


