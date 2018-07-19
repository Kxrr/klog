# Klog

A simple logger wraps [go-logging](https://github.com/op/go-logging).


## Usage

```go
package main

import (
	log "github.com/kxrr/klog"
)

func main() {
	// log.Configure("foo", "Foo - ", os.Stdout)
	log.Printf("hello, world")
}
```
