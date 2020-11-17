package brewchild

import "math"

// SGToPlato Plato = (-1 * 616.868) + (1111.14 * sg) – (630.272 * sg^2) + (135.997 * sg^3)
func SGToPlato(sg float64) (plato float64) {
	plato = -616.868 + (1111.14 * sg) - (630.272 * math.Pow(sg, 2)) + (135.997 * math.Pow(sg, 3))
	return
}
