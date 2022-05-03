package main

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/domain"
)

func main() {
	//config := cfg.NewConfig()
	//server := startup.NewServer(config)
	//server.Start()
	username := "neo4j"
	password := "neo4j"
	database := "neo4j"
	//url := "http://localhost:7474"
	url := "bolt://localhost:7687"
	driver, _ := neo4j.NewDriver(url, neo4j.BasicAuth(username, password, ""))
	session := driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: database,
	})
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

	print("GOTOVO")
	print(followers)
	print(err)
}
