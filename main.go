package main

import (
	"A/internal/handler"
	"A/internal/uidraw"
	"fmt"

	"github.com/jroimartin/gocui"
)

// Данные окна.
var dataView = handler.DataView{
	ViewName:          "main",
	CurrentFocusInput: "input1",
	ListInputName:     []string{"input1", "input2"},
}

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
	maxX := 150
	maxY := 30

	// Создание основного представления
	v, err := g.SetView(dataView.ViewName, 0, 0, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Демонстрация"
		v.Wrap = true
	}

	// Кнопка - СТАРТ
	dataDrawButton := uidraw.DrawDataButton{
		NameSquare:  "square1",
		ColorSquare: gocui.ColorGreen,
		NameButton:  "button1",
		TextButton:  "СТАРТ",
		StartX:      2,
		StartY:      1,
	}
	if err := uidraw.DrawButton(g, dataDrawButton); err != nil {
		return fmt.Errorf("Ошибка отрисовки кнопки: <%w>", err)
	}

	// Кнопка - СТОП
	dataDrawButton = uidraw.DrawDataButton{
		NameSquare:  "square2",
		ColorSquare: gocui.ColorDefault,
		NameButton:  "button2",
		TextButton:  "СТОП",
		StartX:      2,
		StartY:      4,
	}
	if err := uidraw.DrawButton(g, dataDrawButton); err != nil {
		return fmt.Errorf("Ошибка отрисовки кнопки: <%w>", err)
	}

	// Поле ввода-1.
	dataDrawInput := uidraw.DrawDataInput{
		Color:     gocui.ColorCyan,
		NameInput: dataView.ListInputName[0],
		NameTitle: "Данные-1",
		Len:       30,
		StartX:    2,
		StartY:    9,
	}
	if err := uidraw.DrawInput(g, dataDrawInput); err != nil {
		return fmt.Errorf("Ошибка отрисовки ввода: <%w>", err)
	}

	// Поле ввода-2.
	dataDrawInput = uidraw.DrawDataInput{
		Color:     gocui.ColorCyan,
		NameInput: dataView.ListInputName[1],
		NameTitle: "Данные-2",
		Len:       30,
		StartX:    2,
		StartY:    12,
	}
	if err := uidraw.DrawInput(g, dataDrawInput); err != nil {
		return fmt.Errorf("Ошибка отрисовки ввода: <%w>", err)
	}

	// Поле вывода-1.
	dataDrawOutput := uidraw.DrawDataOutput{
		NameOutput: "output1",
		NameTitle:  "",
		Len:        30,
		StartX:     34,
		StartY:     9,
	}
	if err := uidraw.DrawOutput(g, dataDrawOutput); err != nil {
		return fmt.Errorf("Ошибка отрисовки вывода: <%w>", err)
	}

	// Поле вывода-2.
	dataDrawOutput = uidraw.DrawDataOutput{
		NameOutput: "output2",
		NameTitle:  "",
		Len:        30,
		StartX:     34,
		StartY:     12,
	}
	if err := uidraw.DrawOutput(g, dataDrawOutput); err != nil {
		return fmt.Errorf("Ошибка отрисовки вывода: <%w>", err)
	}

	// Пояснение к действию - 1.
	dataDrawGuide := uidraw.DrawDataGuide{
		NameGuide: "Guide1",
		Text:      "TAB - перевод фокуса",
		StartX:    2,
		StartY:    26,
		Weight:    30,
	}
	if err := uidraw.DrawGuide(g, dataDrawGuide); err != nil {
		return fmt.Errorf("Ошибка отрисовки пояснения: <%w>", err)
	}

	// Пояснение к действию - 2.
	dataDrawGuide = uidraw.DrawDataGuide{
		NameGuide: "Guide2",
		Text:      "Enter - фиксация ввода",
		StartX:    34,
		StartY:    26,
		Weight:    30,
	}
	if err := uidraw.DrawGuide(g, dataDrawGuide); err != nil {
		return fmt.Errorf("Ошибка отрисовки пояснения: <%w>", err)
	}

	return nil
}

// Привязки.
func keybindings(g *gocui.Gui) error {

	// Завершение работы.
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, handler.Quit); err != nil {
		return fmt.Errorf("Error: Ошибка Ctrl+C: <%w>", err)
	}

	// Переключение фокуса по Tab.
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, dataView.NextFocus); err != nil {
		return fmt.Errorf("Error: Ошибка Tab: <%w>", err)
	}

	// Enter.
	for _, name := range dataView.ListInputName {
		if err := g.SetKeybinding(name, gocui.KeyEnter, gocui.ModNone, dataView.Enter); err != nil {
			return fmt.Errorf("Error: ошибка обработки нажатия Enter: <%w>, на элементе: <%s>", err, name)
		}
	}

	return nil
}
