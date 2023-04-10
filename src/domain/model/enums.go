package domain

type UserType int
type OrderStatus int

const (
	CUSTOMER UserType = iota
	ADMIN
)

const (
	PREPAIRING OrderStatus = iota
	READY_TO_BE_DELIVERED
	CANCELLED
)
