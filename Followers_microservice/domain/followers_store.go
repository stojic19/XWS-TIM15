package domain

type FollowersStore interface {
	GetFollowing(username string) ([]*User, error)
}
