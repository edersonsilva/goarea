package goarea

import "math"

const Pi = 3.1416

// Pi é uma proporcao numerica definida pela relação entre
// o perimetro de uma circunferencia e seu diametro
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
