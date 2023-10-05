package main

import (
    "fmt"
    "strings"
    //"math/rand"
)

// J'essai de faire en sorte que la fonction pioche un mot aléatoire
// Je vais essayer en utilisant OS 
// var wordDatabase = []string{ 
//     "arbre",
//     "voiture",
//     "champs",
//     "coton",
//     "avions",
//     "poulet",
//     "pizza",
//     "rat",
//     "informatique",
//     "challenge",
// }

// func chooseRandomWord() string {
//     // Générer un index aléatoire pour choisir un mot
//     index := rand.Intn(len(wordDatabase))
//     return wordDatabase[index]
// }

func main() {
    wordToGuess := "avions" 
    guessedWord := make([]string, len(wordToGuess))

    for i := range guessedWord {
        guessedWord[i] = "_"
    }

    attemptsLeft := 6
    guessedLetters := make([]string, 0)

    fmt.Println("Welcome to the Hangman game!")
    fmt.Printf("Guess the word: %s\n", strings.Join(guessedWord, " "))

    for attemptsLeft > 0 { // Tant qu'il reste des "vies" au joueurs 
        fmt.Printf("Attempts left: %d\n", attemptsLeft)
        fmt.Print("Guessed letters: ")
        fmt.Println(strings.Join(guessedLetters, " "))

        var guess string
        fmt.Print("Guess a letter: ")
        fmt.Scanf("%s", &guess)

        if len(guess) != 1 {
            fmt.Println("Please enter a single letter.")
            continue
        }

        if strings.Contains(wordToGuess, guess) {
            for i, letter := range wordToGuess {
                if string(letter) == guess {
                    guessedWord[i] = guess
                }
            }
        } else {
            fmt.Printf("The letter %s is not in the word.\n", guess)
            attemptsLeft--
        }

        guessedLetters = append(guessedLetters, guess)

        if strings.Join(guessedWord, "") == wordToGuess {
            fmt.Println("Congratulations! You have guessed the word:", wordToGuess)
            break
        }

        fmt.Printf("Current word: %s\n", strings.Join(guessedWord, " "))
    }

    if attemptsLeft == 0 {
        fmt.Println("You lost! The word was:", wordToGuess)
    }
}
