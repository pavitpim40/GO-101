package service_test

import (
	"name-of-package/service"
	"testing"
)

func TestInput1ShouldBe1(t *testing.T) {
	expected := "1"
	actual := service.Fizzbuzz(1)
	if actual != expected {
		t.Errorf("%s is not equal to %s", actual, expected)
	}
}