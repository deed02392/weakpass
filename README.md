# Gobrute.

Gobrute is a bruteforce framework in golang，can bruteforce almost everything（ssh, redis, mysql, mongodb ...）with simple config。

---------------------------------------
  * [Features](#features)
  * [Requirements](#requirements)
  * [Installation](#installation)
  * [Usage](#usage)
  * [Testing / Development](#testing--development)
  * [License](#license)


---------------------------------------

## Features
  * Easy to config, Easy to use, Easy to extend.
  * Support ssh, redis, mysql ... on ground.
  * Easy to implement your own bruteforce plugin. [github-bruteforce](#github-bruteforce)

## Usage

```go
import (
    "github.com/gushitong/gobrute"
)

	config := &BruteConfig{
                Dictpath:    "dict/userpass.txt",
                Addrs:     []string{"redis://127.0.0.1:6379"},
        }

        c, err := NewClient(DefaultRedisBruter(), config)

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








