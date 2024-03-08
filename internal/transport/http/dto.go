package http

type SignUpDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateDto struct {
	Email      string `json:"email"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Patronymic string `json:"patronymic"`
	Phone      string `json:"phone"`
}

type SendRecoveryDto struct {
	Email string `json:"email"`
}

type ConfirmRecoveryDto struct {
	Key      string `json:"key"`
	Password string `json:"password"`
}
