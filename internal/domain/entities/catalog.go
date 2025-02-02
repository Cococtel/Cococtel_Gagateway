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

	User struct {
		UserID   string `json:"user_id"`
		Name     string `json:"name,omitempty"`
		Lastname string `json:"lastname,omitempty"`
		Email    string `json:"email,omitempty"`
		Country  string `json:"country,omitempty"`
		Phone    string `json:"phone,omitempty"`
		Image    string `json:"image,omitempty"`
	}

	SuccessfulLogin struct {
		UserID      string `json:"id"`
		Name        string `json:"name,omitempty"`
		DoubleAuth  bool   `json:"double_auth,omitempty"`
		Expiration  string `json:"expiration,omitempty"`
		Token       string `json:"token,omitempty"`
		AccountType string `json:"account_type,omitempty"`
	}
)
