package services

import "log"

// Cfg is a structure holding all data related to the application configuration.
type Cfg struct {
	RepoFolder     string // Folder containing all repos to be analysed.
	FileServerPort int    // Port used by the file server.
}

// Config loads the application configuration.
func Config() *Cfg {
	// Configure the logger
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("Loading Configuration...")

	config := &Cfg{
		RepoFolder:     "..",
		FileServerPort: 3000,
	}

	log.Printf("Config loaded: %+v\n", *config)

	return config
}
