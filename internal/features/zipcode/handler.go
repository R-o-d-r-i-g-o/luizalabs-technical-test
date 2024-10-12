package zipcode

import (
	"luizalabs-technical-test/internal/common"
	"luizalabs-technical-test/pkg/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerImp defines the interface for handling server operations.
// It embeds the server.HandlerImp interface, allowing for extended functionality and custom implementations.
type HandlerImp interface {
	server.HandlerImp
}

// handler struct holds a reference to the service layer.
type handler struct {
	svc ServiceImp
}

// NewHandler creates and returns a new handler instance with the injected service.
func NewHandler(svc ServiceImp) HandlerImp {
	return &handler{svc}
}

// Register sets up the route for retrieving CEP information.
func (h *handler) Register(r *gin.RouterGroup) {
	g := r.Group("/zip-code")
	g.GET("/:zip-code", h.getAddressByZipCode)
}

// getAddressByZipCode handles the request to retrieve CEP information.
func (h *handler) getAddressByZipCode(c *gin.Context) {
	zipCode := c.Param("zip-code")
	zipCode = common.StripNonNumericCharacters(zipCode)

	if isAccepted := common.ValidateZipCode(zipCode); !isAccepted {
		c.JSON(http.StatusBadRequest, server.APIErrorResponse{Error: ErrZipCodeNotFormatted.Error()})
		return
	}

	for {
		response, err := h.svc.GetAddressByZipCode(zipCode)
		if response != nil {
			c.JSON(http.StatusOK, response)
			break
		}

		zipCode = common.AdjustLastNonZeroDigit(zipCode)
		if zipCode == common.EmptyZipCodeValue {
			c.JSON(http.StatusNotFound, server.APIErrorResponse{Error: err.Error()})
			// TODO: log Error here.
			break
		}
	}
}
