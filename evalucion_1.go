package main

import "fmt"

// Variables globales
var productosVendidos []string
var subtotales []float64

// Registrar una venta
func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)

	productosVendidos = append(productosVendidos, nombre)
	subtotales = append(subtotales, subtotal)

	fmt.Printf("Venta registrada. Subtotal: $%.2f\n", subtotal)
}

func main() {

	var opcion int

	for {
		fmt.Println("\n===== TIENDA =====")
		fmt.Println("1. Registrar nueva venta")
		fmt.Println("2. Mostrar estadísticas")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		if opcion == 1 {

			fmt.Println("\n--- Productos disponibles ---")
			fmt.Println("1. Arroz - $1.25")
			fmt.Println("2. Leche - $0.95")
			fmt.Println("3. Pan - $0.50")

			var producto int
			fmt.Print("Seleccione un producto: ")
			fmt.Scan(&producto)

			var nombre string
			var precio float64

			if producto == 1 {
				nombre = "Arroz"
				precio = 1.25
			} else if producto == 2 {
				nombre = "Leche"
				precio = 0.95
			} else if producto == 3 {
				nombre = "Pan"
				precio = 0.50
			} else {
				fmt.Println("Producto no válido.")
				continue
			}

			var cantidad int
			fmt.Print("Ingrese la cantidad: ")
			fmt.Scan(&cantidad)

			RegistrarVenta(nombre, precio, cantidad)

			fmt.Println("\nVolviendo al menú...")

		} else if opcion == 2 {

			fmt.Println("Esta opción todavía no está disponible.")

		} else if opcion == 3 {

			fmt.Println("Programa finalizado.")
			break

		} else {

			fmt.Println("Opción no válida.")
		}
	}
}
