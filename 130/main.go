package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	var gasEngine = engine{
		electric: false,
	}

	var electricMotor = engine{
		electric: true,
	}

	impreza := vehicle{
		engine: engine{false},
		make:   "Subaru",
		model:  "Impreza",
		doors:  4,
		color:  "pearl white",
	}

	crosstrek := vehicle{
		engine: gasEngine,
		make:   "Subaru",
		model:  "Crosstrek",
		doors:  4,
		color:  "pearl white",
	}

	model3 := vehicle{
		engine: electricMotor,
		make:   "Tesla",
		model:  "Model 3",
		doors:  4,
		color:  "blue",
	}

	cars := []vehicle{
		impreza,
		crosstrek,
		model3,
	}
	for _, v := range cars {
		v.car()
	}
}

func (v vehicle) car() {
	fmt.Printf("%s %s with %d doors and color %s.\n", v.make, v.model, v.doors, v.color)
}
