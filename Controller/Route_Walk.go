package Controller

import (
	"github.com/gorilla/mux"
)

func SetRoutesWalk(routes *mux.Router) {
	routes.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		// pathTemplate, err := route.GetPathTemplate()
		// if err == nil {
		// 	logger.Rosa("Rota:", pathTemplate)
		// }
		// pathRegexp, err := route.GetPathRegexp()
		// if err == nil {
		// 	logger.Rosa("Endereço regexp:", pathRegexp)
		// }
		// queriesTemplates, err := route.GetQueriesTemplates()
		// if err == nil {
		// 	logger.Rosa("Consulta templates:", strings.Join(queriesTemplates, ","))
		// }
		// queriesRegexps, err := route.GetQueriesRegexp()
		// if err == nil {
		// 	logger.Rosa("Consulta regexps:", strings.Join(queriesRegexps, ","))
		// }
		// methods, err := route.GetMethods()
		// if err == nil {
		// 	logger.Rosa("Meodo:", strings.Join(methods, ","))
		// }
		return nil
	})
}
