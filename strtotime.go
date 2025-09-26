package utils

import (
	"fmt"
	"strconv"
	"time"
)

var TimeFormats = []string{
	time.Layout,   // 01/02 03:04:05PM '06 -0700
	time.ANSIC,    // Mon Jan _2 15:04:05 2006
	time.UnixDate, // Mon Jan _2 15:04:05 MST 2006
	time.RubyDate, // Mon Jan 02 15:04:05 -0700 2006
	time.RFC822,   // 02 Jan 06 15:04 MST
	time.RFC822Z,  // 02 Jan 06 15:04 -0700
	time.RFC850,   // Monday, 02-Jan-06 15:04:05 MST
	time.RFC1123,  // Mon, 02 Jan 2006 15:04:05 MST
	time.RFC1123Z, // Mon, 02 Jan 2006 15:04:05 -0700
	"2006-01-02T15:04Z",
	"2006-01-02T15:04:5Z",
	"2006-01-02T15:04:05Z",
	"2006-01-02T15:04:05.9Z",
	"2006-01-02T15:04:05.99Z",
	"2006-01-02T15:04:05.999Z",
	"2006-01-02T15:04:05.9999Z",
	"2006-01-02T15:04:05.99999Z",
	"2006-01-02T15:04:05.999999Z",
	"2006-01-02T15:04:05.9999999Z",
	"2006-01-02T15:04:05.99999999Z",
	"2006-01-02T15:04:05.999999999Z",
	time.RFC3339,     // 2006-01-02T15:04:05Z07:00
	time.RFC3339Nano, // 2006-01-02T15:04:05.999999999Z07:00
	time.Kitchen,     // 3:04PM
	time.Stamp,       // Jan _2 15:04:05
	time.StampMilli,  // Jan _2 15:04:05.000
	time.StampMicro,  // Jan _2 15:04:05.000000
	time.StampNano,   // Jan _2 15:04:05.000000000
	time.DateTime,    // 2006-01-02 15:04:05
	time.DateOnly,    // 2006-01-02
	"2006-01",
	"06-1-2",
	"06-1",
	"01-02-2006",
	"01-2006",
	"1-2-06",
	"1-06",

	"2006/01/02",
	"2006/01",
	"01/02/2006",
	"01/2006",
	"06/1/2",
	"06/1",
	"1/2/06",
	"1/06",

	"2006",
	"2006.01.02",
	"2006.01",
	"06.1.2",
	"06.1",
	"01.02.2006",
	"01.2006",
	"1.2.06",
	"1.2.2006",
	"1.2.06",
	"1.06",

	// "20060102130405",
	// "2006010213045",
	// "200601021345",
	// "200601021304",
	// "20060102134",
	// "2006010213",
	// "20060102",
	// "200601",
	// "0612",
	// "612"
	time.TimeOnly, // 15:04:05
	"15:04:5",
	"15:4:5",
	"15:04",
	"15:4",
	"150405",
}

// check if the string is a timestamp
// return true if the string is a timestamp and the time.Time object of the timestamp
// return false if the string is not a timestamp and the zero time.Time object
// input: string timestamp example:
// 	1715224735
// 	1715224735000
// 	1715224735000000
// 	1715224735000000000
// 	-1715224735
//  0.78995300 1715234643
//  1715234643.78995300
func IsTimestamp(str string) (bool, time.Time) {
	if str == "" {
		return false, time.Time{}
	}
	// range the str to check if all the characters are digits
	flag, space, dot := false, int64(0), int64(0)
	nowTime := time.Now()
	if str[0] == '-' {
		flag = true
		str = str[1:]
	}
	for i, c := range str {
		if c == ' ' && space == 0 {
			space = int64(i)
			continue
		} else if c == ' ' && space > 0 {
			return false, time.Time{}
		}
		if c == '.' && dot == 0 {
			dot = int64(i)
			continue
		} else if c == '.' && dot > 0 {
			return false, time.Time{}
		}
		if c < '0' || c > '9' {
			return false, time.Time{}
		}
	}
	// convert the str to int64
	sec, nsec := int64(0), int64(0)
	nowTimeLen := len(fmt.Sprintf("%d", nowTime.Unix()))

	secStr, nsecStr := str, ""
	if space > 0 {
		// if has space and dot then format: nsec sec
		secStr, nsecStr = str[space+1:], str[:space]

		if dot > 0 {
			dotPrefix := nsecStr[:dot]
			if dotPrefix != "0" {
				return false, time.Time{}
			}
			nsecStr = nsecStr[dot+1:]
		}
	} else if dot > 0 && space == 0 {
		// if has only dot then format: sec.nsec
		secStr, nsecStr = str[:dot], str[dot+1:]
	}

	if len(secStr) > nowTimeLen {
		secStr, nsecStr = str[:nowTimeLen], str[nowTimeLen:]
	}

	for len(nsecStr) < 9 {
		nsecStr += "0"
	}

	sec, _ = strconv.ParseInt(secStr, 10, 64)
	if len(nsecStr) > 0 {
		nsec, _ = strconv.ParseInt(nsecStr, 10, 64)
	}

	if flag {
		sec = -sec
	}

	// if timestamp is less than 1e9, it means the timestamp is in second
	// convert the int64 to time.Time
	return true, time.Unix(sec, int64(nsec))
}

// check input str and convert it to time.Time
func StrToTime(str string) (bool, time.Time) {
	// if ok, t := IsTimestamp(str); ok {
	// 	return ok, t
	// }
	for _, format := range TimeFormats {
		if t, err := time.Parse(format, str); err == nil {
			return true, t
		}
	}
	return false, time.Time{}
}
