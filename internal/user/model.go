package user

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username,omitempty"`
	PasswordHash string `json:"-" bson:"password_hash"`
	Email        string `json:"email" bson:"email"`
}

type CreateUserDTO struct {
	Email        string `json:"email"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}
