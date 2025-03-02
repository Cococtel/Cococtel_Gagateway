package entities

type (
	User struct {
		UserID   string `json:"user_id"`
		Name     string `json:"name,omitempty"`
		Lastname string `json:"lastname,omitempty"`
		Email    string `json:"email,omitempty"`
		Country  string `json:"country,omitempty"`
		Phone    string `json:"phone,omitempty"`
		Image    string `json:"image,omitempty"`
		Username string `json:"username,omitempty"`
	}

	SuccessfulLogin struct {
		UserID      string `json:"id"`
		Name        string `json:"name,omitempty"`
		DoubleAuth  bool   `json:"double_auth,omitempty"`
		Expiration  string `json:"expiration,omitempty"`
		Token       string `json:"token,omitempty"`
		AccountType string `json:"account_type,omitempty"`
	}

	UserResponse struct {
		Data User `json:"data"`
	}

	LoginResponse struct {
		Data SuccessfulLogin `json:"data"`
	}
)
