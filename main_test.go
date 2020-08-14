package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber"
)

func makeGetRequestToFunc(t *testing.T, path string, handler fiber.Handler) (*http.Response, error) {
	app := fiber.New()
	app.Get(path, handler)

	req, _ := http.NewRequest("GET", path, nil)
	res, err := app.Test(req, -1)

	if err != nil {
		t.Errorf("error sending test request: %s", err.Error())
		return nil, err
	}

	return res, err
}

func TestGetIndex(t *testing.T) {

	res, err := makeGetRequestToFunc(t, "/", index)

	if res.StatusCode != 200 {
		t.Errorf("expected status code %d, received %d", 200, res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	output := make(map[string]string)
	err = decoder.Decode(&output)

	if err != nil {
		t.Errorf("error decoding response body: %s", err.Error())
	}

	expected := "Olá TDC 2020"
	if output["message"] != expected {
		t.Errorf("index() = %s; want %s", expected, output["message"])
	}
}

func TestGetAbout(t *testing.T) {

	res, err := makeGetRequestToFunc(t, "/about", about)

	if res.StatusCode != 200 {
		t.Errorf("expected status code %d, received %d", 200, res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	output := make(map[string]string)
	err = decoder.Decode(&output)

	if err != nil {
		t.Errorf("error decoding response body: %s", err.Error())
	}

	expected := "Um grande evento online de desenvolvimento"
	if output["message"] != expected {
		t.Errorf("about() = %s; want %s", expected, output["message"])
	}
}
