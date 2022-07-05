package main

type (
	User struct {
		Email     string `json:"email"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}

	DbUser struct {
		UID       string `json:"uid"`
		Email     string `json:"email"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
		CreatedAt string `json:"created_at"`
	}

	Credentials struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
)
