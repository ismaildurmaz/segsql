package utils_test

import (
	"reflect"
	"testing"

	"github.com/ismaildurmaz/segsql/utils"
)

func TestGenericArray_Add(t *testing.T) {
	var a utils.GenericArray[int]

	a.Add(1)
	a.Add(2)

	if a.Len() != 2 {
		t.Fatalf("expected len=2, got %d", a.Len())
	}

	if a.Get(0) != 1 || a.Get(1) != 2 {
		t.Fatalf("unexpected values: %v", a)
	}
}

func TestGenericArray_AddItems(t *testing.T) {
	var a utils.GenericArray[string]

	a.AddItems("a", "b", "c")

	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(a.Copy(), expected) {
		t.Fatalf("expected %v, got %v", expected, a.Copy())
	}
}

func TestGenericArray_AddList(t *testing.T) {
	var a utils.GenericArray[int]

	a.AddList([]int{1, 2})
	a.AddList([]int{}) // should be no-op
	a.AddList([]int{3})

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(a.Copy(), expected) {
		t.Fatalf("expected %v, got %v", expected, a.Copy())
	}
}

func TestGenericArray_SetAndGet(t *testing.T) {
	var a utils.GenericArray[int]
	a.AddItems(1, 2, 3)

	a.Set(1, 99)

	if a.Get(1) != 99 {
		t.Fatalf("expected value at index 1 to be 99, got %d", a.Get(1))
	}
}

func TestGenericArray_Clear(t *testing.T) {
	var a utils.GenericArray[int]
	a.AddItems(1, 2, 3)

	a.Clear()

	if !a.IsEmpty() {
		t.Fatalf("expected array to be empty after Clear")
	}
}

func TestGenericArray_Slice(t *testing.T) {
	var a utils.GenericArray[int]
	a.AddItems(1, 2, 3, 4)

	s := a.Slice(1, 3)

	expected := []int{2, 3}
	if !reflect.DeepEqual(s, expected) {
		t.Fatalf("expected %v, got %v", expected, s)
	}
}

func TestGenericArray_Copy(t *testing.T) {
	var a utils.GenericArray[int]
	a.AddItems(1, 2, 3)

	c := a.Copy()
	c[0] = 99

	// original should not change
	if a.Get(0) != 1 {
		t.Fatalf("expected original slice to be unchanged, got %v", a.Copy())
	}
}

func TestGenericArray_IsEmpty(t *testing.T) {
	var a utils.GenericArray[int]

	if !a.IsEmpty() {
		t.Fatalf("expected empty array")
	}

	a.Add(1)

	if a.IsEmpty() {
		t.Fatalf("expected non-empty array")
	}
}

func TestMergeLists(t *testing.T) {
	a := utils.NewAnyList()
	a.AddItems(1, "a")

	b := utils.NewAnyList()
	b.AddItems(2, "b")

	c := utils.NewAnyList() // empty list should be ignored

	result := utils.MergeLists(a, b, c)

	expected := []any{1, "a", 2, "b"}
	if !reflect.DeepEqual(result.Copy(), expected) {
		t.Fatalf("expected %v, got %v", expected, result.Copy())
	}
}

func TestMergeLists_Empty(t *testing.T) {
	result := utils.MergeLists()

	if !result.IsEmpty() {
		t.Fatalf("expected empty result")
	}
}
