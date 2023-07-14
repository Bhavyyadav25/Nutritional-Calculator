package main

import (
	"fmt"
)

func main(){
	ns := GetNutritionalScore(NutritionalData{
		Energy: EnergyFromKcal(1000),
		Sugars: SugarGram(50),
		SaturatedFattyAcids: SaturatedFattyAcids(70),
		Sodium: SodiumMilligram(500),
		Fruits: FruitsPercent(60),
		Fibre: FibreGram(4),
		Protien: ProtienGram(10),
	}, Food)

	fmt.Printf("Nutrition Score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}