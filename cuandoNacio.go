package main

import (
	"fmt"
	"time"
)

func main() {

	name := ""
	var edad int
	var dias int
	var err error
	fecha := time.Now()

	dia := fecha.Day()
	mes := fecha.Month()
	yearActual := fecha.Year()

	var diaUser int
	var mesUser time.Month

	for {
		fmt.Println("\nHola, ¿cómo te llamas?")
		_, err := fmt.Scanln(&name)
		if err != nil {
			fmt.Println("Error al leer tu nombre:", err)
			continue
		}
		break
	}

	for {
		fmt.Println("\nHola,", name, "¿cuántos años tienes?")
		_, err = fmt.Scanln(&edad)
		if err != nil {
			fmt.Println("Error al leer tu edad:", err)
			continue
		}
		break
	}

	for {
		fmt.Println("\n¿Qué día naciste?")
		_, err = fmt.Scanln(&diaUser)
		if err != nil {
			fmt.Println("Error al leer el día:", err)
			continue
		}
		break
	}

	for {
		fmt.Println("\n¿De qué mes? (número del mes)")
		_, err = fmt.Scanln(&mesUser)
		if err != nil {
			fmt.Println("Error al leer el mes:", err)
			continue
		}
		break
	}

	if diaUser == dia && mesUser == mes {
		fmt.Println("\n¡¡¡ Feliz cumpleaños,", name, "!!!")
		time.Sleep(1 * time.Second)
		fmt.Println("\n¡DISFRUTA DE TU DÍA!")
		return
	}

	var year int
	if (mesUser > mes) || (mesUser == mes && diaUser > dia) {
		year = yearActual - edad - 1
	} else {
		year = yearActual - edad
	}

	dias = diasParaCumpleaños(mesUser, diaUser, mes, dia)
	fmt.Printf("\nNaciste en el año %d. Faltan %d días para tu cumpleaños.\n", year, dias)
	time.Sleep(1 * time.Second)
	fmt.Println("\nHasta la próxima")
}

func diasParaCumpleaños(mesNac time.Month, diaNac int, mesActual time.Month, diaActual int) int {
	yearActual := time.Now().Year()

	fechaProximoCumple := time.Date(yearActual, mesNac, diaNac, 0, 0, 0, 0, time.Local)
	if fechaProximoCumple.Before(time.Now()) {
		fechaProximoCumple = fechaProximoCumple.AddDate(1, 0, 0)
	}

	dias := int(time.Until(fechaProximoCumple).Hours() / 24)
	return dias
}
