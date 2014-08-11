package gosr

import (
	"crypto/md5"
	"encoding/base64"
)

type WContent struct {
	ContentType		string				// What is the Type of the main body (mine: text/json...)
	ContentMD5  	string				// MD5 hash of the body
	Content			string				// What is the main body of the request
}

func NewContent() WContent {
	return WContent{ Content : "" }
}

func ( w * WContent ) Set( content , contentFormat string ) ( * WContent ){
	w.Content     = content
	w.ContentType = contentFormat
	w.ContentMD5  = w.CalculateContentMD5();

	return w
}

func ( w * WContent ) CalculateContentMD5() string {
	var sum string = ""
	if len(w.Content) > 0 {
		d := md5.New()
		d.Write([]byte(w.Content))
		m5 := d.Sum(nil)
		sum = base64.StdEncoding.EncodeToString(m5)
	}
	return sum
}


func ( w * WContent ) Verify() bool {
	if w.Content != "" {
		sum := w.CalculateContentMD5()
		return sum == w.ContentMD5
	}
	return true
}
