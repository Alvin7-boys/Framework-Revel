package controllers

import (
	"Tugas/app/models" // import package model
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) UserId() revel.Result {
	id, _ := strconv.Atoi(c.Params.Form.Get("id"))
	data, err := models.Users(id)
	if err != nil {
		revel.AppLog.Errorf("Failed to get data: %v", err)
		return c.RenderError(err)
	}
	return c.Render(data)
}

func (c App) GetAllUser() revel.Result {
	if err := c.Authenticate(); err != nil {
		return c.RenderJSON("Login Terlebih Dahulu")
	}
	users, err := models.GetAllUsers()
	if err != nil {
		// handling jika terdapat error saat mengambil data dari model
		revel.AppLog.Errorf("Failed to get all users: %v", err)
		return c.RenderError(err)
	}

	// kirim data ke tampilan
	return c.RenderJSON(users)
}
func (c App) DeleteUser(name string) revel.Result {
	err := models.DeleteUserByName(name)
	if err != nil {
		revel.AppLog.Errorf("Failed to delete user: %v", err)
		return c.RenderError(err)
	}
	return c.RenderJSON("User with name " + name + " has been deleted")
}
func (c App) UpdateUser(id int) revel.Result {
	// mengambil data dari request
	name := c.Params.Form.Get("name")
	age, _ := strconv.Atoi(c.Params.Form.Get("age"))
	address := c.Params.Form.Get("address")
	password := c.Params.Form.Get("password")
	email := c.Params.Form.Get("email")

	// memanggil fungsi update di dalam model
	user, err := models.UpdateUserByID(id, name, age, address, password, email)
	if err != nil {
		return c.RenderError(err)
	}
	response := models.NewSuccess("Update user success", user)
	return c.RenderJSON(response)
}
func (c App) CreateUser() revel.Result {
	name := c.Params.Form.Get("name")
	age, _ := strconv.Atoi(c.Params.Form.Get("age"))
	address := c.Params.Form.Get("address")
	password := c.Params.Form.Get("password")
	email := c.Params.Form.Get("email")
	id, _ := strconv.Atoi(c.Params.Form.Get("id"))
	// memanggil fungsi create di dalam model
	user, err := models.CreateUser(name, age, address, password, email, id)
	if err != nil {
		return c.RenderError(err)
	}
	response := models.NewSuccess("Create user success", user)
	return c.RenderJSON(response)
}
func (c App) Login() revel.Result {
	email := c.Params.Form.Get("email")
	password := c.Params.Form.Get("password")

	if email == "" {
		return c.RenderError(errors.New("Email harus diisi."))
	}
	if password == "" {
		return c.RenderError(errors.New("Password harus diisi."))
	}

	// Validasi email dan password di database
	err := models.ValidateUser(email, password)
	if err != nil {
		return c.RenderError(err)
	}

	// Generate token dan set cookie
	cookie, err := models.GenerateTokenAndCookies(email, password)
	if err != nil {
		return c.RenderError(err)
	}
	c.SetCookie(cookie)
	return c.RenderText("Login berhasil.")
}

func (c App) Authenticate() revel.Result {
	// ambil cookie dengan key "auth-token"
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		// jika cookie tidak ditemukan, redirect ke halaman login
		return c.RenderText(err.Error())
	}

	// validasi token dari cookie
	tokenString := cookie.GetValue()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// TODO: ganti dengan secret key yang benar
		return []byte("secret"), nil
	})

	if err != nil {
		// jika token tidak valid, redirect ke halaman login
		return c.RenderText(err.Error())

	}

	if token.Valid {
		// jika token valid, lanjutkan ke halaman yang diminta
		return nil
	} else {
		// jika token tidak valid, redirect ke halaman login
		return c.RenderText(err.Error())
	}

}
func (c App) Logout() revel.Result {
	// Hapus cookie dengan key "token"
	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})

	return c.RenderText("Logout berhasil.")
}
