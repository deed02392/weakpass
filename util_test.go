package gobrute

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestReadUserPass(t *testing.T) {
	pairs := []byte("admin admin\n 123456 123456\n")
	tmpfile, err := ioutil.TempFile("", "tmpfile")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(pairs); err != nil {
		log.Fatal(err)
	}

	c, err := ReadUserPass(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}

	expect := []*Credential{
		{"admin", "admin"}, {"123456", "123456"},
	}

	for i, _ := range expect {
		if c[i].User != expect[i].User || c[i].Pass != expect[i].Pass {
			t.Errorf("Expect: %q, Got: %q", expect[i], c[i])
		}
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
