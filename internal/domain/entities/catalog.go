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

	Ingredient struct {
		ID       string `json:"_id,omitempty"`
		Name     string `json:"name"`
		Quantity string `json:"quantity"`
	}

	Rating struct {
		UserID string  `json:"user_id"`
		Rating float64 `json:"rating"`
	}

	Recipe struct {
		ID            string       `json:"_id"`
		Name          string       `json:"name"`
		Category      string       `json:"category"`
		Ingredients   []Ingredient `json:"ingredients"`
		Instructions  []string     `json:"instructions"`
		CreatorId     string       `json:"creatorId"`
		Rating        float64      `json:"rating"`
		Likes         int          `json:"likes"`
		Liquors       []string     `json:"liquors"`
		CreatedAt     string       `json:"createdAt"`
		Ratings       []Rating     `json:"ratings"`
		Description   string       `json:"description"`
		AverageRating float64      `json:"averageRating"`
	}

	AIRecipe struct {
		CocktailName string       `json:"cocktailName"`
		Ingredients  []Ingredient `json:"ingredients"`
		Steps        []string     `json:"steps"`
		Observations string       `json:"observations"`
	}

	Product struct {
		Name                 string `json:"name"`
		PhotoLink            string `json:"photo_link"`
		Description          string `json:"description"`
		AdditionalAttributes string `json:"additional_attributes"`
		ISBN                 string `json:"isbn"`
	}
)
