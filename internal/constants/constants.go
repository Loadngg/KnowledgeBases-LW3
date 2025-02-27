package constants

type TextEnum int

const (
	WindowName TextEnum = iota
	NnDumpName
	Tab1Name
	Tab2Name
	GenerateNewNnCheckboxLabel
	StartBtnLabel
	ResultLabel
	MissingResult
	HiddenWeightsLabel
	OutputWeightsLabel
	InputEntryPlaceholder
	InvalidInputError
)

var textLabels = map[TextEnum]string{
	WindowName:                 "Базы знаний Лр3",
	NnDumpName:                 "gonn",
	Tab1Name:                   "Использовать",
	Tab2Name:                   "Исходные данные",
	GenerateNewNnCheckboxLabel: "Обучить заново",
	StartBtnLabel:              "Начать",
	ResultLabel:                "Результат: %v",
	MissingResult:              "Результат отсутствует",
	HiddenWeightsLabel:         "Скрытые веса:",
	OutputWeightsLabel:         "Выходные веса:",
	InputEntryPlaceholder:      "0.63 0.25 0.84 1",
	InvalidInputError:          "Неверный ввод начальных значений",
}

func (e TextEnum) String() string {
	if val, ok := textLabels[e]; ok {
		return val
	}
	return "Неизвестный ключ"
}
