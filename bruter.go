package gobrute

import (
	"golang.org/x/crypto/ssh"
	"log"
	"strconv"
)

type Bruter interface {

	//Brute method accept a *Request (as a Bruteforece request to remote host).
	//
	//If success, return a *Response else return a error instead.
	Brute(req *Request) (*Response, error)
}

type SSHBruter struct{}

// Default SSH Bruter
func DefaultSSHBruter() SSHBruter {
	return SSHBruter{}
}

// SSHBruter.

func (r SSHBruter) Brute(req *Request) (*Response, error) {

	sshConfig := &ssh.ClientConfig{
		User: req.User,
		Auth: []ssh.AuthMethod{ssh.Password(req.Pass)},
	}

	Addr := req.Addr + ":" + strconv.Itoa(req.Port)
	_, err := ssh.Dial(req.Protocol, Addr, sshConfig)
	log.Printf("Sending req: %s, User: %s, Pass: %s", req.Addr, req.User, req.Pass)
	if err != nil {
		return nil, err
	}
	log.Printf("[------]Successful req: %s, User: %s, Pass: %s", req.Addr, req.User, req.Pass)
	// create a successful Response.
	var resp = &Response{
		Req: req, User: req.User, Pass: req.Pass,
	}
	return resp, nil
}
