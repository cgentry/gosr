package gosr

import (
	"time"
	"fmt"
	"net/http"
	"math"
)

type WDate time.Time

func NewWDate() WDate {
	return WDate( time.Now().UTC() )
}

func ( tm WDate ) Verify( timeWindow * time.Duration ) error {

	// Now...see what the difference is between NOW and the HTTP date
	diff := math.Abs(time.Now().Sub( time.Time( tm ) ).Minutes())        // We want how far in the past it is...

	if diff > timeWindow.Minutes() {
		return fmt.Errorf("%s - %.0f min. max/%.0f in header",
			"Time is outside of window", timeWindow.Minutes(), diff)
	}
	return nil
}


func ( tm WDate ) Parse( dt string ) ( err error ) {
	var t time.Time
	t, err = http.ParseTime( dt )
	if err == nil {
		tm = WDate( t)
	}
	return
}

func ( tm WDate ) Format( ) string {
	return time.Time( tm ).UTC().Format( http.TimeFormat )
}

func ( tm WDate ) Set() {
	tm = WDate( time.Now().UTC() )
}


