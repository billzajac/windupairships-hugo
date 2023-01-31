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
	// now := t.Format("2006-01-02 15:04:05")

	tetragrams := []string{"𝌆", "𝌇", "𝌈", "𝌉", "𝌊", "𝌋", "𝌌", "𝌍", "𝌎", "𝌏", "𝌐", "𝌑", "𝌒", "𝌓", "𝌔", "𝌕", "𝌖", "𝌗", "𝌘", "𝌙", "𝌚", "𝌛", "𝌜", "𝌝", "𝌞", "𝌟", "𝌠", "𝌡", "𝌢", "𝌣", "𝌤", "𝌥", "𝌦", "𝌧", "𝌨", "𝌩", "𝌪", "𝌫", "𝌬", "𝌭", "𝌮", "𝌯", "𝌰", "𝌱", "𝌲", "𝌳", "𝌴", "𝌵", "𝌶", "𝌷", "𝌸", "𝌹", "𝌺", "𝌻", "𝌼", "𝌽", "𝌾", "𝌿", "𝍀", "𝍁", "𝍂", "𝍃", "𝍄", "𝍅", "𝍆", "𝍇", "𝍈", "𝍉", "𝍊", "𝍋", "𝍌", "𝍍", "𝍎", "𝍏", "𝍐", "𝍑", "𝍒", "𝍓", "𝍔", "𝍕", "𝍖"}

	// choose random from [0,81)
	randInt, err := rand.Int(rand.Reader, big.NewInt(81))
	if err != nil {
		panic(err)
	}

	// Have to convert bigint back to int64 to use in index
	// Remember to add 1 here because the slice starts with position 0
	tetragram := tetragrams[randInt.Int64()]
	tetragram_num := randInt.Int64() + 1
        href := fmt.Sprintf("<a href='https://terebess.hu/english/tao/Wing.html#Kap%02d'>", tetragram_num)
	link_tet := fmt.Sprintf("%s%s</a>", href, tetragram)
	link_num := fmt.Sprintf("%s%0d</a>", href, tetragram_num)
	body := fmt.Sprintf("<html><body><h1>%s</h1>\n<h2>%s</h2>\n%s\n</body></html>", link_tet, link_num, t)
	//fmt.Println("This message will show up in the CLI console.")

	//Headers:         map[string]string{"Content-Type": "text/plain; charset=utf-8"},
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Body:            body,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
