package user

import "golang.org/x/crypto/bcrypt"

func (u *User) SetPassword(newPassword string) error {
    hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    u.password = string(hash) 

    return nil
}

func (u User) ValidatePassword (password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
    return err == nil
}

func CheckPasswordStrength(password string) string {
    return ""
}
