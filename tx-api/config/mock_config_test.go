package config

import "testing"

func TestMockConfig(t *testing.T) {
	cfg := MockConfig()

	if cfg == nil {
		t.Fatal("expected non-nil config")
	}

	db := cfg.Database

	tests := []struct {
		field string
		got   interface{}
		want  interface{}
	}{
		{"Host", db.Host, "localhost"},
		{"Port", db.Port, 5432},
		{"User", db.User, "testuser"},
		{"Password", db.Password, "testpass"},
		{"Name", db.Name, "testdb"},
		{"Schema", db.Schema, "testschema"},
	}

	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("%s: expected %v, got %v", tt.field, tt.want, tt.got)
		}
	}
}
