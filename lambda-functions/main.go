package main

import (
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

//Request - request lambda handler
type Request struct {
}

//Response - response lambda handler
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

//Handler - Calls ahgora-timekeeper Heroku server
func Handler(request Request) (Response, error) {
	resp, err := http.Get("https://ahgora-timekeeper.herokuapp.com/status")
	if err != nil {
		return Response{
			Message: "Error on GET ahgora-timekeeper",
			Ok:      false,
		}, err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{
			Message: "Error on read body message",
			Ok:      false,
		}, err
	}

	return Response{
		Message: string(contents),
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
