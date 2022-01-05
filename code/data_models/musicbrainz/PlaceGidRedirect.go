package dto

import (
	pg "github.com/lib/pq"
)

type PlaceGidRedirect struct {
	GID     sql.NullString `db:"gid"`
	NewID   int            `db:"new_id"`
	Created pg.NullTime    `db:"created"`
}
