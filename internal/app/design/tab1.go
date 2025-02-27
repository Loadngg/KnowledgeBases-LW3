package design

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"lr3/internal/app/network"
	"lr3/internal/constants"
	"lr3/internal/utils"
)

func Tab1(n *network.Network) *fyne.Container {
	generateNew := false

	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder(constants.InputEntryPlaceholder.String())

	generateNewNnCheckbox := widget.NewCheck(
		constants.GenerateNewNnCheckboxLabel.String(),
		func(b bool) { generateNew = b },
	)

	resultLabel := widget.NewLabel(constants.MissingResult.String())
	hiddenWeightsLabel := widget.NewLabel(constants.HiddenWeightsLabel.String())
	hiddenWeightsContentLabel := widget.NewLabel("")
	outputWeightsLabel := widget.NewLabel(constants.OutputWeightsLabel.String())
	outputWeightsContentLabel := widget.NewLabel("")

	startBtn := widget.NewButton(constants.StartBtnLabel.String(), nil)
	startBtn.OnTapped = func() {
		inputStr := inputEntry.Text
		input, err := utils.ParseFloat64Array(inputStr)
		if err != nil {
			resultLabel.SetText(constants.InvalidInputError.String())
			return
		}

		startBtn.Disable()
		nn := n.GetNN(generateNew, func() {
			startBtn.Enable()
		})

		result := nn.Forward(input)
		resultLabel.SetText(fmt.Sprintf(constants.ResultLabel.String(), n.GetResult(result)))

		hiddenWeightsContentLabel.SetText(utils.FormatMatrix(nn.WeightHidden, 4))
		outputWeightsContentLabel.SetText(utils.FormatMatrix(nn.WeightOutput, 4))
	}

	content := container.NewVBox(
		inputEntry,
		generateNewNnCheckbox,
		resultLabel,
		startBtn,
		hiddenWeightsLabel,
		hiddenWeightsContentLabel,
		outputWeightsLabel,
		outputWeightsContentLabel,
	)
	return content
}
