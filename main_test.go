package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber"
)

func TestGetIndex(t *testing.T) {

	app := fiber.New()
	app.Get("/", index)

	req, _ := http.NewRequest("GET", "/", nil)
	res, err := app.Test(req, -1)

	if err != nil {
		t.Errorf("error sending test request: %s", err.Error())
	}

	if res.StatusCode != 200 {
		t.Errorf("expected status code %d, received %d", 200, res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	output := make(map[string]string)
	err = decoder.Decode(&output)

	if err != nil {
		t.Errorf("error decoding response body: %s", err.Error())
	}

	expected := "Olá Mundo"
	if output["message"] != expected {
		t.Errorf("index() = %s; want %s", expected, output["message"])
	}
}
