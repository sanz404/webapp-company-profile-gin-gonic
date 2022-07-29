package seeds

import (
	"backend/helpers"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	mathrandom "math/rand"
	"os"

	"github.com/bxcodec/faker/v3"
	"github.com/gosimple/slug"
	"github.com/kristijorgji/goseeder"
)

func blogSeeder(s goseeder.Seeder) {
	createAbout(s)
	createCategory(s)
	createContent(s)
	createFeature(s)
	createTeam(s)
	createProject(s)
	createProduct(s)
	createUser(s)
	createNotification(s)
	createFaq(s)
	createMessage(s)
	createArticle(s)
}

func createArticle(s goseeder.Seeder) {
	s.DB.Exec(`TRUNCATE TABLE articles`)
	s.DB.Exec(`TRUNCATE TABLE article_categories`)
	for i := 0; i < 100; i++ {
		user_id := mathrandom.Intn(100-1+1) + 1
		title := faker.Sentence()
		slug := slug.Make(title)
		content := faker.Paragraph()
		is_published := 1
		stmt, _ := s.DB.Prepare(`INSERT INTO articles(user_id, title, slug, content, is_published) VALUES (?,?,?,?,?)`)
		_, err := stmt.Exec(user_id, title, slug, content, is_published)
		if err != nil {
			panic(err)
		}
	}
	for i := 0; i < 100; i++ {
		category_id := mathrandom.Intn(2-1+1) + 1
		article_id := mathrandom.Intn(100-1+1) + 1
		stmt, _ := s.DB.Prepare(`INSERT INTO article_categories(category_id, article_id) VALUES (?,?)`)
		_, err := stmt.Exec(category_id, article_id)
		if err != nil {
			panic(err)
		}
	}
}

func createFaq(s goseeder.Seeder) {
	jsonFile, err := os.Open(RootDir() + "/data/json/faqs.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	s.DB.Exec(`TRUNCATE TABLE faqs`)
	for _, item := range result["RECORDS"].([]interface{}) {
		question := item.(map[string]interface{})["question"].(string)
		answer := item.(map[string]interface{})["answer"].(string)
		sort := item.(map[string]interface{})["sort"].(float64)
		is_published := 1
		stmt, _ := s.DB.Prepare(`INSERT INTO faqs(category_id, question, answer, sort, is_published) VALUES (?,?,?,?,?)`)
		_, err := stmt.Exec(mathrandom.Intn(2-1+1)+1, question, answer, sort, is_published)
		if err != nil {
			panic(err)
		}
	}
}

func createMessage(s goseeder.Seeder) {
	jsonFile, err := os.Open(RootDir() + "/data/json/messages.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	s.DB.Exec(`TRUNCATE TABLE messages`)
	for _, item := range result["RECORDS"].([]interface{}) {
		full_name := item.(map[string]interface{})["full_name"].(string)
		email := item.(map[string]interface{})["email"].(string)
		phone := item.(map[string]interface{})["phone"].(string)
		message := item.(map[string]interface{})["message"].(string)
		stmt, _ := s.DB.Prepare(`INSERT INTO messages(full_name, email, phone, message) VALUES (?,?,?,?)`)
		_, err := stmt.Exec(full_name, email, phone, message)
		if err != nil {
			panic(err)
		}
	}
}

func createNotification(s goseeder.Seeder) {
	s.DB.Exec(`TRUNCATE TABLE notifications`)
	for i := 0; i < 100; i++ {
		user_id := mathrandom.Intn(100-1+1) + 1
		subject := faker.Sentence()
		sort_content := faker.Sentence()
		full_content := faker.Paragraph()
		stmt, _ := s.DB.Prepare(`INSERT INTO notifications(user_id, subject, sort_content, full_content) VALUES (?,?,?,?)`)
		_, err := stmt.Exec(user_id, subject, sort_content, full_content)
		if err != nil {
			panic(err)
		}
	}
}

func createUser(s goseeder.Seeder) {
	defaultPassword := "5ecReT!"
	adminUserName := "admin"
	adminEmail := "admin@devel.com"

	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	key := hex.EncodeToString(bytes) //encode key in bytes to string and keep as secret, put in a vault
	encrypted := helpers.Encrypt(defaultPassword, key)

	s.DB.Exec(`TRUNCATE TABLE users`)
	for i := 1; i <= 100; i++ {
		stmt, _ := s.DB.Prepare(`INSERT INTO users(country_id,username,email,phone,address1,address2,city,zip_code,password,password_token,is_admin,status) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`)
		country_id := mathrandom.Intn(243-1+1) + 1
		username := faker.Username()
		email := faker.Email()
		is_admin := 0

		if i == 1 {
			username = adminUserName
			email = adminEmail
			is_admin = 1
		}

		phone := faker.Phonenumber()
		address1 := faker.Sentence()
		address2 := faker.Sentence()
		city := faker.Sentence()
		zip_code := faker.Phonenumber()
		password := encrypted
		password_token := encrypted
		status := 1
		_, err := stmt.Exec(country_id, username, email, phone, address1, address2, city, zip_code, password, password_token, is_admin, status)
		if err != nil {
			panic(err)
		}
	}
}

func createProduct(s goseeder.Seeder) {
	jsonFile, err := os.Open(RootDir() + "/data/json/products.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	s.DB.Exec(`TRUNCATE TABLE products`)
	for _, item := range result["RECORDS"].([]interface{}) {
		name := item.(map[string]interface{})["name"].(string)
		description := item.(map[string]interface{})["description"].(string)
		price := item.(map[string]interface{})["price"].(float64)
		is_published := 1
		stmt, _ := s.DB.Prepare(`INSERT INTO products(category_id,name,price,description,is_published) VALUES (?,?,?,?,?)`)
		_, err := stmt.Exec(mathrandom.Intn(2-1+1)+1, name, price, description, is_published)
		if err != nil {
			panic(err)
		}
	}
}

func createProject(s goseeder.Seeder) {
	jsonFile, err := os.Open(RootDir() + "/data/json/projects.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	s.DB.Exec(`TRUNCATE TABLE projects`)
	for _, item := range result["RECORDS"].([]interface{}) {
		name := item.(map[string]interface{})["name"].(string)
		description := item.(map[string]interface{})["description"].(string)
		is_published := 1
		stmt, _ := s.DB.Prepare(`INSERT INTO projects(category_id,name,description,is_published) VALUES (?,?,?,?)`)
		_, err := stmt.Exec(mathrandom.Intn(2-1+1)+1, name, description, is_published)
		if err != nil {
			panic(err)
		}
	}
}

func createTeam(s goseeder.Seeder) {
	jsonFile, err := os.Open(RootDir() + "/data/json/teams.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	s.DB.Exec(`TRUNCATE TABLE teams`)
	for _, item := range result["RECORDS"].([]interface{}) {
		name := item.(map[string]interface{})["name"].(string)
		position := item.(map[string]interface{})["position"].(string)
		description := item.(map[string]interface{})["description"].(string)
		is_published := 1
		stmt, _ := s.DB.Prepare(`INSERT INTO teams(name,position,description,is_published) VALUES (?,?,?,?)`)
		_, err := stmt.Exec(name, position, description, is_published)
		if err != nil {
			panic(err)
		}
	}
}

func createFeature(s goseeder.Seeder) {
	jsonFile, err := os.Open(RootDir() + "/data/json/features.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	s.DB.Exec(`TRUNCATE TABLE features`)
	for _, item := range result["RECORDS"].([]interface{}) {
		icon := item.(map[string]interface{})["icon"].(string)
		title := item.(map[string]interface{})["title"].(string)
		description := item.(map[string]interface{})["description"].(string)
		is_published := 1
		stmt, _ := s.DB.Prepare(`INSERT INTO features(icon,title, description, is_published) VALUES (?,?,?,?)`)
		_, err := stmt.Exec(icon, title, description, is_published)
		if err != nil {
			panic(err)
		}
	}
}

func createContent(s goseeder.Seeder) {
	jsonFile, err := os.Open(RootDir() + "/data/json/contents.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	s.DB.Exec(`TRUNCATE TABLE contents`)
	for _, item := range result["RECORDS"].([]interface{}) {
		key_name := item.(map[string]interface{})["key_name"].(string)
		key_value := item.(map[string]interface{})["key_value"].(string)
		stmt, _ := s.DB.Prepare(`INSERT INTO contents(key_name,key_value) VALUES (?,?)`)
		_, err := stmt.Exec(key_name, key_value)
		if err != nil {
			panic(err)
		}
	}
}

func createCategory(s goseeder.Seeder) {
	categories := []string{"category_articles", "category_faqs", "category_products", "category_projects"}
	for i := 0; i < len(categories); i++ {
		category := categories[i]
		jsonFile, err := os.Open(RootDir() + "/data/json/" + category + ".json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)

		s.DB.Exec(`TRUNCATE TABLE ` + category)
		for _, item := range result["RECORDS"].([]interface{}) {
			name := item.(map[string]interface{})["name"].(string)
			description := item.(map[string]interface{})["description"].(string)
			stmt, _ := s.DB.Prepare(`INSERT INTO ` + category + `(name, description) VALUES (?,?)`)
			_, err := stmt.Exec(name, description)
			if err != nil {
				panic(err)
			}
		}
	}
}

func createAbout(s goseeder.Seeder) {
	jsonFile, err := os.Open(RootDir() + "/data/json/abouts.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	s.DB.Exec(`TRUNCATE TABLE abouts`)
	for _, item := range result["RECORDS"].([]interface{}) {
		title := item.(map[string]interface{})["title"].(string)
		description := item.(map[string]interface{})["description"].(string)
		is_published := 1
		stmt, _ := s.DB.Prepare(`INSERT INTO abouts(title, description, is_published) VALUES (?,?,?)`)
		_, err := stmt.Exec(title, description, is_published)
		if err != nil {
			panic(err)
		}
	}
}
