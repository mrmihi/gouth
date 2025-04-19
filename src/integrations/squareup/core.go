package squareup

import (
	"github.com/go-resty/resty/v2"
	"goose/src/config"
)

var client *resty.Client

func getClient() *resty.Client {
	if client != nil {
		return client
	}
	client = resty.New().SetDebug(true).SetBaseURL(config.Env.SquareUpUrl).SetAuthToken(config.Env.SquareUpToken)
	return client
}
