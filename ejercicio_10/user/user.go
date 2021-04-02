package user

import "time"

type Users struct {
	Id       int
	Name     string
	DateTime time.Time
	Status   bool
}

// *User es un puntero a la estructura y se hace refencia con this
func (this *Users) AddUser(id int, name string, datetime time.Time, status bool) {
	this.Id = id
	this.Name = name
	this.DateTime = datetime
	this.Status = status
}
