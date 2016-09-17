# Bruteforce framework with Go.

------

Gobrute 是一个用golang实现的密码爆破框架（非直接可用的工具），用于实现对各种网络程序（ssh, redis, mysql, mongodb ...）的密码破解。

> * [Hello World Example](#hello-world-example)
> * [暴力破解SSH](#ssh-bruteforce)
> * [暴力破解Redis](#redis-bruteforce)
> * [暴力破解任意网络程序](#bruteforce-anything)


## Hello World Example

```
import (
    "log"
    "github.com/gushitong/gobrute"
)

bruter = DefaultSSHBruter()
config := &BruteConfig{
                Protocol:    "tcp",
                Port:        22,
                Workers:     10, 
                RequireUser: true,
                RequirePass: true,
                Dictpath:    "dict/userpass.txt",
                Targets:     []string{"127.0.0.1", "192.168.1.15"},
        }
client, err := NewClient(bruter, config)

if err != nil {
        log.Printf("Err: %s", err)
}

tick := time.NewTicker(200 * time.Millisecond)
respch := c.Run()
responses := make([]*Response, 0)
completed := 0
for completed < c.Config.Jobs {
        select {
        case resp := <-respch:
                completed++
                if resp != nil {
                        responses = append(responses, resp)
                        }
        case <-tick.C:

            }
}
tick.Stop()

// responses holds the bruteforce results.
log.Printf("resp: %s", responses)

```

## SSH Bruteforce


## Redis Bruteforce



## Bruteforce Anything.


