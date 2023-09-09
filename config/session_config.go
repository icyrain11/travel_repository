package config

type Session struct {
	Name   string
	Salt   string
	MaxAge int
}

func InitSessionConfig() {

}
