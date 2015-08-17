package main

import (
	"strconv"
	"net/http"
	"time"
	"fmt"
)



func main() {
	lat := "28.611"
	lon := "77.2200"
	
	flag := 1

	for flag==1{
		url := "localhost:8080/data?lat="+lat+"&lon="+lon
		fmt.Println(url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil{
			
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		fmt.Println(resp)
		lng, _ := strconv.ParseFloat(lon, 64)
		lng = lng+0.0001
		lon = strconv.FormatFloat(lng, 'f', 6, 64)
		fmt.Println(lon)
		time.Sleep(3*time.Second)

	}
	
}