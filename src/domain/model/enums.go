package domain

import "fmt"

type UserRole int
type OrderStatus int

const (
	CUSTOMER UserRole = iota
	ADMIN
)

const (
	IN_PREPARATION OrderStatus = iota
	READY_TO_BE_DELIVERED
	CANCELLED
	DELIVERED
)

func (role UserRole) String() string {
	switch role {
	case CUSTOMER:
		return "CUSTOMER"
	case ADMIN:
		return "ADMIN"
	default:
		return fmt.Sprintf("%d", int(role))
	}
}

func (status OrderStatus) String() string {
	switch status {
	case IN_PREPARATION:
		return "IN_PREPARATION"
	case READY_TO_BE_DELIVERED:
		return "READY_TO_BE_DELIVERED"
	case CANCELLED:
		return "CANCELLED"
	case DELIVERED:
		return "DELIVERED"
	default:
		return fmt.Sprintf("%d", int(status))
	}
}
