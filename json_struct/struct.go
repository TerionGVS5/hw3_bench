package json_struct

type User struct {
	Browsers []string `json:"browsers"`
	Email    string `json:"email"`
	Name string `json:"name"`
}