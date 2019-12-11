package handlers

import "testing"

func Test_initHandlers(t *testing.T) {
	if err := InitHandlers("../configs/conf.example.json"); err != nil {
		t.Fatal(err)
	}
	t.Log(Handlers)
}
