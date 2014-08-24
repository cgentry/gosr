package gosr

import (
	"time"
	"fmt"
	"net/http"
)

type WDate struct {
	tm		  time.Time
	tmString  string
}

/*
 * Return a new date value initialised to the current time
 */
func NewWDate() *WDate {
	w := new( WDate )
	w.SetNow()
	return w
}

/*
 * Verify that the timestamp is within the timewindow, given in minutes.
 */
func (d * WDate) Verify( timeWindow int ) error {
	var diff int

	now      := time.Now().UTC()
	thisdate := d.tm.UTC()

	if now.Before(thisdate) {
		diff = int( thisdate.Sub(now).Minutes() )
	}else {
		diff = int(now.Sub(thisdate).Minutes())
	}     // We want how far in the past it is...
	if diff > 0 {
		if diff > timeWindow  {
			return fmt.Errorf("%s - %0d min. max/%d difference",
				"Time is outside of window", timeWindow, diff)
		}
	}
	return nil
}

/*
 * Parse a datetime string and set the timestamp to that value
 * Parm: string that matches http.TimeFormat or a standard
 * Return: error or nil
 */
func (d * WDate) Parse(dt string) ( err error ) {
	var t time.Time

	t, err = http.ParseTime(dt)
	if err == nil {
		d.tm = t.UTC()
		d.tmString = dt
	}
	return
}

/*
 * Format the timestamp into an HTTP-standard time string
 * For checksum generation you want to use SourceTime() not Format()
 */
func (d *WDate) Format() string {
	return d.tm.UTC().Format(http.TimeFormat)
}

/*
 * Set the timestamp to the current date/time
 */
func (d *WDate) SetNow() *WDate {
	now := time.Now()
	return d.SetTime( now )
}

/*
 * Set the timestamp to the UTC value of the date/time passed
 */
func ( d * WDate ) SetTime( newTime time.Time ) *WDate {
	d.tmString = newTime.UTC().Format( http.TimeFormat )
	d.tm = newTime.UTC()
	return d
}

/*
 * Return the timetstamp in guaranteed UTC format
 * (This is like a 'GetTime' command but it makes sure the caller nows it is in UTC)
 */
func (d * WDate ) UTC() time.Time {
	return d.tm
}

/*
 * Return the string that set the time. It could be in a different timezone than Format() would
 * return.
 */
func ( d * WDate ) SourceTime() string {
	return d.tmString
}



