package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/newGid", newGid)

	r.Run(":9999")
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type jsonRpcHttpBase struct {
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  map[string]string `json:"params"`
	Id      string            `json:"params"`
}

func newGid(c *gin.Context) {
	// 生产GID
	body := jsonRpcHttpBase{
		Jsonrpc: "2.0",
		Method:  "dtmserver.NewGid",
		Params:  nil,
		Id:      RandStringRunes(10),
	}

	requestBody, err := json.Marshal(body)
	if err != nil {
		c.JSON(200, err.Error())
		return
	}
	resp, _ := http.Post("http://127.0.0.1:36791", "application/json", bytes.NewBuffer(requestBody))

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Printf("get request: %s \n", res)

	c.JSON(200, res)
	return
}
