package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	switch t := value.(type) {
	case error:
		color.Red(t.Error())
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	default:
		color.Red("Неизвестная ошибка")
	}
}
