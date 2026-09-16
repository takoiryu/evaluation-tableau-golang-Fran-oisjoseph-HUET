package main

import "fmt"

func main() {
	equipe := [6]Soldat{
		{"Arthas", 1200, 250},
		{"Kael", 850, 320},
		{"Thrall", 1500, 180},
		{"Sylvanas", 700, 400},
		{"Garrosh", 1000, 280},
		{"Jaina", 500, 450},
	}
	afficherEquipe(equipe)
	vieplus := trouverPlusDeVie(equipe)
	fmt.Println("Le soldat avec le plus de vie est:", vieplus.nom, "avec :", vieplus.vie, "pv")
	atkplus := trouverPlusDAttaque(equipe)
	fmt.Println("Le soldat avec le plus d'attaque est:", atkplus.nom, "avec :", atkplus.attaque, "d'attaque")
	moyenne := calculerVieMoyenne(equipe)
	fmt.Println("Vie moyenne :", moyenne)
	nbinf := compterFaibles(equipe)
	fmt.Println("Soldats avec moins de 800 PV :", nbinf)
}

type Soldat struct {
	nom     string
	vie     int
	attaque int
}

func afficherEquipe(equipe [6]Soldat) {
	fmt.Println("=== ÉQUIPE ===")
	for i := 0; i < len(equipe); i++ {
		fmt.Println(equipe[i].nom)
		fmt.Println("vie :", equipe[i].vie)
		fmt.Println("atk :", equipe[i].attaque)
		fmt.Printf("\n")
	}
}
func trouverPlusDeVie(equipe [6]Soldat) Soldat {
	mvi := equipe[0]
	for i := 0; i < len(equipe); i++ {
		if equipe[i].vie > mvi.vie {
			mvi = equipe[i]
		}
	}
	return mvi
}
func trouverPlusDAttaque(equipe [6]Soldat) Soldat {
	matk := equipe[0]
	for i := 0; i < len(equipe); i++ {
		if equipe[i].attaque > matk.attaque {
			matk = equipe[i]
		}
	}
	return matk
}
func calculerVieMoyenne(equipe [6]Soldat) float64 {
	somme := 0
	moy := 0.0
	for i := 0; i < len(equipe); i++ {
		somme += equipe[i].vie
	}
	moy = float64(somme) / float64(len(equipe))
	return moy
}
func compterFaibles(equipe [6]Soldat) int {
	cpt := 0
	for i := 0; i < len(equipe); i++ {
		if equipe[i].vie < 800 {
			cpt++
		}
	}
	return cpt
}
