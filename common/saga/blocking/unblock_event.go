package blocking

type UnblockCommandType int8

const (
	UnblockChat UnblockCommandType = iota
	ConfirmUnblock
	RevertUnblock
	UnknownUnblockCommand
)

type UnblockCommand struct {
	BlockerId string
	BlockedId string
	Type      UnblockCommandType
}

type UnblockReplyType int8

const (
	ChatUnblocked UnblockReplyType = iota
	ChatNotUnblocked
	UserUnblocked
	UserNotUnblocked
	UnknownUnblockReply
)

type UnblockReply struct {
	BlockerId string
	BlockedId string
	Type      UnblockReplyType
}
