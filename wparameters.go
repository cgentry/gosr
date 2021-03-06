package gosr

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"
)

type WParameters map[string][]string

func NewWParameters() WParameters {
	return make( WParameters )
}

func (v WParameters) Get(key string) string {
	if v == nil {
		return ""
	}
	vs, ok := v[key]
	if !ok || len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// Set sets the key to value. It replaces any existing
// values.
func (v WParameters) Set(key, value string) {
	v[key] = []string{value}
}

//func (v WParameters) Add(key, value string) {
//	v[key] = append(v[key], value)
//}

// Del deletes the values associated with key.
func (v WParameters) Del(key string) {
	delete(v, key)
}

/*
 * Get a list of all the options, sorted in order
 */
func (w WParameters) SortedKeys() ( optionKeys []string ) {
	for key, _ := range w {
		optionKeys = append(optionKeys, key)
	}
	sort.Strings(optionKeys)

	return
}

func ( w WParameters ) IsPresent( parms []string ) * Error {
	for _,key := range parms {
		if _,found := w[key] ; ! found {
			return NewErrorWithText( http.StatusBadRequest , "Missing " + key + " parameter")
		}
	}
	return nil
}

/*
 * Join will take all the strings for a single key and join them with a semi-colon
 */
func (w WParameters) Join(key string) ( joined string) {
	if _,ok := w[key]; ok {
		joined = strings.Join( w[key] , `;`)
	}
	return
}

var typeTime 			reflect.Type
var typeString  		reflect.Type
var typeStringArray		reflect.Type
var typeStringMap		reflect.Type

func (w WParameters) Add(key string , value interface{}) ( err error ) {

	typeOf := reflect.TypeOf( value )
	switch typeOf{
	default:
		err = fmt.Errorf( "Value of type '%s' is invalid" ,reflect.TypeOf(value) )

	case typeString :
		w[ key ] = append( w[key] , value.(string) )

	case typeStringArray :
		s := reflect.ValueOf(value)

		for i := 0; i < s.Len(); i++ {
			if err = w.Add( key , s.Index(i).Interface() ); err != nil {
				break
			}
		}

	//case typeStringMap :

	}

	return
}

func init() {
	typeString 		= reflect.TypeOf( "" )
	typeStringArray = reflect.TypeOf( []string{})
	typeStringMap   = reflect.TypeOf( map[string]string{})
	typeTime 		= reflect.TypeOf( time.Time{} )
}
