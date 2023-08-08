package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main1() {
	_, err := initConf()
	if err != nil {
		fmt.Printf("Reading configuration error: %v\n", err)
		return
	}
	http.HandleFunc("/", redirectHandler)
	fmt.Printf("Starting server at port 8080\n")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

func initConf() (map[string]interface{}, error) {
	// get redirect conf path
	confPath := os.Getenv("AFF_FWD_CONF_PATH")
	if confPath == "" {
		return nil, fmt.Errorf("AFF_FWD_CONF_PATH is not set")
	}
	// read the configuration file
	config, err := mapReader(confPath)
	return config, err
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {

}

func mapReader(confPath string) (map[string]interface{}, error) {
	f, err := os.Open(confPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Could not read the configuration file: %v\n", err)
		return nil, err
	}

	// Parse the JSON into a map
	var config map[string]interface{}
	err = json.Unmarshal(content, &config)
	if err != nil {
		fmt.Printf("Could not parse the configuration file: %v\n", err)
		return nil, err
	}

	return config, nil
}
