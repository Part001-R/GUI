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
