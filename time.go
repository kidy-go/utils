// time.go kee > 2021/03/18

package utils

import (
	"time"
	"reflect"
	"fmt"
)

var (
	TimeBase = time.Date(1582, time.October, 15, 0, 0, 0, 0, time.UTC).Unix()
)

type DateTime struct {
	Year    int
	Month   int
	Weekday int
	Day     int
	Hour    int
	Minute  int
	Second  int
	Usec    int
	Time    time.Time
}

// Parse time|datetime to DateTime
func Parse[T string|time.Time](in T) DateTime {
	var (
		t   time.Time
		err error
	)
	// Parse time
	refv := reflect.ValueOf(in)
	if refv.Kind() == reflect.String {
		t, err = time.Parse("2006-01-02 15:04:05", refv.Interface().(string))
		if err != nil {
			panic(err)
		}
	} else if refv.Kind() == reflect.Struct {
		t = refv.Interface().(time.Time)
	} else {
		panic("Invalid type")
	}
	return DateTime{
		Year:    t.Year(),
		Month:   int(t.Month()),
		Weekday: int(t.Weekday()),
		Day:     t.Day(),
		Hour:    t.Hour(),
		Minute:  t.Minute(),
		Second:  t.Second(),
		Usec:    t.Nanosecond() / 1000,
		Time:    t,
	}
}


func NewTime() DateTime {
	t := time.Now()
	return DateTime{
		Year:    t.Year(),
		Month:   int(t.Month()),
		Weekday: int(t.Weekday()),
		Day:     t.Day(),
		Hour:    t.Hour(),
		Minute:  t.Minute(),
		Second:  t.Second(),
		Usec:    t.Nanosecond() / 1000,
	}
}

func Now() DateTime {
	return NewTime()
}

func (d DateTime) Unix() int64 {
	return d.Time.Unix()
}

func (d DateTime) UnixNano() int64 {
	return d.Time.UnixNano()
}

func (d DateTime) Add(dur int, typ string) DateTime {
	switch typ {
	case "day":
		return d.AddDays(dur)
	case "month":
		return d.AddMonths(dur)
	case "year":
		return d.AddYears(dur)
	case "hour":
		return d.AddHours(dur)
	case "minute":
		return d.AddMinutes(dur)
	case "second":
		return d.AddSeconds(dur)
	default:
		panic("Invalid type")
	}
}

func (d DateTime) Sub(dur int, typ string) DateTime {
	switch typ {
	case "day":
		return d.SubDays(dur)
	case "month":
		return d.SubMonths(dur)
	case "year":
		return d.SubYears(dur)
	case "hour":
		return d.SubHours(dur)
	case "minute":
		return d.SubMinutes(dur)
	case "second":
		return d.SubSeconds(dur)
	default:
		panic("Invalid type")
	}
}

func (d DateTime) AddDays(days int) DateTime {
	return Parse(d.Time.AddDate(0, 0, days))
}

func (d DateTime) AddMonths(months int) DateTime {
	return Parse(d.Time.AddDate(0, months, 0))
}

func (d DateTime) AddYears(years int) DateTime {
	return Parse(d.Time.AddDate(years, 0, 0))
}

func (d DateTime) AddHours(hours int) DateTime {
	return Parse(d.Time.Add(time.Hour * time.Duration(hours)))
}

func (d DateTime) AddMinutes(minutes int) DateTime {
	return Parse(d.Time.Add(time.Minute * time.Duration(minutes)))
}

func (d DateTime) AddSeconds(seconds int) DateTime {
	return Parse(d.Time.Add(time.Second * time.Duration(seconds)))
}

func (d DateTime) SubDays(days int) DateTime {
	return d.AddDays(-days)
}

func (d DateTime) SubMonths(months int) DateTime {
	return d.AddMonths(-months)
}

func (d DateTime) SubYears(years int) DateTime {
	return d.AddYears(-years)
}

func (d DateTime) SubHours(hours int) DateTime {
	return d.AddHours(-hours)
}

func (d DateTime) SubMinutes(minutes int) DateTime {
	return d.AddMinutes(-minutes)
}

func (d DateTime) SubSeconds(seconds int) DateTime {
	return d.AddSeconds(-seconds)
}

func (d DateTime) Format(format string) string {
	format = ParseFormat(format)
	fmt.Println(format)
	return d.Time.Format(format)
}

// 转换 YYYYMMddHHmmss... 格式 为 time.Time 可用格式
func ParseFormat(dFormat string) string {
	foramts := map[string]string{
		"Y": "2006",	// 4位数年份
		"y": "06",		// 2位数年份
		"m": "01",		// 月份,有前导零
		"M": "Jan",		// 月份缩写
		"n": "1",		// 月份,没有前导零
		"F": "January",	// 月份全称
		"d": "02",		// 日期,有前导零
		"j": "2",		// 日期,没有前导零
		"D": "Mon",		// 星期缩写
		"l": "Monday",	// 星期全称
		"N": "1",		// 星期,ISO-8601数字表示
		"w": "0",		// 星期,数字表示
		"W": "01",		// 一年中的第几周,每周从星期天开始
		"H": "15",		// 24小时制,有前导零
		"h": "03",		// 12小时制,有前导零
		"i": "04",		// 分钟,有前导零
		"s": "05",		// 秒,有前导零
		"u": "000000",	// 微秒
		"a": "pm",		// 小写的上午和下午
		"A": "PM",		// 大写的上午和下午
		"c": "2006-01-02T15:04:05-07:00", // ISO 8601格式
		"r": "Mon, 02 Jan 2006 15:04:05 -0700", // RFC 2822格式
		"U": "1136239445",	// Unix时间戳
		"T": "MST",	// 时区
		"z": "-07:00", // 时区偏移
		"t": "2",	// 一个月中的天数,无前导零
		"L": "1",	// 是否为闰年
		"e": "America/New_York", // 时区标识符 如：America/New_York
		"O": "-0700",	// 与UTC时间的偏移量
		"P": "-07:00",	// 与UTC时间的偏移量
		"Z": "-25200", // 时区偏移量/秒数
		"B": "@400",	// 不包括闰秒的Swatch Internet时间
	}
	var foramt string
	for _, v := range dFormat {
		if _, ok := foramts[string(v)]; ok {
			foramt += foramts[string(v)]
		} else {
			foramt += string(v)
		}
	}

	return foramt
}






