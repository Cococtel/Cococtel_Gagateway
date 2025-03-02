package dtos

type (
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

	User struct {
		Name     *string `json:"name"`
		Lastname *string `json:"lastname,omitempty"`
		Phone    *string `json:"phone,omitempty"`
		Email    *string `json:"email,omitempty"`
		Username *string `json:"username,omitempty"`
		Image    *string `json:"image,omitempty"`
	}
)
