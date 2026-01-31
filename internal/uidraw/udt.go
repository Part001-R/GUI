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

// Данные для отрисовки ввода.
type DrawDataInput struct {
	Color     gocui.Attribute // Цвет заливки
	NameInput string          // Имя поля
	NameTitle string          // Имя заголовка
	Len       int             // Длиина поля
	StartX    int             // Cартовая координата Х
	StartY    int             // Стартовая координата Y
}

// Данные для отрисовки вывода.
type DrawDataOutput struct {
	NameOutput string // Имя поля
	NameTitle  string // Имя заголовка
	Len        int    // Длиина поля
	StartX     int    // Cартовая координата Х
	StartY     int    // Стартовая координата Y
}

// Данные для отриcовки пояснения.
type DrawDataGuide struct {
	NameGuide string // Имя поля
	Weight    int    // Ширина
	Text      string // Текст пояснения
	StartX    int    // Cартовая координата Х
	StartY    int    // Стартовая координата Y
}
