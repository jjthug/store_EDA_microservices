package handlers

import (
	"context"
	"database/sql"

	"EDA_GO/internal/am"
	"EDA_GO/internal/di"
	"EDA_GO/search/internal/constants"
)

func RegisterIntegrationEventHandlersTx(container di.Container) (err error) {
	rawMsgHandler := am.MessageHandlerFunc(func(ctx context.Context, msg am.IncomingMessage) (err error) {
		ctx = container.Scoped(ctx)
		defer func(tx *sql.Tx) {
			if p := recover(); p != nil {
				_ = tx.Rollback()
				panic(p)
			} else if err != nil {
				_ = tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

		return di.Get(ctx, constants.IntegrationEventHandlersKey).(am.MessageHandler).HandleMessage(ctx, msg)
	})

	subscriber := container.Get(constants.MessageSubscriberKey).(am.MessageSubscriber)

	return RegisterIntegrationEventHandlers(subscriber, rawMsgHandler)
}
