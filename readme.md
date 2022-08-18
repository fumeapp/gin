# Fume adapter for Gin

[![Go Reference][1]][2]
[![GoCard][3]][4]

[1]: https://pkg.go.dev/badge/github.com/fumeapp/gin.svg
[2]: https://pkg.go.dev/github.com/fumeapp/gin
[3]: https://goreportcard.com/badge/github.com/fumeapp/gin
[4]: https://goreportcard.com/report/github.com/fumeapp/gin

This is a simple adapter for [Gin](https://github.com/gin-gonic/gin) that allows you to deploy your application using Fume.

```go
package main

import (
	fume "github.com/fumeapp/gin"
	"github.com/gin-gonic/gin"
)


func main() {
	routes := gin.New()
	routes.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hello World"}) })
	fume.Start(routes, fume.Options{})
}
```

## Options

| Option | Description | Default |
| ------ | ----------- | ------- |
| `Port` | Port to listen on | `8080` |
| `Host` | Host to listen on | `localhost` |
