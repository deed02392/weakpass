package gobrute

import (
	"errors"
	"strings"
	"time"
)

// A Request requesents an bruteforce request to a remote host send by a Client.
type Request struct {

	// Addr is the target address. eg: (redis://127.0.0.1:6379/1
	Network string
	Address string

	// Timeout is the maximum amount of time a dial will wait for
	// a connect to complete. If Deadline is also set, it may fail
	// earlier.
	//
	// The default is no timeout.
	Timeout time.Duration

	Options map[string]string
}

func NewRequest(rawaddr string, timeout time.Duration, options map[string]string) (*Request, error) {
	if strings.HasPrefix(rawaddr, "redis://") || strings.HasPrefix(rawaddr, "mysql://") || strings.HasPrefix(rawaddr, "ssh://") {
		address := strings.Split(rawaddr, "://")[1]
		return &Request{
			Network: "tcp", Address: address, Timeout: timeout, Options: options,
		}, nil
	}
	return nil, errors.New("Unsupport addr format.")
}
