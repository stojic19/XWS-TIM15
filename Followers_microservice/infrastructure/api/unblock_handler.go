package api

import (
	"github.com/stojic19/XWS-TIM15/Followers_microservice/application"
	"github.com/stojic19/XWS-TIM15/common/saga/blocking"
	saga "github.com/stojic19/XWS-TIM15/common/saga/messaging"
)

type UnblockCommandHandler struct {
	service           *application.FollowersService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUnblockCommandHandler(orderService *application.FollowersService, publisher saga.Publisher, subscriber saga.Subscriber) (*UnblockCommandHandler, error) {
	o := &UnblockCommandHandler{
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

func (handler *UnblockCommandHandler) handle(command *blocking.UnblockCommand) {
	reply := blocking.UnblockReply{
		BlockedId: command.BlockedId,
		BlockerId: command.BlockerId,
	}
	switch command.Type {
	case blocking.ConfirmUnblock:
		//potvrditi blok
		handler.service.ConfirmUnblock(command.BlockerId, command.BlockedId)
		reply.Type = blocking.UserUnblocked
	case blocking.RevertUnblock:
		//revertuj blok
		handler.service.RevertUnblock(command.BlockerId, command.BlockedId)
		reply.Type = blocking.UserNotUnblocked
	default:
		reply.Type = blocking.UnknownUnblockReply
	}

	if reply.Type != blocking.UnknownUnblockReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
