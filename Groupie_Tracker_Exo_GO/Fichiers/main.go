package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var choice int
	var filename string

	fmt.Println("Menu :")
	fmt.Println("1. Récupérer tout le texte contenu dans un fichier.txt")
	fmt.Println("2. Ajouter du texte dans ce fichier depuis le go.")
	fmt.Println("3. Supprimer tout le contenu du fichier")
	fmt.Println("4. Remplacer le contenu par du texte donné par l'utilisateur")

	fmt.Print("Choisissez une option (1-4) : ")
	fmt.Scan(&choice)

	fmt.Print("Entrez le nom du fichier : ")
	fmt.Scan(&filename)

	switch choice {
	case 1:
		// Lire tout le contenu du fichier et l'afficher
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier :", err)
			return
		}
		fmt.Println("Contenu du fichier :", string(data))

	case 2:
		// Demander à l'utilisateur le texte à ajouter
		fmt.Print("Entrez le texte à ajouter : ")
		reader := bufio.NewReader(os.Stdin)
		newText, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erreur lors de la lecture de l'entrée utilisateur :", err)
			return
		}

		// Ajouter le texte à la fin du fichier
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture du fichier :", err)
			return
		}
		defer f.Close()

		_, err = fmt.Fprintln(f, newText)
		if err != nil {
			fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
			return
		}
		fmt.Println("Texte ajouté avec succès !")

	case 3:
		// Supprimer tout le contenu du fichier
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Println("Le fichier n'existe pas :", err)
			return
		}

		err := os.WriteFile(filename, []byte(""), 0644)
		if err != nil {
			fmt.Println("Erreur lors de la suppression du contenu du fichier :", err)
			return
		}
		fmt.Println("Contenu du fichier supprimé avec succès !")

	case 4:
		// Écrire un nouveau contenu dans le fichier
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Println("Le fichier n'existe pas :", err)
			return
		}

		// Demander à l'utilisateur le texte de remplacement
		fmt.Print("Entrez le texte de remplacement : ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		newText := scanner.Text()

		// Écrire le nouveau texte dans le fichier
		err := os.WriteFile(filename, []byte(newText), 0644)
		if err != nil {
			fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
			return
		}
		fmt.Println("Contenu du fichier remplacé avec succès !")

	default:
		fmt.Println("Choix invalide.")
	}
}
