package mailing

// Message представляет собой структуру для хранения данных сообщения.
type Message struct {
	Topic string `json:"topic"` // Topic содержит тему сообщения.
	Body  string `json:"body"`  // Body содержит текстовое содержимое сообщения.
	HTML  string `json:"html"`  // HTML содержит HTML-версию сообщения, если она есть.
}

// To представляет получателя сообщения с указанием типа (например, "email" или "telegram") и адреса получателя.
type To struct {
	Type      string `json:"type"`      // Type указывает тип мессенджера получателя (например, "email" или "telegram").
	Recipient string `json:"recipient"` // Recipient содержит адрес получателя сообщения.
}

// InData представляет собой структуру для хранения входных данных сообщения,
// включая получателей и само сообщение.
type InData struct {
	To      []To    `json:"to"`      // To содержит массив получателей сообщения.
	Message Message `json:"message"` // Message содержит данные о сообщении (тема, текст, HTML).
}
