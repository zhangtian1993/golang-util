package db

import (
	"testing"
)

func TestLoadEnvOpenConfig(t *testing.T) {
	config := LoadEnvOpenConfig(&OpenConfig{
		Debug: true,
	})
	if config.Debug != true {
		t.Fatal("config debug error")
	}
	if LoadEnvOpenConfig(nil) == nil {
		t.Fatal("load env config with nil error")
	}
}

func TestOpenPostgres(t *testing.T) {
	config := LoadEnvOpenConfig(&OpenConfig{
		Debug: true,
	})
	db1, err1 := OpenPostgres(config)
	if err1 != nil {
		t.Fatal("err1", err1)
	}
	db1.Close()

	config.Port = "55432"
	_, err2 := OpenPostgres(config)
	if err2 == nil {
		t.Fatal("open wrong db port error")
	}
	t.Log("err2", err2)
}
