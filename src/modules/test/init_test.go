package test

import (
	"modules/app"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Init()
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
func TestInit(t *testing.T) {
	if app.Development.Testing == false {
		t.Fatal(app.Development.Testing)
	}
}
