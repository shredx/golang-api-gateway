package pubsub

import (
	"flag"
	"os"
)

/*
 * This file contains the definition and implementation of configuration management in the application
 */

//APIUrl is url of the api to be accessed
var APIUrl = flag.String("api-url", "http://127.0.0.1:8085", "URL of the api to be access. Eg. http://127.0.0.1:8085")

//APITokenHeader is the header key with which the token is stored in the request
var APITokenHeader = flag.String("api-token-header", "token", "Header key with which the token is stored in the request")

//RedisURL is the url of the redis instance
var RedisURL = flag.String("redis-url", "redis:6379", "URL with which redis can be accessed. Eg. redis:6379")

//RedisPassword is the password with which the redis instance can be accessed
var RedisPassword = flag.String("redis-password", "", "Password to access the redis instance")

//RedisBlockChannelName is the name of the redis pubsub channel through which the messages to block the requests come
var RedisBlockChannelName = flag.String("redis-block-channel", "BLOCK", "Name of the Redis pubsub channel through which the message for blocking requests come")

//RedisAPIUsageChannelName is the name of the redis pubsub channel through which
var RedisAPIUsageChannelName = flag.String("redis-api-usage-channel", "API_USAGE", "Name of the Redis pubsub channel through which API usage info has to be passed")

//RedisResetChannelName is the name of the Redis pubsub channel through which the message for reseting the api usage message come
var RedisResetChannelName = flag.String("redis-reset-channel", "RESET", "Name of the Redis pubsub channel through which the message for reseting the api usage message come")

func init() {
	flag.Parse()

	//We will also accept the configuration through environment variables

	//API_URL
	val := os.Getenv("API_URL")
	if len(val) != 0 {
		*APIUrl = val
	}
	//API_TOKEN_HEADER
	val = os.Getenv("API-TOKEN-HEADER")
	if len(val) != 0 {
		*APITokenHeader = val
	}
	//REDIS_URL
	val = os.Getenv("REDIS_URL")
	if len(val) != 0 {
		*RedisURL = val
	}
	//REDIS_PASSWORD
	val = os.Getenv("REDIS_PASSWORD")
	if len(val) != 0 {
		*RedisPassword = val
	}
	//REDIS_BLOCK_CHANNEL
	val = os.Getenv("REDIS_BLOCK_CHANNEL")
	if len(val) != 0 {
		*RedisBlockChannelName = val
	}
	//REDIS_USAGE_CHANNEL
	val = os.Getenv("REDIS_USAGE_CHANNEL")
	if len(val) != 0 {
		*RedisAPIUsageChannelName = val
	}
	//REDIS_RESET_CHANNEL
	val = os.Getenv("REDIS_RESET_CHANNEL")
	if len(val) != 0 {
		*RedisResetChannelName = val
	}
}
