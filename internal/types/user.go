package types

const UserContextKey = "user"

type AuthenticatedUser struct {
	Id       int
	Email    string
	LoggedIn bool
}
