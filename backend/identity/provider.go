package identity

// Provider is a specialised database concerned exclusively
// with managing Users and a source of authority for their
// authentication.
type Provider interface {
	IsEditable() bool
	IsFederated() bool

	CreateUser(user *User) (*User, error)
	ReadUser(username string) (*User, error)
	UpdateUser(username string, user *User) (*User, error)

	LoginUser(user *User, password string) bool
	ChangePassword(user *User, password string) error

	ConsumeEndpoint(payload []byte) error
}

type UserNotFound error
