package kafka_connect

import (
	"bytes"
	"github.com/rafaelorencini/connector-api-cdc/domain"
	"log"
	"net/http"
	"time"
)

const endpoint = "http://localhost:8083/connectors"

type KafkaConnect struct {
}

func NewKafkaConnect() domain.KafkaConnectInterface {
	return new(KafkaConnect)
}

func (c *KafkaConnect) SendRequest(method string, payload string) *http.Response {
	client := &http.Client{Timeout: 10 * time.Second}
	var jsonData = []byte(payload)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	// Close the connection to reuse it
	//defer response.Body.Close()
	//
	//body, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	log.Fatalf("Couldn't parse response body. %+v", err)
	//}

	return response
}
