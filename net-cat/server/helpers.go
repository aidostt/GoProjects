package server

func IsinAscii(text string) bool {
	for _, v := range text {
		if v < ' ' || v >= '~' {
			return false
		}
	}
	return true
}

func IsClientExist(stat string) bool {
	for n := range clients {
		if n == stat {
			return true
		}
	}
	return false
}