// Представление графических элементов.
package uidraw

import "github.com/jroimartin/gocui"

// Данные для отрисовки кнопки.
type DrawDataButton struct {
	NameSquare  string          // Имя прямоугольника
	ColorSquare gocui.Attribute // Цвет заливки прямоугольника
	NameButton  string          // Имя кнопки
	TextButton  string          // Текст для кнопки
	StartX      int             // Cартовая координата Х
	StartY      int             // Стартовая координата Y
}
