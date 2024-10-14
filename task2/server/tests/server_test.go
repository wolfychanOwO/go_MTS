package server

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	srv "main/task2/server"
)

func TestVersionHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	w := httptest.NewRecorder()

	srv.VersionHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK; got %v", res.Status)
	}

	body := w.Body.String()
	if body != "v1.0.0\n" {
		t.Fatalf("Expected version 'v1.0.0'; got %s", body)
	}
}

func TestDecodeHandler_Success(t *testing.T) {
	input := `{"inputString": "` + base64.StdEncoding.EncodeToString([]byte("hello")) + `"}`
	req := httptest.NewRequest(http.MethodPost, "/decode", strings.NewReader(input))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	srv.DecodeHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK; got %v", res.Status)
	}

	// Проверяем содержимое ответа
	var output map[string]string
	err := json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	expected := "hello"
	if output["outputString"] != expected {
		t.Fatalf("Expected '%s'; got '%s'", expected, output["outputString"])
	}
}

// Тест для /decode с некорректной строкой base64
func TestDecodeHandler_Fail(t *testing.T) {
	// Создаем запрос с невалидной base64 строкой
	input := `{"inputString": "invalid_base64"}`
	req := httptest.NewRequest(http.MethodPost, "/decode", strings.NewReader(input))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	srv.DecodeHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("Expected status BadRequest; got %v", res.Status)
	}
}

// Тест для /hard-op
func TestHardOpHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hard-op", nil)
	w := httptest.NewRecorder()

	start := time.Now()
	srv.HardOpHandler(w, req)
	duration := time.Since(start)

	res := w.Result()
	defer res.Body.Close()

	// Проверка, что запрос занял от 10 до 20 секунд
	if duration < 10*time.Second || duration > 20*time.Second {
		t.Fatalf("Expected delay between 10 and 20 seconds; got %v", duration)
	}

	// Проверка статуса ответа (может быть 200 или 500)
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Expected status 200 or 500; got %v", res.Status)
	}
}
