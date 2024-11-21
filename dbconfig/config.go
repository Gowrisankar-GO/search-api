package dbconfig

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {

	if err := godotenv.Load(); err != nil {

		log.Fatal("failed to load env")
	}

}

func DbConn() *sql.DB {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", dsn)

	if err != nil {

		log.Fatal(err)
	}

	return db
}

func MigrateTable(db *sql.DB) error {

	_, err := db.Exec("create table users (id serial primary key,name text not null,phone_number text not null,country text null)")

	if err != nil && fmt.Sprint(err)!=`pq: relation "users" already exists` {

		fmt.Println("table migration error: ", err)

		return err
	}

	return nil
}

func PopulateUser(db *sql.DB) error {

	names := []string{"jhon", "mike", "antony", "Bean", "Michael", "Karan", "pradeep"}

	countries := []string{"India", "USA", "Russia", "Germany", "Pakistan", "Canada"}

	for i := 0; i < 1000; i++ {

		rand.Seed(time.Now().UnixNano())

		name := names[rand.Intn(len(names))]

		country := countries[rand.Intn(len(countries))]

		phone := GeneratePhoneNumber()

		_, err := db.Exec("insert into users (name,country,phone_number) values ($1,$2,$3)", name, country, phone)

		if err != nil {

			fmt.Println("data insert error: ", err)

			return err
		}

	}

	return nil
}

func CreateIndexAndExtensions(db *sql.DB)error{

	_,err := db.Exec("create extension fuzzystrmatch")

	if err != nil && fmt.Sprint(err)!=`pq: extension "fuzzystrmatch" already exists` {

		fmt.Println("fuzzystrmatch extension creation error: ", err)

		return err
	}

	_,err = db.Exec("create extension pg_trgm")

	if err != nil && fmt.Sprint(err)!=`pq: extension "pg_trgm" already exists` {

		fmt.Println("pg_trgm extension creation error: ", err)

		return err
	}

	_,err = db.Exec("create index idx_users_name on users using gin (name gin_trgm_ops)")

	if err != nil && fmt.Sprint(err)!=`pq: relation "idx_users_name" already exists` {

		fmt.Println("index creation error: ", err)

		return err
	}

	return nil
}

func GeneratePhoneNumber() string {

	rand.Seed(time.Now().UnixNano())

	countryCode := "+" + fmt.Sprintf("%03d", rand.Intn(999)+1)

	areaCode := rand.Intn(900) + 100

	localNumber := rand.Intn(9000000) + 1000000

	phone := fmt.Sprintf("%s %03d%07d", countryCode, areaCode, localNumber)

	return phone
}
