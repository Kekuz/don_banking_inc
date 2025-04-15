package domain

type Currency int

const (
	RUB Currency = iota
	EUR
	YER
	BRL
	TMT
	CNY
	UNKNOWN
)

func (c Currency) String() string {
	switch c {
	case RUB:
		return "RUB"
	case EUR:
		return "EUR"
	case YER:
		return "YER"
	case BRL:
		return "BRL"
	case TMT:
		return "TMT"
	case CNY:
		return "CNY"
	default:
		return "wrong currency"
	}
}
