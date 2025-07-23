package main

import (
	"testing"
)

func Test_generateComponents(t *testing.T) {
	_, _, err := parseOpenAPISpec("testdata/openapi.yaml")
	if err != nil {
		t.Errorf("parseOpenAPISpec() error = %v", err)
		return
	}
	_, err = generateComponents(spec, "testdata", false)
	if err != nil {
		t.Errorf("generateComponents() error = %v", err)
		return
	}
}

func Test_generateClient(t *testing.T) {
	_, _, err := parseOpenAPISpec("testdata/openapi.yaml")
	if err != nil {
		t.Errorf("parseOpenAPISpec() error = %v", err)
		return
	}
	_, err = generateClient(spec, "testdata", false)
	if err != nil {
		t.Errorf("generateClient() error = %v", err)
		return
	}
}
