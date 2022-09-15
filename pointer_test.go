package canihave_test

import (
	"testing"

	"github.com/afmahmuda/canihave"
	"github.com/stretchr/testify/assert"
)

func TestGetValOr(t *testing.T) {

	t.Run("string_non_nil", func(t *testing.T) {
		arg := "marco"
		got := canihave.ValueOfOrDefault(&arg, "polo")
		assert.Equal(t, got, "marco")
	})

	t.Run("string_nil", func(t *testing.T) {
		got := canihave.ValueOfOrDefault(nil, "polo")
		assert.Equal(t, got, "polo")
	})

	t.Run("int_non_nil", func(t *testing.T) {
		var arg int = 99
		got := canihave.ValueOfOrDefault(&arg, 10)
		assert.Equal(t, got, 99)
	})

	t.Run("int_nil", func(t *testing.T) {
		got := canihave.ValueOfOrDefault(nil, int(10))
		assert.Equal(t, got, 10)
	})

	t.Run("struct_non_nil", func(t *testing.T) {

		type people struct {
			name string
			age  int
		}

		var arg = people{name: "barbara", age: 83}
		got := canihave.ValueOfOrDefault(&arg, people{name: "denden", age: 87})
		assert.Equal(t, got, people{name: "barbara", age: 83})
	})

	t.Run("struct_nil", func(t *testing.T) {

		type people struct {
			name string
			age  int
		}

		got := canihave.ValueOfOrDefault(nil, people{name: "denden", age: 87})
		assert.Equal(t, got, people{name: "denden", age: 87})

	})
}

func TestFromValue(t *testing.T) {

	t.Run("string", func(t *testing.T) {
		got := canihave.PointerOf("polo")
		want := "polo"
		assert.Equal(t, got, &want)
	})

	t.Run("int", func(t *testing.T) {
		got := canihave.PointerOf(44)
		want := 44
		assert.Equal(t, got, &want)
	})

	t.Run("struct", func(t *testing.T) {
		type people struct {
			name string
			age  int
		}
		got := canihave.PointerOf(people{name: "jojo", age: 18})
		want := &people{name: "jojo", age: 18}
		assert.Equal(t, got, want)
	})
}

func TestValueOfOrEmpty(t *testing.T) {
	t.Run("string_non_nil", func(t *testing.T) {
		arg := "marco"
		got := canihave.ValueOfOrEmpty(&arg)
		assert.Equal(t, got, "marco")
	})

	t.Run("string_nil", func(t *testing.T) {
		var arg *string
		got := canihave.ValueOfOrEmpty(arg)
		assert.Equal(t, got, "")
	})

	t.Run("int_non_nil", func(t *testing.T) {
		var arg int = 99
		got := canihave.ValueOfOrEmpty(&arg)
		assert.Equal(t, got, 99)
	})

	t.Run("int_nil", func(t *testing.T) {
		var arg *int
		got := canihave.ValueOfOrEmpty(arg)
		assert.Equal(t, got, 0)
	})

	t.Run("struct_non_nil", func(t *testing.T) {

		type people struct {
			name string
			age  int
		}

		var arg = &people{name: "barbara", age: 83}
		got := canihave.ValueOfOrEmpty(arg)
		assert.Equal(t, got, people{name: "barbara", age: 83})
	})

	t.Run("struct_nil", func(t *testing.T) {

		type people struct {
			name string
			age  int
		}

		var arg *people

		got := canihave.ValueOfOrEmpty(arg)
		assert.Equal(t, got, people{})

	})
}
