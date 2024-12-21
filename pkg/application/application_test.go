package application

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalcHandler_Success(t *testing.T) {
	// Prepare a valid request body
	body := `{"expression": "2+2"}`
	expected := "result: 4.000000"
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", strings.NewReader(body))
	w := httptest.NewRecorder()
	// Call the handler
	CalcHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll((res.Body))
	if err != nil {
		t.Errorf("Error:%v", err)
	}

	// Check the response status code
	// Check the response body
	if string(data) != expected {
		t.Errorf("expected result: 4.000000 but got %v", string(data))
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, but got %d", http.StatusOK, w.Code)
	}
}
func TestCalcHandler_Success2(t *testing.T) {
	// Prepare a valid request body
	// Prepare a valid request body
	body := `{"expression": "1/2"}`
	expected := "result: 0.500000"
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", strings.NewReader(body))
	w := httptest.NewRecorder()
	// Call the handler
	CalcHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll((res.Body))
	if err != nil {
		t.Errorf("Error:%v", err)
	}

	// Check the response status code
	// Check the response body
	if string(data) != expected {
		t.Errorf("expected result: 0.500000 but got %v", string(data))
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status")
	}
}

func TestCalcHandler_InvalidExpression(t *testing.T) {
	// Prepare a request body with an invalid expression
	body := `{"expression": "invalid"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", strings.NewReader(body))
	w := httptest.NewRecorder()

	// Call the handler
	CalcHandler(w, req)

	// Check the response status code
	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status %d, but got %d", http.StatusUnprocessableEntity, w.Code)
	}

	// Check the response body
	expectedBody := "error: Expression is not valid"
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Errorf("expected body to contain %q, but got %q", expectedBody, w.Body.String())
	}
}

func TestCalcHandler_MethodNotAllowed(t *testing.T) {
	// Prepare a request with an invalid method (GET instead of POST)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/calculate", nil)
	w := httptest.NewRecorder()

	// Call the handler
	CalcHandler(w, req)

	// Check the response status code
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status %d, but got %d", http.StatusMethodNotAllowed, w.Code)
	}

	// Check the response body
	expectedBody := "error: Method not allowed"
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Errorf("expected body to contain %q, but got %q", expectedBody, w.Body.String())
	}
}

func TestCalcHandler_InternalServerError(t *testing.T) {
	body := `{"expression": "7/0"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", strings.NewReader(body))
	w := httptest.NewRecorder()

	// Call the handler
	CalcHandler(w, req)

	// Check the response status code
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, but got %d", http.StatusInternalServerError, w.Code)
	}

	// Check the response body
	expectedBody := "error: Internal server error"
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Errorf("expected body to contain %q, but got %q", expectedBody, w.Body.String())
	}
}