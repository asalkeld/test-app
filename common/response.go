package common

import (
	"fmt"

	"github.com/nitrictech/go-sdk/faas"
)

// Updates context with error information
func HttpResponse(ctx *faas.HttpContext, message string, status int) (*faas.HttpContext, error) {
	fmt.Println(message)
	ctx.Response.Body = []byte(message)
	ctx.Response.Status = status
	return ctx, nil
}
