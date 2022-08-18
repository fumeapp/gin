package fume

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambdaV2

type Options struct {
	// optional - hostname to listen on (default: localhost)
	Host string 
	// optional - port to run on (default: 8080)
	Port int 
}

func Start(routes *gin.Engine, options Options) {

	defaults := &Options{
		Host: "localhost",
		Port: 8080,
	}

	if options.Port != 0 { defaults.Port = options.Port }
	if options.Host != "" { defaults.Host = options.Host }

	if os.Getenv("_HANDLER") != "" {
		ginLambda = ginadapter.NewV2(routes)
		lambda.Start(Handler)
	} else {
		server := &http.Server{
			Addr:    fmt.Sprintf("%s:%d", defaults.Host, defaults.Port),
			Handler: routes,
		}
		server.ListenAndServe()
	}
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}