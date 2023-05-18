package dao

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"pis/domain"
)

type TicketsDAO interface {
	CreateTicket(ctx context.Context, ticket *domain.Ticket) error
	UpdateTicket(ctx context.Context, ticket *domain.Ticket) error
	DeleteTicket(ctx context.Context, ticketID uuid.UUID) error
	GetTicketByID(ctx context.Context, ticketID uuid.UUID) (*domain.Ticket, error)
}

type MySQLTicketsDAO struct {
	db *sql.DB
}

func NewMySQLTicketsDAO(db *sql.DB) *MySQLTicketsDAO {
	return &MySQLTicketsDAO{
		db: db,
	}
}

func (dao *MySQLTicketsDAO) CreateTicket(ctx context.Context, ticket *domain.Ticket) error {
	query := "INSERT INTO Tickets (id, cruise_id, passenger_id, passenger_class, bonuses) VALUES (?, ?, ?, ?, ?)"
	_, err := dao.db.ExecContext(ctx, query, ticket.Id, ticket.CruiseID, ticket.PassengerID, ticket.PassengerClass, ticket.Bonuses)
	return err
}

func (dao *MySQLTicketsDAO) UpdateTicket(ctx context.Context, ticket *domain.Ticket) error {
	query := "UPDATE Tickets SET cruise_id = ?, passenger_id = ?, passenger_class = ?, bonuses = ? WHERE id = ?"
	_, err := dao.db.ExecContext(ctx, query, ticket.CruiseID, ticket.PassengerID, ticket.PassengerClass, ticket.Bonuses, ticket.Id)
	return err
}

func (dao *MySQLTicketsDAO) DeleteTicket(ctx context.Context, ticketID uuid.UUID) error {
	query := "DELETE FROM Tickets WHERE id = ?"
	_, err := dao.db.ExecContext(ctx, query, ticketID)
	return err
}

func (dao *MySQLTicketsDAO) GetTicketByID(ctx context.Context, ticketID uuid.UUID) (*domain.Ticket, error) {
	query := "SELECT id, cruise_id, passenger_id, passenger_class, bonuses FROM Tickets WHERE id = ?"
	row := dao.db.QueryRowContext(ctx, query, ticketID)

	ticket := &domain.Ticket{}
	err := row.Scan(&ticket.Id, &ticket.CruiseID, &ticket.PassengerID, &ticket.PassengerClass, &ticket.Bonuses)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Ticket not found
		}
		return nil, err
	}

	return ticket, nil
}
