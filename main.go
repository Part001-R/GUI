package main

import (
	"A/internal/handler"
	"A/internal/logfile"
	"A/internal/uidraw"
	"fmt"
	"log"
	"time"

	"github.com/jroimartin/gocui"
)

// Объект - счётчик тиков.
var objCounter = handler.DataObjCounter{
	Status: handler.StageStop,
	Value:  0,
	ChCmd:  make(chan int),
}

// Данные окна.
var dataView = handler.DataView{
	ViewName:          "main",
	CurrentFocusInput: "input1",
	ListInputName:     []string{"input1", "input2"},
	ObjCounter:        &objCounter,
}

func main() {

	lgrTxt, err := logfile.New("testLog.txt")
	if err != nil {
		log.Fatalf("Ошибка создания логгера:<%v>", err)
	}
	objCounter.LogFile = lgrTxt
	dataView.LogFile = lgrTxt

	// -------------------------------------------------

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

	//
	// -------------------------------------------
	//

	// Создание основного представления
	maxX := 150
	maxY := 30
	v, err := g.SetView(dataView.ViewName, 0, 0, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Демонстрация"
		v.Wrap = true
	}

	//
	// -------------------------------------------
	//

	// Кнопка - СТАРТ
	dataButtonStart := uidraw.DrawDataButton{
		NameSquare:  "square1",
		ColorSquare: gocui.ColorGreen,
		NameButton:  "button1",
		TextButton:  "СТАРТ",
		StartX:      2,
		StartY:      1,
	}
	if err := uidraw.DrawButton(g, dataButtonStart); err != nil {
		return fmt.Errorf("Ошибка отрисовки кнопки: <%w>", err)
	}
	objCounter.NameEnStart = dataButtonStart.NameSquare

	// Кнопка - СТОП
	dataButtonStop := uidraw.DrawDataButton{
		NameSquare:  "square2",
		ColorSquare: gocui.ColorDefault,
		NameButton:  "button2",
		TextButton:  "СТОП",
		StartX:      2,
		StartY:      4,
	}
	if err := uidraw.DrawButton(g, dataButtonStop); err != nil {
		return fmt.Errorf("Ошибка отрисовки кнопки: <%w>", err)
	}
	objCounter.NameEnStop = dataButtonStop.NameSquare

	//
	// -------------------------------------------
	//

	// Поле ввода-1.
	dataInput1 := uidraw.DrawDataInput{
		Color:     gocui.ColorCyan,
		NameInput: dataView.ListInputName[0],
		NameTitle: "Данные-1",
		Len:       30,
		StartX:    2,
		StartY:    9,
	}
	if err := uidraw.DrawInput(g, dataInput1); err != nil {
		return fmt.Errorf("Ошибка отрисовки ввода: <%w>", err)
	}

	// Поле ввода-2.
	dataInput2 := uidraw.DrawDataInput{
		Color:     gocui.ColorCyan,
		NameInput: dataView.ListInputName[1],
		NameTitle: "Данные-2",
		Len:       30,
		StartX:    2,
		StartY:    12,
	}
	if err := uidraw.DrawInput(g, dataInput2); err != nil {
		return fmt.Errorf("Ошибка отрисовки ввода: <%w>", err)
	}

	//
	// -------------------------------------------
	//

	// Поле вывода-1.
	dataOutput1 := uidraw.DrawDataOutput{
		NameOutput: "output1",
		NameTitle:  "",
		Len:        30,
		StartX:     34,
		StartY:     9,
	}
	if err := uidraw.DrawOutput(g, dataOutput1); err != nil {
		return fmt.Errorf("Ошибка отрисовки вывода: <%w>", err)
	}

	// Поле вывода-2.
	dataOutput2 := uidraw.DrawDataOutput{
		NameOutput: "output2",
		NameTitle:  "",
		Len:        30,
		StartX:     34,
		StartY:     12,
	}
	if err := uidraw.DrawOutput(g, dataOutput2); err != nil {
		return fmt.Errorf("Ошибка отрисовки вывода: <%w>", err)
	}

	//
	// -------------------------------------------
	//

	// Пояснение к действию - TAB.
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

	// Пояснение к действию - Enter.
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

	// Пояснение к действию - СТАРТ.
	dataDrawGuide = uidraw.DrawDataGuide{
		NameGuide: "Guide3",
		Text:      "Ctrl+A - СТАРТ",
		StartX:    66,
		StartY:    26,
		Weight:    30,
	}
	if err := uidraw.DrawGuide(g, dataDrawGuide); err != nil {
		return fmt.Errorf("Ошибка отрисовки пояснения: <%w>", err)
	}

	// Пояснение к действию - СТОП.
	dataDrawGuide = uidraw.DrawDataGuide{
		NameGuide: "Guide4",
		Text:      "Ctrl+B - СТОП",
		StartX:    98,
		StartY:    26,
		Weight:    30,
	}
	if err := uidraw.DrawGuide(g, dataDrawGuide); err != nil {
		return fmt.Errorf("Ошибка отрисовки пояснения: <%w>", err)
	}

	//
	// -------------------------------------------
	//

	// Поле вывода счётчика.
	dataOutputCnt := uidraw.DrawDataOutput{
		NameOutput: "outputCnt",
		NameTitle:  "",
		Len:        10,
		StartX:     34,
		StartY:     2,
	}
	if err := uidraw.DrawOutput(g, dataOutputCnt); err != nil {
		return fmt.Errorf("Ошибка отрисовки вывода: <%w>", err)
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

	// СТАРТ.
	if err := g.SetKeybinding("", gocui.KeyCtrlA, gocui.ModNone, objCounter.Start); err != nil {
		return fmt.Errorf("Error: Ошибка CtrlA: <%w>", err)
	}

	// СТОП.
	if err := g.SetKeybinding("", gocui.KeyCtrlB, gocui.ModNone, objCounter.Stop); err != nil {
		return fmt.Errorf("Error: Ошибка CtrlB: <%w>", err)
	}

	// Enter.
	for _, name := range dataView.ListInputName {
		if err := g.SetKeybinding(name, gocui.KeyEnter, gocui.ModNone, dataView.Enter); err != nil {
			return fmt.Errorf("Error: ошибка обработки нажатия Enter: <%w>, на элементе: <%s>", err, name)
		}
	}

	// Обновление данных вида.
	go func() {
		for {
			g.Update(func(g *gocui.Gui) error {
				if err := dataView.Update(g); err != nil {
					return fmt.Errorf("Error: ошибка обновления индикаторов: <%w>", err)
				}
				return nil
			})
			time.Sleep(100 * time.Millisecond)
		}
	}()

	return nil
}
