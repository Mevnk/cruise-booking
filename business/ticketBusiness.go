package business

import (
	"context"
	"github.com/google/uuid"
	"pis/dao"
	"pis/domain"
	"pis/pkg/customErrors"
)

type TicketLogic interface {
	CreateTicket(ctx context.Context, params TicketParams) (ticketId uuid.UUID, err error)
	GetTicketData(ctx context.Context, id uuid.UUID) (ticket *domain.Ticket, err error)
	ExpendTicket(ctx context.Context, id uuid.UUID) (err error)
}

type ticketLogic struct {
	ticketDAO dao.TicketsDAO
	cruiseDAO dao.CruisesDAO
}

func (tl ticketLogic) CreateTicket(ctx context.Context, params TicketParams) (ticketId uuid.UUID, err error) {
	_, err = tl.cruiseDAO.GetCruiseByID(ctx, params.CruiseID)
	if err != nil {
		return uuid.Nil, customErrors.NewCustomError(customErrors.Cruise, customErrors.NotFound)
	}

	ticket := domain.NewTicket(params.CruiseID, params.PassengerID, params.PassengerClass, params.Bonuses)

	if err = tl.ticketDAO.CreateTicket(ctx, ticket); err != nil {
		return uuid.Nil, customErrors.NewCustomError(customErrors.Ticket, customErrors.Creation)
	}

	return ticket.Id, err
}

func (tl ticketLogic) GetTicketData(ctx context.Context, id uuid.UUID) (ticket *domain.Ticket, err error) {
	ticket, err = tl.ticketDAO.GetTicketByID(ctx, id)
	if err != nil {
		return nil, customErrors.NewCustomError(customErrors.Ticket, customErrors.NotFound)
	}

	return ticket, err
}

func (tl ticketLogic) ExpendTicket(ctx context.Context, id uuid.UUID) (err error) {
	err = tl.ticketDAO.DeleteTicket(ctx, id)
	if err != nil {
		return customErrors.NewCustomError(customErrors.Ticket, customErrors.Deletion)
	}
	return nil
}

type TicketParams struct {
	CruiseID       uuid.UUID
	PassengerID    uuid.UUID
	PassengerClass string
	Bonuses        string
}
