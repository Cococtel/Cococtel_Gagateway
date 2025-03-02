package dtos

type (
	Liquor struct {
		Name                 string `json:"name"`
		EAN                  int    `json:"EAN,omitempty"`
		Category             string `json:"category,omitempty"`
		Description          string `json:"description,omitempty"`
		AdditionalAttributes string `json:"additional_attributes,omitempty"`
	}

	Ingredient struct {
		Name     string `json:"name"`
		Quantity string `json:"quantity"`
	}

	Rating struct {
		UserID string `json:"user_id"`
		Rating int    `json:"rating"`
	}

	Recipe struct {
		Name         string       `json:"name"`
		Category     string       `json:"category"`
		Ingredients  []Ingredient `json:"ingredients"`
		Instructions []string     `json:"instructions"`
		CreatorId    string       `json:"creatorId"`
		Description  string       `json:"description"`
	}
)
