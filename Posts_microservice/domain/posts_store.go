package domain

type PostsStore interface {
	GetAll() ([]*Post, error)
	Create(*Post) error
}
