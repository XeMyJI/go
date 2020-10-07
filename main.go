package main

import (
	"errors"
	"fmt"
	"math"
)

const pi = 3.1415

func main() {
	printCircleArea(2)
	printRectangleArea(4,4)
	printTriangle(13,14,15)
}

func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Радиус круга: %d см.\n", radius)
	fmt.Println("Формула для расчета площади круга: A=πr2\n")

	fmt.Printf("Площадь круга: %f см. кв. \n------------\n\n", circleArea)
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус круга не может быть отрицательным!\n----------")
	}

	return float32(radius) * float32(radius) * pi, nil
}

func printRectangleArea(a,b int){
	S, err := calcRectangleArea(a,b)
	P, err := calcRectanglePerimeter(a,b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Прямоугольник со сторонами a, b: \n Сторона а:=", a)
	fmt.Println("Сторона b:=" , b, "\n")
	fmt.Printf("Площадь прямоугольника: %.2f см2", S)
	fmt.Printf("\nПериметр прямоугольника: %.2f см", P)
}

func calcRectangleArea(a,b int)(float32, error) {
	if a <= 0  {
		return float32(0), errors.New("Длина стороны 'a' не должно быть отрицательным и равной 0!")
	}
	if b <= 0 {
		return float32(0), errors.New("Длина стороны 'b' не должно быть отрицательным и равной 0!")
	}
	return float32(a) * float32(b),nil
}

func calcRectanglePerimeter(a,b int)(float32, error) {
	if a <= 0  {
		return float32(0), errors.New("Длина стороны 'a' не должно быть отрицательным и равной 0!")
	}
	if b <= 0 {
		return float32(0), errors.New("Длина стороны 'b' не должно быть отрицательным и равной 0!")
	}
	return float32(a) + float32(b) * 2,nil
}

func printTriangle(a,b,c int) {
	Pt, err := calcTrianglePerimeter(a,b,c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Треугольник со стораними a:%v b:%v c:%v", a, b, c)
	fmt.Printf("Периметр треугольника P: %f см.", Pt)
}

func calcTrianglePerimeter(a, b, c int) (float32, error){
	if a <= 0  {
		return float32(0), errors.New("Длина стороны 'a' не должно быть отрицательным и равной 0!")
	}

	return 
}