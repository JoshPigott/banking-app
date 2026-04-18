package models

type AccountType string

const (
	Everyday  AccountType = "everyday"
	Saver     AccountType = "saver"
	KiwiSaver AccountType = "kiwiSaver"
)

func (a AccountType) IsValid() bool {
	switch a {
	case Everyday, Saver, KiwiSaver:
		return true
	default:
		return false
	}
}

func (a AccountType) GetTableName() string {
	return string(a) + "account"
}

type AccountBalance struct {
	AccountType AccountType
	Balance     float64
}
