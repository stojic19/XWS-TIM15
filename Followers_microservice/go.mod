module github.com/stojic19/XWS-TIM15/Followers_microservice

go 1.18

replace github.com/stojic19/XWS-TIM15/common => ../common

require (
	github.com/neo4j/neo4j-go-driver/v4 v4.4.2
	github.com/stojic19/XWS-TIM15/common v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.46.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0 // indirect
	github.com/nats-io/nats.go v1.16.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/tamararankovic/microservices_demo/common v0.0.0-20220326142530-97bfd7810e53 // indirect
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220429170224-98d788798c3e // indirect
)
