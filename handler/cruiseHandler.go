package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pis/business"
	"pis/handler/dto"
)

type CruiseHandler struct {
	validator   *validator.Validate
	cruiseLogic business.CruiseLogic
}

func NewCruiseHandler(cruiseLogic business.CruiseLogic, validator *validator.Validate) *CruiseHandler {
	return &CruiseHandler{cruiseLogic: cruiseLogic, validator: validator}
}

// AddCruise
//
// @Summary Add a new cruise
// @Tags Cruise
// @Accept json
// @Param input body dto.AddCruiseDTO true "New cruise data"
// @Success 201 {object} dto.ReturnIdDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /cruise [post]
func (h *CruiseHandler) AddCruise(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	cruiseDto := new(dto.AddCruiseDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := ctx.BodyParser(cruiseDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(cruiseDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	id, err := h.cruiseLogic.CreateCruise(ctx.Context(), toAddCruiseParams(cruiseDto))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.ReturnIdDTO{Id: id})
}

// GetCruiseById
//
// @Summary Get a cruise by ID
// @Tags cruise
// @Accept json
// @Param id path string true "Cruise ID"
// @Success 200 {object} dto.CruiseDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /cruise/{id} [get]
func (h *CruiseHandler) GetCruiseById(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	cruise, err := h.cruiseLogic.GetCruiseData(ctx.Context(), uuid.MustParse(idDto.Id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(mapDtoCruise(cruise))
}

// DeleteCruise
//
// @Summary Delete a Cruise from a hotel
// @Tags Cruise
// @Accept json
// @Param id path string true "Cruise ID"
// @Success 200 {object} dto.SuccessDTO
// @Failure 400 {object} dto.ErrorDTO
// @Failure 500 {object} dto.ErrorDTO
// @Router /cruise/{id} [delete]
func (h *CruiseHandler) DeleteCruise(ctx *fiber.Ctx) error {
	idDto := new(dto.GetByIdDTO)
	if err := ctx.ParamsParser(idDto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.validator.Struct(idDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	if err := h.cruiseLogic.DeleteCruise(ctx.Context(), uuid.MustParse(idDto.Id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDTO{Message: err.Error()})
	}
	return ctx.JSON(dto.SuccessDTO{Message: "Deleted"})
}
