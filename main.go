package goarea

import "math"

// Pi é uma proporção numérica relacionada a circunferencias
const Pi = 3.1416

// Circ responsavel por calcular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect is reponsible to calculate a rectangle area
func Rect(base, altura float64) float64 {
	return base * altura
}

// Not visible
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
