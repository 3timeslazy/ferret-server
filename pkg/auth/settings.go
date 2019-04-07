package auth

type Settings struct {
	RootPassword []byte
}

func NewDefaultSettings() Settings {
	return Settings{
		RootPassword: []byte{},
	}
}
