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
		fmt.Println("vie", equipe[i].vie)
		fmt.Println("atk", equipe[i].attaque)
		fmt.Printf("\n")
	}
}
