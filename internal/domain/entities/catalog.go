package entities

type (
	Liquor struct {
		ID                   string `json:"_id"`
		Name                 string `json:"name,omitempty"`
		EAN                  int    `json:"EAN,omitempty"`
		Category             string `json:"category,omitempty"`
		Description          string `json:"description,omitempty"`
		AdditionalAttributes string `json:"additional_attributes,omitempty"`
	}

	Recipe struct {
		ID           string   `json:"_id"`
		Name         string   `json:"name,omitempty"`
		Ingredients  []string `json:"ingredients,omitempty"`
		Instructions string   `json:"instructions,omitempty"`
		Category     string   `json:"category,omitempty"`
		CreatedAt    string   `json:"createdAt,omitempty"`
		Liquors      []string `json:"liquors,omitempty"`
	}

	AIRecipe struct {
		CocktailName string       `json:"cocktailName"`
		Ingredients  []Ingredient `json:"ingredients"`
		Steps        []string     `json:"steps"`
		Observations string       `json:"observations"`
	}

	Ingredient struct {
		Name     string `json:"name"`
		Quantity string `json:"quantity"`
	}

	Product struct {
		Name                 string `json:"name"`
		PhotoLink            string `json:"photo_link"`
		Description          string `json:"description"`
		AdditionalAttributes string `json:"additional_attributes"`
		ISBN                 string `json:"isbn"`
	}
)
