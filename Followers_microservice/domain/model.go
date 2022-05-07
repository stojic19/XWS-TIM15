package domain

type User struct {
	Id               string
	Following        []User
	Followers        []User
	FollowRequests   []User
	FollowerRequests []User
}
