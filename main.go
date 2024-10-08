package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/context/ctxhttp"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := "https://security.access.redhat.com/data/csaf/v2/advisories/2024/rhba-2024_1246.json"

	resp, err := ctxhttp.Get(ctx, http.DefaultClient, url)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatalf("Request timed out: %v", err)
		}
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Println("Response:")
	fmt.Println(string(body))
}
