package internal

var colors = map[string]string{
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",
	"gray":   "\033[37m",
	"white":  "\033[38m",
	"reset":  "\033[0m",
}

func colorize(s, c string) string {
	return colors[c] + s + colors["reset"]
}
