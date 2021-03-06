package domain

import (
	"time"
)

type User struct {
	Id       string
	Username string
}

type UserFollowerInfoList struct {
	Users []*UserFollowerInfo
}

type UserFollowerInfo struct {
	Id                  string
	Username            string
	Name                string
	Gender              string
	StartDate           time.Time
	ReverseRelationship string
}

type PostUsersInfoList struct {
	Posts []*PostUsersInfo
}

type PostUsersInfo struct {
	Id         string
	Title      string
	Content    Content
	CreateTime time.Time
	Owner      *UserPostInfo
	Comments   []*CommentPostInfo
	Likes      []*UserPostInfo
	Dislikes   []*UserPostInfo
}

type Content struct {
	Text   string
	Links  []string
	Images []string
}

type UserPostInfo struct {
	Id                   string
	Username             string
	Name                 string
	Gender               string
	DateOfBirth          time.Time
	OutgoingRelationship string
	IngoingRelationship  string
}

type CommentPostInfo struct {
	Owner      *UserPostInfo
	Content    string
	CreateTime time.Time
}

type JobOffersUsersInfoList struct {
	JobOffers []*JobOfferUsersInfo
}

type JobOfferUsersInfo struct {
	Id           string
	Position     string
	Description  string
	Requirements string
	IsActive     bool
	Subscribers  []*UserJobOfferInfo
}

type UserJobOfferInfo struct {
	Id          string
	Username    string
	Name        string
	Gender      string
	DateOfBirth time.Time
}
