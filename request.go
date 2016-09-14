package gobrute

// A Request requesents an bruteforce request to a remote host send by a Client.
type Request struct {

	// Addr is the target address. eg: (redis://127.0.0.1:6379/1
	Addr     string
	Port     int
	Protocol string
	User     string
	Pass     string
}

// NewRequest returns a new brutefore request suitable for use with Client.Do
//
//
func NewRequest(addr string, port int, protocol string, user string, pass string) *Request {
	return &Request{
		Addr: addr, Port: port, Protocol: protocol, User: user, Pass: pass,
	}
}
