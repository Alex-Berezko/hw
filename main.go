package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID     int
	name   string
	status string
}

var allTask []Task

func main() {

	task1 := Task{1, "Создать в коде несколько задач (хардкод)", "новая"}
	task2 := Task{2, "Реализовать вывод списка всех задач в консоль", "в работе"}
	task3 := Task{3, "Реализовать фильтрацию задач по статусу", "завершена"}
	scanner := bufio.NewScanner(os.Stdin)
	allTask = []Task{task1, task2, task3}
	AllTasks(allTask)
	Menu(allTask, scanner)

}

func FilteringStatus(task []Task, scanner *bufio.Scanner) {
	result := []string{}

	fmt.Println("Вы вошли в меню фильтрования по статусу")
	fmt.Println("Новая, в работе, завершено")

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Ошибка чтения:", err)
		}
	}

	for i := 0; i < len(task); i++ {
		if task[i].status == strings.ToLower(scanner.Text()) {
			result = append(result, task[i].name)
		}
	}
	if len(result) != 0 {
		fmt.Printf("cписок задач - %s\n", result)
	} else {
		fmt.Println("Задачи с таким статусом не найдены")
		fmt.Println("Доступные статусы: новая, в работе, завершена")
	}
}

func AllTasks(task []Task) {
	for i := 0; i < len(task); i++ {
		fmt.Printf("номер %v, название %s,  статус %s\n", task[i].ID, task[i].name, task[i].status)
	}
}

func Menu(allTask []Task, scanner *bufio.Scanner) {

	for {
		fmt.Println("Меню:")
		fmt.Println("1. Показать все")
		fmt.Println("2. Фильтровать по статусу")
		fmt.Println("3. Выйти")
		fmt.Println("Введите число:")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Ошибка чтения:", err)
			}
			return
		}

		switch scanner.Text() {
		case "1":
			AllTasks(allTask)
		case "2":
			FilteringStatus(allTask, scanner)
		case "3":
			os.Exit(0)
		}
	}

}
