package main

import (
    "fmt"
    "net/http"
    "os"
    "bufio"
    "strconv"
)

func fuzzer(url string, fileName string) {
    // Ouverture du fichier contenant la liste de mots
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Lecture de chaque mot dans le fichier
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := scanner.Text()

        // Envoi d'une requête HTTP pour chaque mot en utilisant la syntaxe "IP:port/mot" ou "URL:port/mot"
        resp, err := http.Get(fmt.Sprintf("%s/%s", url, word))
        if err != nil {
            fmt.Println(err)
            continue
        }
        defer resp.Body.Close()

        // Affichage du statut de la réponse
        fmt.Println(resp.Status)
    }
}

func main() {
    // Récupération de l'IP ou de l'URL en entrée
    url := os.Args[1]
    var port string
    var fileName string
    // vérifie si on a un port en entrée 
    if len(os.Args) > 2 {
        if _, err := strconv.Atoi(os.Args[2]); err == nil {
            port = os.Args[2]
            fileName = os.Args[3]
        } else {
            fileName = os.Args[2]
        }
    }
    // si on a un port en entrée on le concatene avec l'url ou l'ip
    if port != "" {
        url = fmt.Sprintf("%s:%s", url, port)
    }
    // Appel de la fonction fuzzer
    fuzzer(url, fileName)
}
