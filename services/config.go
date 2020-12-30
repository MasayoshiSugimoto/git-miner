package services

import "log"

type Cfg struct {
	RepoFolder     string
	FileServerPort int
}

func Config() *Cfg {
	// Configure the logger
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("Loading Configuration...")

	config := &Cfg{
		RepoFolder:     ".",
		FileServerPort: 3000,
	}

	log.Printf("Config loaded: %+v\n", *config)

	return config
}
