package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)


const emailFile = "handing_errors/example.txt"
func writeEmailToFile(email []byte) {
	err := os.WriteFile(emailFile, email, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getEmailFromFile() (string, error) {
	emailListByte, err := os.ReadFile(emailFile)
	if err != nil {
		// log.Fatal(err)
		return "", errors.New("Failed to read file.")
	}
	emailListString := string(emailListByte)
	return emailListString, nil

 

}

 
func main() {
	// email := "Jayne_Kuhic@sydney.com"

	// writeEmailToFile([]byte(email))

	// emailList := []string{
	// 	"Eliseo@gardner.biz",
	// 	"Nikita@garfield.biz",
	// 	"Lew@alysha.tv",
	// 	"Hayden@althea.biz",
	// }

	// emailListString := strings.Join(emailList, "\n")
	// writeEmailToFile([]byte(emailListString))

	s, err := getEmailFromFile()
	if err!= nil {
		// log.Fatal(err)
		fmt.Println(err)
		fmt.Println("--------------------")
		// return
		panic("Can't continue, sorry.")
	}
	splitStringSlice  := strings.Split(s, ",")
	fmt.Println(splitStringSlice)
	fmt.Println("123")
}