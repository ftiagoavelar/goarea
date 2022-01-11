package goarea

import (
	"math"
)

// Pi é a proporção definida entre o perímetro da
// circunferência e o diâmetro
const Pi = 3.1416

// Circ é responsável para calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula a área do retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2.0
}

