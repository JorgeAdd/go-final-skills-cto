package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/JorgeAdd/go-final-skills-cto/cryptoAPI/internal/service"
	"github.com/aws/aws-lambda-go/events"
)

func CryptoController(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// router := mux.NewRouter()
	// //Get crypto general info
	// router.PathPrefix("/crypto").HandlerFunc(service.GetCryptoInfo).Methods("GET")

	// //Get crypto general info
	// router.PathPrefix("/cryptoByDate").HandlerFunc(service.GetCryptoFilterDate).Methods("GET")

	// //Get crypto general info
	// router.PathPrefix("/cryptoByBook").HandlerFunc(service.GetCryptoFilterBook).Methods("GET")

	// log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))

	// return router

	if request.HTTPMethod == "GET" {
		body := request.Body
		fmt.Println("bod-1")
		fmt.Println(body)
		fmt.Println("bod0")
		// fmt.Println(json.Marshal(body).usd_book)
		fmt.Println("bod")
		fmt.Println(request.PathParameters)
		fmt.Println("bod2")
		fmt.Println(request.QueryStringParameters)
		// req := httptest.NewRequest(request.HTTPMethod, request.Path, body)
		var stringResponse []service.Payload

		switch request.Path {
		case "/crypto":
			stringResponse = service.GetCryptoInfo()
		case "/cryptobydate":
			stringResponse = service.GetCryptoFilterDate(body)
		case "/cryptobybook":
			stringResponse = service.GetCryptoFilterBook(body)
		default:
			fmt.Println("default")
		}

		// Response
		savedString, errorJSON := json.Marshal(stringResponse)
		ApiResponse := events.APIGatewayProxyResponse{Body: string(savedString), StatusCode: 200}
		return ApiResponse, errorJSON
	} else {
		err := errors.New("method Not Allowed")
		ApiResponse := events.APIGatewayProxyResponse{Body: "Method not OK", StatusCode: 502}
		return ApiResponse, err

	}
}
