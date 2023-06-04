package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pis/business"
	"pis/handler/dto"
)

type TicketHandler struct {
	validator   *validator.Validate
	ticketLogic business.TicketLogic
}

func NewTicketHandler(ticketLogic business.TicketLogic, validator *validator.Validate) *TicketHandler {
	return &TicketHandler{ticketLogic: ticketLogic, validator: validator}
}

// AddTicket
//
// @Summary Add a new cruise
// @Tags Ticket
// @Accept json
// @Param input body dto.AddTicketDTO true "New Ticket data"
// @Success 201 {object} dto.ReturnIdDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /cruise [post]
func (h *TicketHandler) AddTicket(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	ticketDTO := new(dto.AddTicketDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := ctx.BodyParser(ticketDTO); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(ticketDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	id, err := h.ticketLogic.CreateTicket(ctx.Context(), toAddTicketParams(ticketDTO))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.ReturnIdDTO{Id: id})
}

// GetTicketById
//
// @Summary Get a Ticket by ID
// @Tags Ticket
// @Accept json
// @Param id path string true "Ticket ID"
// @Success 200 {object} dto.TicketDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /cruise/{id} [get]
func (h *TicketHandler) GetTicketById(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	ticket, err := h.ticketLogic.GetTicketData(ctx.Context(), uuid.MustParse(idDto.Id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(mapDtoTicket(ticket))
}

// DeleteTicket
//
// @Summary Delete a Ticket from a hotel
// @Tags Cruise
// @Accept json
// @Param id path string true "Ticket ID"
// @Success 200 {object} dto.SuccessDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /cruise/{id} [delete]
func (h *TicketHandler) DeleteTicket(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.ticketLogic.ExpendTicket(ctx.Context(), uuid.MustParse(idDto.Id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(dto.SuccessDTO{Message: "Deleted"})
}
