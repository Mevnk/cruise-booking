package handler

import (
	"github.com/google/uuid"
	"pis/business"
	"pis/domain"
	"pis/handler/dto"
)

func toSignUpParams(dto *dto.SignUpRequestDTO) business.SignUpParams {
	return business.SignUpParams{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func toSignInParams(dto *dto.SignInRequestDTO) business.SignInParams {
	return business.SignInParams{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func mapDtoUser(user *domain.User) *dto.UserDTO {
	return &dto.UserDTO{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role: &dto.RoleDTO{
			Id:   user.Role.Id,
			Name: user.Role.Name,
		},
	}
}

func toUpdateRoleParams(userId uuid.UUID, roleDto *dto.UpdateRoleDTO) business.UpdateUserRoleParams {
	return business.UpdateUserRoleParams{UserId: userId, Role: roleDto.Role}
}

func mapDtoUsers(users []*domain.User) []*dto.UserDTO {
	var dtoUsers []*dto.UserDTO
	for _, user := range users {
		dtoUsers = append(dtoUsers, mapDtoUser(user))
	}
	return dtoUsers
}

func toAddCruiseParams(dto *dto.AddCruiseDTO) business.CruiseParams {
	return business.CruiseParams{
		DepartureDate: dto.DepartureDate,
		Price:         dto.Price,
		Route:         dto.Route,
		NofPorts:      dto.NumberOfPorts,
		Duration:      dto.Duration,
	}
}

func toAddTicketParams(dto *dto.AddTicketDTO) business.TicketParams {
	return business.TicketParams{
		CruiseID:       dto.CruiseID,
		PassengerID:    dto.PassengerID,
		PassengerClass: dto.PassengerClass,
		Bonuses:        dto.Bonuses,
	}
}

func toAddExcursionParams(dto *dto.AddExcursionDTO) business.ExcursionParams {
	return business.ExcursionParams{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
	}
}

func mapDtoCruise(cruise *domain.Cruise) *dto.CruiseDTO {
	return &dto.CruiseDTO{
		ID:            cruise.Id,
		DepartureDate: cruise.DepartureDate,
		Price:         cruise.Price,
		Route:         cruise.Route,
		NumberOfPorts: cruise.NofPorts,
		Duration:      cruise.Duration,
	}
}

func mapDtoTicket(ticket *domain.Ticket) *dto.TicketDTO {
	return &dto.TicketDTO{
		ID:             ticket.Id,
		CruiseID:       ticket.CruiseID,
		PassengerID:    ticket.PassengerID,
		PassengerClass: ticket.PassengerClass,
		Bonuses:        ticket.Bonuses,
	}
}
