package xhttp

type Logger interface {
	Debug(string)
	Debugf(string, ...any)
	Fatal(string)
	Fatalf(string, ...any)
	Panic(string)
	Panicf(string, ...any)
}
