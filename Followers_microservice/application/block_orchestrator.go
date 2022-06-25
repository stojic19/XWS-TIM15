package application

import (
	"github.com/stojic19/XWS-TIM15/common/saga/blocking"
	saga "github.com/stojic19/XWS-TIM15/common/saga/messaging"
)

type BlockOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewBlockOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*BlockOrchestrator, error) {
	o := &BlockOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *BlockOrchestrator) Start(blockerId string, blockedId string) error {
	event := blocking.BlockCommand{
		BlockerId: blockerId,
		BlockedId: blockedId,
		Type:      blocking.BlockChat,
	}
	return o.commandPublisher.Publish(event)
}

func (o *BlockOrchestrator) handle(reply *blocking.BlockReply) {
	command := blocking.BlockCommand{}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != blocking.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *BlockOrchestrator) nextCommandType(reply blocking.BlockReplyType) blocking.BlockCommandType {
	switch reply {
	case blocking.ChatBlocked:
		return blocking.ConfirmBlock
	case blocking.ChatNotBlocked:
		return blocking.RevertBlock
	default:
		return blocking.UnknownCommand
	}
}
