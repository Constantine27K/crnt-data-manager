package crnt_auth_service

const (
	EntityAdmin = "admin"
	EntityUser  = "user"
)

type Payload struct {
	Username string
	Role     string
	Team     string
}
