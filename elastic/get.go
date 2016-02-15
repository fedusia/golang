package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	elasticHost     = "elk-102.ix.km"
	clusterStatusOk = "green"
	excludeFile     = "./elastic_exlude.txt"
)

type myCluster struct {
	Status string `json:"status"`
}

type myConfig struct {
	ExcludedFileNames []string
}

// type myIndicies struct {
// 	indices map[string]map[string]string `json: indices`
// }

var (
	cluster myCluster
	config  myConfig
	f       interface{}
)

func main() {
	bytes, err := ioutil.ReadFile(excludeFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	if len(config.ExcludedFileNames) == 0 {
		log.Fatalln("Exclude file is empty")
	}

	ClusterStatusUrl := fmt.Sprint("http://" + elasticHost + ":9200/_cluster/health?pretty")
	resp, err := http.Get(ClusterStatusUrl)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Response status: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Couldn't read data: %v", body)
	}
	err = json.Unmarshal(body, &cluster)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	if cluster.Status != clusterStatusOk {
		log.Fatalf("Cluster status is: %s!!! Do nothing ", cluster.Status)
	}

	ClusterIndicesUrl := fmt.Sprint("http://" + elasticHost + ":9200/_stats?pretty")
	resp, err = http.Get(ClusterIndicesUrl)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Response status: %v", resp.Status)
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Couldn't read data: %v", body)
	}
	fmt.Println(string(body))

	err = json.Unmarshal(body, &f)
	if err != nil {
		log.Fatalf("error:", err)
	}
	for key, value := range f{} {
		fmt.Println(key)
	}
}
