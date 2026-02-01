package handler

import (
	"A/internal/logfile"
	"A/internal/uidraw"
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

// Данные вида.
type DataView struct {
	ViewName          string           // Имя вида
	CurrentFocusInput string           // Имя поля ввода, на который установлен фокус.
	ListInputName     []string         // Список имён полей ввода.
	ObjCounter        *DataObjCounter  // Объект счётчика
	LogFile           *logfile.LogFile // Логгер
}

// Выход. Возвращается ошибка.
//
// Параметры:
//
//	g - указатель на Gui.
//	v - указатель на View.
func Quit(g *gocui.Gui, v *gocui.View) error {

	return gocui.ErrQuit
}

// Перевод фокуса. Возвращается ошибка.
//
// Параметры:
//
//	g - указатель на Gui.
//	v - указатель на View.
func (i *DataView) NextFocus(g *gocui.Gui, v *gocui.View) error {

	// Перевод фокуса
	switch i.ViewName {
	case "main": // Окно регистрации.
		switch i.CurrentFocusInput {
		case "input1":
			i.CurrentFocusInput = "input2"
		case "input2":
			i.CurrentFocusInput = "input1"
		default:
		}

	default:
		return nil
	}

	_, err := g.SetCurrentView(i.CurrentFocusInput)
	if err != nil {
		return fmt.Errorf("Ошибка в функции SetCurrentView: <%w>", err)
	}

	// Обновление
	var fields []string
	if i.ViewName == "main" {
		fields = []string{"input1", "input2"}
	}

	for _, name := range fields {
		view, err := g.View(name)
		if err != nil {
			return fmt.Errorf("Ошибка установки фокуса: <%w>", err)
		}
		if view != nil {
			uidraw.SetFocusStyle(view, i.CurrentFocusInput, name)
		}
	}

	// Установка курсора в конец текущей строки
	currentView, err := g.View(i.CurrentFocusInput)
	if err != nil {
		return fmt.Errorf("Ошибка в функции View: <%w>", err)
	}
	if currentView != nil {
		buffer := strings.TrimSuffix(currentView.Buffer(), "\n")
		cursorX := len(buffer)
		currentView.SetCursor(cursorX, 0)
	}

	return nil
}

// Фиксация ввода. Возвращается ошибка.
//
// Параметры:
//
//	g - указатель на Gui.
//	v - указатель на View.
func (i *DataView) Enter(g *gocui.Gui, v *gocui.View) error {

	// Получаем содержимое поля ввода
	inputView, err := g.View(i.CurrentFocusInput)
	if err != nil {
		return fmt.Errorf("не удалось получить вид ввода: <%w>", err)
	}

	// Читаем содержимое
	inputContent := inputView.Buffer()
	inputContent = strings.ReplaceAll(inputContent, "\n", "")

	// Получаем вид вывода

	switch i.CurrentFocusInput {
	case i.ListInputName[0]:

		outputView, err := g.View("output1")
		if err != nil {
			return fmt.Errorf("не удалось получить вид вывода: <%w>", err)
		}
		outputView.Clear()
		outputView.Write([]byte(inputContent))

	case i.ListInputName[1]:

		outputView, err := g.View("output2")
		if err != nil {
			return fmt.Errorf("не удалось получить вид вывода: <%w>", err)
		}
		outputView.Clear()
		outputView.Write([]byte(inputContent))

	default:
	}

	return nil
}

// Обновление данных вида. Возвращается ошибка.
//
// Параметры:
//
//	g - указатель на Gui.
func (i *DataView) Update(g *gocui.Gui) error {

	// Обработка данных счёта.
	outputView, err := g.View("outputCnt")
	if err != nil {
		return fmt.Errorf("не удалось получить вид вывода: <%w>", err)
	}
	outputView.Clear()
	outputView.Write([]byte(fmt.Sprintf("%d", i.ObjCounter.Value)))

	//
	// Обновления индикаторов размерешения
	//

	// Если процесс работает
	if i.ObjCounter.Status == StageRun {

		// Обновление для СТАРТ
		elStart, err := g.View(i.ObjCounter.NameEnStart)
		if err != nil {
			return fmt.Errorf("не удалось получить вид вывода: <%w>", err)
		}
		elStart.BgColor = gocui.ColorDefault

		// Обновление для СТОП
		elStop, err := g.View(i.ObjCounter.NameEnStop)
		if err != nil {
			return fmt.Errorf("не удалось получить вид вывода: <%w>", err)
		}
		elStop.BgColor = gocui.ColorGreen
	}

	// Если процесс остановлен
	if i.ObjCounter.Status == StageStop {

		// Обновление для СТАРТ
		elStart, err := g.View(i.ObjCounter.NameEnStart)
		if err != nil {
			return fmt.Errorf("не удалось получить вид вывода: <%w>", err)
		}
		elStart.BgColor = gocui.ColorGreen

		// Обновление для СТОП
		elStop, err := g.View(i.ObjCounter.NameEnStop)
		if err != nil {
			return fmt.Errorf("не удалось получить вид вывода: <%w>", err)
		}
		elStop.BgColor = gocui.ColorDefault
	}

	return nil
}
