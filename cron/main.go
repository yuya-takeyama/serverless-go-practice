package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.CloudWatchEvent) {
	buf := new(bytes.Buffer)
	j := json.NewEncoder(buf)
	j.Encode(event)
	fmt.Printf("Event payload: %s\n", buf.String())
}

func main() {
	lambda.Start(handler)
}
