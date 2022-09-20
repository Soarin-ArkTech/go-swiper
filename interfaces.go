package main

type Services interface {
	GrabHostname() string
	GrabAddress() string
	GrabUsername() string
	GrabPassword() string
	backupCMD() []string
}

type BackupList struct {
	VyOS  `json:"vyOS, omitempty"`
	Linux `json:"Linux, omitempty"`
}
