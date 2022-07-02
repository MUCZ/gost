package main

import (
	"gost/client"
	"testing"
)

func TestMain(m *testing.M) {
	// go server.Start()
}

func TestSet(t *testing.T) {
	ret, err := client.Post("test")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestGet(t *testing.T) {
	ret, err := client.Get("test")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
