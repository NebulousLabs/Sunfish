package main

import (
	"bytes"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"net/http"
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		sf.logger.Println("ERROR: Could not encode user, %s", newUser)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
