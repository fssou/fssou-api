package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"in.francl.api/internal/fssou"
	"log"
)

func init() {
	log.Println("Lambda cold start")
}

func main() {
	lambda.Start(fssou.Handler)
}
