//go:build !arm
// +build !arm

package rpistringsarray

import "strings"

// ArrayToString allows calling uname -m on the operating system
// and uses int8 as an argument, which is needed to convert a
// int8 array to string. see https://pkg.go.dev/syscall#Utsname
func ArrayToString(x [65]int8) string {
	var buf [65]byte
	for i, b := range x {
		buf[i] = byte(b)
	}
	str := string(buf[:])
	if i := strings.Index(str, "\x00"); i != -1 {
		str = str[:i]
	}
	return str
}
