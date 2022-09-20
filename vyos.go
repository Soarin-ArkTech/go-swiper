package main

type VyOS struct {
	Hostname string `json:"Hostname"`
	Address  string `json:"Address"`
	Username string `json:"User"`
	Pass     string `json:"Passwd"`
}

func (in VyOS) GrabHostname() string {
	return in.Hostname
}

func (in VyOS) GrabAddress() string {
	return in.Address
}

func (in VyOS) GrabUsername() string {
	return in.Username
}

func (in VyOS) GrabPassword() string {
	return in.Pass
}

func (in VyOS) backupCMD() []string {

	// Grab & format date
	dateParsed := parseDate()

	// Our commands to run in SSH session
	vyCmds := []string{
		"source /opt/vyatta/etc/functions/script-template",
		"run show configuration commands >" + in.GrabHostname() + "-" + dateParsed + ".txt",
		" ",
	}

	return vyCmds
}
