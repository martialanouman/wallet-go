package main

type Currency int

const (
	currencyUnknown Currency = iota
	XOF
	EUR
	USD
	currencyEnd
)

func currencyCode(c Currency) string {
	switch c {
	case XOF:
		return "XOF"
	case EUR:
		return "EUR"
	case USD:
		return "USD"
	default:
		return "UNKNOWN"
	}
}

func minorUnits(c Currency) int {
	switch c {
	case XOF:
		return 0
	case EUR:
	case USD:
		return 2
	default:
		return -1
	}
}
