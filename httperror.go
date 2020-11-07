package gower

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/situmorangbastian/gower/models"
)

type errResponse struct {
	Message string `json:"message"`
}

// ErrMiddleware returns custom middleware for Fiber that generate HTTP error response
// with HTTP status code.
func ErrMiddleware(ctx *fiber.Ctx, err error) error {
	errResponse := errResponse{
		Message: errors.Cause(err).Error(),
	}

	// Retreive the custom response if it's an fiber.*Error
	if e, ok := errors.Cause(err).(*fiber.Error); ok {
		return ctx.Status(e.Code).JSON(errResponse)
	}

	// Check error based on error type
	switch errors.Cause(err).(type) {
	case models.ConstraintError:
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	case models.NotFoundError:
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	case models.ConflictError:
		return ctx.Status(http.StatusConflict).JSON(errResponse)
	default:
		log.Error(err)
		errResponse.Message = "Internal Server Error"
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
}
