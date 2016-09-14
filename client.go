package gobrute

import (
	"fmt"
	"os"
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
	workers := c.Config.Workers
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
	}

	return responses
}

func (c *Client) Run() <-chan *Response {
	credentials, err := ReadUserPass(c.Config.Dictpath)
	if err != nil {
		fmt.Printf("Error reading dictfile.")
		os.Exit(1)
	}

	requests := make([]*Request, 0)
	if len(c.Config.Targets) <= 0 {
		fmt.Printf("No targets specificed. Exiting")
		os.Exit(1)
	}

	for _, t := range c.Config.Targets {
		for _, credential := range credentials {
			req := &Request{
				Addr:     t,
				Protocol: c.Config.Protocol,
				Port:     c.Config.Port,
				User:     credential.User,
				Pass:     credential.Pass,
			}
			requests = append(requests, req)
		}
	}
	c.Config.Jobs = len(requests)
	return c.DoBatch(requests...)
}
