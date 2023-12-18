package optional

import (
	"fmt"
	"reflect"
)

type Optional[T any] struct {
	value *T
}

func NewOptional[T any](value T) *Optional[T] {
	return &Optional[T]{
		value: &value,
	}
}

func EmptyOptional[T any]() *Optional[T] {
	return &Optional[T]{}
}

func (o *Optional[T]) IsPresent() bool {
	return o.value != nil
}

func (o *Optional[T]) IsEmpty() bool {
	return o.value == nil
}

func (o *Optional[T]) Get() (T, error) {
	if o.value == nil {
		valueType := reflect.TypeOf(o.value).Elem()
		zeroValue := reflect.Zero(valueType).Interface().(T)
		return zeroValue, fmt.Errorf("optional is empty")
	} else {
		return *o.value, nil
	}
}

func (o *Optional[T]) GetOrElse(defaultValue T) T {
	if o.value == nil {
		return defaultValue
	}
	return *o.value
}
