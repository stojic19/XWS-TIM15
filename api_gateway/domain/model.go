package domain

import "time"

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
