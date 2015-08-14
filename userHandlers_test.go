package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	logDir = "testLogs"
	dbName = "testSunfish"
	sf     *Sunfish
	url    string
	client http.Client
)

func addUser(userInfo *UserSignup) (resp *http.Response, err error) {
	marshalledJson, _ := json.Marshal(userInfo)
	req, err := http.NewRequest("POST", url+"user/", bytes.NewReader(marshalledJson))
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func TestAddUser(t *testing.T) {
	var user User
	userInfo := UserSignup{
		Email:    "test@example.com",
		Username: "testUser",
		Fullname: "Test User",
		Password: "password1234",
	}
	res, err := addUser(&userInfo)

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Error: expected %d. Received: %d", http.StatusCreated, res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &user)
	if err != nil {
		t.Errorf("Error unmarshalling JSON into user. %s %s", body, err)
	}
}

// TestBlankUserName test that servers sends bad request with no password field
func TestBlankUserName(t *testing.T) {
	userInfo := UserSignup{
		Email:    "test1@example.com",
		Username: "",
		Fullname: "Test User 1",
		Password: "password1234",
	}
	res, err := addUser(&userInfo)

	if err != nil {
		t.Errorf("Error adding user: %s", userInfo)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Error: expected %d. Received: %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestBlankEmail(t *testing.T) {
	userInfo := UserSignup{
		Email:    "",
		Username: "testUser2",
		Fullname: "Test User 2",
		Password: "password1234",
	}
	res, err := addUser(&userInfo)

	if err != nil {
		t.Errorf("Error adding user: %s", userInfo)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Error: expected %d. Received: %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestBlankPassword(t *testing.T) {
	userInfo := UserSignup{
		Email:    "test3@example.com",
		Username: "testUser3",
		Fullname: "Test User 3",
		Password: "",
	}
	res, err := addUser(&userInfo)

	if err != nil {
		t.Errorf("Error adding user: %s", userInfo)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Error: expected %d. Received: %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestMain(m *testing.M) {
	sf = NewSunfish(logDir, dbName)

	// Drop db so we have a clean db for testing
	sf.DB.DropDatabase()
	server := httptest.NewServer(sf.Router)

	url = server.URL + "/api/"

	err := m.Run()

	sf.Close()
	server.Close()
	os.Exit(err)
}
