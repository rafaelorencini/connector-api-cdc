package kafka_connect

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

const endpoint = "http://localhost:8083/connectors"

type KafkaConnect struct {
}

func (c *KafkaConnect) httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func (c *KafkaConnect) sendRequest(client *http.Client, method string, payload string) *http.Response {
	var jsonData = []byte(payload)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

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
