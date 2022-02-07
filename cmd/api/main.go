package main

import (
	"github.com/JorgeAdd/go-final-skills-cto/cryptoAPI/internal/controller"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(controller.CryptoController)
}
