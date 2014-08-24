package gosr

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)


type WResponse struct {
	IsVerified		bool
	Signature		string				// The HMAC signature for all elements

	Timestamp   	*WDate				// When this request was made
	Content			WContent			// Data for content

	Parameters		*WParameters			// A map, keyed by a string of an array of strings
	// Parameters can have multiple values

}


func NewWResponse() *WResponse {
	w := new( WResponse )
	w.Initialise()
	return w
}

func( w * WResponse ) Initialise() {
	w.IsVerified = false
	w.Timestamp  = NewWDate()
	w.Parameters = new(WParameters)
}

/* ------------------------------------------------
 * Getters and setters are here
 * ------------------------------------------------
 */

// GetSignature - Return the signature that was sent to us
func ( w * WResponse ) GetSignature() string {
	return w.Signature
}

/* ------------------------------------------------
 * Interface requirements
 * ------------------------------------------------
 */

// CalculateSignature - Create an HMAC signed value from the data in the request block
// This does not use the full content body but only the content signature for the hash
func ( w * WResponse ) CalculateSignature( secret []byte ) string {

	mac := hmac.New( sha256.New , secret )						// Setup with secret key
	mac.Write( []byte( w.Timestamp.SourceTime()))				// + in date string
	mac.Write( []byte( w.Content.Signature))					// + MD5 value of content (as stored)
	mac.Write( []byte( w.Content.ContentType))					// + Content-Type

	// Now...parameter strings
	sortedKeys := w.Parameters.SortedKeys()
	for _,key := range sortedKeys {
		mac.Write( []byte( key + ":" + w.Parameters.Join( key )))
	}

	return base64.StdEncoding.EncodeToString(mac.Sum( nil ) )
}

func( w * WResponse ) Sign( secret []byte ) * WResponse {
	w.Content.Sign()
	w.Signature = w.CalculateSignature( secret )
	return w
}

// Verify - Verify that the signature in the request block is the same as one calculated
func ( w * WResponse ) Verify( secret []byte  , timeWindow int ) error {
	w.IsVerified = false
	if len( secret ) == 0 {
		return errors.New( "Secret cannot be zero-length" )
	}
	w.VerifyElements( timeWindow )
	sig := w.CalculateSignature( secret )

	if ! hmac.Equal([]byte(w.Signature ), []byte( sig)){
		return errors.New("Signature doesn't verify")
	}
	w.IsVerified = true
	return nil
}
/*
 * Any base elements that need to be verified should be done here.
 * This will verify just the timestamp and contents
 */
func ( w * WResponse ) VerifyElements( timeWindow int ) ( err error ) {
	if err = w.Timestamp.Verify( timeWindow); err == nil {
		if ! w.Content.Verify() {
			err = errors.New( "Content checksum doesn't match")
		}else if w.Signature == "" {
			err = errors.New( "No Signature is present")
		}
	}
	return
}

