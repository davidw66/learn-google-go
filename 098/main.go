package main

import "fmt"

func main() {
	xs := []string{"Almond Biscotti Café",
		"Banana Pudding ",
		"Balsamic Strawberry (GF)",
		"Bittersweet Chocolate (GF)",
		"Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)",
		"Browned Butter Cookie Dough",
		"Chocolate Covered Black Cherry (GF)",
		"Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)",
		"Coffee (GF)",
		"Cookies & Cream",
		"Fresh Basil (GF)",
		"Garden Mint Chip (GF)",
		"Lavender Lemon Honey (GF)",
		"Lemon Bar",
		"Madagascar Vanilla (GF)",
		"Matcha (GF)",
		"Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy Chocolate Peanut Butter (GF, V)",
		"Non-Dairy Coconut Matcha (GF, V)",
		"Non-Dairy Sunbutter Cinnamon (GF, V)",
		"Orange Cream (GF) ",
		"Peanut Butter Cookie Dough",
		"Raspberry Sorbet (GF, V)",
		"Salty Caramel (GF)",
		"Slate Slate Different",
		"Strawberry Lemonade Sorbet (GF, V)",
		"Vanilla Caramel Blondie",
		"Vietnamese Cinnamon (GF)",
		"Wolverine Tracks (GF)"}

	xs = append(xs, "The bomb")
	fmt.Println(xs)

	fmt.Println(len(xs), cap(xs))

	for _, v := range xs {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", xs)

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(xi)

	xi = append(xi[0:2], xi[3:5]...)
	fmt.Println(xi)

	xi2 := make([]int, 0, 100)
	xi2 = append(xi2, 1, 2, 3, 4)
	fmt.Println(xi2)
}
