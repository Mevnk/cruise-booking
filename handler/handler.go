package handler

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"pis/domain"
	"pis/handler/dto"
)

type handler struct {
	userHandler      *UserHandler
	cruiseHandler    *CruiseHandler
	excursionHandler *ExcursionHandler
	ticketHandler    *TicketHandler
}

func NewHandler(userHandler *UserHandler,
	cruiseHandler *CruiseHandler,
	excursionHandler *ExcursionHandler,
	ticketHandler *TicketHandler) *handler {
	return &handler{
		userHandler:      userHandler,
		cruiseHandler:    cruiseHandler,
		excursionHandler: excursionHandler,
		ticketHandler:    ticketHandler,
	}
}

func (h *handler) InitRoutes(app *fiber.App) {
	// auth handlers
	app.Post("/sign-in", h.userHandler.SignIn)
	app.Post("/sign-up", h.userHandler.SignUp)

	// cruise handlers
	app.Post("/cruise", h.isManager, h.cruiseHandler.AddCruise)
	app.Get("/cruise/:id", h.cruiseHandler.GetCruiseById)
	app.Delete("/cruise/:id", h.isManager, h.cruiseHandler.DeleteCruise)

	// excursion handlers
	app.Post("/excursion", h.isManager, h.excursionHandler.AddExcursion)
	app.Delete("/excursion/:id", h.isManager, h.excursionHandler.DeleteExcursion)

	// ticket handlers
	app.Post("/ticket", h.isManager, h.ticketHandler.AddTicket)
	app.Get("/ticket/:id", h.ticketHandler.GetTicketById)
	app.Delete("/ticket/:id", h.isManager, h.ticketHandler.DeleteTicket)

	// swagger handler
	app.Get("/swagger/*", swagger.HandlerDefault)
}

func (h *handler) isManager(ctx *fiber.Ctx) error {
	role := ctx.Cookies("role")
	if role == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ErrorDTO{Message: "Unauthorized"})
	}
	if role != domain.MANAGER && role != domain.ADMIN {
		return ctx.Status(fiber.StatusForbidden).JSON(dto.ErrorDTO{Message: "Access denied"})
	}
	return ctx.Next()
}

func (h *handler) isAdmin(ctx *fiber.Ctx) error {
	role := ctx.Cookies("role")
	if role == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ErrorDTO{Message: "Unauthorized"})
	}
	if role != domain.ADMIN {
		return ctx.Status(fiber.StatusForbidden).JSON(dto.ErrorDTO{Message: "Access denied"})
	}
	return ctx.Next()
}
