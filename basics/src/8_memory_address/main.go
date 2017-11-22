package main
import "fmt"

func main() {
	var value int
	fmt.Println("Cuantas pizzas comiste hoy?")
	fmt.Println(&value)
	fmt.Scan(&value)
	peso := value * 33
	fmt.Println("Tu peso por comer ", value, " pizzas hoy es: ", peso)
}