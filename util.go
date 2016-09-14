package gobrute

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
)

type Counter struct {
	v   int
	mux sync.Mutex
}

func (c *Counter) Dec() {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access map c.v.
	c.v--
	c.mux.Unlock()
}

type Credential struct {
	User string
	Pass string
}

/**
 * read credential from file.
 */
func ReadUserPass(path string) ([]*Credential, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	pairs := make([]*Credential, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var pair = strings.Split(strings.Trim(scanner.Text(), " "), " ")
		var c = &Credential{pair[0], pair[1]}
		pairs = append(pairs, c)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return pairs, nil
}
