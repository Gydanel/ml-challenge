package util

import "errors"

func Float32Sqrt(x float32) float32 {
	return x * x
}

func Cramer(m [2][3]float32) (x, y float32, err error) {
	/* Regla de cramer para resolver sistema de ecuaciones lineales.
	 * Si el determinante es igual a cero entonces el sistema no tiene una unica solucion.
	 */
	var det = m[0][0]*m[1][1] - m[0][1]*m[1][0]
	if det == 0 {
		return 0, 0, errors.New("system is not SDC")
	}
	return (m[0][2]*m[1][1] - m[1][2]*m[0][1]) / det, (m[0][0]*m[1][2] - m[1][0]*m[0][2]) / det, nil
}
