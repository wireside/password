package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) []byte {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	
	return data

}

func WriteFile (content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println("Запись успешна")
}
