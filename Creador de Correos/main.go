package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
)

func main() {
    datos, error := os.Open("Calificaciones.csv")
    if error != nil {
        log.Fatal(error)
    }
    defer datos.Close()

    csvReader := csv.NewReader(datos)
    for {
        guarda, error := csvReader.Read()
        if error == io.EOF {
            break
        }
        if error != nil {
            log.Fatal(error)
        }
        fmt.Printf("%+v\n", guarda)
    }
}
