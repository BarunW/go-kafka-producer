package test

import (
	"fmt"
	es "source/cmd/eventSource"
	"testing"
)

func TestGenerateData(t *testing.T) {
	u, err := es.NewUserInteractionData()
	if err != nil {
		t.Fail()
	}
	fmt.Println(u)
}
