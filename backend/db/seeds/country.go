package seeds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

func countrySeeder(s goseeder.Seeder) {

	jsonFile, err := os.Open(RootDir() + "/data/json/countries.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	s.DB.Exec(`TRUNCATE TABLE countries`)
	for _, item := range result["ref_country_codes"].([]interface{}) {
		code := item.(map[string]interface{})["alpha2"].(string)
		name := item.(map[string]interface{})["country"].(string)
		stmt, _ := s.DB.Prepare(`INSERT INTO countries(code, name) VALUES (?,?)`)
		_, err := stmt.Exec(code, name)
		if err != nil {
			panic(err)
		}
	}
}
