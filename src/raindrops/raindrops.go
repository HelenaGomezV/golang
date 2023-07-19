package raindrops

import "strconv"

func Convert(number int) string {
	result := ""

	factors := [3]int{3, 5, 7}
	sounds := [3]string{"Pling", "Plang", "Plong"}

	for i := 0; i < 3; i++ {
		if number%factors[i] == 0 {
			result += sounds[i]
		}
	}

	if result == "" {
		result = strconv.Itoa(number)
	}

	return result
}
