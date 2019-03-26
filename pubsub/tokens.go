package pubsub

import (
	"fmt"

	"github.com/shredx/golang-redis-rate-limiter/app/models"
)

/*
 * This file contains the implementations to authenticate tokens
 */

//Token is wrapper around rate request with aditional Out channel to respond back
type Token struct {
	models.RateRequest                         //RateRequest
	Out                chan models.RateRequest //Out is the channel through requests can be responded
}

//NewToken will generate a new token request
func NewToken(keyHash string) Token {
	return Token{RateRequest: models.RateRequest{KeyHash: keyHash, Type: models.BLOCK}, Out: make(chan models.RateRequest)}
}

//TokenChan is the channel for getting the token information
var TokenChan = make(chan Token)

func init() {
	go Tokens(TokenChan)
}

//Tokens go routine will cater to retieve the token information from redis cache
func Tokens(ch chan Token) {
	/*
	 * We will store the tokens in a map.
	 * We will then go into an infinite go routine
	 * If the token is not available in the map, will fetch it from the redis cache
	 */
	//cache for storing the tokens
	cache := map[string]bool{}

	//going into an infinite for loop to serve token requests
	for {
		req := <-ch

		//got a request try to get the key from cache
		val := cache[req.KeyHash]
		if !val {
			//since key is not in cache, try redis
			ke, err := Client.Get(req.KeyHash).Result()
			if err == nil {
				val = true
				cache[req.KeyHash] = true
			}
			fmt.Println(ke, req.KeyHash, err)
		}

		//returning the response
		if val {
			req.Type = models.OK
		}
		req.Out <- req.RateRequest
	}
}
