package main

import "fmt"

type superhero struct {
	name      string
	power     int
	saveCount int
}

type supervillan struct {
	name         string
	power        int
	peopleharmed int
}

type HT interface {
	utility() int
}

func main() {
	superman := superhero{"Superman", 100, 20}
	batman := superhero{"Batman", 30, 4}

	superman.power = superman.utility()
	batman.power = batman.utility()

	println(superman.name, "has power", superman.power)
	println(batman.name, "has power", batman.power)

	// Create a supervillan
	joker := supervillan{"Joker", 10, 100}
	harley := supervillan{"Harley", 5, 20}

	joker.power = joker.utility()
	harley.power = harley.utility()

	println(joker.name, "has power", joker.power)
	println(harley.name, "has power", harley.power)

	HitSqaud := []HT{superman, batman, joker, harley}

	for _, h := range HitSqaud {
		fmt.Println("Utility of", h, "is", h.utility())
	}

	totalUtility := 0
	for _, h := range HitSqaud {
		totalUtility += h.utility()
	}
	fmt.Println("Total utility of HitSqaud is", totalUtility)

}

func (s superhero) utility() int {
	return s.saveCount * s.power
}

func (v supervillan) utility() int {
	return v.peopleharmed / v.power
}
