package go_happy_birthday

import (
	"net/http"
	"time"
)

type Employee struct {
	ID          int
	Name        string
	Birthday    time.Time
	Email       string
	Telegram    string
	Subscribers []int
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	// Здесь должен быть код для получения списка сотрудников из базы данных
	// и возвращения его в формате JSON.
}

func subscribe(w http.ResponseWriter, r *http.Request) {
	// Здесь должен быть код для подписки/отписки от оповещения о дне рождения.
}

func notifyBirthday() {
	// Здесь должен быть код для оповещения о дне рождения.
}
