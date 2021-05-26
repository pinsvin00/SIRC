package main

type Message struct {
	text string `json:"text"`
	sender string `json:"sender"`
	time string `json:"time"`
	color string `json:"color"`
}

type IRC struct {
	id int `json:"id"`
	counter int
	guid_nickname map[string]string
	guid_color map[string]string
	nicknames []string
	subscribers int
	message_queue map[string][]Message `json:"messages"`
}