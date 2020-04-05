package db

// Информация.
type info struct {
	Text string
}

// Расписание.
type schedule struct {
	Weekday int
	Name    string
	Teacher string
	Number  string
	Time    string
}

// Получить расписание в string.
func (s schedule) toString() string {
	return "Предмет: " + s.Name + "\n" +
		"Преподаватель: " + s.Teacher + "\n" +
		"Время: " + s.Time + "\n" +
		"Ссылка: https://us04web.zoom.us/j/" + s.Number + "\n\n"
}

// Получить информацию в string.
func (i info) toString() string {
	return i.Text
}
