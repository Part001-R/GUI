package handler

import (
	"A/internal/logfile"
	"log"
	"time"

	"github.com/jroimartin/gocui"
)

// Представление объекта.
type DataObjCounter struct {
	Name        string           // Имя объекта
	NameEnStart string           // Индикатор разрешения обработки
	NameEnStop  string           // Индикатор разрешения обработки
	Status      int              // Статус объекта
	Value       int              // Значение
	ChCmd       chan int         // Команда.
	LogFile     *logfile.LogFile // Логгер
}

// Запуск объекта.
func (d *DataObjCounter) Start(g *gocui.Gui, v *gocui.View) error {

	if err := d.LogFile.Write("Принята команда СТАРТ"); err != nil {
		log.Fatalf("Ошибка добавления записи в лог")
	}

	if d.Status != StageStop {
		return nil
	}
	d.Status = StageRun

	go runCounter(d)

	return nil
}

// Остановка объекта.
func (d *DataObjCounter) Stop(g *gocui.Gui, v *gocui.View) error {

	if err := d.LogFile.Write("Принята команда СТОП"); err != nil {
		log.Fatalf("Ошибка добавления записи в лог")
	}

	if d.Status != StageRun {
		return nil
	}
	d.ChCmd <- CmdStop

	return nil
}

// Функционал объекта.
func runCounter(d *DataObjCounter) {

	if err := d.LogFile.Write("Обработчик начал работу"); err != nil {
		log.Fatalf("Ошибка добавления записи в лог")
	}

	value := 0
	tck := time.NewTicker(1 * time.Second)
	defer tck.Stop()

	done := false
	for !done {
		select {
		case <-tck.C:
			value++
			d.Value = value

		case cmd := <-d.ChCmd:
			if cmd == CmdStop {
				done = true
			}
		}
	}

	d.Status = StageStop

	if err := d.LogFile.Write("Обработчик завершил работу"); err != nil {
		log.Fatalf("Ошибка добавления записи в лог")
	}
}
