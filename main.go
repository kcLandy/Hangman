package main

import (
    "fmt"
    "strings"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    ReadFile()
    wordToFind := ReadFile()[rand.Intn(len(ReadFile()))] // choisis un mot aléatoire dans le word.txt
    guessedWord := make([]string, len(wordToFind))

    for i := range guessedWord {
        guessedWord[i] = "_"
    }

    attemptsLeft := 9 // nombre d'essais
    guessedLetters := make([]string, 0)

    fmt.Println("Welcome to the Hangman game!")
    fmt.Printf("Guess the word: %s\n", strings.Join(guessedWord, " "))

    for attemptsLeft > 0 { 
        fmt.Printf("Attempts left: %d\n", attemptsLeft) // nous rappelle combien d'essai il nous reste
        fmt.Print("Guessed letters: ")
        fmt.Println(strings.Join(guessedLetters, " "))

        var guess string
        fmt.Print("Enter a letter: ")
        fmt.Scanf("%s", &guess) // vérifie si la lettre choisie correspond à l'une du mot

        if len(guess) != 1 {
            fmt.Println("just one letter is allowed") // On ne peut donner qu'une lettre a la fois 
            continue
         }

        guess = strings.ToUpper(guess) 

        if strings.Contains(wordToFind, guess) {
            for i, letter := range wordToFind {
                if string(letter) == guess {
                    guessedWord[i] = guess
                }
            }
        } else {
            fmt.Printf("The letter %s is not in the word.\n", guess)
            attemptsLeft--
            draw_hangman(9 - attemptsLeft)
        }

        guessedLetters = append(guessedLetters, guess)

        if strings.Join(guessedWord, "") == wordToFind {
            fmt.Println("Congratulations! You have guessed the word:", wordToFind)
            break
        }

        fmt.Printf("Current word: %s\n", strings.Join(guessedWord, " "))
    }

    if attemptsLeft == 0 {
        fmt.Println("You lost! The word was:", wordToFind)
    }
}