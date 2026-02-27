package main

import "fmt"

type notifier interface {
	notify() error
}

type user struct {
	name  string
	email string
}

func (u *user) notify() error {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
	return nil
}

type admin struct {
	user
	level string
}

func (a *admin) notify() error {
	fmt.Printf("Sending admin email to %s<%s> with level %s\n", a.name, a.email, a.level)
	return nil
}

func sendNotification(n notifier) error {
	return n.notify()
}

func main() {
	admin := &admin{
		user: user{
			name:  "Janet Jones",
			email: "janet@email.com",
		},
		level: "super",
	}
	sendNotification(admin)
	admin.user.notify()
	admin.notify()
}
