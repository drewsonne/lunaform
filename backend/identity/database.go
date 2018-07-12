package identity

import (
	"errors"
	"github.com/drewsonne/lunarform/backend/database"
)

// NewDatabaseIdentityProvider is not yet implemented and will return an error
func NewDatabaseIdentityProvider(db database.Database) (idp Provider, err error) {
	return nil, errors.New("Database IdP Not Implemented")
}
