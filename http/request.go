package http

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"github.com/cgentry/gosr"
	"io/ioutil"
	"net/http"
	"strings"
	"sort"
)

type Request struct {
	gosr.WRequest
	RawQuery string
	RawParam []string
}



/* ===============================================================
 *               PRIVATE FUNCTIONS
 * ===============================================================
 */
func (s * Request) decodeAuth(r * http.Request) error {

	parts := strings.SplitN(r.Header.Get(`Authorization`), ":", 2)
	if len(parts) < 2 {
		return errors.New(TOKEN_INCOMPLETE)
	}
	s.User = strings.TrimSpace(parts[0])
	s.Signature = strings.TrimSpace(parts[1])
	if len(s.WRequest.User) == 0 || len(s.Signature) == 0 {
		return errors.New(TOKEN_MISSING_PARM)
	}
	return nil
}
func (s * Request) getUri(r * http.Request) string {
	// Reconstitued version
	// The RawURI does not include the fragment so we need to build it here...
	val := r.URL.Path
	if r.URL.RawQuery != "" {
		val = val+`?`+r.URL.RawQuery
	}
	if r.URL.Fragment != "" {
		val = val+`#`+r.URL.Fragment
	}
	return val
}

/* ===============================================================
 *               PUBLIC FUNCTIONS
 * ===============================================================
 */

func NewRequest() * Request {
	rtn :=  new( Request )
	return rtn
}

/*
 * The HTTP header date can be either encoded as Date: or Timestamp:
 * The preferred method is Timestamp: as some systems will force Date: to be set.
 */
func (s * Request) GetHttpDateString(r * http.Request) ( string , error ) {
	requestDate := r.Header.Get(HEADER_TIMESTAMP)        // Header has "Timestamp:"
	if len(requestDate) == 0 {                    // Umm..NO
		requestDate = r.Header.Get(HEADER_DATE)        // Header has "Date:" ?

		if len(requestDate) == 0 {
			return "", errors.New(TIMESTAMP_MISSING)
		}
	}
	return requestDate, nil
}

/*
 * This will copy the contents of the http request over to the wrequest
 */
func (s * Request) CopyContent(r * http.Request ) ( err error ) {
	defer r.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(r.Body)
	if err == nil {
		s.Content.Content = string(body)					// Copy the body over
		s.Content.ContentType = r.Header.Get(HEADER_TYPE)   // Content type
		s.Content.Signature = r.Header.Get(HEADER_MD5)      // Signature of content
		if !s.Content.Verify() {                            // + and verify
			err = errors.New(MD5_MISMATCH)
		}
	}
	return
}

// Copy the list of desired headers from http.request to our parameters array
func ( s * Request ) CopyParameters( r * http.Request, extraPrefix string ) {
	s.RawParam = nil
	r.ParseForm()
	// First, copy desired parameters from the header block
	for key,value := range r.Header {
		if strings.HasPrefix( key , PARAMETER_QUERY ){
			s.RawParam = append( s.RawParam , key + ":" + strings.Join( value , `,`) )
			s.Parameters.Add( key , value )
		}
	}
	// URL.Values may have multiple values per-key

	for key,value := range r.Form {
		s.Parameters.Add( key , value )
	}
	sort.Strings( s.RawParam )
	return
}


/**
 * Create a signature value from the request, user and secret and body
 */
func ( s * Request ) CalculateSignature( secret []byte  ) ( string , error ){

	if len( secret ) == 0 {
		return "", errors.New( SECRET_INVALID )
	}

	mac := hmac.New( sha256.New , secret )					// Setup with secret key
	mac.Write( []byte( strings.TrimSpace(s.User) ) )		// Add in user ID
	mac.Write( []byte( s.Timestamp.SourceTime() ) )			// Add in DATE

	mac.Write( []byte( s.Content.CalculateSignature() ))	// Add in the MD5 calculate value
	mac.Write( []byte( s.Content.ContentType))				// Add in Content-Type from header

	mac.Write( []byte( s.Action ) )							// Get/Put
	mac.Write( []byte( s.Operation ))						// fragment
	mac.Write( []byte( s.RawQuery ))						// a=b&c=d....

	for _ , v := range s.RawParam {							// Add in all of the raw parameters
		mac.Write( []byte( v ) )
	}

	return base64.StdEncoding.EncodeToString(mac.Sum( nil ) ), nil
}

func (s * Request) Verify(secret []byte , timeWindow int ) ( err error ) {

	if s.User == "" {
		err = errors.New(TOKEN_MISSING_PARM)
	}else if err = s.VerifyElements(timeWindow) ; err == nil {
		var computed string
		if computed, err = s.CalculateSignature(secret); err == nil {
			if !hmac.Equal([]byte(s.Signature), []byte(computed)) {
				err = errors.New(SIGNATURE_INVALID)
			}
		}

	}
	s.IsVerified = ( err == nil )
	return
}

/*
 * Decode will copy all the relevant information over to the request
 * and move the body over. All the data is set at once.
 * This will perform the verification at the same time
 */
func (s * Request) Decode(r * http.Request , extraPrefix string ) ( err error ) {
	s.IsVerified = false
	if err = s.decodeAuth(r); err == nil {                            // Set UserID and Signature
		var dt string
		if dt, err = s.GetHttpDateString(r) ; err == nil {            // Get the time string
			if err = s.Timestamp.Parse(dt); err == nil {			  // ..and set the timestamp
				if err = s.CopyContent(r)    ; err == nil {           // Set Content,type,
					s.CopyParameters( r , extraPrefix )				  // Copy all the parameters
					
					s.Action    = r.Method							  // Get/Put....
					s.Operation = r.URL.Path				      	  // Path

					// Add in the 'extra' parameters
				}
			}
		}
	}
	return
}



