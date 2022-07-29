package seeds

import (
	"github.com/bxcodec/faker/v3"
	"github.com/kristijorgji/goseeder"
)

func contactSeeder(s goseeder.Seeder) {
	s.DB.Exec(`TRUNCATE TABLE contacts`)
	for i := 0; i < 100; i++ {
		stmt, _ := s.DB.Prepare(`INSERT INTO contacts(name, email, phone, website, address) VALUES (?,?,?,?,?)`)
		_, err := stmt.Exec(faker.Name(), faker.Email(), faker.Phonenumber(), faker.DomainName(), faker.Paragraph())
		if err != nil {
			panic(err)
		}
	}

}
