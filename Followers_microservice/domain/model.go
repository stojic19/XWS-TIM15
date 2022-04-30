package domain

type User struct {
	username         string
	following        []User
	followers        []User
	followRequests   []User
	followerRequests []User
}
