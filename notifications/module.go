package notifications

import (
	"context"

	"EDA_GO/customers/customerspb"
	"EDA_GO/internal/am"
	"EDA_GO/internal/amotel"
	"EDA_GO/internal/amprom"
	"EDA_GO/internal/jetstream"
	pg "EDA_GO/internal/postgres"
	"EDA_GO/internal/postgresotel"
	"EDA_GO/internal/registry"
	"EDA_GO/internal/system"
	"EDA_GO/internal/tm"
	"EDA_GO/notifications/internal/application"
	"EDA_GO/notifications/internal/constants"
	"EDA_GO/notifications/internal/grpc"
	"EDA_GO/notifications/internal/handlers"
	"EDA_GO/notifications/internal/postgres"
	"EDA_GO/ordering/orderingpb"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	inboxStore := pg.NewInboxStore(constants.InboxTableName, svc.DB())
	messageSubscriber := am.NewMessageSubscriber(
		jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger()),
		amotel.OtelMessageContextExtractor(),
		amprom.ReceivedMessagesCounter(constants.ServiceName),
	)
	customers := postgres.NewCustomerCacheRepository(
		constants.CustomersCacheTableName,
		postgresotel.Trace(svc.DB()),
		grpc.NewCustomerRepository(svc.Config().Rpc.Service(constants.CustomersServiceName)),
	)

	// setup application
	app := application.New(customers)
	integrationEventHandlers := handlers.NewIntegrationEventHandlers(
		reg, app, customers,
		tm.InboxHandler(inboxStore),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, svc.RPC()); err != nil {
		return err
	}
	if err = handlers.RegisterIntegrationEventHandlers(messageSubscriber, integrationEventHandlers); err != nil {
		return err
	}

	return nil
}
