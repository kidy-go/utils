// str.go kee > 2021/03/18

package utils

import (
	"sync/atomic"
	"time"
)

type UUID [16]byte

func GenUUID() UUID {
	var (
		uuid         UUID
		clockSeq     uint32
		hardwareAddr []byte
		aTime        = time.Now()
	)
	utcTime := aTime.In(time.UTC)
	t := uint64(utcTime.Unix()-TimeBase)*10000000 + uint64(utcTime.Nanosecond()/100)
	uuid[0], uuid[1], uuid[2], uuid[3] = byte(t>>24), byte(t>>16), byte(t>>8), byte(t)
	uuid[4], uuid[5] = byte(t>>40), byte(t>>32)
	uuid[6], uuid[7] = byte(t>>56)&0x0F, byte(t>>48)

	clock := atomic.AddUint32(&clockSeq, 1)
	uuid[8] = byte(clock >> 8)
	uuid[9] = byte(clock)

	copy(uuid[10:], hardwareAddr)

	uuid[6] |= 0x10 // set version to 1 (time based uuid)
	uuid[8] &= 0x3F // clear variant
	uuid[8] |= 0x80 // set to IETF variant

	return uuid
}

func (u UUID) formatString() []byte {
	var offsets = [...]int{0, 2, 4, 6, 9, 11, 14, 16, 19, 21, 24, 26, 28, 30, 32, 34}
	const hexString = "0123456789abcdef"
	r := make([]byte, 36)
	for i, b := range u {
		r[offsets[i]] = hexString[b>>4]
		r[offsets[i]+1] = hexString[b&0xF]
	}
	r[8] = '-'
	r[13] = '-'
	r[18] = '-'
	r[23] = '-'
	return r
}

func (u UUID) String() string {
	return string(u.formatString())
}

func (u UUID) ShortString() string {
	return string(u.formatString()[:8])
}
