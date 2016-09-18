# Bruteforce framework with Go.

Gobrute 是一个用golang实现的密码爆破框架，用于实现对各种网络程序（ssh, redis, mysql, mongodb ...）的密码破解。

## Hello World Example

```
import (
    "github.com/gushitong/gobrute"
)

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
                // Process err.
        }

        c.Start()

        for {
                if !c.IsFinished() {
                        time.Sleep(200)
                } else {
                        break
                }
        }

        c.GetResult()


```








