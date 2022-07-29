package seeds

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/kristijorgji/goseeder"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func init() {
	goseeder.Register(contactSeeder)
	goseeder.Register(countrySeeder)
	goseeder.Register(blogSeeder)
}
