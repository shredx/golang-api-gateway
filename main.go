package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shredx/golang-api-gateway/pubsub"
	"github.com/shredx/golang-redis-rate-limiter/app/models"
)

var client = &http.Client{}

func handler(w http.ResponseWriter, req *http.Request) {
	/*
	 * First we need to authozire the token
	 * We need to get the token and check whether the token usage limits has reached or not
	 * If the limits have been reached , we will return unauthorized header
	 * If not we will create a request and make a do request to api
	 * concurrently we will also send the api usage message
	 */
	//authorizing the token
	key := req.Header.Get(*pubsub.APITokenHeader)
	if len(key) == 0 {
		//no api token. Will return unotherized
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Provide api access token in header" + *pubsub.APITokenHeader))
		return
	}
	tok := pubsub.NewToken(key)
	pubsub.TokenChan <- tok
	tokA := <-tok.Out
	//if tokA is having a type BLOCK it means token not found
	if tokA.Type == models.BLOCK {
		//no api token. Will return unotherized
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Provide api token is invalid"))
		return
	}

	//checking whether the limit has been reached
	re := pubsub.NewRequest(models.RateRequest{KeyHash: key, Type: models.STATUS})
	pubsub.BlackListChannel <- re
	res := <-re.Out
	if res.Type == models.BLOCK {
		//limit has reached
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("Your api usage limit for the given token has been expired"))
		return
	}

	//Now we make and send the request
	// we need to buffer the body if we want to read it here and send it
	// in the request.
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// you can reassign the body if you need to parse it as multipart
	req.Body = ioutil.NopCloser(bytes.NewReader(body))

	// create a new url from the raw RequestURI sent by the client
	url := fmt.Sprintf("%s%s", *pubsub.APIUrl, req.RequestURI)
	proxyReq, err := http.NewRequest(req.Method, url, bytes.NewReader(body))

	// We may want to filter some headers, otherwise we could just use a shallow copy
	// proxyReq.Header = req.Header
	proxyReq.Header = make(http.Header)
	for h, val := range req.Header {
		proxyReq.Header[h] = val
	}

	//concurrently informing the redis about the usage
	go func() {
		pubsub.Client.Publish(*pubsub.RedisAPIUsageChannelName, key)
	}()

	//making the request to the proxy
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, err.Error())
		return
	}
	w.Header().Set("content-type", resp.Header.Get("content-type"))
	w.Write(b)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
