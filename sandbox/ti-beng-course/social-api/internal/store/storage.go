package store

import (
	"context"
	"database/sql"
)

// app.store.Posts.Create()

type Storage struct {
	Posts interface {
		Create(context.Context) error
	}
	Users interface {
		Create(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage {
		Posts: &PostsStore{db},
		Users : &UserStore{db},
	}
}