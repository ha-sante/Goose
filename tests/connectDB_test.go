package testing

import (
	"log"
	"testing"

	"os"

	"github.com/Henry-Asante/goose"
	"github.com/joho/godotenv"
)

var databaseURL string;

func TestConnectDatabase(t *testing.T){

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL = os.Getenv("MLABURL"); // set the database

	_ , err = goose.ConnectDatabase(databaseURL); // returns a connection or error
	
	if err != nil{
		t.Errorf("TestConnectDatabase Failed")
	}
}