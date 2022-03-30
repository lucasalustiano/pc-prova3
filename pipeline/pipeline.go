package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	var dir string
	var wg sync.WaitGroup

	fmt.Print("\nDigite o path do diretório: ")
	fmt.Scanf("%s", &dir)
	fmt.Println("--------------------------------------------------")

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() { // Só interessa arquivos
			wg.Add(1)
			go func(path string) { // Cria uma goroutine para cada arquivo
				defer wg.Done()
				data, err := ioutil.ReadFile(path)

				if err != nil {
					return
				}

				if len(data) > 0 && data[0]%2 == 0 {
					fmt.Println("Valor do primeiro byte:", data[0], "- Caminho:", path)
				}
			}(path)

			return nil
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	wg.Wait()
}
