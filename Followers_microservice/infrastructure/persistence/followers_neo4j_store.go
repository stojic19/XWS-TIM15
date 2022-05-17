package persistence

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/domain"
	"io"
	"log"
	"time"
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

func (store *FollowersStore) GetFollows(id string) ([]*domain.User, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	followers, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH (:User {userId:$id})-[follow:FOLLOWING]->(followed:User) RETURN followed, follow",
			map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		results := []*domain.User{}
		for records.Next() {
			record := records.Record()
			id, _ := record.Get("followed")
			relationship, _ := record.Get("follow")
			user := domain.User{
				Id:           id.(dbtype.Node).Props["userId"].(string),
				TimeOfFollow: time.Time(relationship.(dbtype.Relationship).Props["timeStarted"].(dbtype.LocalDateTime)),
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

func (store *FollowersStore) GetFollowers(id string) ([]*domain.User, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	followers, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH (:User {userId:$id})<-[follow:FOLLOWING]-(follower:User) RETURN follower, follow",
			map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		results := []*domain.User{}
		for records.Next() {
			record := records.Record()
			id, _ := record.Get("follower")
			relationship, _ := record.Get("follow")
			user := domain.User{
				Id:           id.(dbtype.Node).Props["userId"].(string),
				TimeOfFollow: time.Time(relationship.(dbtype.Relationship).Props["timeStarted"].(dbtype.LocalDateTime)),
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

func (store *FollowersStore) GetFollowRequests(id string) ([]*domain.User, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	followers, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH (:User {userId:$id})-[follow:REQUESTING_FOLLOW]->(followed:User) RETURN followed, follow",
			map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		results := []*domain.User{}
		for records.Next() {
			record := records.Record()
			id, _ := record.Get("followed")
			relationship, _ := record.Get("follow")
			user := domain.User{
				Id:           id.(dbtype.Node).Props["userId"].(string),
				TimeOfFollow: time.Time(relationship.(dbtype.Relationship).Props["timeSent"].(dbtype.LocalDateTime)),
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

func (store *FollowersStore) GetFollowerRequests(id string) ([]*domain.User, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	followers, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH (:User {userId:$id})<-[follow:REQUESTING_FOLLOW]-(follower:User) RETURN follower, follow",
			map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		results := []*domain.User{}
		for records.Next() {
			record := records.Record()
			id, _ := record.Get("follower")
			relationship, _ := record.Get("follow")
			user := domain.User{
				Id:           id.(dbtype.Node).Props["userId"].(string),
				TimeOfFollow: time.Time(relationship.(dbtype.Relationship).Props["timeSent"].(dbtype.LocalDateTime)),
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
func (store *FollowersStore) GetRelationship(followerId string, followedId string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	relationship, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MATCH (:User {userId:$followedId})<-[relationship]-(:User {userId:$followerId}) RETURN type(relationship)",
			map[string]interface{}{"followedId": followedId, "followerId": followerId})
		if err != nil {
			return nil, err
		}
		for records.Next() {
			record := records.Record()
			relationship := record.Values[0]
			results := relationship
			return results, nil
		}
		return "NO RELATIONSHIP", nil
	})
	if err != nil {
		return "", err
	}
	return relationship.(string), nil
}

func (store *FollowersStore) Follow(followerId string, followedId string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		currentTime := neo4j.LocalDateTime(time.Now())
		result, err := tx.Run(
			"MERGE (followed:User {userId: $followedId}) "+
				"ON CREATE SET followed.userId = $followedId "+
				"MERGE (follower:User {userId: $followerId}) "+
				"ON CREATE SET follower.userId = $followerId "+
				"MERGE (followed) <- [fol:FOLLOWING] - (follower) "+
				"ON CREATE SET fol.timeStarted = $timeStarted",
			map[string]interface{}{"followedId": followedId, "followerId": followerId,
				"timeStarted": currentTime})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return "Failed to follow: " + followerId + " -> " + followedId, err
	}
	return session.LastBookmark(), nil
}

func (store *FollowersStore) FollowRequest(followerId string, followedId string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		currentTime := neo4j.LocalDateTime(time.Now())
		result, err := tx.Run(
			"MERGE (followed:User {userId: $followedId}) "+
				"ON CREATE SET followed.userId = $followedId "+
				"MERGE (follower:User {userId: $followerId}) "+
				"ON CREATE SET follower.userId = $followerId "+
				"MERGE (followed) <- [req:REQUESTING_FOLLOW] - (follower) "+
				"ON CREATE SET req.timeSent = $timeSent",
			map[string]interface{}{"followedId": followedId, "followerId": followerId,
				"timeSent": currentTime})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return "Failed to create follow request: " + followerId + " -> " + followedId, err
	}
	return session.LastBookmark(), nil
}

func (store *FollowersStore) ConfirmFollow(followerId string, followedId string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		currentTime := neo4j.LocalDateTime(time.Now())
		result, err := tx.Run(
			"MATCH (followed:User {userId: $followedId}) "+
				"MATCH (follower:User {userId: $followerId}) "+
				"MATCH (followed) <- [followRequest:REQUESTING_FOLLOW] - (follower) "+
				"MERGE (followed) <- [fol:FOLLOWING] - (follower) "+
				"ON CREATE SET fol.timeStarted = $timeStarted "+
				"DELETE followRequest",
			map[string]interface{}{"followedId": followedId, "followerId": followerId,
				"timeStarted": currentTime})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return "Failed to create follow request: " + followerId + " -> " + followedId, err
	}
	return session.LastBookmark(), nil
}

func (store *FollowersStore) Unfollow(followerId string, followedId string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"MATCH (followed:User {userId: $followedId}) "+
				"MATCH (follower:User {userId: $followerId}) "+
				"MATCH (followed) <- [follow:FOLLOWING] - (follower) "+
				"DELETE follow",
			map[string]interface{}{"followedId": followedId, "followerId": followerId})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return "Failed to follow: " + followerId + " -> " + followedId, err
	}
	return session.LastBookmark(), nil
}

func (store *FollowersStore) RemoveFollowRequest(followerId string, followedId string) (string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: store.databaseName,
	})
	defer unsafeClose(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"MATCH (followed:User {userId: $followedId}) "+
				"MATCH (follower:User {userId: $followerId}) "+
				"MATCH (followed) <- [followRequest:REQUESTING_FOLLOW] - (follower) "+
				"DELETE followRequest",
			map[string]interface{}{"followedId": followedId, "followerId": followerId})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return "Failed to create follow request: " + followerId + " -> " + followedId, err
	}
	return session.LastBookmark(), nil
}

func unsafeClose(closeable io.Closer) {
	if err := closeable.Close(); err != nil {
		log.Fatal(fmt.Errorf("could not close resource: %w", err))
	}
}
