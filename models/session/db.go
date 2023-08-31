package session

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/vitorwdson/hercules/models/user"
)

func GetByUUID(db *sql.DB, id uuid.UUID) (*Session, error) {
	row := db.QueryRow(`
        SELECT
            s.id,
            s.created_at,
            s.revoked,
            u.id,
            u.username,
            u.name,
            u.nickname
        FROM
            sessions s
        LEFT JOIN
            users u
                ON u.id = s.user_id
        WHERE
            s.id = $1;            
    `, id)

	session := Session{}
    session.User = &user.User{}
	err := row.Scan(
		&session.ID,
		&session.CreatedAt,
		&session.Revoked,
		&session.User.ID,
		&session.User.Username,
		&session.User.Name,
		&session.User.Nickname,
	)

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *Session) Save(db *sql.DB) error {
	if s.ID != uuid.Nil {
		_, err := db.Exec(`
            UPDATE
                sessions
            SET
                revoked = $1,
            WHERE
                id = $2;
        `, s.Revoked, s.ID)
		if err != nil {
			return err
		}
	} else {
		err := db.QueryRow(`
            INSERT INTO
                sessions (
                    user_id,
                    created_at,
                    revoked
                )
            VALUES (
                $1,
                $2,
                $3,
            ) RETURNING id;
        `, s.User.ID, s.CreatedAt, s.Revoked).Scan(&s.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
