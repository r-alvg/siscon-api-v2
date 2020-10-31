package response

type Model struct {
	Payload interface{} `json:"payload"`
	Links   []Link      `json:"links"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}
