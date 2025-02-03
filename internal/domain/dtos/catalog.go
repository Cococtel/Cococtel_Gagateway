package dtos

type (
	Liquor struct {
		Name                 string `json:"name"`
		EAN                  int    `json:"EAN,omitempty"`
		Category             string `json:"category,omitempty"`
		Description          string `json:"description,omitempty"`
		AdditionalAttributes string `json:"additional_attributes,omitempty"`
	}

	Recipe struct {
		Name         string   `json:"name"`
		Ingredients  []string `json:"ingredients,omitempty"`
		Instructions string   `json:"instructions,omitempty"`
		Category     string   `json:"category,omitempty"`
		Liquors      []string `json:"liquors,omitempty"`
	}
)
