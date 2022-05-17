package domain

type FollowersStore interface {
	GetFollows(id string) ([]*User, error)
	GetFollowers(id string) ([]*User, error)
	GetFollowRequests(id string) ([]*User, error)
	GetFollowerRequests(id string) ([]*User, error)
	GetRelationship(followerId string, followedId string) (string, error)
	Follow(followerId string, followedId string) (string, error)
	FollowRequest(followerId string, followedId string) (string, error)
	ConfirmFollow(followerId string, followedId string) (string, error)
	Unfollow(followerId string, followedId string) (string, error)
	RemoveFollowRequest(followerId string, followedId string) (string, error)
}
