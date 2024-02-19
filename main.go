package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// RandomWalk представляет собой структуру для генерации случайного блуждания.
type RandomWalk struct {
	NumPoints int       // NumPoints - количество точек, которые необходимо сгенерировать.
	XValues   []float64 // XValues - координаты X точек случайного блуждания.
	YValues   []float64 // YValues - координаты Y точек случайного блуждания.
}

// FillWalk заполняет случайное блуждание точками.
func (rw *RandomWalk) FillWalk() {
	for len(rw.XValues) < rw.NumPoints {
		// Генерация направления и длины шага по оси X.
		xDirection := rand.Intn(2)*2 - 1
		xDistance := rand.Float64() * 0.5 // Уменьшаем диапазон шагов для X
		xStep := float64(xDirection) * xDistance

		// Генерация направления и длины шага по оси Y.
		yDirection := rand.Intn(2)*2 - 1
		yDistance := rand.Float64() * 0.5 // Уменьшаем диапазон шагов для Y
		yStep := float64(yDirection) * yDistance

		// Проверка, что шаг не равен нулю по обеим осям.
		if xStep == 0 && yStep == 0 {
			continue
		}

		// Вычисление следующей координаты X и Y.
		nextX := rw.XValues[len(rw.XValues)-1] + xStep
		nextY := rw.YValues[len(rw.YValues)-1] + yStep

		// Добавление координат в списки XValues и YValues.
		rw.XValues = append(rw.XValues, nextX)
		rw.YValues = append(rw.YValues, nextY)
	}
}

func main() {
	// Создание объекта RandomWalk с заданными параметрами.
	rw := RandomWalk{
		NumPoints: 200000,       // Количество точек для генерации.
		XValues:   []float64{0}, // Начальное значение координаты X.
		YValues:   []float64{0}, // Начальное значение координаты Y.
	}
	rw.FillWalk() // Заполнение случайного блуждания точками.

	// Создание нового графика.
	p := plot.New()

	// Установка заголовка и подписей осей.
	p.Title.Text = "Random Walk"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Создание точек для графика.
	points := make(plotter.XYs, len(rw.XValues))
	for i := range points {
		points[i].X = rw.XValues[i]
		points[i].Y = rw.YValues[i]
	}

	// Создание графического элемента Scatter и добавление его к графику.
	s, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	s.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255} // Установка цвета точек в синий.
	p.Add(s)

	// Сохранение графика в файл.
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "mygraph.png"); err != nil {
		panic(err)
	}

	// Вывод сообщения о сохранении файла.
	fmt.Println("Graph saved as mygraph.png")
}
