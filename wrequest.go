package gosr

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"sort"
	"strings"
	"time"
)

/*
 * This is the Service Request definition for remote reqeuests into a service provider.
 * The source of the request can come from anywhere, but all the data must be filled in
 * in order to pass the data around.
 */

type WRequest struct {
	User			string				// The user-id making the call
	Action			string				// What is the request being made
	Subaction   	string				// This is not usually used, but would be the 'fragment' in a URL
	Signature		string				// The HMAC of the request

	Parameters  	map[string]string   // Action parameters

	Timestamp   	WDate				// When this request was made

	Content			WContent			// Data for content

}



func ( w *WRequest ) Append( m map[string]string ) *WRequest {
	for key, value := range m {
		w.Parameters[ key ] = value
	}
	return w
}



func ( w * WRequest ) CalculateSignature( secret []byte , timeWindow time.Duration) ( string , error ){
	if len( secret ) == 0 {
		return "", errors.New( "Secret cannot be zero-length" )
	}
	if err := w.Timestamp.Verify( timeWindow ); err != nil {
		return "" , err
	}
	if ! w.Content.Verify() {
		return "" , errors.New( "Content checksum doesn't match")
	}

	mac := hmac.New( sha256.New , secret )						// Setup with secret key
	mac.Write( []byte( strings.TrimSpace(w.User) ) )			// + Add in user ID
	mac.Write( []byte( w.Timestamp.Format()))					// + in date string
	mac.Write( []byte( w.Content.ContentMD5()))					// + MD5 calculate value
	mac.Write( []byte( w.Content.ContentType))					// + Content-Type f
	mac.Write( []byte( w.Action ))								// + What is the action?
	mac.Write( []byte( w.SubAction ))
	//															... add in all the parameters
	optionKeys := keys( w.Parameters )
	sort.String( optionKeys )
	for _,key := range optionKeys {
		mac.Write( []byte( key + w.Parameters[key]))
	}

	return base64.StdEncoding.EncodeToString(mac.Sum( nil ) ), nil
}

func ( w * WRequest ) VerifySignature( secret []byte ) error {
	sig,err := w.CalculateSignature( secret  )
}
