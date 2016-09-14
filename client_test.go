package gobrute

import (
	"testing"
	"time"
)

func TestSSHClient(t *testing.T) {
	var bruter = DefaultSSHBruter()
	var config = DefaultSSHConfig()
	c, err := NewClient(bruter, config)

	if err != nil {
		t.Error(err)
	}

	// Test SSH Client Do.

	req := &Request{Addr: "127.0.0.1", Protocol: "tcp", Port: 22, User: "admin", Pass: "admin"}
	resp, err := c.Do(req)
	if resp != nil || err == nil {
		t.Fail()
	}

	// Test SSH Client DoAsync.
	respch := c.DoAsync(req)
	resp = <-respch
	if resp != nil {
		t.Fail()
	}

	// Test SSH Client DoBatch.

	reqs := []*Request{
		&Request{Addr: "127.0.0.1", Protocol: "tcp", Port: 22, User: "123456", Pass: "123456"},
		&Request{Addr: "127.0.0.1", Protocol: "tcp", Port: 22, User: "1234567", Pass: "1234567"},
		&Request{Addr: "127.0.0.1", Protocol: "tcp", Port: 22, User: "12345678", Pass: "12345678"},
	}
	tick := time.NewTicker(200 * time.Millisecond)
	respch = c.DoBatch(reqs...)
	responses := make([]*Response, 0)
	completed := 0
	for completed < len(reqs) {
		select {
		case resp := <-respch:
			completed++
			if resp != nil {
				responses = append(responses, resp)
			}
		case <-tick.C:

		}
	}
	tick.Stop()
	if len(responses) != 0 {
		t.Fail()
	}

}

func TestSSHClientRun(t *testing.T) {
	config := &BruteConfig{
		Protocol:    "tcp",
		Port:        22,
		Workers:     10,
		RequireUser: true,
		RequirePass: true,
		Dictpath:    "dict/userpass.txt",
		Targets:     []string{"127.0.0.1", "192.168.1.15"},
	}
	c, err := NewClient(DefaultSSHBruter(), config)

	if err != nil {
		t.Error(err)
	}

	tick := time.NewTicker(200 * time.Millisecond)
	respch := c.Run()
	responses := make([]*Response, 0)
	completed := 0
	for completed < c.Config.Jobs {
		select {
		case resp := <-respch:
			completed++
			if resp != nil {
				responses = append(responses, resp)
			}
		case <-tick.C:

		}
	}
	tick.Stop()
	if len(responses) != 1 {
		t.Fail()
	}

	var resp = responses[0]
	if resp.User != "ruby" || resp.Pass != "521332" {
		t.Fail()
	}
}
