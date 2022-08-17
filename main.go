package fume

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambdaV2

type Options struct {
	dev bool // development mode
	port int // port to run on (default 8080)
}

func Start(routes *gin.Engine, options *Options) {

	defaults := &Options{
		dev: false,
		port: 8080,
	}

	if options.dev {
		defaults.dev = options.dev
	}

	if options.port != 0 {
		defaults.port = options.port
	}

	if options.dev {
		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", defaults.port),
			Handler: routes,
		}
		server.ListenAndServe()
	} else {
		ginLambda = ginadapter.NewV2(routes)
		lambda.Start(Handler)
	}
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}