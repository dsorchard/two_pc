package common

type Operation int

const (
	NoOp Operation = iota
	PutOp
	DelOp
	RecoveryOp
)

func (s Operation) String() string {
	switch s {
	case PutOp:
		return "PUT"
	case DelOp:
		return "DEL"
	case RecoveryOp:
		return "RECOVERY"
	case NoOp:
		return "NOOP"
	default:
		panic("unhandled default case")
	}
	return "INVALID"
}

func ParseOperation(s string) Operation {
	switch s {
	case "PUT":
		return PutOp
	case "DEL":
		return DelOp
	case "RECOVERY":
		return RecoveryOp
	}
	return NoOp
}

// ----------------------------------------------------------------------

type TxState int

const (
	NoState TxState = iota
	Started
	Prepared
	Committed
	Aborted
)

func (s TxState) String() string {
	switch s {
	case Started:
		return "STARTED"
	case Prepared:
		return "PREPARED"
	case Committed:
		return "COMMITTED"
	case Aborted:
		return "ABORTED"
	default:
		panic("unhandled default case")
	}
	return "INVALID"
}
func ParseTxState(s string) TxState {
	switch s {
	case "STARTED":
		return Started
	case "PREPARED":
		return Prepared
	case "COMMITTED":
		return Committed
	case "ABORTED":
		return Aborted
	}
	return NoState
}

//----------------------------------------------------------------------

type Tx struct {
	Id    string
	Key   string
	Op    Operation
	State TxState
}
