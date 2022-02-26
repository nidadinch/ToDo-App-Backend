package model

type Item struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type ItemsResponse []Item

type ItemAddRequest struct {
	Text string `json:"text"`
}
