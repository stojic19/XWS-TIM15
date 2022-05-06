package persistence

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/domain"
	"io"
	"log"
)

type FollowersStore struct {
	driver       neo4j.Driver
	databaseName string
}

func NewFollowersStore(driver *neo4j.Driver, dbName string) *FollowersStore {
	return &FollowersStore{
		driver:       *driver,
		databaseName: dbName,
	}
}

func (store *FollowersStore) GetFollows(username string) ([]*domain.User, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	followers, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH (:User {username:$username})-[:FOLLOWING]->(followed:User) RETURN followed",
			map[string]interface{}{"username": username})
		if err != nil {
			return nil, err
		}
		results := []*domain.User{}
		for records.Next() {
			record := records.Record()
			username, _ := record.Get("followed")
			user := domain.User{
				Username: username.(dbtype.Node).Props["username"].(string),
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

func (store *FollowersStore) GetFollowers(username string) ([]*domain.User, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	followers, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH (:User {username:$username})<-[:FOLLOWING]-(follower:User) RETURN follower",
			map[string]interface{}{"username": username})
		if err != nil {
			return nil, err
		}
		results := []*domain.User{}
		for records.Next() {
			record := records.Record()
			username, _ := record.Get("follower")
			user := domain.User{
				Username: username.(dbtype.Node).Props["username"].(string),
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

func (store *FollowersStore) Follow(followerUsername string, followedUsername string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"MERGE (followed:User {username: $followedUsername}) "+
				"ON CREATE SET followed.username = $followedUsername "+
				"MERGE (follower:User {username: $followerUsername}) "+
				"ON CREATE SET follower.username = $followerUsername "+
				"MERGE (followed) <- [:FOLLOWING] - (follower)",
			map[string]interface{}{"followedUsername": followedUsername, "followerUsername": followerUsername})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return "Failed to follow: " + followerUsername + " -> " + followedUsername, err
	}
	return session.LastBookmark(), nil
}

func (store *FollowersStore) FollowRequest(followerUsername string, followedUsername string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"MERGE (followed:User {username: $followedUsername}) "+
				"ON CREATE SET followed.username = $followedUsername "+
				"MERGE (follower:User {username: $followerUsername}) "+
				"ON CREATE SET follower.username = $followerUsername "+
				"MERGE (followed) <- [:REQUESTING_FOLLOW] - (follower)",
			map[string]interface{}{"followedUsername": followedUsername, "followerUsername": followerUsername})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return "Failed to create follow request: " + followerUsername + " -> " + followedUsername, err
	}
	return session.LastBookmark(), nil
}

func (store *FollowersStore) ConfirmFollow(followerUsername string, followedUsername string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"MERGE (followed:User {username: $followedUsername}) "+
				"ON CREATE SET followed.username = $followedUsername "+
				"MERGE (follower:User {username: $followerUsername}) "+
				"ON CREATE SET follower.username = $followerUsername "+
				"MATCH (followed) <- [followRequest:REQUESTING_FOLLOW] - (follower)"+
				"MERGE (followed) <- [:FOLLOWING] - (follower)"+
				"DELETE followRequest",
			map[string]interface{}{"followedUsername": followedUsername, "followerUsername": followerUsername})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return "Failed to create follow request: " + followerUsername + " -> " + followedUsername, err
	}
	return session.LastBookmark(), nil
}

func unsafeClose(closeable io.Closer) {
	if err := closeable.Close(); err != nil {
		log.Fatal(fmt.Errorf("could not close resource: %w", err))
	}
}
