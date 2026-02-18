package util

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

func IsAdministrator() bool {
	var token windows.Token
	// Open the current process token
	err := windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_QUERY, &token)
	if err != nil {
		return false
	}
	defer token.Close()

	// Check for elevation
	var isElevated uint32
	var length uint32
	err = windows.GetTokenInformation(token, windows.TokenElevation, (*byte)(unsafe.Pointer(&isElevated)), uint32(unsafe.Sizeof(isElevated)), &length)
	if err != nil {
		return false
	}

	return isElevated != 0
}
