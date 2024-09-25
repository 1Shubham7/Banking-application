package util

// supported currencies
const (
	INR  = "INR"
	USD  = "USD"
	EUR  = "EUR"
	CAD  = "CAD"
	YUAN = "YUAN"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, INR, YUAN, EUR, CAD:
		return true
	}
	return false
}
