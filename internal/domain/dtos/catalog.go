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

	Register struct {
		Name     *string `json:"name"`
		Lastname *string `json:"lastname,omitempty"`
		Phone    *string `json:"phone,omitempty"`
		Email    *string `json:"email,omitempty"`
		Image    *string `json:"image,omitempty"`
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
		Type     *string `json:"type,omitempty"`
	}

	Login struct {
		User     *string `json:"user,omitempty"`
		Password *string `json:"password,omitempty"`
		Type     *string `json:"type,omitempty"`
	}
)
