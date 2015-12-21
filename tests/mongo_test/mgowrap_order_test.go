package tests

import (
	"os"
	"testing"
)

func setup() {
	println("setup")
}

func teardown() {
	println("teardown")
}

func TestMain(m *testing.M) {

	setup()

	retCode := m.Run()

	teardown()

	os.Exit(retCode)
}

func TestA(t *testing.T) {
	println("A")
}

func TestZ(t *testing.T) {
	println("Z")
}

func TestD(t *testing.T) {
	println("D")
}
