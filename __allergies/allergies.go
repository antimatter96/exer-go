package allergies

var subsToValue = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

var valueToSubs = map[int]string{
	1: "eggs",
	2: "peanuts",
	3: "shellfish",
	4: "strawberries",
	5: "tomatoes",
	6: "chocolate",
	7: "pollen",
	8: "cats",
}

// AllergicTo takes a score and substance and determines if allergic to it
func AllergicTo(score uint, subs string) bool {
	if score&subsToValue[subs] > 0 {
		return true
	}
	return false
}

// Allergies returns all substances that person is allergic to
func Allergies(score uint) []string {
	toRet := make([]string, 0)
	if score < 0 {
		return toRet
	}
	for i := 0; score != 0 && i < 8; i++ {
		if score%2 == 1 {
			toRet = append(toRet, valueToSubs[i+1])
		}
		score /= 2
	}
	return toRet
}
