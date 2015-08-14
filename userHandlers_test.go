package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	logDir = "testLogs"
	dbName = "testSunfish"
	sf     *Sunfish
	url    string
)

func TestAddUser(t *testing.T) {
	var user User
	userString := `{
		"username": "testUser",
		"email": "test@example.com",
		"fullname": "Test User",
		"password": "password1234"
	}`

	req, err := http.NewRequest("POST", url+"user/", strings.NewReader(userString))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		sf.logger.Println(err)
	}

	res, err := http.DefaultClient.Do(req)

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Error: expected %d. Received: %d", http.StatusCreated, res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &user)
	if err != nil {
		t.Errorf("Error unmarshalling JSON into user. %s %s", body, err)
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
