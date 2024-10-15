package auth

import (
	"luizalabs-technical-test/pkg/errors"
	"luizalabs-technical-test/pkg/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

// swagAuthenticateUserResponse is used to work around Swagger's lack of support for Go generics.
type swagAuthenticateUserResponse = server.APIResponse[AuthenticateUserResponse]

// HandlerImp defines the interface for handling server operations.
// It embeds the server.HandlerImp interface, allowing for extended functionality and custom implementations.
type HandlerImp interface {
	server.HandlerImp
}

// handler struct is an empty implementation of the Handler interface.
type handler struct {
	service ServiceImp
}

// NewHandler creates and returns a new handler instance.
func NewHandler(service ServiceImp) HandlerImp {
	return &handler{service}
}

// Register sets up the route for retrieving auth information.
func (h *handler) Register(r *gin.RouterGroup) {
	g := r.Group("/auth")
	g.POST("/register", h.postRegister)
	g.POST("/login", h.postLogin)
}

// postRegister registers a new user.
//
//	@Summary		Register a new user
//	@Description	Registers a new user with the provided information.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body	PostRegisterPayload	true	"User registration data"
//	@Success		201		"User successfully registered"
//	@Failure		400		{object}	server.APIErrorResponse	"Bad request"
//	@Failure		500		{object}	server.APIErrorResponse	"Internal server error"
//	@Router			/v1/auth/register [post]
func (h *handler) postRegister(c *gin.Context) {
	var payload PostRegisterPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, server.APIErrorResponse{
			Error: ErrInvalidCredentials.WithErr(err).Error(),
			Code:  ErrInvalidCredentials.Code,
		})
		return
	}

	if err := h.service.RegisterUser(payload.ToUserEntity()); err != nil {
		c.JSON(http.StatusInternalServerError, server.APIErrorResponse{
			Error: err.Error(),
			Code:  err.(errors.ErrorImp).CodeStr(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

// postLogin authenticates the user and returns a JWT token.
//
//	@Summary		Authenticate user and return a JWT token
//	@Description	Authenticates the user with the provided credentials and returns a JWT token.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		PostLoginPayload				true	"User credentials"
//	@Success		202		{object}	swagAuthenticateUserResponse	"Token generated successfully"
//	@Failure		400		{object}	server.APIErrorResponse			"Bad request"
//	@Failure		401		{object}	server.APIErrorResponse			"Unauthorized"
//	@Router			/v1/auth/login [post]
func (h *handler) postLogin(c *gin.Context) {
	var payload PostLoginPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, server.APIErrorResponse{
			Error: ErrInvalidCredentials.WithErr(err).Error(),
			Code:  ErrInvalidCredentials.Code,
		})
		return
	}

	jwt, err := h.service.AuthenticateUser(payload.ToPostLoginPayloadToInput())
	if err != nil {
		c.JSON(http.StatusUnauthorized, server.APIErrorResponse{
			Error: err.Error(),
			Code:  err.(errors.ErrorImp).CodeStr(),
		})
		return
	}

	c.JSON(http.StatusAccepted, swagAuthenticateUserResponse{
		Data: AuthenticateUserResponse{jwt},
	})
}
