package gosr

import (
	"crypto/md5"
	"encoding/base64"
)

type WContent struct {
	ContentType		string				// What is the Type of the main body (mine: text/json...)
	Signature  		string				// MD5 hash of the body
	Content			string				// What is the main body of the request
}

func NewContent() *WContent {
	return new( WContent )
}

/*
 * Set the content and format. Calculate the checksum
 */
func ( w * WContent ) Set( content , contentFormat string ) ( * WContent ){
	w.ContentType = contentFormat
	w.SetContent( content )

	return w
}

func ( w * WContent ) SetContent( content string ) * WContent {
	w.Content = content
	w.Signature = w.CalculateSignature()
	return w
}

func ( w * WContent ) GetContent( content string ) string {
	return w.Content
}

func ( w * WContent ) SetContentType( contentType string ) *WContent {
	w.ContentType = contentType
	return w
}
func ( w * WContent ) GetContentType( contentType string ) string {
	return w.ContentType
}

/*
 * Calculate the signature of the content.
 * Return: base64-encoded MD5 hash of the body
 */
func ( w * WContent ) CalculateSignature() string {
	var sum string = ""
	if len(w.Content) > 0 {
		d := md5.New()
		d.Write([]byte(w.Content))
		m5 := d.Sum(nil)
		sum = base64.StdEncoding.EncodeToString(m5)
	}
	return sum
}

// Sign - sign the content
func ( w *WContent ) Sign() * WContent {
	w.Signature = w.CalculateSignature()
	return w
}

/*
 * Verify that the signature checksum is valid. If content is empty, true is returned.
 * Return: bool - true if valid, false if not
 */
func ( w * WContent ) Verify() bool {
	if w.Content != "" {
		sum := w.CalculateSignature()
		return sum == w.Signature
	}
	return true
}
