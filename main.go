package main

import (
	"fmt"
)

func main() {
	var num1, num2, result float64
	var operator int
	fmt.Print("İlk sayıyı girin: ")
	fmt.Scanln(&num1)
	result = num1
	for {
		fmt.Println("Bir operatör seçin:")
		fmt.Println("1-Toplama")
		fmt.Println("2-Çıkartma")
		fmt.Println("3-Çarpma")
		fmt.Println("4-Bölme")
		fmt.Println("5-Üsalma")
		fmt.Println("6-Karekök alma")
		if result != num1 {
			fmt.Println("0-Sonucu göster")
		}
		for {
			fmt.Scanln(&operator)
			if operator >= 0 && operator <= 6 {
				break
			}
			fmt.Println("Lütfen 0-6 aralığında bir sayı girin.")
		}
		if operator == 0 {
			break
		}
		if operator != 6 {
			fmt.Print("Sonraki sayıyı girin: ")
			fmt.Scanln(&num2)
		}
		switch operator {
		case 1:
			result += num2
		case 2:
			result -= num2
		case 3:
			result *= num2
		case 4:
			result /= num2
		case 5:
			result = pow(result, num2)
		case 6:
			result = sqrt(result)
		default:
			fmt.Println("Geçersiz operatör")
		}
	}
	fmt.Printf("Sonuç: %v\n", result)
}

func pow(x, y float64) float64 {
	result := 1.0
	for i := 0; i < int(y); i++ {
		result *= x
	}
	return result
}

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
