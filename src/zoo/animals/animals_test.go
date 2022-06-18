package animals

import (
	"testing"
)

// test1: ElephantEeed
func TestElephantFeed(t *testing.T) {
	expect := "Grass"
	actual := ElephantFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

// test2: MonkeyFeed
func TestMonkeyFeed(t *testing.T) {
	expect := "Banana"
	actual := MonkeyFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

// test3: RabbitFeed
func TestRabbitFeed(t *testing.T) {
	expect := "Carrot"
	actual := RabbitFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
