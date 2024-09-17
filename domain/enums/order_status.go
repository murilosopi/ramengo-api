package enums

type OrderStatus uint8

const (
	ConfirmedStatus OrderStatus = iota + 1
	PreparingStatus
	ReadyStatus
	CancelledStatus
)

func IsValidOrderStatus(status int) bool {
	switch OrderStatus(status) {
	case ConfirmedStatus, PreparingStatus, ReadyStatus, CancelledStatus:
		return true
	}

	return false
}

func (s OrderStatus) String() string {
	switch s {
	case ConfirmedStatus:
		return "New"
	case PreparingStatus:
		return "Preparing"
	case ReadyStatus:
		return "Ready"
	case CancelledStatus:
		return "Cancelled"
	default:
		return "Unknow"
	}
}
