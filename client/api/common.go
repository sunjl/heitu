package api

import (
	"github.com/parnurzeal/gorequest"
)

var (
	superAgent *gorequest.SuperAgent
	address    string
	token      string
)

func InitConfig(addressIn, tokenIn string) {
	superAgent = gorequest.New()
	address = addressIn
	token = tokenIn
}
