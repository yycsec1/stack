package stack

type Stack[T any] struct {
	items []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) Push(item T) {
	stack.items = append(stack.items, item)
}

func (stack *Stack[T]) Pop() (T, bool) {
	var returnValue T
	length := len(stack.items)
	if length > 0 {
		returnValue, stack.items = stack.items[length-1], stack.items[:(length-1)]
		return returnValue, true
	}
	return returnValue, false
}

func (stack *Stack[T]) Peek() (T, bool) {
	var returnValue T
	length := len(stack.items)
	if length > 0 {
		returnValue = stack.items[length-1]
		return returnValue, true
	}
	return returnValue, false
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

func (stack *Stack[T]) Len() int {
	return len(stack.items)
}
