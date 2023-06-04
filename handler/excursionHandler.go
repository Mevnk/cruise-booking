package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pis/business"
	"pis/handler/dto"
)

type ExcursionHandler struct {
	validator      *validator.Validate
	excursionLogic business.ExcursionLogic
}

func NewExcursionHandler(excursionLogic business.ExcursionLogic, validator *validator.Validate) *ExcursionHandler {
	return &ExcursionHandler{excursionLogic: excursionLogic, validator: validator}
}

func (h *ExcursionHandler) AddExcursion(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	excursionDTO := new(dto.AddExcursionDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := ctx.BodyParser(excursionDTO); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(excursionDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	id, err := h.excursionLogic.CreateExcursion(ctx.Context(), toAddExcursionParams(excursionDTO))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.ReturnIdDTO{Id: id})
}

func (h *ExcursionHandler) DeleteExcursion(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.excursionLogic.DeleteExcursion(ctx.Context(), uuid.MustParse(idDto.Id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(dto.SuccessDTO{Message: "Deleted"})
}
