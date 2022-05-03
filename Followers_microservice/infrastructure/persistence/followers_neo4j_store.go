package persistence

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/domain"
	"io"
	"log"
)

const (
	DATABASE   = "catalogue"
	COLLECTION = "product"
)

type FollowersStore struct {
	driver neo4j.Driver
}

func NewFollowersStore(driver *neo4j.Driver) *FollowersStore {
	return &FollowersStore{
		driver: *driver,
	}
}

func (store *FollowersStore) GetFollowing(username string) ([]*domain.User, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: "followers",
	})
	defer unsafeClose(session)

	followers, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH (:User {username:$username})<-[:FOLLOWING]-(follower) RETURN follower",
			map[string]interface{}{"username": username})
		if err != nil {
			return nil, err
		}
		var results []*domain.User
		for records.Next() {
			record := records.Record()
			username, _ := record.Get("username")
			user := domain.User{
				Username: username.(string),
			}
			results = append(results, &user)
		}
		return results, nil
	})
	if err != nil {
		return nil, err
	}
	return followers.([]*domain.User), nil
}

func unsafeClose(closeable io.Closer) {
	if err := closeable.Close(); err != nil {
		log.Fatal(fmt.Errorf("could not close resource: %w", err))
	}
}
