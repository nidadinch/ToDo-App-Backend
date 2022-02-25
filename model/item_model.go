package model

type Item struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type ItemsResponse map[int]string
