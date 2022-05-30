package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("can't add shape: out of the shapesCapacity range")
	}

	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i > b.shapesCapacity-1 {
		return nil, errors.New("can't get shape by Index: it would go out of the shapesCapacity range")
	} else if b.shapes[i] == nil {
		return nil, errors.New("can't get shape by Index: such shape doesn't exist")
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	returnedShape, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("extractByIndex error: %w", err)
	}

	// b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	b.shapes[i] = b.shapes[len(b.shapes)-1]
	b.shapes = b.shapes[:len(b.shapes)-1]
	return returnedShape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	returnedShape, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("extractByIndex error: %w", err)
	}

	b.shapes[i] = shape
	return returnedShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var result float64
	for _, v := range b.shapes {
		result += v.CalcPerimeter()
	}

	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var result float64
	for _, v := range b.shapes {
		result += v.CalcArea()
	}

	return result
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var result []Shape
	for _, v := range b.shapes {
		if _, ok := v.(*Circle); ok {
			continue
		} else if _, ok = v.(Circle); ok {
			continue
		}
		result = append(result, v)
	}

	if len(result) == len(b.shapes) {
		return errors.New("circles are not exist in the list")
	}

	b.shapes = result
	return nil
}
