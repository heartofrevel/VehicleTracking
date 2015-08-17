package main

import (
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
)

var lat string 
var lon string
var flag int = 0



type Data struct{
	Latitude string
	Longitude string
}



func main() {
	http.Handle("/", http.FileServer(http.Dir("frontend")))
	http.HandleFunc("/tracker", htmlHandler)
	http.HandleFunc("/data", dataHandler)
	http.ListenAndServe(":80", nil)
}

func htmlHandler(w http.ResponseWriter, r *http.Request){
	if flag == 0{
		flag = 1
	}else{
		flag = 0
	}
	var trackingData Data
	trackingData.Latitude = lat
	trackingData.Longitude = lon
	byteResponse, _ := json.Marshal(trackingData)
	w.Write(byteResponse)
	
}


func dataHandler(w http.ResponseWriter, r *http.Request){
	urlv, _ := url.ParseQuery(r.URL.RawQuery)
	lat = urlv.Get("lat")
	lon = urlv.Get("lon")
	fmt.Println(lon)
}

