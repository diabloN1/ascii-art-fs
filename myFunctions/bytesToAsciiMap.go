package myfunctions

func BytesToAsciiMap(style []byte) map[int]string {
	chars := make(map[int]string)
	line := 1
	next := 9
	char := ""
	intValue := 32

	//Range byte by byte standard to fill each ascii separatly in the map.
	for i := 1; i < len(style) - 1; i++ {
		if i < len(style) - 2 {
		 	if style[i] == '\n' {
				line++
			} else if line == next+1 {
				next += 9
				chars[intValue] = char[:len(char)-2]
				intValue++
				char = ""
			}
			char += string(style[i])
		} else {
			char += string(style[i])
			chars[intValue] = char
		}
	}
	return chars
}
