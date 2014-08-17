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

func NewWDate() WDate {
	return WDate(time.Now().UTC())
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
	d.original = dt
	t, err = http.ParseTime(dt)
	if err == nil {
		d.tm = (WDate)(t.UTC())
	}
	return
}

/*
 * Format the timestamp into an HTTP-standard time string
 */
func (d *WDate) Format() string {
	return d.tm.UTC().Format(http.TimeFormat)
}

/*
 * Set the timestamp to the current date/time
 */
func (d *WDate) Set() WDate {
	d.tm = WDate(time.Now().UTC())
	return d
}

/*
 * Set the timestamp to the UTC value of the date/time passed
 */
func ( d * WDate ) SetTime( newTime time.Time ) *WDate {
	d.tm = WDate( newTime.UTC() )
	return d
}

/*
 * Return the timetstamp in guaranteed UTC format
 */
func (d * WDate ) UTC() time.Time {
	return ((time.Time)(d.tm)).UTC()
}

func ( d * WDate ) SourceTime() string {
	return d.tmString
}



