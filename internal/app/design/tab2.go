package design

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"lr3/internal/app/network/data"
)

func Tab2(data *data.NNData) *widget.Table {
	tableData := data.GetMergedData()
	table := widget.NewTableWithHeaders(
		func() (int, int) {
			return len(tableData), len(tableData[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("0")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(fmt.Sprintf("%v", tableData[i.Row][i.Col]))
		},
	)
	table.UpdateHeader = func(id widget.TableCellID, o fyne.CanvasObject) {
		l := o.(*widget.Label)
		if id.Row < 0 {
			col := id.Col + 1
			if col == 5 {
				l.SetText("Y")
				return
			}
			l.SetText(fmt.Sprintf("X%v", col))
		} else if id.Col < 0 {
			l.SetText(strconv.Itoa(id.Row + 1))
		} else {
			l.SetText("")
		}
	}
	return table
}
