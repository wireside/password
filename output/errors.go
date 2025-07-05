package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	errValue, ok := value.(error)
	if ok {
		color.Red(errValue.Error())
		return
	}
	
	strValue, ok := value.(string)
	if ok {
		color.Red(strValue)
		return
	}
	
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ошибки: %d", intValue)
		return
	}
	
	color.Red("Неизвестная ошибка")
}
