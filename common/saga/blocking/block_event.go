package blocking

type BlockCommandType int8

const (
	BlockChat BlockCommandType = iota
	ConfirmBlock
	RevertBlock
	UnknownCommand
)

type BlockCommand struct {
	BlockerId string
	BlockedId string
	Type      BlockCommandType
}
type BlockReplyType int8

const (
	ChatBlocked BlockReplyType = iota
	ChatNotBlocked
	UserBlocked
	UserNotBlocked
	UnknownReply
)

type BlockReply struct {
	BlockerId string
	BlockedId string
	Type      BlockReplyType
}
