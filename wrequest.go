package gosr

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
)

/*
 * This is the Service Request definition for remote reqeuests into a service provider.
 * The source of the request can come from anywhere, but all the data must be filled in
 * in order to pass the data around.
 */

type WRequest struct {
	IsVerified 		bool				// Signature is verified
	Signature		string				// The HMAC signature for all elements

	User			string				// The user-id making the call

	Action			string				// Get/Put/Delete/ etc
	Operation       string				// What task do you want done? Register, Store, etc.

	Timestamp   	WDate				// When this request was made
	Content			WContent			// Data for content

	Parameters		WParameters			// A map, keyed by a string of an array of strings
										// Parameters can have multiple values

}

func NewWRequest() WRequest {
	return WRequest{ IsVerified : false }
}

func ( w * WRequest ) GetUser() string {
	return strings.TrimSpace(w.User)
}


func ( w * WRequest ) CalculateSignature( secret []byte ) ( string , error ){

	mac := hmac.New( sha256.New , secret )						// Setup with secret key
	mac.Write( []byte( w.GetUser() ) )							// + Add in user ID
	mac.Write( []byte( w.Timestamp.SourceTime()))				// + in date string
	mac.Write( []byte( w.Content.Signature))					// + MD5 value of content (as stored)
	mac.Write( []byte( w.Content.ContentType))					// + Content-Type
	mac.Write( []byte( w.Action ))								// + Action string
	mac.Write( []byte( w.Operation ))							// + what operation to perform

	// Now...parameter strings
	sortedKeys := w.Parameters.SortedKeys()
	for _,key := range sortedKeys {
		mac.Write( []byte( key + ":" + w.Parameters.JoinValues( key )))
	}

	return base64.StdEncoding.EncodeToString(mac.Sum( nil ) ), nil
}

/*
 * Any base elements that need to be verified should be done here.
 * This will verify just the timestamp and contents
 */
func ( w * WRequest ) VerifyElements( timeWindow int ) ( err error ) {
	if err = w.Timestamp.Verify( timeWindow); err == nil {
		if ! w.Content.Verify() {
			err = errors.New( "Content checksum doesn't match")
		}else if w.Signature == "" {
			err = errors.New( "No Signature is present")
		}
	}
	return
}

func ( w * WRequest ) Verify( secret []byte  , timeWindow int ) error {
	w.IsVerified = false
	if len( secret ) == 0 {
		return errors.New( "Secret cannot be zero-length" )
	}
	w.VerifyElements( timeWindow )
	sig,err := w.CalculateSignature( secret )
	if err != nil {
		return err
	}
	if ! hmac.Equal([]byte(w.Signature ), []byte( sig)){
		return errors.New("Signature doesn't verify")
	}
	w.IsVerified = true
	return nil
}
