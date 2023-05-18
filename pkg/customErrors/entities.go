package customErrors

type Entity string

const (
	User            Entity = "USER"
	Ticket                 = "TICKET"
	Cruise                 = "CRUISE"
	Ship                   = "SHIP"
	Role                   = "ROLE"
	Excursion              = "EXCURSION"
	CruiseExcursion        = "CRUISE-EXCURSION BIND"
)
