package gosr


import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	//"time"
)

func TestAdd_Single_string( t * testing.T ){
	w := NewWParameters()


	Convey( "Simple add" , t , func(){
		err := w.Add( "hi" , "value")
		So( err , ShouldBeNil )
		So( w.Join("hi") , ShouldEqual , "value")
	})
}

func TestAdd_Array_Strings( t * testing.T ){
	w := NewWParameters()
	mstrings := []string{ "one","two","three"}


	Convey( "array add" , t , func(){
		err := w.Add( "hi" , mstrings )
		So( err , ShouldBeNil )
		So( w.Join("hi") , ShouldEqual , "one;two;three")
	})
}

func TestAdd_Map_Strings_Should_fail( t * testing.T ){
	w := NewWParameters()
	mstrings := map[string][]string{ "err" : {"one","two","three"} }

	Convey( "Fail" , t , func(){
		err := w.Add( "hi" , mstrings )
		So( err.Error() , ShouldEqual,"Value of type 'map[string][]string' is invalid" )
	})
}

func TestAdd_Join_Returns_Blank( t * testing.T ){
	w := NewWParameters()

	Convey( "Fail" , t , func(){
		So( w.Join("key") , ShouldEqual, "" )
	})
}


