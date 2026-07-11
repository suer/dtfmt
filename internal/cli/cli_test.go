package cli

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestRunTimestamp(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := Run([]string{"1700000000"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("code = %d, want 0, stderr = %s", code, stderr.String())
	}

	var doc map[string]interface{}
	if err := json.Unmarshal(stdout.Bytes(), &doc); err != nil {
		t.Fatalf("invalid JSON output: %v", err)
	}
	input, ok := doc["input"].(map[string]interface{})
	if !ok {
		t.Fatalf("missing input field in %v", doc)
	}
	if input["type"] != "timestamp" {
		t.Errorf("input.type = %v, want timestamp", input["type"])
	}
}

func TestRunHelp(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := Run([]string{"-h"}, &stdout, &stderr)
	if code != 0 {
		t.Errorf("code = %d, want 0", code)
	}
}

func TestRunNoArgs(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := Run([]string{}, &stdout, &stderr)
	if code != 2 {
		t.Errorf("code = %d, want 2", code)
	}
}

func TestRunTooManyArgs(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := Run([]string{"a", "b"}, &stdout, &stderr)
	if code != 2 {
		t.Errorf("code = %d, want 2", code)
	}
}

func TestRunUnrecognizedInput(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := Run([]string{"not a valid input at all"}, &stdout, &stderr)
	if code != 1 {
		t.Errorf("code = %d, want 1", code)
	}
	if stderr.Len() == 0 {
		t.Error("expected stderr message on failure")
	}
}
