package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Message struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Content      string
	MsgId        int
}

type Response struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Content      string
}

func main() {
	http.HandleFunc("/", wechatHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func wechatHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Read request body error: ", err)
		return
	}
	fmt.Println("Received message: ", string(body))

	var msg Message
	err = xml.Unmarshal(body, &msg)
	if err != nil {
		log.Println("XML unmarshal error: ", err)
		return
	}

	response := Response{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   int(time.Now().Unix()),
		MsgType:      "text",
		Content:      "你好，我是ChatAI，你的消息已收到！",
	}

	respXML, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Println("XML marshal error: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(respXML)
}
