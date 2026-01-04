package utils

type GenericArray[T any] []T

type AnyList = GenericArray[any]

func NewAnyList() AnyList {
	return AnyList{}
}

func (a *GenericArray[T]) Add(item T) {
	*a = append(*a, item)
}

func (a *GenericArray[T]) AddItems(items ...T) {
	*a = append(*a, items...)
}

func (a *GenericArray[T]) AddList(items []T) {
	if len(items) == 0 {
		return
	}
	*a = append(*a, items...)
}

func (a *GenericArray[T]) Len() int {
	return len(*a)
}

func (a *GenericArray[T]) Get(i int) T {
	return (*a)[i]
}

func (a *GenericArray[T]) Set(i int, item T) {
	(*a)[i] = item
}

func (a *GenericArray[T]) Clear() {
	*a = (*a)[:0]
}

func (a *GenericArray[T]) Slice(start, end int) []T {
	return (*a)[start:end]
}

func (a *GenericArray[T]) Copy() []T {
	return append([]T{}, *a...)
}

func (a *GenericArray[T]) IsEmpty() bool {
	return len(*a) == 0
}

func MergeLists(lists ...AnyList) AnyList {
	result := NewAnyList()
	for _, list := range lists {
		result.AddList(list)
	}

	return result
}
