package iteration

const repeatCount = 5

func Repeat(character string) string {
	// var repeated strings.Builder
	// for i := 0; i < repeatCount; i++ {
	// 	repeated.WriteString(character)
	// }
	// return repeated.String()

	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
