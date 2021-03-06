package storage

// Storage exposes a group off operations to work with placehodlers values
type Storage interface {
	Init()
	GetLabels(appname string) []string
	GetApps() []string
	SetOption(appname string, label string, variable string, value string)
	GetOption(appname string, label string, variable string) string
	GetOptions(appname string, label string) map[string]string
	Close()
}

// CreateSampleData insert a set off sample data into the storage
func CreateSampleData(s Storage) {
	s.SetOption("app1", "prod", "DATASOURCE_URL", "jdbc:mysql://proddatabaseserver:3306/shcema?profileSQL=true")
	s.SetOption("app1", "", "DATASOURCE_USERNAME", "GOPHER")
	s.SetOption("app1", "dev", "DATASOURCE_USERNAME", "GOPHER-dev")
	s.SetOption("app1", "prod", "DATASOURCE_USERNAME", "GOPHER-prod")
	s.SetOption("app1", "", "DATASOURCE_PASSWORD", "FOOBAR")
	s.SetOption("app1", "dev", "DATASOURCE_PASSWORD", "LOREM")
	s.SetOption("app1", "prod", "DATASOURCE_PASSWORD", "IPSUM")

	s.SetOption("app2", "", "DATASOURCE_PASSWORD", "DOCKER-PASS")
	s.SetOption("app2", "dev", "DATASOURCE_PASSWORD", "DEV-PASS")
}
