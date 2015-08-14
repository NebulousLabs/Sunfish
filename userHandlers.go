package main

import (
	"bytes"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func (sf *Sunfish) AddUser(w http.ResponseWriter, r *http.Request) {
	var newUserInfo UserSignup
	var newUser User
	const maxUserSize = 1 << 10
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxUserSize))
	r.Body.Close()

	if err != nil {
		sf.logger.Println("ERROR: Could not read body of request.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &newUserInfo); err != nil {
		sf.logger.Println("ERROR could not unmarshall JSON into userInfo. %s", body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser.Id = bson.NewObjectId()
	newUser.Active = true
	newUser.Admin = false

	// Check for empty fields
	if newUserInfo.Username == "" || newUserInfo.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check password meets requirements
	// Requirements will be min length of 8 and contain alpha and number for now
	var alpha = "abcdefghijklmnopqrstuvwxyz"
	var numbers = "1234567890"
	var password = newUserInfo.Password

	if len(password) < 8 || !strings.ContainsAny(password, alpha) || !strings.ContainsAny(password, numbers) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check for used username and emails
	// TODO switch to count instead of actually getting the users
	var usernameUsers []User
	var emailUsers []User
	err = sf.DB.C("users").Find(
		bson.M{"username": newUserInfo.Username}).All(&usernameUsers)
	err = sf.DB.C("users").Find(
		bson.M{"email": newUserInfo.Email}).All(&emailUsers)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// if we found an account with the same email or username fail
	if len(usernameUsers) > 0 || len(emailUsers) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser.Username = newUserInfo.Username
	newUser.Email = newUserInfo.Email
	newUser.Fullname = newUserInfo.Fullname
	// Use bcrypt to hash password at level 10 difficulty
	passwordBytes := bytes.NewBufferString(newUserInfo.Password)
	passwordHash, err := bcrypt.GenerateFromPassword(passwordBytes.Bytes(), 10)
	newUser.PasswordHash = passwordHash

	err = sf.DB.C("users").Insert(newUser)
	if err != nil {
		sf.logger.Println("ERROR: Could not save new user record into database.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sf.logger.Println("Successfully added a new user. %s", newUser.Id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		sf.logger.Println("ERROR: Could not encode user, %s", newUser)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
