package stack

import (
	"reflect"
	"slices"
	"testing"
)

type customType string

var (
	linkedStackType  = reflect.TypeFor[linkedStack[customType]]()
	nodePointerType  = reflect.TypeFor[*node[customType]]()
	nodeType         = reflect.TypeFor[node[customType]]()
	customStringType = reflect.TypeFor[customType]()
)

var (
	headNameMatchFunc    = func(name string) bool { return name == "head" || name == "first" }
	elementNameMatchFunc = func(name string) bool { return name == "element" || name == "value" || name == "item" }
)

func TestNode_Fields(t *testing.T) {
	t.Run("Node has a field to store the next node", func(t *testing.T) {
		nextField, ok := nodeType.FieldByName("next")
		if !ok {
			t.Fatal(`node should have a field called "next"`)
		}

		nextType := nextField.Type
		if nextType.Kind() != reflect.Pointer {
			t.Fatal(`"next" field should be a pointer`)
		}

		if nextType.Elem() != nodePointerType.Elem() {
			t.Errorf(`"next" field should be a pointer of type "T" to a node struct`)
		}
	})

	t.Run("Node has a field to store the value", func(t *testing.T) {
		valueField, ok := nodeType.FieldByNameFunc(elementNameMatchFunc)
		if !ok {
			t.Fatal(`node should have a field called "value"`)
		}

		if valueField.Type != customStringType {
			t.Fatal(`The field to store the value should be of type "T"`)
		}
	})
}

func TestStack_Fields(t *testing.T) {
	t.Run("Stack has a field to store a reference to the first element", func(t *testing.T) {
		headField, ok := linkedStackType.FieldByNameFunc(headNameMatchFunc)
		if !ok {
			t.Fatal(`linkedStack should have a field called "head" or "first"`)
		}

		headType := headField.Type
		if headType.Kind() != reflect.Pointer {
			t.Fatal(`"head" (or "first") field should be a pointer`)
		}

		if headType.Elem() != nodePointerType.Elem() {
			t.Errorf(`"head" (or "first") should be a pointer of type "T" to a node struct`)
		}
	})

	t.Run("Stack has a field to store the size", func(t *testing.T) {
		sizeField, ok := linkedStackType.FieldByName("size")
		if !ok {
			t.Fatal(`linkedStack should have a field called "size"`)
		}
		if sizeField.Type.Kind() != reflect.Int {
			t.Errorf(`"size" should be of type "int"`)
		}
	})
}

func TestStack_Push(t *testing.T) {
	table := []struct {
		name           string
		elementsToPush []customType
		expected       []customType
		errMsg         string
	}{
		{
			name:           "Push adds element into head when stuck is empty",
			elementsToPush: []customType{"hello"},
			expected:       []customType{"hello"},
			errMsg: `Ensure that the pushed element is correctly assigned as the head's value 
when the stack is initially empty.`,
		},
		{
			name:           "Push correctly sets the next pointers and maintains the order",
			elementsToPush: []customType{"hello", "there", "day", "world"},
			expected:       []customType{"world", "day", "there", "hello"},
			errMsg: `unexpected order of elements in the stack.
Check that the 'next' pointers are correctly updated 
and the last pushed element is at the top of the stack`,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[customType]()
			for _, e := range tt.elementsToPush {
				stack.Push(e)
			}

			headField := reflect.ValueOf(stack).Elem().FieldByNameFunc(headNameMatchFunc)
			if !headField.IsValid() || headField.IsZero() {
				t.Fatal("head is not set. Ensure that the 'Push' method correctly updates the head pointer")
			}

			elementField := headField.Elem().FieldByNameFunc(elementNameMatchFunc)
			if !elementField.IsValid() {
				t.Fatal("value in head is not set. Check that the value is assigned to the head node")
			}

			head := reflectedHead(headField)
			stackSlice := slices.Collect(head.Walk)
			if !slices.Equal(stackSlice, tt.expected) {
				t.Error(tt.errMsg)
			}
		})
	}
}

type reflectedHead reflect.Value

func (h reflectedHead) Walk(yield func(element customType) bool) {
	h.walk(yield)
}

func (h reflectedHead) walk(yield func(element customType) bool) bool {
	value := reflect.Value(h)
	if value.IsNil() {
		return true
	}

	elementField := value.Elem().FieldByNameFunc(elementNameMatchFunc)
	if !elementField.IsValid() {
		return true
	}

	yield(getUnexportedField(elementField))
	value = value.Elem().FieldByName("next")
	reflectedHead(value).walk(yield)

	return false
}

func getUnexportedField(field reflect.Value) customType {
	value := reflect.NewAt(field.Type(), field.Addr().UnsafePointer())
	return value.Elem().Interface().(customType)
}
