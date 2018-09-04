package space

type Planet string

const earthDays float64 = 31557600

var earthOrbitalYears = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.6151972,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.44749,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(ageInSeconds float64, planet Planet) float64 {
	return ageInSeconds / (earthDays * earthOrbitalYears[planet])
}
