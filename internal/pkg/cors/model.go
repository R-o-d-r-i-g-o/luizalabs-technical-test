package cors

import (
	"net/http"
	"time"
)

var (
	// maxAge defines the maximum duration (in seconds) that the results
	// of a preflight request can be cached by the browser.
	maxAge = int(3 * time.Hour)

	// allowedOrigins is a list of origins that are permitted to access the resources.
	// The wildcard "*" allows requests from any origin.
	allowedOrigins = []string{"*"}

	// allowedMethods specifies the HTTP methods that are allowed for cross-origin requests.
	allowedMethods = []string{
		http.MethodGet,
		http.MethodPut,
		http.MethodPost,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
	}

	// allowedHeaders lists the headers that can be included in cross-origin requests.
	allowedHeaders = []string{
		"Origin",
		"Content-Type",
		"Accept",
		"Authorization",
		"Cache-Control",
		"X-Requested-With",
		"User-Agent",
		"X-Cache-Control",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Methods",
		"Access-Control-Request-Method",
		"Access-Control-Request-Headers",
	}

	// exposedHeaders lists the headers that are exposed to the browser and can be accessed in JavaScript.
	exposedHeaders = []string{
		"Content-Length",
	}
)
