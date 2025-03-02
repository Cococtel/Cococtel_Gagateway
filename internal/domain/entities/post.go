package entities

type (
	Interaction struct {
		Type      int    `json:"type"`
		Value     string `json:"value"`
		UserId    string `json:"userId"`
		CreatedAt string `json:"createdAt"`
	}

	Post struct {
		ID           string        `json:"_id"`
		UrlImage     string        `json:"urlImage"`
		Title        string        `json:"title"`
		Content      string        `json:"content"`
		Author       string        `json:"author"`
		CreatedAt    string        `json:"createdAt"`
		Interactions []Interaction `json:"interactions"`
	}
)
