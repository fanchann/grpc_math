package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestAdd(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:8080/add/1/1", nil)

	client := http.Client{}
	response, _ := client.Do(request)
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	success, ok := responseBody["success"].(bool)
	if !ok || !success {
		t.Fatalf("Expected success: true, but got: %v", responseBody["success"])
	}

	result, ok := responseBody["data"].(map[string]interface{})["result"].(float64)
	if !ok || result != 2 {
		t.Fatalf("Expected result: 2, but got: %v", result)
	}
}

func TestMultiply(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:8080/mul/2/2", nil)

	client := http.Client{}
	response, _ := client.Do(request)
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	success, ok := responseBody["success"].(bool)
	if !ok || !success {
		t.Fatalf("Expected success: true, but got: %v", responseBody["success"])
	}

	result, ok := responseBody["data"].(map[string]interface{})["result"].(float64)
	if !ok || result != 4 {
		t.Fatalf("Expected result: 4, but got: %v", result)
	}
}

func TestDivide(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:8080/div/20/4", nil)

	client := http.Client{}
	response, _ := client.Do(request)
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	success, ok := responseBody["success"].(bool)
	if !ok || !success {
		t.Fatalf("Expected success: true, but got: %v", responseBody["success"])
	}

	result, ok := responseBody["data"].(map[string]interface{})["result"].(float64)
	if !ok || result != 5 {
		t.Fatalf("Expected result: 5, but got: %v", result)
	}
}

func TestReduce(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:8080/red/50/40", nil)

	client := http.Client{}
	response, _ := client.Do(request)
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	success, ok := responseBody["success"].(bool)
	if !ok || !success {
		t.Fatalf("Expected success: true, but got: %v", responseBody["success"])
	}

	result, ok := responseBody["data"].(map[string]interface{})["result"].(float64)
	if !ok || result != 10 {
		t.Fatalf("Expected result: 10, but got: %v", result)
	}
}
