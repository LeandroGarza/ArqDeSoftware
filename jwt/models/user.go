package models

type User struct {
	Name     string `json:"name"`
	Password string `json:"password, omitempty"` //osea si esta vacio q no lo vuelva json
	Role     string `json:"role"`
}
