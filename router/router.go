package router

import (
	"customer-onboarding/handlers"
	"net/http"
)

func SetUpRoutes() {
	groupsHandler := http.HandlerFunc(handlers.GroupsHandler)
	usersHandler := http.HandlerFunc(handlers.UsersHandler)
	http.Handle("/groups/", groupsHandler)
	http.Handle("/users/", usersHandler)
}
