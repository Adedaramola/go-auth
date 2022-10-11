package auth

import "github.com/jmoiron/sqlx"

type AuthManager struct {
	db *sqlx.DB
}

func (a *AuthManager) Login() bool {

	return true
}
