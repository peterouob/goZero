package test

import (
	"encoding/json"
	"fmt"
	"github.com/peterouob/goZero/define"
	"github.com/peterouob/goZero/helper"
	"testing"
)

var userServiceAddr = "http://127.0.0.1:8888"

var M map[string]interface{}

func TestUserLogin(t *testing.T) {
	m := define.M{
		"username": "peter",
		"password": "123456",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/user/login", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
