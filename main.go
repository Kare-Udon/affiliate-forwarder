package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Sites struct {
	Sites []Site `json:"sites"`
}

type Site struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	sites, err := initConf()
	if err != nil {
		fmt.Printf("Reading configuration error: %v\n", err)
		return
	}
	// Convert the sites to a map
	siteMap := make(map[string]Site)
	for _, site := range sites.Sites {
		siteMap[site.Name] = site
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the site name from the URL
		siteName := r.URL.Path[1:]
		if siteName == "favicon.ico" {
			// Ignore favicon requests
			return
		}
		// Get the site from the map
		site, ok := siteMap[siteName]
		if !ok {
			fmt.Printf("Site %s not found\n", siteName)
			return
		}
		// Redirect to the site
		http.Redirect(w, r, site.Url, http.StatusMovedPermanently)
	})
	fmt.Printf("Starting server at port 8080\n")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

func initConf() (Sites, error) {
	// get redirect conf path
	confPath := os.Getenv("AFF_FWD_CONF_PATH")
	if confPath == "" {
		return Sites{}, fmt.Errorf("AFF_FWD_CONF_PATH is not set")
	}
	jsonFile, err := os.Open(confPath)
	defer jsonFile.Close()
	if err != nil {
		return Sites{}, fmt.Errorf("Error opening JSON file: %v", err)
	}
	// read the configuration file
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return Sites{}, fmt.Errorf("Error reading JSON file: %v", err)
	}
	var sites Sites
	err = json.Unmarshal(byteValue, &sites)
	if err != nil {
		return Sites{}, fmt.Errorf("Error parsing JSON file: %v", err)
	}
	return sites, nil
}
