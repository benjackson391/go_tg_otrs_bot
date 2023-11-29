package handler

import (
	"fmt"
	"strconv"
	"tg_bot/config"
	"tg_bot/internal/common"
	"tg_bot/internal/logger"
	"tg_bot/internal/models"
	"tg_bot/internal/otrs"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState) {
	if isMediaMessage(update.Message) {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Некорректное вложение, приложите Файл")
		bot.Send(msg)
	} else if update.Message.Document != nil {
		HandleDocument(update, bot, userData)
	} else {
		length := len(update.Message.Text)
		threshold := config.MessageMaxLength
		if length > threshold {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Cообщение слишком длинное (%d символов, максимальная длинна сообщения - %d символов)", length, threshold)))
		} else {
			switch userData.Action {
			case "waiting_for_title":

				userData.Topic = update.Message.Text

				logger.Debug(fmt.Sprintf("[%s:%s] input title: `%s`", userData.UserName, userData.Action, userData.Topic))

				userData.Action = "waiting_for_comment"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите описание заявки")
				bot.Send(msg)

			case "waiting_for_comment":
				userData.Description = update.Message.Text

				logger.Debug(fmt.Sprintf("[%s:%s] input description: `%s`", userData.UserName, userData.Action, userData.Description))

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Прикрепить файл?")
				inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Да", "attach_file"),
						tgbotapi.NewInlineKeyboardButtonData("Нет", "create"),
					),
				)
				msg.ReplyMarkup = inlineKeyboard
				bot.Send(msg)
			}
		}
	}
}

func HandleDocument(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState) {
	logger.Debug(fmt.Sprintf("[%s:%s] handle document: `%s`", userData.UserName, userData.Action, update.Message.Document.FileName))

	doc := update.Message.Document

	if doc.FileSize > config.FileSizeLimit {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, config.BigFileMessage))
	} else {
		documentBytes, err := common.GetFileContent(doc.FileID, bot)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, config.ErrorMsq1))
		} else {
			_, text := otrs.CreateOrUpdateTicket(bot, userData, doc, &documentBytes)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			buttons := []models.Button{
				{Title: "Назад", Callback: "start"},
			}
			msg.ReplyMarkup = common.GetInlineKeyboard(buttons)
			bot.Send(msg)
			common.CleanUpUserData(userData)
		}
	}

}

func handleAuthoriseMessage(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState) {
	logger.Debug("message.handleAuthoriseMessage")
	if userData.Action != "" {
		switch userData.Action {
		case "waiting_for_auth_email":
			Email := update.Message.Text
			userData.Email = Email
			Code := common.GenerateSixDigitNumber()
			userData.Code = int(Code)

			resp, err := otrs.OtrsConfirm("confirm_account", models.OtrsConfirmRequest{
				Code:  Code,
				Email: Email,
			})

			var template = "Адрес %s не найден в OTRS!\nПопробуйте другой адрес"
			if err != nil {
				template = "Произошла ошибка при подтверждении %s, попробуйте позднее"
			} else if resp.Sent == 1 {
				template = "Проверочный код отправлен на %s\nДля подтверждения адреса электронной почты введите код из письма"
				userData.Action = "waiting_for_auth_email_code"
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(template, Email))
			bot.Send(msg)
		case "waiting_for_auth_email_code":
			Code := update.Message.Text
			text := "Код неверный, попробуйте еще раз"
			CodeInt, _ := strconv.Atoi(Code)

			if userData.Code == CodeInt {
				text = "Код верный"
				// resp, err
				otrs.OtrsConfirm("confirm_account", models.OtrsConfirmRequest{
					Email:         userData.Email,
					TelegramLogin: update.Message.From.UserName,
				})
				common.CleanUpUserData(userData)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

			if userData.Code == CodeInt {
				msg.ReplyMarkup = common.GetInitialKeyboard()
			}
			bot.Send(msg)

		default:
		}
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, config.WelcomeMessage+"\n\nВведите свой адрес электронный почты, который зарегистрирован в OTRS")
		userData.Action = "waiting_for_auth_email"
		bot.Send(msg)
	}
}

func isMediaMessage(msg *tgbotapi.Message) bool {
	return msg.Photo != nil ||
		msg.Video != nil ||
		msg.Audio != nil ||
		msg.Voice != nil ||
		msg.Animation != nil ||
		msg.Game != nil ||
		msg.Sticker != nil ||
		msg.VideoNote != nil ||
		msg.Contact != nil ||
		msg.Location != nil ||
		msg.Venue != nil
}
