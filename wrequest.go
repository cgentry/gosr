package gosr

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"sort"
	"strings"
)

/*
 * This is the Service Request definition for remote reqeuests into a service provider.
 * The source of the request can come from anywhere, but all the data must be filled in
 * in order to pass the data around.
 */

type WRequest struct {
	IsVerified 		bool
	User			string				// The user-id making the call
	Action			string				// What is the request being made
	Subaction   	string				// This is not usually used, but would be the 'fragment' in a URL
	Signature		string				// The HMAC of the request

	Parameters  	map[string]string   // Action parameters
	Timestamp   	WDate				// When this request was made
	Content			WContent			// Data for content

}

func NewWRequest() WRequest {
	return WRequest{ IsVerified : false }
}

func ( w *WRequest ) Append( m map[string]string ) *WRequest {
	for key, value := range m {
		w.Parameters[ key ] = value
	}
	return w
}

func ( w * WRequest ) GetAction() string {
	return w.Action + w.Subaction
}

func ( w * WRequest ) GetParameterString()( parm string ) {
	var optionKeys []string

	for key,_ := range w.Parameters {
		optionKeys = append( optionKeys , key )
	}
	sort.Strings( optionKeys )
	for _,key := range optionKeys {
		parm = parm + key + w.Parameters[key]
	}
	return
}

func ( w * WRequest ) GetUser() string {
	return strings.TrimSpace(w.User)
}


func ( w * WRequest ) CalculateSignature( secret []byte ) ( string , error ){

	mac := hmac.New( sha256.New , secret )						// Setup with secret key
	mac.Write( []byte( w.GetUser() ) )							// + Add in user ID
	mac.Write( []byte( w.Timestamp.Format()))					// + in date string
	mac.Write( []byte( w.Content.Signature))					// + MD5 calculate value of content
	mac.Write( []byte( w.Content.ContentType))					// + Content-Type
	mac.Write( []byte( w.GetAction() ))							// + Action string
	mac.Write( []byte( w.GetParameterString()))					// + add in all the parameters
	return base64.StdEncoding.EncodeToString(mac.Sum( nil ) ), nil
}

func ( w * WRequest ) Verify( secret []byte , timeWindow int ) error {
	s.IsVerified = false
	if len( secret ) == 0 {
		return errors.New( "Secret cannot be zero-length" )
	}
	if err := w.Timestamp.Verify( timeWindow); err != nil {
		return err
	}
	if ! w.Content.Verify() {
		return errors.New( "Content checksum doesn't match")
	}
	sig,err := w.CalculateSignature( secret )
	if err != nil {
		return err
	}
	if ! hmac.Equal([]byte(w.Signature ), []byte( sig)){
		return errors.New("Signature doesn't verify")
	}
	s.IsVerified = true
	return nil
}
