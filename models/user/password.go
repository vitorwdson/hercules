package user

func (u *User) SetPassword(newPassword string) {
    // TODO: encrypt password before saving
    u.password = newPassword 
}

func (u User) ValidatePassword (password string) bool {
    // TODO: check against encrypted password
    return password == u.password
}

