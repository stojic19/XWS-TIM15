package api

import (
	"github.com/stojic19/XWS-TIM15/Followers_microservice/application"
	"github.com/stojic19/XWS-TIM15/common/saga/blocking"
	saga "github.com/stojic19/XWS-TIM15/common/saga/messaging"
)

type BlockCommandHandler struct {
	service           *application.FollowersService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewBlockCommandHandler(orderService *application.FollowersService, publisher saga.Publisher, subscriber saga.Subscriber) (*BlockCommandHandler, error) {
	o := &BlockCommandHandler{
		service:           orderService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *BlockCommandHandler) handle(command *blocking.BlockCommand) {
	reply := blocking.BlockReply{
		BlockedId: command.BlockedId,
		BlockerId: command.BlockerId,
	}
	switch command.Type {
	case blocking.ConfirmBlock:
		//potvrditi blok
		handler.service.ConfirmBlock(command.BlockerId, command.BlockedId)
		reply.Type = blocking.UserBlocked
	case blocking.RevertBlock:
		//revertuj blok
		handler.service.RevertBlock(command.BlockerId, command.BlockedId)
		reply.Type = blocking.UserNotBlocked
	default:
		reply.Type = blocking.UnknownReply
	}

	if reply.Type != blocking.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
