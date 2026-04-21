package domain

type BankAccountType string

const (
	Everyday  BankAccountType = "everyday"
	Saver     BankAccountType = "saver"
	KiwiSaver BankAccountType = "kiwiSaver"
)

func (a BankAccountType) IsValid() bool {
	switch a {
	case Everyday, Saver, KiwiSaver:
		return true
	default:
		return false
	}
}

func (a BankAccountType) CanWithdraw() bool {
	switch a {
	case Everyday, Saver:
		return true
	default:
		return false
	}
}

func (a BankAccountType) GetTableName() string {
	return string(a) + "Account"
}

type AccountBalance struct {
	BankAccountType BankAccountType
	Balance         float64
}

type TransferRequest struct {
	SessionID   string
	AccountFrom BankAccountType
	AccountTo   BankAccountType
	AmountCents int
	// Derived / computed later
	AccountFromTable string
	AccountToTable   string
	UserID           string
}

type PaymentRequest struct {
	SessionID        string
	AccountFrom      BankAccountType
	ReceiverUsername string
	AmountCents      int
	// Derived / computed later
	AccountFromTable string
	AccountToTable   string
	UserID           string
	ReceiveUserID    string
}
