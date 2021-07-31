package goconfig

import (
	"log"
	"os"
	"testing"
)

func init() {
	if err := os.Setenv("TESTLOAD_STRING", "hello-world"); err != nil {
		log.Fatal(err)
	}
	if err := os.Setenv("TESTLOAD_BOOL", "true"); err != nil {
		log.Fatal(err)
	}
}

func TestLoad(t *testing.T) {
	var test struct {
		TestString string `config:"TESTLOAD_STRING"`
		TestBool   bool   `config:"TESTLOAD_BOOL"`
	}
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "ALL TYPES",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Load(&test); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			if test.TestString != "hello-world"{
				t.Errorf("invalid TESTLOAD_STRING value")
			}
		})
	}
}
