package apperror

type ErrorType int

const (
	IncorrectClientIdException ErrorType = iota
	ClientNotFoundException
	AccountNotFoundException
	WrongInputValueException
	NegativeInputValueException
	InsufficientFundsException
)

func (i ErrorType) String() string {
	switch i {
	case IncorrectClientIdException:
		return "Неверный номер клиента"
	case ClientNotFoundException:
		return "Нет клиента с таким ID"
	case AccountNotFoundException:
		return "Нет счета с таким ID"
	case WrongInputValueException:
		return "Вы ввели неверное значение"
	case InsufficientFundsException:
		return "Вы пытаетесь снять средства, превышающие остаток"
	case NegativeInputValueException:
		return "Вы ввели отрицательное число"
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
