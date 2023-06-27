# ginerror

```bash
go get -u github.com/zhuweiyou-go/ginerror
```

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuweiyou-go/ginerror"
)

func main() {
	r := gin.Default()
	// default
	r.Use(ginerror.Handler())

	// or custom
	r.Use(ginerror.Handler(func(err *gin.Error) any {
		return gin.H{"error": err.Error()}
	}))
}
```
