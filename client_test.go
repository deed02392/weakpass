package gobrute

import (
	"testing"
	"time"
)

// Use Redis Brutefore as example.
func TestClientStart(t *testing.T) {
	config := &BruteConfig{
		NumWorker: 100,
		Dictpath:  "dict/userpass.txt",
		Addrs:     []string{"redis://127.0.0.1:6379"},
	}

	c, err := NewClient(DefaultRedisBruter(), config)

	if err != nil {
		t.Error(err)
	}

	c.Start()

	for {
		if !c.IsFinished() {
			t.Logf("Progress: %v", c.GetProgress())
			time.Sleep(200)
		} else {
			t.Log("Done.")
			break
		}
	}

	responses := c.GetResult()

	if len(responses) != 1 {
		t.Fail()
	}

	var resp = responses[0]

	if resp.Request.Options["Pass"] != "foobared" {
		t.Fail()
	}

}
