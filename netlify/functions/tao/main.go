package main

import (
	"context"
	"fmt"

	"time"
	"crypto/rand"
	"math/big"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	t := time.Now()
	tetragrams := []string{"𝌆", "𝌇", "𝌈", "𝌉", "𝌊", "𝌋", "𝌌", "𝌍", "𝌎", "𝌏", "𝌐", "𝌑", "𝌒", "𝌓", "𝌔", "𝌕", "𝌖", "𝌗", "𝌘", "𝌙", "𝌚", "𝌛", "𝌜", "𝌝", "𝌞", "𝌟", "𝌠", "𝌡", "𝌢", "𝌣", "𝌤", "𝌥", "𝌦", "𝌧", "𝌨", "𝌩", "𝌪", "𝌫", "𝌬", "𝌭", "𝌮", "𝌯", "𝌰", "𝌱", "𝌲", "𝌳", "𝌴", "𝌵", "𝌶", "𝌷", "𝌸", "𝌹", "𝌺", "𝌻", "𝌼", "𝌽", "𝌾", "𝌿", "𝍀", "𝍁", "𝍂", "𝍃", "𝍄", "𝍅", "𝍆", "𝍇", "𝍈", "𝍉", "𝍊", "𝍋", "𝍌", "𝍍", "𝍎", "𝍏", "𝍐", "𝍑", "𝍒", "𝍓", "𝍔", "𝍕", "𝍖"}

	// choose random from [0,81)
	randInt, err := rand.Int(rand.Reader, big.NewInt(81))
	if err != nil {
		panic(err)
	}

	// have to convert bigint back to int64 to use in index
	tetragram := tetragrams[randInt.Int64()]
	// Remember to add 1 here because the slice starts with position 0
	tetragram_int := randInt.Int64() + 1
	// now := t.Format("2006-01-02 15:04:05")

	body := fmt.Sprintf("%d %s %s\n", tetragram_int, tetragram, t)
	fmt.Println("This message will show up in the CLI console.")

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain; charset=utf-8"},
		Body:            body,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
