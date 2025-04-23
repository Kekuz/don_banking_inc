package domain

type Currency int

const (
	UNKNOWN Currency = iota
	RUB
	EUR
	USD
	GBP
)

func (c Currency) String() string {
	switch c {
	case RUB:
		return "Рубли"
	case EUR:
		return "Евро"
	case USD:
		return "Доллары"
	case GBP:
		return "Фунты"
	default:
		return "Неверное значение"
	}
}

func (c Currency) StringAcronym() string {
	switch c {
	case RUB:
		return "RUB"
	case EUR:
		return "EUR"
	case USD:
		return "USD"
	case GBP:
		return "GBP"
	default:
		return "nil"
	}
}

func ToCurrency(s string) Currency {
	switch s {
	case "RUB":
		return RUB
	case "EUR":
		return EUR
	case "USD":
		return USD
	case "GBP":
		return GBP
	default:
		return UNKNOWN
	}
}
