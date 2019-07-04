package sms_test

import (
	"github.com/qiniu/api.v7/auth"

	"github.com/qiniu/api.v7/sms"
)

var manager *sms.Manager

func init() {
	accessKey := "DP4FyFXIHuThsAqZec6ykFkqjMy6EmSzC1Amd3hd"
	secretKey := "BoG_hT6idwA85rFQ4vpmJGiHXzVOur9RtQm6RtaQ"

	mac := auth.New(accessKey, secretKey)
	manager = sms.NewManager(mac)
}
