package domain

type PostsStore interface {
	GetAll() ([]*Post, error)
	Create(post *Post) error
}
