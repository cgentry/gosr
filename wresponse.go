package gosr

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"net/http"
)


type WResponse struct {
	Status			*Error				// General error return
	StatusText		string				// Text response

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
	w.Status	 = http.StatusOK
	w.StatusText = http.StatusText( http.StatusOK)
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
func (w * WResponse) Verify(secret []byte  , timeWindow int) error {
	w.IsVerified = false
	if len(secret) == 0 {
		w.Status = NewErrorWithText(http.StatusInternalServerError, "Secret cannot be zero-length")
	}else {
		w.VerifyElements(timeWindow)
		sig := w.CalculateSignature(secret)

		if !hmac.Equal([]byte(w.Signature), []byte(sig)) {
			w.Status = NewError(http.StatusUnauthorized)
		}else {
			w.IsVerified = true
		}
	}
	return w.Status
}
/*
 * Any base elements that need to be verified should be done here.
 * This will verify just the timestamp and contents
 */
func ( w * WResponse ) VerifyElements( timeWindow int ) error  {
	if err := w.Timestamp.Verify( timeWindow); err != nil {
		w.Status = NewErrorWithText( http.StatusUnauthorized , err.Error() )
	}else{
		if ! w.Content.Verify() {
			w.Status = NewErrorWithText( http.StatusUnauthorized , "Content checksum doesn't match")
		}else if w.Signature == "" {
			w.Status = NewErrorWithText( http.StatusBadRequest , "No Signature is present")
		}
	}
	return w.Status
}

