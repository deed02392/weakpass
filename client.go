package gobrute

import (
	"fmt"
	"os"
	"time"
)

// A Client is a bruteforce client.
//
// Clients are safe for concurrent use by mutiple goroutines.
type Client struct {

	// A bruter describes who implements a brute method.
	//
	// Detailly, the bruter accept a *Request and return a *Response if success else return error.
	Bruter Bruter

	// BruteConfig describes How the client act during the bruteforce. (eg: Num of Workers)
	Config *BruteConfig

	NumJob int

	CompletedJob int

	Finished bool

	Responses []*Response

	Shutdown bool
}

// NewClient returns a client with custom configation.
//
//
func NewClient(bruter Bruter, config *BruteConfig) (*Client, error) {
	return &Client{
		Bruter: bruter, Config: config,
	}, nil
}

// Client.Do is a warpper for Bruter.Do
func (c *Client) Do(req *Request) (*Response, error) {
	bruter := Bruter(c.Bruter)
	return bruter.Brute(req)
}

// Client.DoAsync
func (c *Client) DoAsync(req *Request) <-chan *Response {
	r := make(chan *Response, 1)
	go func() {
		resp, err := c.Do(req)
		if err == nil {
			r <- resp
		}
		close(r)
	}()
	return r
}

// Client.DoBatch
func (c *Client) DoBatch(reqs ...*Request) <-chan *Response {
	workers := c.Config.NumWorker
	if workers == 0 {
		workers = len(reqs)
	}

	responses := make(chan *Response, workers)
	workerDone := make(chan bool, workers)
	producer := make(chan *Request, 0)

	go func() {
		for i := 0; i < len(reqs); i++ {
			producer <- reqs[i]
		}
		close(producer)

		for i := 0; i < workers; i++ {
			<-workerDone
		}
		close(responses)
	}()

	for i := 0; i < workers; i++ {
		go func(i int) {
			for req := range producer {
				resp := <-c.DoAsync(req)
				responses <- resp
			}
			workerDone <- true
		}(i)

		if c.Shutdown {
			// Stop start new goroutine after get the shutdown signal.
			break
		}
	}

	return responses
}

//

func (c *Client) _start() <-chan *Response {
	credentials, err := ReadUserPass(c.Config.Dictpath)
	if err != nil {
		fmt.Printf("Error reading dictfile.")
		os.Exit(1)
	}

	requests := make([]*Request, 0)
	if len(c.Config.Addrs) <= 0 {
		fmt.Printf("No targets specificed. Exiting")
		os.Exit(1)
	}

	for _, addr := range c.Config.Addrs {
		for _, credential := range credentials {
			options := map[string]string{"User": credential.User, "Pass": credential.Pass}
			req, err := NewRequest(addr, c.Config.Timeout, options)
			if err == nil {
				requests = append(requests, req)
			}
		}
	}
	c.NumJob = len(requests)
	return c.DoBatch(requests...)
}

// Call a client's start method means start a bruteforce task.
//
// You check task progress by: client.Progress
// And check whether task is finished by: client.Finiehed.

func (c *Client) Start() {

	t := time.NewTicker(200 * time.Millisecond)

	responses := make([]*Response, 0)

	respch := c._start()

	c.CompletedJob = 0

	go func() {
		for c.CompletedJob < c.NumJob {
			select {
			case resp := <-respch:
				c.CompletedJob++
				if resp != nil {
					responses = append(responses, resp)
				}
			case <-t.C:
				// pass
			}
		}
		t.Stop()
		c.Finished = true
		c.Responses = responses
	}()
}

/*
 Wrapper for client.Progress
*/
func (c *Client) GetProgress() float32 {
	return float32(c.CompletedJob) / float32(c.NumJob)
}

/*
Wrapper for client.Finished
*/
func (c *Client) IsFinished() bool {
	return c.Finished
}

/*
Wrapper for client.Result
*/
func (c *Client) GetResult() []*Response {
	return c.Responses
}
