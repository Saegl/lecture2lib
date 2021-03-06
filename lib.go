package lecture2lib

import (
	"errors"
	"github.com/aidarkhanov/nanoid/v2"
	"math"
	"unicode"
)

func ChangeCharCase(char rune) rune {
	if unicode.IsLower(char) {
		return unicode.ToUpper(char)
	} else {
		return unicode.ToLower(char)
	}
}

func FindQuadraticRoots(a, b, c float64) (float64, float64, error) {
	d := b*b - 4*a*c
	if d < 0 {
		return -1, -1, errors.New("discriminant is lower than zero")
	}
	d = math.Sqrt(d)
	r1 := (-b + d) / (2 * a)
	r2 := (-b - d) / (2 * a)
	return r1, r2, nil
}

func GenerateUUID() string {
	uuid, _ := nanoid.New()
	return uuid
}