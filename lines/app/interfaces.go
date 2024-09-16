package app

import (
	linesHttp "lines/lines/http"
)

// App is the interface that all apps must implement.
type App interface {
	// Initialise is called to initialise the app.
	Initialise() error
	// RegisterHTTPRoutes is called to register the app's HTTP routes.
	RegisterHTTPRoutes(engine linesHttp.HttpEngine)
	// RegisterEventHandlers is called to register the app's event handlers.
	RegisterEventHandlers()
	// Close is called to close the app.
	Close()
}
