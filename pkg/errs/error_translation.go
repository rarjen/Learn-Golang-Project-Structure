package errs

import (
	"fmt"
	"template-ulamm-backend-go/utils"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func ErrorTranslation(errCode string, lang string) *Error {
	// Create a Localizer for the specified language
	localizer := i18n.NewLocalizer(utils.GetLanguageBundle(), lang)

	title, err := localizer.LocalizeMessage(&i18n.Message{
		ID: fmt.Sprintf("%s_title", errCode),
	})
	if err != nil {
		title = errCode
	}

	message, err := localizer.LocalizeMessage(&i18n.Message{
		ID: fmt.Sprintf("%s_message", errCode),
	})
	if err != nil {
		message = errCode
	}

	return &Error{
		errCode: errCode,
		title:   title,
		message: message,
	}
}
