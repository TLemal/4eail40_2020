package main

import "fmt"

type Shape interface {
	String() string
}

type Rectangle struct {
	Width  int
	Length int
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Square of width %d and length %d", r.Width, r.Length)
}

type Circle struct {
	Radius int
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle of radius %d", c.Radius)
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
	}
}
