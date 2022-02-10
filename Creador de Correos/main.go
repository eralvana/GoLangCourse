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
        // fmt.Printf("%+v\n", guarda)
				fmt.Println("Hola", guarda[3], guarda[1], guarda[2], ":\n \n Espero que estés muy bien.\n \n Aquí te envío los resultados de tus evaluaciones parciales:\n \n Evaluación parcial 1:", guarda[4], "\n Evaluación parcial 2:", guarda[5], "\n Evaluación parcial 3:", guarda[6],"\n Evaluación parcial 4:", guarda[7], "\n \n Promedio:", guarda[8], "\n \n Nosotros, considerando el desempeño que tuviste en el curso, te asignamos la calificación de", guarda[9])
    }
}
