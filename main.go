package main

import (
	"A/internal/uidraw"
	"fmt"

	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		panic(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

// Отрисовка главного экрана.
func layout(g *gocui.Gui) error {
	maxX := 50
	maxY := 20

	// Создание основного представления
	if v, err := g.SetView("main", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Кнопка"
		v.Wrap = true
	}

	// Кнопка - СТОП
	dataDrawButton := uidraw.DrawDataButton{
		NameSquare:  "square2",
		ColorSquare: gocui.ColorDefault,
		NameButton:  "button2",
		TextButton:  "СТОП",
		StartX:      2,
		StartY:      5,
	}
	if err := uidraw.DrawButton(g, dataDrawButton); err != nil {
		return fmt.Errorf("Ошибка отрисовки кнопки: <%w>", err)
	}

	// Кнопка - СТАРТ
	dataDrawButton = uidraw.DrawDataButton{
		NameSquare:  "square1",
		ColorSquare: gocui.ColorGreen,
		NameButton:  "button1",
		TextButton:  "СТАРТ",
		StartX:      2,
		StartY:      2,
	}
	if err := uidraw.DrawButton(g, dataDrawButton); err != nil {
		return fmt.Errorf("Ошибка отрисовки кнопки: <%w>", err)
	}

	// Кнопка - КВИТИРОВАНИЕ
	dataDrawButton = uidraw.DrawDataButton{
		NameSquare:  "square3",
		ColorSquare: gocui.ColorGreen,
		NameButton:  "button3",
		TextButton:  "КВИТИРОВАНИЕ",
		StartX:      9,
		StartY:      15,
	}
	if err := uidraw.DrawButton(g, dataDrawButton); err != nil {
		return fmt.Errorf("Ошибка отрисовки кнопки: <%w>", err)
	}

	return nil
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", 'q', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}); err != nil {
		return err
	}
	return nil
}
