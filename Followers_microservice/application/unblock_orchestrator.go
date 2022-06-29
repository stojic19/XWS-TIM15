package application

import (
	"github.com/stojic19/XWS-TIM15/common/saga/blocking"
	saga "github.com/stojic19/XWS-TIM15/common/saga/messaging"
)

type UnblockOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewUnblockOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*UnblockOrchestrator, error) {
	o := &UnblockOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *UnblockOrchestrator) Start(blockerId string, blockedId string) error {
	event := blocking.UnblockCommand{
		BlockerId: blockerId,
		BlockedId: blockedId,
		Type:      blocking.UnblockChat,
	}
	return o.commandPublisher.Publish(event)
}

func (o *UnblockOrchestrator) handle(reply *blocking.UnblockReply) {
	command := blocking.UnblockCommand{
		BlockerId: reply.BlockerId,
		BlockedId: reply.BlockedId,
	}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != blocking.UnknownUnblockCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *UnblockOrchestrator) nextCommandType(reply blocking.UnblockReplyType) blocking.UnblockCommandType {
	switch reply {
	case blocking.ChatUnblocked:
		return blocking.ConfirmUnblock
	case blocking.ChatNotUnblocked:
		return blocking.RevertUnblock
	default:
		return blocking.UnknownUnblockCommand
	}
}
