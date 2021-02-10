package read_json

import (
	"os"
	"path"
	"reflect"
	"testing"
)

func getMock(filename string) string {
	pwdPath, _ := os.Getwd()
	mockPath := path.Join(pwdPath, "mocks", filename)
	return mockPath
}

func correctJsonTest(t *testing.T) {
	mockPath := getMock("correct.json")
	result, err := ReadJson(mockPath)
	if err != nil {
		t.Errorf("Read method failed with %s", err)
	}

	expected := map[string]interface{}{
		"lorem": "ipsum",
		"dolor": "sit",
	}

	isEqual := reflect.DeepEqual(result, expected)

	if !isEqual {
		t.Errorf("Expected result doesn't match. Expected %s | Received: %s", expected, result)
	}

}

func correctJsonWithIntTest(t *testing.T) {
	mockPath := getMock("with_int.json")
	result, err := ReadJson(mockPath)
	if err != nil {
		t.Errorf("Read method failed with %s", err)
	}

	expected := map[string]interface{}{
		"lorem": "ipsum",
		"dolor": "sit",
		"amet":  0.0,
	}

	isEqual := reflect.DeepEqual(result, expected)

	if !isEqual {
		t.Errorf("Expected result doesn't match. Expected %s | Received: %s", expected, result)
	}

}

func incorrectJsonTest(t *testing.T) {
	mockPath := getMock("incorrect.json")
	result, err := ReadJson(mockPath)
	if result != nil {
		t.Errorf("Value returned differs from nil. Received: %s", result)
	}

	if err == nil {
		t.Errorf("Error expected, got \"nil\"")
	}

	errorType := reflect.TypeOf(err).String()

	if errorType != "*json.SyntaxError" {
		t.Errorf("Incorrect error returned. Received: %s", errorType)
	}
}

func notJsonTest(t *testing.T) {
	mockPath := getMock("not_json")
	result, err := ReadJson(mockPath)
	if result != nil {
		t.Errorf("Value returned differs from nil. Received: %s", result)
	}

	if err == nil {
		t.Errorf("Error expected, got \"nil\"")
	}

	errorType := reflect.TypeOf(err).String()

	if errorType != "*json.SyntaxError" {
		t.Errorf("Incorrect error returned. Received: %s", errorType)
	}
}

func missingJsonTest(t *testing.T) {
	mockPath := getMock("not_existing_json.json")
	result, err := ReadJson(mockPath)
	if result != nil {
		t.Errorf("Value returned differs from nil. Received: %s", result)
	}

	if err == nil {
		t.Errorf("Error expected, got \"nil\"")
	}

	errorType := reflect.TypeOf(err).String()

	if errorType != "*os.PathError" {
		t.Errorf("Incorrect error returned. Received: %s", errorType)
	}
}

func TestReadJson(t *testing.T) {
	t.Run("Correct json file", correctJsonTest)
	t.Run("Correct json file with int", correctJsonWithIntTest)
	t.Run("Incorrect json file", incorrectJsonTest)
	t.Run("Not a Json file", notJsonTest)
	t.Run("Missing a Json file", missingJsonTest)
}
