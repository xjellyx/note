package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/dashboard", dashboardPage)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<html><head><title>Home Page</title></head><body><h1>Welcome to my website!</h1><p><a href="/login">Click here to login</a></p></body></html>`)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<html><head><title>Login Page</title></head><body><h1>Login to my website</h1><form><input type="text" name="username" placeholder="Username"><br><br><input type="password" name="password" placeholder="Password"><br><br><input type="submit" value="Login"></form></body></html>`)
}

func dashboardPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<html><head><title>Dashboard</title></head><body><h1>Welcome to your dashboard</h1><p>Here you can manage your account.</p><ul><li><a href="#">Profile</a></li><li><a href="#">Settings</a></li><li><a href="#">Logout</a></li></ul></body></html>`)
}
