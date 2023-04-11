package domain

type UserType int
type OrderStatus int

const (
	CUSTOMER UserType = iota
	ADMIN
)

const (
	IN_PREPARATION OrderStatus = iota
	READY_TO_BE_DELIVERED
	CANCELLED
)
