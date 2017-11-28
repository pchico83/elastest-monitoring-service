// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"../restapi/operations"
	"../restapi/operations/announcements"
	"../restapi/operations/flush"
	"../restapi/operations/health"
	"../restapi/operations/monitoring_machine"
	"../restapi/operations/publishers"
	"../restapi/operations/subscribers"
	"../implementation"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../swagger.yaml

func configureFlags(api *operations.MonitoringAsAServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MonitoringAsAServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.XMLProducer = runtime.XMLProducer()

	api.MonitoringMachineDeleteMoMHandler = monitoring_machine.DeleteMoMHandlerFunc(func(params monitoring_machine.DeleteMoMParams) middleware.Responder {
		return middleware.NotImplemented("operation monitoring_machine.DeleteMoM has not yet been implemented")
	})
	api.FlushFlushHandler = flush.FlushHandlerFunc(func(params flush.FlushParams) middleware.Responder {
		return middleware.NotImplemented("operation flush.Flush has not yet been implemented")
	})
	api.MonitoringMachineGetDeployedMoMsHandler = monitoring_machine.GetDeployedMoMsHandlerFunc(func(params monitoring_machine.GetDeployedMoMsParams) middleware.Responder {
		return middleware.NotImplemented("operation monitoring_machine.GetDeployedMoMs has not yet been implemented")
	})
	api.HealthGetEnvironmentHandler = health.GetEnvironmentHandlerFunc(func(params health.GetEnvironmentParams) middleware.Responder {
		return implementation.GetEnvironment()
	})
	api.HealthGetHealthHandler = health.GetHealthHandlerFunc(func(params health.GetHealthParams) middleware.Responder {
		return implementation.GetHealth()
	})
	api.MonitoringMachineGetMoMByIDHandler = monitoring_machine.GetMoMByIDHandlerFunc(func(params monitoring_machine.GetMoMByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation monitoring_machine.GetMoMByID has not yet been implemented")
	})
	api.MonitoringMachinePostMoMHandler = monitoring_machine.PostMoMHandlerFunc(func(params monitoring_machine.PostMoMParams) middleware.Responder {
		return implementation.PostMOM(params)
	})
	api.PublishersPublishEventsHandler = publishers.PublishEventsHandlerFunc(func(params publishers.PublishEventsParams) middleware.Responder {
		return middleware.NotImplemented("operation publishers.PublishEvents has not yet been implemented")
	})
	api.AnnouncementsRegisterHandler = announcements.RegisterHandlerFunc(func(params announcements.RegisterParams) middleware.Responder {
		return middleware.NotImplemented("operation announcements.Register has not yet been implemented")
	})
	api.SubscribersSubscribeElastestEndpointsHandler = subscribers.SubscribeElastestEndpointsHandlerFunc(func(params subscribers.SubscribeElastestEndpointsParams) middleware.Responder {
		return middleware.NotImplemented("operation subscribers.SubscribeElastestEndpoints has not yet been implemented")
	})
	api.SubscribersSubscribeElasticSearchHandler = subscribers.SubscribeElasticSearchHandlerFunc(func(params subscribers.SubscribeElasticSearchParams) middleware.Responder {
		return implementation.SubscribeES(params)
	})
	api.SubscribersSubscribeRabbitMQHandler = subscribers.SubscribeRabbitMQHandlerFunc(func(params subscribers.SubscribeRabbitMQParams) middleware.Responder {
		return implementation.SubscribeRMQ(params)
	})
	api.AnnouncementsUnregisterHandler = announcements.UnregisterHandlerFunc(func(params announcements.UnregisterParams) middleware.Responder {
		return middleware.NotImplemented("operation announcements.Unregister has not yet been implemented")
	})
	api.SubscribersUnsubscribeHandler = subscribers.UnsubscribeHandlerFunc(func(params subscribers.UnsubscribeParams) middleware.Responder {
		return middleware.NotImplemented("operation subscribers.Unsubscribe has not yet been implemented")
	})
	api.MonitoringMachineUpdateMoMHandler = monitoring_machine.UpdateMoMHandlerFunc(func(params monitoring_machine.UpdateMoMParams) middleware.Responder {
		return middleware.NotImplemented("operation monitoring_machine.UpdateMoM has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
        go implementation.OpenAndLoop()
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}