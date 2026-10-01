package main

import "fmt"

// Slices globales
var productosVendidos []string
var subtotales []float64

// Función para registrar una venta
func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)

	productosVendidos = append(productosVendidos, nombre)
	subtotales = append(subtotales, subtotal)

	fmt.Println("Venta registrada correctamente.")
	fmt.Printf("Producto: %s\n", nombre)
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
}

// Función para mostrar estadísticas
func MostrarEstadisticas() {
	if len(productosVendidos) == 0 {
		fmt.Println("\nNo existen ventas registradas.")
		return
	}

	var total float64

	fmt.Println("\n===== ESTADÍSTICAS =====")

	for i := 0; i < len(productosVendidos); i++ {
		fmt.Printf("%d. %s - $%.2f\n",
			i+1,
			productosVendidos[i],
			subtotales[i])

		total += subtotales[i]
	}

	fmt.Printf("Total recaudado: $%.2f\n", total)
}

// Función principal
func main() {

	for {
		fmt.Println("\n===== SISTEMA DE VENTAS =====")
		fmt.Println("1. Registrar nueva venta")
		fmt.Println("2. Mostrar estadísticas")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {

		case 1:
			fmt.Println("\n===== PRODUCTOS DISPONIBLES =====")
			fmt.Println("1. Arroz - $1.25")
			fmt.Println("2. Leche - $0.95")
			fmt.Println("3. Pan - $0.50")

			var producto int
			fmt.Print("Seleccione el producto: ")
			fmt.Scan(&producto)

			var nombre string
			var precio float64

			switch producto {

			case 1:
				nombre = "Arroz"
				precio = 1.25

			case 2:
				nombre = "Leche"
				precio = 0.95

			case 3:
				nombre = "Pan"
				precio = 0.50

			default:
				fmt.Println("Producto no válido.")
				continue
			}

			var cantidad int
			fmt.Print("Ingrese la cantidad vendida: ")
			fmt.Scan(&cantidad)

			if cantidad <= 0 {
				fmt.Println("La cantidad debe ser mayor que 0.")
				continue
			}

			RegistrarVenta(nombre, precio, cantidad)

		case 2:
			MostrarEstadisticas()

		case 3:
			fmt.Println("\nPrograma finalizado.")
			return

		default:
			fmt.Println("Opción no válida.")
		}
	}
}
