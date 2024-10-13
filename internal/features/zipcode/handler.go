package zipcode

import (
	"luizalabs-technical-test/internal/pkg/formatter"
	"luizalabs-technical-test/internal/pkg/validator"
	"luizalabs-technical-test/pkg/middleware"
	"luizalabs-technical-test/pkg/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

// swagGetAddressByZipCodeResponse is used to work around Swagger's lack of support for Go generics.
type swagGetAddressByZipCodeResponse = server.APIResponse[GetAddressByZipCodeResponse]

// HandlerImp defines the interface for handling server operations.
// It embeds the server.HandlerImp interface, allowing for extended functionality and custom implementations.
type HandlerImp interface {
	server.HandlerImp
}

// handler struct holds a reference to the service layer.
type handler struct {
	svc        ServiceImp
	cacheLayer middleware.Middleware
}

// NewHandler creates and returns a new handler instance with the injected service.
func NewHandler(svc ServiceImp, cacheMiddleware middleware.Middleware) HandlerImp {
	return &handler{
		svc,
		cacheMiddleware,
	}
}

// Register sets up the route for retrieving ZipCode information.
func (h *handler) Register(r *gin.RouterGroup) {
	g := r.Group("/address")
	g.GET("/:zip-code", h.cacheLayer.Middleware(), h.getAddressByZipCode)
}

// getAddressByZipCode handles the request to retrieve CEP information.
//
//	@Summary		Retrieve CEP information by ZIP code
//	@Description	Get address details using a provided ZIP code. Returns a structured response with address data or error information.
//	@Tags			Address
//	@Accept			json
//	@Produce		json
//	@Param			zip-code	path		string	true	"ZIP Code"
//	@Success		200			{object}	swagGetAddressByZipCodeResponse
//	@Success		302			{object}	swagGetAddressByZipCodeResponse	"Cached value retrieved"
//	@Failure		400			{object}	server.APIErrorResponse			"Invalid ZIP code format"
//	@Failure		404			{object}	server.APIErrorResponse			"ZIP code not found"
//	@Router			/v1/address/{zip-code} [get]
func (h *handler) getAddressByZipCode(c *gin.Context) {
	zipCode := c.Param("zip-code")
	zipCode = formatter.StripNonNumericCharacters(zipCode)

	if isAccepted := validator.ValidateZipCode(zipCode); !isAccepted {
		c.JSON(http.StatusBadRequest, server.APIErrorResponse{Error: ErrZipCodeNotFormatted.Error()})
		return
	}

	for {
		response, err := h.svc.GetAddressByZipCode(zipCode)
		if response != nil {
			c.JSON(http.StatusOK, swagGetAddressByZipCodeResponse{Data: *response})
			break
		}

		zipCode = formatter.AdjustLastNonZeroDigit(zipCode)
		if zipCode == validator.EmptyZipCodeValue {
			// TODO: log Error here.
			c.JSON(http.StatusNotFound, server.APIErrorResponse{Error: err.Error()})
			break
		}
	}
}
