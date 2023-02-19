package space

import "math"

type Planet string

const (
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Earth   Planet = "Earth"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)

var periods = map[Planet]float64{
	Mercury: 76005.3024,
	Venus:   194139.072,
	Earth:   315581.4976,
	Mars:    593542.944,
	Jupiter: 3743357.76,
	Saturn:  9295966.08,
	Uranus:  26307031.65,
	Neptune: 52004185.92,
}

func Age(seconds float64, planet Planet) float64 {
	period, exists := periods[planet]
	if !exists {
		return -1
	}

	return math.Round(seconds/period) / 100
}
