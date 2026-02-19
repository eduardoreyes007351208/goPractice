package main

import "fmt"

type cube struct {
	len int
}

func (s *cube) area () int {
	return s.len * s.len * 6
}

func (s *cube) volume () int {
	return s.len * s.len * s.len
}

func changeSize (c *cube, side int) {
	c.len = side
}

func main () {
		s := cube {len: 4}

		fmt.Printf("the sides of the cube are: %d inches\n", s.len)
		fmt.Printf("area: %d inches squared\n", s.area())
		fmt.Printf("volume: %d inches cubed\n", s.volume())

		for i := 5; i < 10; i++ {
			changeSize(&s, i)
			fmt.Printf("new side length is: %d inches\n", s.len)
			fmt.Printf("new area: %d inches squared\n", s.area())
			fmt.Printf("new volume: %d inches cubed\n", s.volume())
			
		} 

}