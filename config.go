package gobrute

import (
	"time"
)

type Rule int

const (
	// Stop immediately after has the first bruteforce result.
	STOP_AFTER_ONE_SUCCESS Rule = iota

	// Stop after each target addr has a bruteforce result.
	STOP_AFTER_ALL_SUCCESS

	// Stop after all passwords has been tried.
	STOP_AFTER_ALL_PASSWORD_TRIED
)

type BruteConfig struct {

	// Addrs contains a bruteforce target list.
	//
	// It must be formed as full patterns: (redis://127.0.0.1:6379/1, myql://localhost:3306/db_name)
	Addrs []string

	// Num workers when bruteforce.
	NumWorker int

	// The dictionary path. The dictformat is: (username password each line.)
	Dictpath string

	// Timeout config for each bruteforce request. Default is nil.
	Timeout time.Duration

	// Interval between each bruteforce request. Default is nil.
	Interval time.Duration

	//Stop After First Result.
	Mode Rule
}
