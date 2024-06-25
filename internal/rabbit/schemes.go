package rabbit

type Message struct {
	Topic string `json:"topic"`
	Body  string `json:"body"`
	HTML  string `json:"html"`
}

type To struct {
	Type      string `json:"type"`
	Recipient string `json:"recipient"`
}

type InData struct {
	To      []To    `json:"to"`
	Message Message `json:"message"`
}
