package domain

type Currency int

const (
	UNKNOWN Currency = iota
	RUB
	EUR
	USD
	GBP
)

var CurrencyName = map[Currency]string{
	UNKNOWN: "Неверное значение",
	RUB:     "Рубли",
	EUR:     "Евро",
	USD:     "Доллары",
	GBP:     "Фунты",
}

func (c Currency) String() string {
	s, isExist := CurrencyName[c]
	if isExist {
		return s
	} else {
		return CurrencyName[UNKNOWN]
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
