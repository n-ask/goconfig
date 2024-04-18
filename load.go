package goconfig

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const defaultSeparator = ","

// Load updates the fields of the given struct with values from environment variables.
// It takes a pointer to a struct as an argument and modifies its fields based on the "env" tags defined for each field.
// The function supports fields of types string, bool, int, uint, float, and slice.
// For string fields, the corresponding environment variable is directly assigned to the field.
// For bool fields, the value is parsed from the environment variable using strconv.ParseBool().
// For int, uint, and float fields, the values are parsed from the environment variable using strconv.ParseInt() and strconv.ParseUint().
// For slice fields, the value is split based on a separator defined in the "sep" tag, or defaulting to a comma if the tag is not present.
// The function returns an error if there is a failure in loading any of the environment variables.
//
// Example usage:
//
//	type Config struct {
//	    Name  string `env:"APP_NAME"`
//	    Port  int    `env:"PORT"`
//	    Admin bool   `env:"ENABLE_ADMIN"`
//	    Roles []string `env:"ALLOWED_ROLES" sep:";"`
//	}
//
// var cfg Config
//
//	if err := Load(&cfg); err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(cfg.Name)
// fmt.Println(cfg.Port)
// fmt.Println(cfg.Admin)
// fmt.Println(cfg.Roles)
func Load(data interface{}) error {
	v := reflect.ValueOf(data).Elem()
	if v.Kind() != reflect.Struct {
		return nil
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		key := t.Field(i).Tag.Get("env")
		if key == "" {
			continue
		}
		if !f.CanSet() {
			log.Default().Printf("cannot set field %s\n", key)
			continue
		}
		s := os.Getenv(key)
		var err error
		switch f.Kind() {
		case reflect.String:
			f.SetString(s)
		case reflect.Bool:
			var b bool
			if s == "" {
				f.SetBool(false)
			} else if b, err = strconv.ParseBool(s); err == nil {
				f.SetBool(b)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var i int64
			if s == "" {
				f.SetInt(0)
			} else if i, err = strconv.ParseInt(s, 10, 64); err == nil {
				f.SetInt(i)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var i uint64
			if s == "" {
				f.SetUint(0)
			} else if i, err = strconv.ParseUint(s, 10, 64); err == nil {
				f.SetUint(i)
			}
		case reflect.Float32, reflect.Float64:
			var n float64
			if s == "" {
				f.SetFloat(0.0)
			} else if n, err = strconv.ParseFloat(s, 64); err == nil {
				f.SetFloat(n)
			}
		case reflect.Slice:
			sep := t.Field(i).Tag.Get("sep")
			if len(sep) == 0 {
				sep = defaultSeparator
			}
			if s == "" {
				f.Set(reflect.ValueOf([]string{}))
			} else {
				f.Set(reflect.ValueOf(strings.Split(s, sep)))
			}

		default:
			log.Default().Printf("Unsupported field type: %s\n", f.Kind().String())
			continue
		}
		if err != nil {
			return fmt.Errorf("failed to load environment variable '%s': %v", key, err)
		}
	}
	return nil
}
