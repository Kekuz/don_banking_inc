package apperror

type ErrorType int

const (
	IncorrectClientId ErrorType = iota
	NoClientWithThatId
	NoAccountsWithThatId
)

func (i ErrorType) String() string {
	switch i {
	case IncorrectClientId:
		return "Неверный номер клиента"
	case NoClientWithThatId:
		return "Нет клиента с таким ID"
	case NoAccountsWithThatId:
		return "Нет счета с таким ID"
	default:
		return "Неизвестная ошибка"
	}
}

type AppError struct {
	Err error
	ErrorType
	// Сообщение об ошибке, которое понятно пользователю.
	Message string
	//Выполняемая операция, содержит имя метода или функции
	Op string
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) UnWrap() error {
	return e.Err
}

func New(err error, errorType ErrorType, message, op string) *AppError {
	return &AppError{
		Err:       err,
		ErrorType: errorType,
		Message:   message,
		Op:        op,
	}
}
