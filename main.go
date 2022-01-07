package main

import (
	"fmt"
	"twitter/creds"
	"twitter/http"
	"twitter/http/client"
	"twitter/http/handler"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	creds.Init("./creds/auth.json")

	kwrd := "Test Driven Development"
	netClient := client.Default()
	hdl := handler.TweetHandler{handler.TweetURLReq{Keyword: kwrd}}

	res, _ := http.ReqTweets(hdl, netClient)
	x, _, _ := terminal.GetSize(0)
	fmt.Println("============================================================")
	for _, v := range res.Tweets {
		fmt.Printf("CONTENT IS : \033[34m	%s\n", v.Content)
		for i := 0; i < x; i++ {
			fmt.Print("=")
		}
	}
	fmt.Println("============================================================")
	fmt.Println("\033[0m")
}
