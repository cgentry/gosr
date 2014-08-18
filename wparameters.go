package gosr

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

type WParameters map[string][]string

func NewWParameters() WParameters {
	return WParameters{}
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

/*
 * Each parameter can be multiple values. This will return a string with all the values
 * comma separated.
 */
func (w WParameters) JoinValues(key string) ( joined string) {
	for _, value := range (w)[key] {
		if joined == "" {
			joined = value
		}else {
			joined = joined+","+value
		}
	}
	return
}

//typeTime := reflect.TypeOf( time.Time{} )
var typeString  reflect.Type
var typeStringArray	reflect.Type
var 
func (w WParameters) Add(key string , value interface{}) ( err error ) {

	typeOf := reflect.TypeOf( value )
	switch typeOf{
	default:
		fmt.Printf( "Type is %s\n" , reflect.TypeOf(value) )
	case typeString :
		fmt.Printf( "type is a string\n")

	case reflect.TypeOf( time.Time{}):
		fmt.Printf( "type is time\n");

	}
	switch reflect.TypeOf(value).Kind() {

	default:
		err = fmt.Errorf( "Value of type '%T' is invalid" , value )

	case reflect.String :
		w[key] = append(w[ key ], value.(string))

	case reflect.Slice:
		s := reflect.ValueOf(value)

		for i := 0; i < s.Len(); i++ {
			if err = w.Add( key , s.Index(i).Interface() ); err != nil {
				break
			}
		}
	}

	return
}

func init() {
	typeString = reflect.TypeOf( "" )
}
