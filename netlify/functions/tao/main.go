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
	// https://en.wiktionary.org/wiki/Taixuanjing
	// Tai Xuan Jing tetragrams
	//
	// There are 81 permutations of the following three lines, given 4 lines/grams
	// ------ Heaven, --  -- Earth, - - - Man
	// 3^4 = 81

	t := time.Now()
	// now := t.Format("2006-01-02 15:04:05")
	tetragram_index := 0 // This will be the base10 sum of the monograms

	tetragrams := []string{"𝌆", "𝌇", "𝌈", "𝌉", "𝌊", "𝌋", "𝌌", "𝌍", "𝌎", "𝌏", "𝌐", "𝌑", "𝌒", "𝌓", "𝌔", "𝌕", "𝌖", "𝌗", "𝌘", "𝌙", "𝌚", "𝌛", "𝌜", "𝌝", "𝌞", "𝌟", "𝌠", "𝌡", "𝌢", "𝌣", "𝌤", "𝌥", "𝌦", "𝌧", "𝌨", "𝌩", "𝌪", "𝌫", "𝌬", "𝌭", "𝌮", "𝌯", "𝌰", "𝌱", "𝌲", "𝌳", "𝌴", "𝌵", "𝌶", "𝌷", "𝌸", "𝌹", "𝌺", "𝌻", "𝌼", "𝌽", "𝌾", "𝌿", "𝍀", "𝍁", "𝍂", "𝍃", "𝍄", "𝍅", "𝍆", "𝍇", "𝍈", "𝍉", "𝍊", "𝍋", "𝍌", "𝍍", "𝍎", "𝍏", "𝍐", "𝍑", "𝍒", "𝍓", "𝍔", "𝍕", "𝍖"}

	// Simple random
	// choose random from [0,81)
	// randBigInt, err := rand.Int(rand.Reader, big.NewInt(81))
	// if err != nil {
	// 	panic(err)
	// }

	// More authentic random to build the tetragram (tetra is four)
	// Choose a random int 0-2 for each "gram" of the tetragram, then sum it to the previous with a weighting
	// To convert from trinary/ternary (base 3) to base 10, multiply each part by it's order of magnitude
	// i.e. The first or bottom line gets a weight of 27, the second, 9, then 3, then 1
	// TODO: Am I doing this backwards? Would it be more authentic to give the last rand trinary the heaviest weight?
	for i := 0; i <= 3; i++ {
		// choose random up from [0,3)
		randBigInt, err := rand.Int(rand.Reader, big.NewInt(3))
		if err != nil {
			panic(err)
		}
		randInt := int(randBigInt.Int64())

		switch i {
			case 0:
				tetragram_index = tetragram_index + randInt*27
			case 1:
				tetragram_index = tetragram_index + randInt*9
			case 2:
				tetragram_index = tetragram_index + randInt*3
			case 3:
				tetragram_index = tetragram_index + randInt*1
		}
	}

	// Have to convert bigint back to int64 to use in index
	tetragram := tetragrams[tetragram_index]

	// Add 1 here because the slice starts with position 0
	tetragram_num := tetragram_index + 1

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
