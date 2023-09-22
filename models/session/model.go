package session

import (
	"time"

	"github.com/google/uuid"
	"github.com/vitorwdson/hercules-go/models/user"
)

type Session struct {
    ID uuid.UUID
    User *user.User
    CreatedAt time.Time
    Revoked bool
}
