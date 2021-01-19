package map2struct

import (
	"log"
	"testing"
)

type User struct {
	Name  string
	ID    int
	Email string `json:"e-mail"`
	*Profile
}

type Profile struct {
	Gender string
}

func TestScan(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	m := map[string]interface{}{}
	m["Name"] = "test"
	m["ID"] = 100
	m["e-mail"] = `@`
	m["Gender"] = `male`
	user := &User{Profile: &Profile{}}
	err := Scan(user, m, `json`)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("==>", user.ID, user.Name, user.Email, user.Gender)
}
