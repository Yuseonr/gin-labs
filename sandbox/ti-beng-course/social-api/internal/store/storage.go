package store

import (
	"context"
	"database/sql"
)

// app.store.Posts.Create()

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage {
		Posts: &PostStore{db},
		Users : &UserStore{db},
	}
}