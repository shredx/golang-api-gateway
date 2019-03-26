package pubsub

import (
	"github.com/go-redis/redis"
	"github.com/shredx/golang-redis-rate-limiter/app/models"
)

/*
 * This file contains the pubsub implementation of the api gateway with redis
 */

//Client is the redis client for the redis
var Client *redis.Client

//Request is wrapper around rate request with aditional Out channel to respond back
type Request struct {
	models.RateRequest                         //RateRequest
	Out                chan models.RateRequest //Out is the channel through requests can be responded
}

//NewRequest creates a new request for a given rate request
func NewRequest(r models.RateRequest) Request {
	return Request{r, make(chan models.RateRequest)}
}

//BlackListChannel is the channel through which black listed requests can be send to the blacklist go routine
var BlackListChannel = make(chan Request)

func init() {
	InitRedis()
	go BlackList(BlackListChannel)
}

//InitRedis will init the redis client
func InitRedis() {
	/*
	 * We will init the connection with the redis message broker
	 */
	Client = redis.NewClient(&redis.Options{
		Addr:     *RedisURL,
		Password: *RedisPassword, // no password set
		DB:       0,              // use default DB
	})
}

//BlackList is the go routine which maintains the blacklisted tokens and updates it
func BlackList(ch chan Request) {
	/*
	 * We will go into infinite loop waiting for the rate request channel to pop requests
	 * When requests arrive, depending upon the type of the request, we will process them
	 * If the request is to block the token we will add the black listed map
	 * If the request is to reset the token usage, we will remove the token from blacklist if existing
	 * If the request is to get the status of the token, if the token is not blacklisted we will return ok
	 */
	//map to store the blacklisted tokens
	blackListed := map[string]bool{}

	//going into the infinite loop waiting for the request channel
	for {
		//waiting for the request
		req := <-ch

		//based on the type of the requet will process it
		switch req.Type {
		case models.BLOCK:
			//if the request is to block, we will add the keyhash to the black listed map
			blackListed[req.KeyHash] = true
		case models.RESET:
			//if the request is to reset, delete the keyhash from the blacklisted list
			delete(blackListed, req.KeyHash)
		case models.STATUS:
			//if the request is to get the status we will c-heck whether the keyhash exist in the black list map
			//if exist we will return block else ok as request type in the output channel
			t := models.OK
			if blackListed[req.KeyHash] {
				t = models.BLOCK
			}
			go SendToChannel(req.Out, models.RateRequest{KeyHash: req.KeyHash, Type: t})
		}
	}
}

//SendToChannel will send arequest to a channel. This function comes handy in situations where requests has
//to be send async.
func SendToChannel(ch chan models.RateRequest, req models.RateRequest) {
	ch <- req
}

func init() {
	InitPubSub()
}

//InitPubSub inits the pubsub to communicate with the rate limiter through redis
func InitPubSub() {
	/*
	 * We will init the subcriber and store the block channel
	 * Will init the subscriber and store the reset channel
	 * Will init the usage go routine
	 * Will init the reste go routine
	 * Will init the rate process go routine
	 */
	//init he subscriber for usage channel
	sub := Client.Subscribe(*RedisBlockChannelName)
	sub.Receive()
	models.UsageChannel = sub.Channel()

	//initing the reset channel and storing it
	subR := Client.Subscribe(*RedisResetChannelName)
	subR.Receive()
	models.ResetChannel = subR.Channel()

	//initing the api token block go routine
	go HandleMessages(models.UsageChannel, models.BLOCK)

	//initing the api usage reset go routine
	go HandleMessages(models.ResetChannel, models.RESET)
}

//HandleMessages go routine will handle the pubsub messages from redis
//t is the type of requests that needed to be catered
func HandleMessages(ch <-chan *redis.Message, t models.Type) {
	/*
	 * We will go into an infinite for loop
	 * Will wait for usage request
	 * When usage message comes in we will create a request and pass it to the blacklist channel
	 */
	for {
		m := <-ch
		BlackListChannel <- Request{RateRequest: models.RateRequest{KeyHash: m.Payload, Type: t}}
	}
}
