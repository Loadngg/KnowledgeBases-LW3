package design

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"

	"lr3/internal/app/network"
	"lr3/internal/constants"
)

func MustLoad(n *network.Network) *container.AppTabs {
	tab1 := Tab1(n)
	tab2 := Tab2(n.Data)

	tab1Container := container.NewTabItemWithIcon(constants.Tab1Name.String(), theme.ComputerIcon(), tab1)
	tab2Container := container.NewTabItemWithIcon(constants.Tab2Name.String(), theme.GridIcon(), tab2)

	tabs := container.NewAppTabs(
		tab1Container,
		tab2Container,
	)

	return tabs
}
