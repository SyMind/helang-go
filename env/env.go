package env

type Env struct {
	Values map[string]U8
}

func NewEnv() Env {
	env := Env{
		Values: map[string]U8{},
	}
	return env
}
