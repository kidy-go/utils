package utils_test

import (
	"fmt"
	"testing"
	"time"
)

// func TestAA(t *testing.T) {
// 	d := utils.NewTime()
//
// 	formats := []string{
// 		"Y", "y", "m", "M", "n", "F", "d", "j", "D", "l", "N", "w", "W", "H", "h",
// 		"i", "s", "u", "a", "A", "c", "r", "U", "T", "z", "t", "L", "e", "O", "P", "Z", "B",
// 	}
//
// 	fmt.Println("Format:")
//
// 	for _, format := range formats {
// 		fmt.Printf("%s: %s\n", format, d.Format(format))
// 	}
// }
//

func TestTime(t *testing.T) {
	dft := time.Unix(1234567890, 0)
	fmt.Println("Default: ", dft.Format("2006-01-02 15:04:05"))
	var exampes = []struct {
		ts  int64
		out string
	}{
		{1234567890, "2009-02-13T23:31:30Z"},
		{12345678909, "2009-02-13T23:31:30.9Z"},
		{123456789098, "2009-02-13T23:31:30.98Z"},
		{1234567890987, "2009-02-13T23:31:30.987Z"},
		{12345678909876, "2009-02-13T23:31:30.9876Z"},
		{123456789098765, "2009-02-13T23:31:30.98765Z"},
		{1234567890987654, "2009-02-13T23:31:30.987654Z"},
		{12345678909876543, "2009-02-13T23:31:30.9876543Z"},
		{123456789098765432, "2009-02-13T23:31:30.98765432Z"},
		{1234567890987654321, "2009-02-13T23:31:30.987654321Z"},
	}
	fmt := "2006-01-02T15:04:05.999999999Z"
	//fmt.Println("Nano: ", time.Nanosecond)
	//fmt.Println("Micro: ", time.Microsecond)
	//fmt.Println("Milli: ", time.Millisecond)
	//fmt.Println("Second: ", time.Second)
	for _, example := range exampes {
		ts := example.ts
		t.Log("TS: ", ts)

		if ts < 1e10 {
			dt := time.Unix(ts, 0).UTC().Format(fmt)
			if dt != example.out {
				t.Errorf("Sec %s, got %s", example.out, dt)
			}
		} else if ts < 1e11 {
			sec, nsec := ts/1e1, ts%1e1*1e8
			dt := time.Unix(sec, nsec).UTC().Format(fmt)
			if dt != example.out {
				t.Errorf("Nano %s, got %s (%v, %v)", example.out, dt, sec, nsec)
			}

		}

		// if ts < 1e10 {
		// 	dt := time.Unix(ts, 0).UTC().Format(fmt)
		// 	if dt != example.out {
		// 		t.Errorf("Sec %s, got %s", example.out, dt)
		// 	}
		// } else if ts < 1e11 {
		// 	sec, nsec := ts/1e1, ts%1e1*1e8
		// 	dt := time.Unix(sec, nsec).UTC().Format(fmt)
		// 	if dt != example.out {
		// 		t.Errorf("Deci %s, got %s (%v,%v)", example.out, dt, sec, nsec)
		// 	}

		// } else if ts < 1e12 {
		// 	sec, nsec := ts/1e2, ts%1e2*1e7
		// 	dt := time.Unix(sec, nsec).UTC().Format(fmt)
		// 	if dt != example.out {
		// 		t.Errorf("Centi %s, got %s (%v,%v)", example.out, dt, sec, nsec)
		// 	}
		// } else if ts < 1e13 {
		// 	//sec, nsec := ts/1e2, ts%1e2*1e6
		// 	sec, nsec := ts/1e3, ts%1e3*1e6
		// 	dt := time.Unix(sec, nsec).UTC().Format(fmt)
		// 	if dt != example.out {
		// 		t.Errorf("Milli %s, got %s (%v,%v)", example.out, dt, sec, nsec)
		// 	}
		// } else if ts < 1e14 {
		// 	sec, nsec := ts/1e4, ts%1e4*1e5
		// 	dt := time.Unix(sec, nsec).UTC().Format(fmt)
		// 	if dt != example.out {
		// 		t.Errorf("Micro %s, got %s (%v,%v)", example.out, dt, sec, nsec)
		// 	}
		// } else if ts < 1e15 {
		// 	sec, nsec := ts/1e5, ts%1e5*1e4
		// 	dt := time.Unix(sec, nsec).UTC().Format(fmt)
		// 	if dt != example.out {
		// 		t.Errorf("Micro %s, got %s (%v,%v)", example.out, dt, sec, nsec)
		// 	}
		// } else if ts < 1e16 {
		// 	sec, nsec := ts/1e6, ts%1e6*1e3
		// 	dt := time.Unix(sec, nsec).UTC().Format(fmt)
		// 	if dt != example.out {
		// 		t.Errorf("Micro %s, got %s (%v,%v)", example.out, dt, sec, nsec)
		// 	}
		// } else if ts < 1e17 {
		// 	sec, nsec := ts/1e7, ts%1e7*1e2
		// 	dt := time.Unix(sec, nsec).UTC().Format(fmt)
		// 	if dt != example.out {
		// 		t.Errorf("Micro %s, got %s (%v,%v)", example.out, dt, sec, nsec)
		// 	}
		// } else if ts <= 1e18 {
		// 	sec, nsec := ts/1e8, ts%1e8*1e1
		// 	dt := time.Unix(sec, nsec).UTC().Format(fmt)
		// 	if dt != example.out {
		// 		t.Errorf("Nano %s, got %s (%v, %v)", example.out, dt, sec, nsec)
		// 	}
		// } else {
		// 	dt := time.Unix(0, ts).UTC().Format(fmt)
		// 	if dt != example.out {
		// 		t.Errorf("Expected %s, got %s => %v", example.out, dt, ts)
		// 	}
		// }
		//fmt.Println("Unix: ", d.Format("2006-01-02 15:04:05"))
	}
}
