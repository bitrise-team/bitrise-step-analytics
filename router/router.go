package router

import (
	"github.com/bitrise-io/api-utils/httpresponse"
	"github.com/bitrise-team/bitrise-step-analytics/configs"
	"github.com/bitrise-team/bitrise-step-analytics/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/rs/cors"
)

// New ...
func New(config configs.ConfigModel) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	commonMiddleware := alice.New(
		cors.AllowAll().Handler,
	)

	r.Handle("/", commonMiddleware.ThenFunc(service.RootHandler))
	r.Handle("/log-analytics", commonMiddleware.Then(
		httpresponse.InternalErrHandlerFuncAdapter(service.AnalyticsLogHandler))).Methods("POST")

	return r
}
