package main

import (
	"log"
	"net/http"
	"time"

	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*
{"messageType":"user","messageId":929831286,"channelUrl":"CoinoneChartChat","createdAt":1500426596685,"updatedAt":0,"channelType":"open","message":"가족도팔아","data":"","_sender":{"nickname":"746기빵포","profileUrl":"https://sendbird.com/main/img/profiles/profile_19_512px.png","userId":"205c422b93637ff2b28a082b992e23986","lastSeenAt":0,"connectionStatus":"nonavailable"},"reqId":"","customType":"","translations":{}}
*/
type Message struct {
	Type             string `json:"messageType"`
	ID               int    `json:"messageId"`
	ChannelUrlstring string `json:"channelUrl"`
	ChannelType      string `json:"channelType"`
	Message          string `json:"message"`
	CreatedAt        int    `json:"createdAt"`
}

var (
	chatCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "chat_count",
		Help: "chat count",
	})
)

func init() {
	prometheus.MustRegister(chatCounter)
}

func main() {
	done := make(chan bool)

	go run_prom()
	go run_chat()

	<-done // Block forever
}

func run_prom() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9200", nil)
}

func run_chat() {
	url := gosocketio.GetUrl("push.coinone.co.kr", 443, true)
	t := transport.GetDefaultWebsocketTransport()
	c, err := gosocketio.Dial(url, "/chat", t)

	if err != nil {
		log.Fatal(err)
	}

	c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("Connected")
	})

	c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Fatal("Disconnected")
	})

	c.On("previous_messages", func(h *gosocketio.Channel, messages []Message) {
	})

	c.On("new_message", func(h *gosocketio.Channel, message Message) {
		chatCounter.Inc()
	})

	for {
		time.Sleep(1 * time.Second)
	}

	c.Close()

	log.Println(" [x] Complete")
}
