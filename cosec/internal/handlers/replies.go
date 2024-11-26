package handlers

import (
	"EDA_GO/cosec/internal"
	"EDA_GO/cosec/internal/models"
	"EDA_GO/internal/am"
	"EDA_GO/internal/registry"
	"EDA_GO/internal/sec"
)

func NewReplyHandlers(reg registry.Registry, orchestrator sec.Orchestrator[*models.CreateOrderData], mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewReplyHandler(reg, orchestrator, mws...)
}

func RegisterReplyHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(internal.CreateOrderReplyChannel, handlers, am.GroupName("cosec-replies"))
	return err
}
