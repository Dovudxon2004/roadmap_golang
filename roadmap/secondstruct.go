package roadmap

import (
	"strconv"
)

type TaskList struct {
	ID        int
	name      string
	status    string
	priority  string
	createdAt string
	createdBy string
	dueDate   string
}

func NewTaskList(a int, b string, c string, d string, e string, f string, g string) *TaskList {
	return &TaskList{
		a,
		b,
		c,
		d,
		e,
		f,
		g,
	}
}

func (a *TaskList) Update(field string, value string) {
	switch field {
	case "ID":
		a.ID, _ = strconv.Atoi(value)
	case "name":
		a.name = value
	case "status":
		a.status = value
	case "priority":
		a.priority = value
	case "createdAt":
		a.createdAt = value
	case "createdBy":
		a.createdBy = value
	case "dueDate":
		a.dueDate = value
	}
}

func (a *TaskList) Get(field string) any {
	switch field {
	case "ID":
		return a.ID
	case "name":
		return a.name
	case "status":
		return a.status
	case "priority":
		return a.priority
	case "createdAt":
		return a.createdAt
	case "createdBy":
		return a.createdBy
	case "dueDate":
		return a.dueDate
	default:
		return 0
	}
}

func (a *TaskList) GetAll() (int, string, string, string, string, string, string) {
	return a.ID, a.name, a.status, a.priority, a.createdAt, a.createdBy, a.dueDate
}

func (a *TaskList) Delete(field string) {
	switch field {
	case "ID":
		a.ID = 0
	case "name":
		a.name = ""
	case "status":
		a.status = ""
	case "priority":
		a.priority = ""
	case "createdAt":
		a.createdAt = ""
	case "createdBy":
		a.createdBy = ""
	case "dueDate":
		a.dueDate = ""
	}
}
