# ⌨️ mgserver

Mgserver is a Go library 📚 for create http server, it's inspired by GIN. You can use it as a GIN server.

## Installation

Use the command to install mgserver.

```bash
go get github.com/milligoshant/mgserver
```
or
```bash
go get github.com/milligoshant/mgserver@v0.0.3
```

## Usage

```go
package main

import (
	"log"

	"github.com/milligoshant/mgserver"
)

func main() {
	var port = ":8080"

	// Initialize Gin router
	r := mgserver.Default()

	// Sample route
	r.GET("/", func(c *mgserver.Context) {
		c.JSON(200, mgserver.H{
			"message": "Server is running",
		})
	})

	r.Static("/static", "./static")

	err := r.
		WithETCD("MgserverTest2", "http://localhost:2379").
		Run(port)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

```

## Contributing

You can write us for contributing👏🏻.
