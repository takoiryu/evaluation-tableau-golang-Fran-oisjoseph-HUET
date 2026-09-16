package main

import "fmt"

func main() {
	var degats int
	var nbatk int
	nbatkinv := 0
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
	for peutContinuer(equipe) == true { //pour le systeme d'erreur j'ai piqué un truc sur stackoverflow
		for {
			fmt.Printf("Nombre d'attaques ennemies : ")
			_, err := fmt.Scan(&nbatk)
			if err != nil || nbatk <= 0 {
				fmt.Println("NAN un entier positif stp")
				nbatk = 0
				var vid string
				fmt.Scanln(&vid)
				continue
			}
			break
		}
		for nbatkinv < nbatk {
			for {
				fmt.Printf("Attaque %d : ", nbatkinv+1)
				_, err := fmt.Scan(&degats)
				if err != nil || degats < 0 {
					fmt.Println("toujours pas un entier positif")
					var vid2 string
					fmt.Scanln(&vid2)
					continue
				}
				break
			}
			attaquerEquipe(&equipe, degats)
			afficherEtat(equipe)
			nbatkinv++
		}
		peutContinuer(equipe)
		fmt.Println("L'équipe peut continuer le combat !")
	}
	fmt.Println("Tous les soldats sont KO... \n La bataille est terminée !")
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
		if equipe[i].vie < 800 && equipe[i].vie > 0 {
			cpt++
		}
	}
	return cpt
}
func attaquerEquipe(equipe *[6]Soldat, degats int) {
	for i := 0; i < len(equipe); i++ {
		(*equipe)[i].vie -= degats
	}
}
func afficherEtat(equipe [6]Soldat) {
	fmt.Println("=== ÉTAT DE L'ÉQUIPE ===")
	for i := 0; i < len(equipe); i++ {
		fmt.Printf(equipe[i].nom)
		if equipe[i].vie <= 0 {
			fmt.Println(" est KO !")
		} else {
			fmt.Println("vie :", equipe[i].vie)
		}
	}
}
func compterVivants(equipe [6]Soldat, index int) int {
	if index >= len(equipe) {
		return 0
	}
	nbviv := 0
	if equipe[index].vie > 0 {
		nbviv++
	}
	return nbviv
}
func peutContinuer(equipe [6]Soldat) bool {
	mort := 0
	vivant := 0
	for i := 0; i < len(equipe); i++ {
		if equipe[i].vie <= 0 {
			mort++
		} else {
			vivant++
		}
	}
	fmt.Println("Nombre de soldats vivants :", vivant)
	fmt.Println("Nombre de soldats KO :", mort)
	if vivant == 0 {
		return false
	} else {
		return true
	}
}
