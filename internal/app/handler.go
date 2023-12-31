package app

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

const serverURL string = "http://localhost:8080/"

var idMap = make(map[string]string)

func RootRouter(res http.ResponseWriter, req *http.Request) {
	// Check path and method to route to corresponding function
	if req.URL.Path == "/" && req.Method == http.MethodPost {
		postShortUrl(res, req)
	} else if req.Method == http.MethodGet && req.URL.Path != "/" {
		getOrigPageRedir(res, req)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}

func postShortUrl(res http.ResponseWriter, req *http.Request) {
	// Read original URL from request body
	bodyURL, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	origURL := string(bodyURL)
	// Verify Original URL validity
	_, err = url.ParseRequestURI(origURL)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get ID for Original URL
	id := genterateID(origURL)
	shortURL := serverURL + id
	// If ID does not exist - write that to idMap
	if _, exists := idMap[id]; !exists {
		idMap[id] = origURL
	}
	// Write Shortened URL to HTTP response
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(shortURL))
	// Write HTTP responce status and headers
	//fmt.Println(serverURL + id)
}

func getOrigPageRedir(res http.ResponseWriter, req *http.Request) {
	// Get hash from GET request URL Path
	hash := strings.TrimPrefix(req.URL.Path, "/")
	// Check if has exists in idMap (aka original URL is stored)
	origUrl, exists := idMap[hash]
	if exists {
		res.Header().Set("Location", origUrl)
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
