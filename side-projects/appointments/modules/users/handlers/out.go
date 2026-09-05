package handlers

type CreateUserOut struct {
	Username string
	ID       int64 `json:"id,string"`
	CPF      string
	Email    string
}
