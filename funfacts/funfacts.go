package funfacts

type FunFacts struct {
	Terra []string
	Sun   []string
	Luna  []string
}

func GetFunFacts(about string) []string {
	funFacts := FunFacts{
		Sun:   []string{"The temperature in the core of the sun is 15,000,000 degrees Celsius", "The temperature of the outer layer of the Sun is 5778 Kelvin"},
		Luna:  []string{"Temperature of the Moon's surface at night is -183 degrees Celsius", "Temperature of the Moon's surface during the day is 106 degrees Celsius"},
		Terra: []string{"Highest temperature measured on the Earth's surface is 134 degrees Fahrenheit, 56.7 degrees Celsius, 329.82 Kelvin", "Lowest temperature measured on the Earth's surface is -89.4 degrees Celsius", "Temperature in the Earth's inner core is 9392 Kelvin"},
	}

	if about == "terra" {
		return funFacts.Terra
	}
	if about == "sun" {
		return funFacts.Sun
	}
	if about == "luna" {
		return funFacts.Luna
	} else {
		return []string{}
	}
}
