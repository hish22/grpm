package util

import (
	"log"

	"golang.org/x/sys/windows"
)

func IsAdministrator() bool {
	var sid *windows.SID
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid,
	)
	if err != nil {
		log.Fatalf("SID Error: %s", err)
	}
	defer windows.FreeSid(sid)

	token, err := windows.OpenCurrentProcessToken()
	if err != nil {
		log.Fatalf("Token Error: %s", err)
	}

	member, err := token.IsMember(sid)
	if err != nil {
		log.Fatalf("Membership Error: %s", err)
	}

	return member
}
