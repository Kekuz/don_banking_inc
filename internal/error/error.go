package apperror

type ErrorType int

const (
	IncorrectClientIdException ErrorType = iota
	ClientNotFoundException
	AccountNotFoundException
	WrongInputValueException
	NegativeInputValueException
	InsufficientFundsException
	AccountDeleteException
	AccountEditException
	AccountNotSelectedException
)

func (i ErrorType) String() string {
	switch i {
	case IncorrectClientIdException:
		return "Номер клиента неверный"
	case ClientNotFoundException:
		return "Клиент с таким ID не найден"
	case AccountNotFoundException:
		return "Счет с таким ID не найден"
	case WrongInputValueException:
		return "Введено неверное значение"
	case InsufficientFundsException:
		return "Снимаемые средства, превышающие остаток"
	case NegativeInputValueException:
		return "Введено отрицательное число"
	case AccountDeleteException:
		return "Ошибка удаления счета"
	case AccountEditException:
		return "Ошибка изменения счета"
	case AccountNotSelectedException:
		return "Счет не выбран"
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
