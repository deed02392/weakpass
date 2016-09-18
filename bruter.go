package gobrute

import (
	//	"database/sql"
	"github.com/garyburd/redigo/redis"
	//	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"log"
	//	"strconv"
	"time"
)

type Bruter interface {

	//Brute method accept a *Request (as a Bruteforece request to remote host).
	//
	//If success, return a *Response else return a error instead.
	Brute(req *Request) (*Response, error)
}

//SSH Bruteforce
type SSHBruter struct{}

// Default SSH Bruter
func DefaultSSHBruter() SSHBruter {
	return SSHBruter{}
}

// SSHBruter.

func (r SSHBruter) Brute(req *Request) (*Response, error) {

	sshConfig := &ssh.ClientConfig{
		User: req.Options["User"],
		Auth: []ssh.AuthMethod{ssh.Password(req.Options["Pass"])},
	}

	_, err := ssh.Dial(req.Network, req.Address, sshConfig)

	log.Printf("Sending req: %v", req)
	if err != nil {
		return nil, err
	}
	log.Printf("[------]Successful req: %v", req)

	// create a successful Response.
	var resp = &Response{
		Request: req, Time: time.Now(),
	}
	return resp, nil
}

// Redis Bruteforce
type RedisBruter struct{}

func DefaultRedisBruter() RedisBruter {
	return RedisBruter{}
}

func (r RedisBruter) Brute(req *Request) (*Response, error) {
	option := redis.DialPassword(req.Options["Pass"])
	_, err := redis.Dial(req.Network, req.Address, option)

	log.Printf("Sending req: %v", req)
	if err != nil {
		// log.Printf("err %s", err)
		return nil, err
	}
	log.Printf("[-------]Successful req: %v", req)
	var resp = &Response{
		Request: req, Time: time.Now(),
	}
	return resp, nil
}
