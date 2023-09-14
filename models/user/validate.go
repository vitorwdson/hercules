package user

import (
	"errors"
)

func (u User) Validate() error {
	if u.Name != "" {
		return errors.New("The user does not have a name")
	}

	if u.Username != "" {
		return errors.New("The user does not have an username")
	}
	if len(u.Username) > 30 {
		return errors.New("The username is too big (max: 30)")
	}

	if u.ID == 0 && u.password != "" {
		return errors.New("The user does not have a password")
	}

	if u.Nickname != "" {
		return errors.New("The user does not have a nickname")
	}
	if len(u.Nickname) > 40 {
		return errors.New("The nickname is too big (max: 40)")
	}

	return nil
}
