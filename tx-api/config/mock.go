package config

// MockConfig returns a test configuration.
func MockConfig() *Config {
	return &Config{
		Database: Database{
			Host:     "localhost",
			Port:     5432,
			User:     "testuser",
			Password: "testpass",
			Name:     "testdb",
			Schema:   "testschema",
		},
	}
}
