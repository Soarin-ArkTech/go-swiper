package transport

import (
	"io"
)

type IConnectCreds interface {
	GrabServiceName() string
	GrabAddress() string
	GrabUser() string
	GrabPassword() string
}

type IConnectedJob interface {
	IConnectCreds
	RunCommands(input InputTaker)
}

type InputTaker interface {
	GetInputPipe() io.WriteCloser
}
