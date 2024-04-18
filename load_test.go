package goconfig

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func init() {
	if err := os.Setenv("TESTLOAD_STRING", "hello-world"); err != nil {
		log.Fatal(err)
	}
	if err := os.Setenv("TESTLOAD_BOOL", "true"); err != nil {
		log.Fatal(err)
	}
	if err := os.Setenv("TESTLOAD_LIST", "str1||str2||str3"); err != nil {
		log.Fatal(err)
	}
}

func TestLoad(t *testing.T) {
	var test struct {
		TestString string   `env:"TESTLOAD_STRING"`
		TestBool   bool     `env:"TESTLOAD_BOOL"`
		TestList   []string `env:"TESTLOAD_LIST" sep:"||"`
	}
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "ALL TYPES",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Load(&test); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			if test.TestString != "hello-world" {
				t.Errorf("invalid TESTLOAD_STRING value")
			}
			if test.TestBool != true {
				t.Errorf("invalid TESTLOAD_BOOL value")
			}
			if !reflect.DeepEqual(test.TestList, []string{"str1", "str2", "str3"}) {
				t.Errorf("invalid TESTLOAD_LIST value: %v", test.TestList)
			}
		})
	}
}
