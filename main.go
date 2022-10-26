package main

import "fmt"

type profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func PowerUp(nilai profile, multiplier int) profile {
	var PowerUp profile

	PowerUp.Name = nilai.Name
	PowerUp.Health = nilai.Health + (nilai.Health * multiplier)
	PowerUp.Power = nilai.Power + (nilai.Power * multiplier)
	PowerUp.Exp = nilai.Exp + (nilai.Exp * multiplier)

	return PowerUp

}

func MakeProfile(Name string) profile {
	var profile profile

	if Name == "Sasuke" {
		profile.Name = "Sasuke"
	}
	if Name == "Sasuke" {
		profile.Health = 200
	}
	if Name == "Sasuke" {
		profile.Power = 100
	}
	if Name == "Sasuke" {
		profile.Exp = 0
	}

	//
	if Name == "Goku" {
		profile.Name = "Goku"
	}
	if Name == "Goku" {
		profile.Health = 400
	}
	if Name == "Goku" {
		profile.Power = 300
	}
	if Name == "Goku" {
		profile.Exp = 100
	}
	if Name == "Goku" {
		profile.Name = "Goku"
	}

	//

	if Name == "Naruto" {
		profile.Health = 150
	}
	if Name == "Naruto" {
		profile.Power = 200
	}
	if Name == "Naruto" {
		profile.Exp = 50
	}
	return profile
}

func main() {
	profile := MakeProfile("Goku")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power : ", profile.Power)
	fmt.Println("Exp : ", profile.Exp)
	fmt.Println("=== Heroes Power Up===")

	PowerUp := PowerUp(profile, 2)
	fmt.Println("Name : ", PowerUp.Name)
	fmt.Println("Health : ", PowerUp.Health)
	fmt.Println("Power : ", PowerUp.Power)
	fmt.Println("Exp : ", PowerUp.Exp)

}
