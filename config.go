package gobrute

type BruteConfig struct {
	//
	Protocol string

	//
	Port int

	// Num workers when bruteforce.
	Workers int

	//
	RequireUser bool

	//
	RequirePass bool

	//
	Dictpath string

	//
	Targets []string

	//
	Jobs int
}

// NewBruteConfig method.
//
//
func NewBruteConfig(protocol string, port int, workers int, requireUser bool, requirePass bool, dictpath string) *BruteConfig {
	return &BruteConfig{
		Protocol: protocol, Port: port, Workers: workers, RequireUser: requireUser, RequirePass: requirePass, Dictpath: dictpath,
	}
}

func DefaultSSHConfig() *BruteConfig {
	return NewBruteConfig("tcp", 22, 100, true, true, "dict/userpass.txt")
}
