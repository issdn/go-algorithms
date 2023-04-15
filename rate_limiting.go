package main

import (
	"crypto/rand"
	"encoding/base64"
)

type FixedWindow struct {
	Interval     uint16
	UserRequests map[string]uint16
}

func (fw *FixedWindow) AddUser() string {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	randomToken := base64.StdEncoding.EncodeToString(randomBytes)
	fw.UserRequests[randomToken] = 1
	return randomToken
}

func (fw *FixedWindow) Request(userId string) (string, uint16) {
	token := ""
	var requests uint16 = 1
	if _requests, ok := fw.UserRequests[userId]; ok {
		fw.UserRequests[userId]++
		requests = _requests + 1
	} else {
		token = fw.AddUser()
	}
	return token, requests
}
