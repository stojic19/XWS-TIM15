package domain

type User struct {
	Username         string
	Following        []User
	Followers        []User
	FollowRequests   []User
	FollowerRequests []User
}
