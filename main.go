package main

import (
	//Importaciones de go (vienen incluidas al instalar)
	"context"
	"fmt"
	"os"
	"strings"

	//importaciones externas (descargadas)
	"github.com/aws/aws-lambda-go/events"
	lambda02 "github.com/aws/aws-lambda-go/lambda"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipal/awsgo"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipal/bd"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipal/handlers"
	/*
		//Importaciones de go (vienen incluidas al instalar)
		"fmt"
		"os"
		"context"
		"strings"

		//importaciones externas (descargadas)
		"github.com/aws/aws-lambda-go/events"
		lambda02 "github.com/aws/aws-lambda-go/lambda"

		//importaciones personalizadas (creadas desde cero)
		"github.com/PedroAntonioKira/ecommerceEscomPrincipal/awsgo"
		"github.com/PedroAntonioKira/ecommerceEscomPrincipal/bd"
		"github.com/PedroAntonioKira/ecommerceEscomPrincipal/models"
		"github.com/PedroAntonioKira/ecommerceEscomPrincipal/handlers"
	*/)

func main() {
	lambda02.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		panic("Error en los aprametros, debe enviar 'SecretName', 'UrlPrefix'")
	}

	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	bd.ReadSecret()

	status, message := handlers.Manejadores(path, method, body, header, request)

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResp,
	}

	return res, nil

}

func ValidoParametros() bool {
	var traeParametro bool

	_, traeParametro = os.LookupEnv("SecretName")
	if !traeParametro {
		fmt.Println("Algo fallo en la parte de SecretName")
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		fmt.Println("Algo fallo en la parte de UrlPrefix")
		return traeParametro
	}

	return traeParametro
}
