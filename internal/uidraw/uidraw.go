// Функции с отрисовкой элементов.
package uidraw

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

// Отрисовка кнопки. Возвращается ошибка.
//
// Параметры:
//
//	g - графический интерфейс.
//	data - данные для отрисовки.
func DrawButton(g *gocui.Gui, data DrawDataButton) error {

	// Вычиление конечный координат.
	buttonXEnd, buttonYEnd := data.StartX+30, data.StartY+2
	squareXEnd, squareYEnd := data.StartX+3, data.StartY+2

	// Кнопка
	if vv, err := g.SetView(data.NameButton, data.StartX, data.StartY, buttonXEnd, buttonYEnd); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		vv.Wrap = true

		// Вычисление позиции для вывода текста.
		txt := []rune(data.TextButton)

		lenButton := buttonXEnd - data.StartX + 3
		if len(txt) > lenButton {
			return fmt.Errorf("Превышена длинна текста")
		}

		pos := (lenButton / 2) - (len(txt) / 2)
		if _, err := vv.Write([]byte(fmt.Sprintf("%s%s", strings.Repeat(" ", pos-1), data.TextButton))); err != nil {
			return err
		}
	}
	// Квадрат
	if sv, err := g.SetView(data.NameSquare, data.StartX, data.StartY, squareXEnd, squareYEnd); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		sv.Wrap = false
		sv.BgColor = data.ColorSquare
	}

	return nil
}

// Отрисовка поля ввода. Возвращается ошибка.
//
// Параметры:
//
//	g - указатель на Gui
//	data - данные для отрисовки.
func DrawInput(g *gocui.Gui, data DrawDataInput) error {

	if v, err := g.SetView(data.NameInput, data.StartX, data.StartY, data.StartX+data.Len, data.StartY+2); err != nil {
		if err != gocui.ErrUnknownView {
			return fmt.Errorf("функция SetView, вернула ошибку: <%w>", err)
		}
		v.Title = data.NameTitle
		v.Editable = true
		v.Wrap = true
		v.Frame = true
		v.BgColor = gocui.ColorBlack
		v.SelBgColor = data.Color
		v.SelFgColor = gocui.ColorBlack

		v.Write([]byte(".........."))
	}

	return nil
}

// Отрисовка поля вывода. Возвращается ошибка.
//
// Параметры:
//
//	g - указатель на Gui
//	data - данные для отрисовки.
func DrawOutput(g *gocui.Gui, data DrawDataOutput) error {

	if v, err := g.SetView(data.NameOutput, data.StartX, data.StartY, data.StartX+data.Len, data.StartY+2); err != nil {
		if err != gocui.ErrUnknownView {
			return fmt.Errorf("функция SetView, вернула ошибку: <%w>", err)
		}
		v.Title = data.NameTitle
		v.Editable = false
		v.Wrap = true
		v.Frame = true
		v.BgColor = gocui.ColorDefault
		v.SelBgColor = gocui.ColorDefault
		v.SelFgColor = gocui.ColorDefault

		v.Write([]byte(""))
	}

	return nil
}

// Установка фокуса.
//
// Параметры:
//
//	v - указатель на View
//	focusName - имя элемента с фокусом.
//	name - имя текущего элемента.
func SetFocusStyle(v *gocui.View, focusName, name string) {
	if name == focusName {
		// Активное поле: яркая рамка + контрастное выделение текста
		v.Frame = true
		v.Highlight = true
		v.SelFgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorCyan
		v.FgColor = gocui.ColorCyan
		v.BgColor = gocui.ColorBlack
	} else {
		// Неактивное поле: сдержанный стиль
		v.Frame = true
		v.Highlight = false
		v.SelFgColor = gocui.ColorWhite
		v.SelBgColor = gocui.ColorDefault
		v.FgColor = gocui.ColorWhite
		v.BgColor = gocui.ColorBlack
	}
}

// Отрисовка пояснений. Возвращается ошибка.
//
// Параметры:
//
//	c - указатель на экземпляр сервиса.
//	data - данные для отрисовки.
func DrawGuide(g *gocui.Gui, data DrawDataGuide) error {

	xEnd := data.StartX + data.Weight
	yEnd := data.StartY + 2

	if v, err := g.SetView(data.NameGuide, data.StartX, data.StartY, xEnd, yEnd); err != nil {
		if err != gocui.ErrUnknownView {
			return fmt.Errorf("функция SetView, вернула ошибку: <%w>", err)
		}
		v.Editable = false
		v.Wrap = true
		v.Frame = true
		v.BgColor = gocui.ColorDefault
		v.FgColor = gocui.ColorWhite

		// Вычисление позиции для вывода текста.
		txt := []rune(data.Text)

		lenGuide := xEnd - data.StartX + 3
		if len(txt) > lenGuide {
			return fmt.Errorf("Превышена длинна текста")
		}

		pos := (lenGuide / 2) - (len(txt) / 2)
		if _, err := v.Write([]byte(fmt.Sprintf("%s%s", strings.Repeat(" ", pos-2), data.Text))); err != nil {
			return err
		}
	}

	return nil
}
