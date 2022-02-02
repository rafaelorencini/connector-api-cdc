package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

//type Connector struct {
//	Name   string `json:"name"`
//	Config Config `json:"config"`
//}
//
//type Config struct {
//	ConnectorClass                       string `json:"connector.class"`
//	taskMax                              string `json:"tasks.max"`
//	DatabaseName                         string `json:"database.hostname"`
//	DatabasePort                         string `json:"database.port"`
//	DatabaseUser                         string `json:"database.user"`
//	DatabasePassword                     string `json:"database.password"`
//	DatabaseServerID                     string `json:"database.server.id"`
//	DatabaseServerName                   string `json:"database.server.name"`
//	DatabaseInclude                      string `json:"database.include"`
//	DatabaseHistoryKafkaBootstrapServers string `json:"database.history.kafka.bootstrap.servers"`
//	DatabaseHistoryKafkaTopic            string `json:"database.history.kafka.topic"`
//}

func main() {
	//teste := &Config{ConnectorClass: "yrdy", taskMax: "wgfsdfg", DatabaseName: "gsdfg", DatabasePort: "gsdfgsd", DatabaseUser: "gsdfgsd", DatabasePassword: "gsdfg", DatabaseServerID: "gsdfgs", DatabaseServerName: "gsdfg", DatabaseInclude: "gsdfg", DatabaseHistoryKafkaBootstrapServers: "gsdfgd", DatabaseHistoryKafkaTopic: "fgsdfgd"}
	//emp := &Connector{Name: "testeste", Config: *teste}
	//e, err := json.Marshal(emp)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(e))

	//Run(3000)
	mytest()

	buf, err := ioutil.ReadFile("/Users/rafael.lorencini/connector-api-v1/defaults/mysql/sink.yaml")
	if err != nil {
		fmt.Println(err)
	}

	data := make(map[string]string)

	err = yaml.Unmarshal(buf, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data["behavior_on_null_values"], err)

}

func mytest() {

	client := http.Client{}
	url := "http://127.0.0.1:8083/connectors"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`
{
  "name": "inventory-connector",  
  "config": {  
    "connector.class": "io.debezium.connector.mysql.MySqlConnector",
    "tasks.max": "1",  
    "database.hostname": "mysql",  
    "database.port": "3306",
    "database.user": "debezium",
    "database.password": "dbz",
    "database.server.id": "184054",  
    "database.server.name": "dbserver1",  
    "database.include.list": "inventory",  
    "database.history.kafka.bootstrap.servers": "kafka:9092",  
    "database.history.kafka.topic": "schema-changes.inventory"  
  }
}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Unable to reach the server.")
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("body=", string(body))
	}
}
