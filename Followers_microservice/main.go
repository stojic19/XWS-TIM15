package main

import (
	"github.com/stojic19/XWS-TIM15/Followers_microservice/startup"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
	/*//username := "neo4j"
	//password := "neo4j"
	//database := "neo4j"
	//url := "http://localhost:7474"
	url := "neo4j://followers_db:7687"
	driver, _ := neo4j.NewDriver(url, neo4j.BasicAuth("neo4j", "neo4j", ""))
	defer driver.Close()
	/*session := driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close()
	followers, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("MATCH (user:User) return user.username", map[string]interface{}{})
		//records, err := tx.Run(
		//	"MATCH (:User {username:$username})<-[:FOLLOWING]-(follower) RETURN follower",
		//	map[string]interface{}{"username": username})
		if err != nil {
			return nil, err
		}
		results := []*domain.User{}
		for records.Next() {
			record := records.Record()
			username, _ := record.Get("user.username")
			user := domain.User{
				Username: username.(string),
			}
			results = append(results, &user)
		}
		return results, nil
	})*/
	/*session := driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()
	followers, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("CREATE (n:User { username: $username}) RETURN n.username", map[string]interface{}{
			"username": "ralo",
		})
		//records, err := tx.Run(
		//	"MATCH (:User {username:$username})<-[:FOLLOWING]-(follower) RETURN follower",
		//	map[string]interface{}{"username": username})
		if err != nil {
			return nil, err
		}
		results := []*domain.User{}
		for records.Next() {
			record := records.Record()
			username, _ := record.Get("n.username")
			user := domain.User{
				Username: username.(string),
			}
			results = append(results, &user)
		}
		return results, nil
	})
	if err != nil {
		print("PANIKA")
		panic(err)
	}
	print("GOTOVO!!")
	print((followers.([]*domain.User))[0].Username)
	print(len(followers.([]*domain.User)))
	print(err)*/
}
