package main

import (
	"fmt"
)

type task struct {
	title       string
	description string
	done        bool
}

type task_manager struct {
	tasks []*task
}

func (tm *task_manager) add_task(title, description string) {
	var new_task = &task{
		title:       title,
		description: description,
		done:        false,
	}
	tm.tasks = append(tm.tasks, new_task)
	fmt.Printf("Задача добавлена: %s\n", title)
}

func (tm *task_manager) mark_as_done(i int) {
	if i >= 0 && i < len(tm.tasks) {
		tm.tasks[i].done = true
		fmt.Printf("Задача выполнена: %s\n", tm.tasks[i].title)
	} else {
		fmt.Println("Некорректная задача")
	}
}

func (tm *task_manager) remove_task(i int) {
	if i >= 0 && i < len(tm.tasks) {
		tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
		fmt.Printf("Задача удалена: %s\n", tm.tasks[i].title)
	} else {
		fmt.Println("Некорректная задача")
	}
}

func (tm *task_manager) filter_tasks() []*task {
	completed := make([]*task, 0)
	for _, tsk := range tm.tasks {
		if tsk.done {
			completed = append(completed, tsk)
		}
	}
	fmt.Printf("Количество выполненных задач: %d\n", len(completed))
	return completed
}

func main() {
	task_mgr := task_manager{}
	task_mgr.add_task("Решить домашку", "Открыть тетрадь, прочитать задание, написать решение")
	task_mgr.add_task("Обед", "Взять тарелку и ложку, наложить, сесть за стол")
	task_mgr.mark_as_done(0)
	task_mgr.filter_tasks()
	task_mgr.remove_task(0)
}
