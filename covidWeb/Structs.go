package main


type News struct{
	Title string `json:"title"`
	Date string `json:"date"`
	ShortText string `json:"shortText"`
	URL string `json:"url"`
}

type ReturnType struct{
	NewsList []News
}