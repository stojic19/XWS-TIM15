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
	BlockPending(blockerId string, blockedId string) (string, error)
	ConfirmBlock(blockerId string, blockedId string) (string, error)
	RevertPendingBlock(blockerId string, blockedId string) (string, error)
	UnblockPending(blockerId string, blockedId string) (string, error)
	ConfirmUnblock(blockerId string, blockedId string) (string, error)
	RevertPendingUnblock(blockerId string, blockedId string) (string, error)
	GetBlocked(id string) ([]*User, error)
	GetBlockers(id string) ([]*User, error)
	GetRecommended(id string) ([]*User, error)
}
