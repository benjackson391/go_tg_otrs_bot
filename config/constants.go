package config

const (
	WelcomeMessage        = "Добро пожаловать в систему постановки заявок EFSOL"
	WelcomeDescription    = "Вы можете создать, обновить или проверить статус вашей заявки.Для остановки бота просто напишите /stop, для повторного запуска нажмите /start"
	LeaveMessage          = "До встречи."
	TicketUpdatedMessage  = "Вашe обращение #%s обновлено"
	TicketCreatedTemplate = "Вашe обращение принято. Номер заявки #%s"
	UnknownCommand        = "Unknown command."
	FileSizeLimit         = 20 * 1024 * 1024 // 20 MB
	BigFileMessage        = "Файл слишком большой. Приложите файл меннее 20 МБ"
	ErrorMsq1             = "Произошла ошибка при получении файла."
	NotEnoughData         = "Ошибка: недостаточно данных, чтобы создать/обновить заявку!"
	VoteTemplate          = "Обращение %s оценено!"
	UpdateTitle           = "Выберите заявку, которую необходимо обновить"
	MessageMaxLength      = 4096
	// MESSAGE_FILE_LARGE = "Загружен файл большого объёма, его обработка может занять длительное время. Пожалуйста подождите."
	// FILE_LIMIT         = 5 * 1048576
	// FILE_BIG_THRESHOLD = 3000000
)
