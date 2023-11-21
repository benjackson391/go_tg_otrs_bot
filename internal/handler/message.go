package handler

import (
	"fmt"
	"strconv"
	"tg_bot/config"
	"tg_bot/internal/common"
	"tg_bot/internal/logger"
	"tg_bot/internal/models"
	"tg_bot/internal/otrs"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleMessage(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState) {
	logger.Debug("message.HandleMessage")
	switch userData.CurrentState {
	case "waiting_for_request_topic":
		userData.Topic = update.Message.Text
		userData.CurrentState = "waiting_for_comment"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите описание заявки")
		bot.Send(msg)

	case "waiting_for_comment":
		userData.Description = update.Message.Text
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

func handleAuthoriseMessage(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState) {
	logger.Debug("message.handleAuthoriseMessage")
	if userData.CurrentState != "" {
		switch userData.CurrentState {
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
				userData.CurrentState = "waiting_for_auth_email_code"
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
		userData.CurrentState = "waiting_for_auth_email"
		bot.Send(msg)
	}
}
