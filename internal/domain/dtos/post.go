package dtos

type (
	Post struct {
		UrlImage string `json:"urlImage"`
		Title    string `json:"title"`
		Content  string `json:"content"`
		Author   string `json:"author"`
	}
)
