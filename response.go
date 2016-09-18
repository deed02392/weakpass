package gobrute

import (
	"time"
)

type Response struct {
	Request *Request
	Time    time.Time
}
