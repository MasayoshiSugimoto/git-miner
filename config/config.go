package config

type Cfg struct {
	RepoFolder string
}

func Config() *Cfg {
	config := &Cfg{
		RepoFolder: ".",
	}
	return config
}
