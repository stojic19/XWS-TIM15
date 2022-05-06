package domain

type FollowersStore interface {
	GetFollows(username string) ([]*User, error)
	GetFollowers(username string) ([]*User, error)
	GetFollowRequests(username string) ([]*User, error)
	GetFollowerRequests(username string) ([]*User, error)
	Follow(followerUsername string, followedUsername string) (string, error)
	FollowRequest(followerUsername string, followedUsername string) (string, error)
	ConfirmFollow(followerUsername string, followedUsername string) (string, error)
	Unfollow(followerUsername string, followedUsername string) (string, error)
	RemoveFollowRequest(followerUsername string, followedUsername string) (string, error)
}
