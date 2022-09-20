package main

type Linux struct {
	Hostname string `json:"hostname"`
	Address  string `json:"address"`
	Username string `json:"user"`
	Pass     string `json:"Passwd"`
}

func (in Linux) GrabHostname() string {
	return in.Hostname
}

func (in Linux) GrabAddress() string {
	return in.Address
}

func (in Linux) GrabUsername() string {
	return in.Username
}

func (in Linux) GrabPassword() string {
	return in.Pass
}

func (in Linux) backupCMD() []string {
	dateParsed := parseDate()

	LinuxCmds := []string{
		"lscpu >" + in.GrabHostname() + "-" + dateParsed + ".txt",
		" ",
	}

	return LinuxCmds
}
