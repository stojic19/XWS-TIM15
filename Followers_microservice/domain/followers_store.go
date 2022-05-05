package domain

type FollowersStore interface {
	GetFollowing(username string) ([]*User, error)
	Follow(followerUsername string, followedUsername string) (string, error)
	FollowRequest(followerUsername string, followedUsername string) (string, error)
}
