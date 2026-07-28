package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

// post merepresentasikan satu baris di tabel posts (model data / struct tag)
// struct tag `json:"..."` dipakai saat struct ini diserialize ke json response
type Post struct {
	ID       int64    `json:"id"`
	Content  string   `json:"content"`
	Title    string   `json:"title"`
	UserID   int64    `json:"user_id"`
	Tags     []string `json:"tags"`
	CreateAt string   `json:"created_at"` 
	UpdateAt string   `json:"updated_at"`
}

// postsstore membungkus koneksi *sql.db untuk operasi crud di tabel posts
// pattern ini disebut repository / data access layer
type PostStore struct {
	db *sql.DB
}

// create menjalankan insert ke tabel posts, lalu memindai balik id, created_at, dan updated_at
// menggunakan returning agar tidak perlu select lagi setelah insert
// post dikirim sebagai pointer karena method ini akan mengisi field id, created_at, updated_at
func (s *PostStore) Create(ctx context.Context, post *Post) error {
	// $1..$4 adalah placeholder postgres (parameterized query) untuk mencegah sql injection
	// berbeda dengan mysql yang pakai ?
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	// queryrowcontext dipakai karena query hanya mengembalikan 1 baris
	// ctx memungkinkan query dibatalkan kalau request http sudah timeout atau client disconnect
	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags), // pq.array mengkonversi []string go ke format postgres array {val1,val2}
	).Scan(
		&post.ID,       // scan menulis nilai dari returning ke field struct (pakai pointer)
		&post.CreateAt,
		&post.UpdateAt,
	)
	if err != nil {
		return err // error dikembalikan ke caller (handler), bukan dipanic
	}
	return nil
}
