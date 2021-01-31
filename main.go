package gorea

import "math"

//Pi é uma proporcao numerica definida pela relacao entre o perimetro de uma circunferencia e seu diametro
const Pi = 3.1416

//Circulo e responsavel por calcular a area da circunferencia
func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Retangulo e responsavel por calcula a area de um retangulo
func Retangulo(base, altura float64) float64 {
	return base * altura
}

//Nao e visivel - privada
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
