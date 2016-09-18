package gobrute

import (
	"testing"
	"time"
)

func TestRedisBruter(t *testing.T) {
	config := &BruteConfig{
		Protocol:    "tcp",
		Port:        6379,
		Workers:     100,
		RequireUser: false,
		RequirePass: true,
		Dictpath:    "dict/userpass.txt",
		Targets:     []string{"127.0.0.1"},
	}

	c, err := NewClient(DefaultRedisBruter(), config)

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

	if resp.User != "" || resp.Pass != "foobared" {
		t.Fail()
	}
}
