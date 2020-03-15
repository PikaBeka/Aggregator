package main

import(
	"fmt"
	"net/http"
	"io/ioutil"

)

func GetAllNews() (ReturnType, error) {
	var data ReturnType
	var news []News 
	resp, err := http.Get("http://localhost:5000/akmola")
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return data, err1
	}
	json.Unmarshal([]byte(body), &news)

	data.NewsList = news
	return data, nil
}