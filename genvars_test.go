package genvars

import (
	"os"
	"testing"
)

func Test_CheckProductionMode(t *testing.T) {
	os.Setenv("GENVARS-TEST_APP_ENV", "PRODUCTION")

	m, err := NewManager("GENVARS-TEST")
	if err != nil {
		t.Error(err)
	}
	if !m.IsProduction() {
		t.Errorf("Wrong enviroment set, should be: PRODUCTION, Got: %v", m.Getenv("APP_ENV"))
	}
}

func Test_CheckDevelopmentMode(t *testing.T) {
	os.Setenv("GENVARS-TEST_APP_ENV", "DEVELOPMENT")

	m, err := NewManager("GENVARS-TEST")
	if err != nil {
		t.Error(err)
	}
	if !m.IsDevelopment() {
		t.Errorf("Wrong enviroment set, should be: DEVELOPMENT, got: %v", m.Getenv("APP_ENV"))
	}
}

func Test_GetProductionVar(t *testing.T) {
	os.Setenv("GENVARS-TEST_APP_ENV", "PRODUCTION")
	os.Setenv("TEST_ENV", "TEST")

	m, err := NewManager("GENVARS-TEST")
	if err != nil {
		t.Error(err)
	}
	if m.Getenv("TEST_ENV") != "TEST" {
		t.Errorf("Wrong var value, should be: TEST, got %v", m.Getenv("TEST_ENV"))
	}
}

func Test_GetDevelopmentVar(t *testing.T) {
	os.Setenv("GENVARS-TEST_APP_ENV", "DEVELOPMENT")
	os.Setenv("GENVARS-TEST_TEST_ENV", "TEST")

	m, err := NewManager("GENVARS-TEST")
	if err != nil {
		t.Error(err)
	}
	if m.Getenv("TEST_ENV") != "TEST" {
		t.Errorf("Wrong var value, should be: TEST, got %v", m.Getenv("TEST_ENV"))
	}
}

func Test_DevelopemtConfigurationOptions(t *testing.T) {
	os.Setenv("GENVARS-TEST_APP_ENV", "DEV")
	os.Setenv("GENVARS-TEST_TEST_ENV", "TEST")

	m, err := NewManager("GENVARS-TEST", ManagerOptions{
		DevTagValue:  "DEV",
		ProdTagValue: "PROD",
	})
	if err != nil {
		panic(err)
	}

	if m.Getenv("TEST_ENV") != "TEST" {
		t.Errorf("Wrong var value, should be: TEST, got %v", m.Getenv("TEST_ENV"))
	}
}

func Test_ProductionConfigurationOptions(t *testing.T) {
	os.Setenv("GENVARS-TEST_APP_ENV", "PROD")
	os.Setenv("GENVARS-TEST_TEST_ENV", "TEST")

	m, err := NewManager("GENVARS-TEST", ManagerOptions{
		DevTagValue:  "DEV",
		ProdTagValue: "PROD",
	})
	if err != nil {
		panic(err)
	}

	if m.Getenv("TEST_ENV") != "TEST" {
		t.Errorf("Wrong var value, should be: TEST, got %v", m.Getenv("TEST_ENV"))
	}
}
