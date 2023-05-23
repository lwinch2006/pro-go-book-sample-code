package main

import (
	"chapter22/constants"
	"chapter22/models"
	"chapter22/utils"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func WorkWithFiles1() {
	utils.Printfln("WorkWithFiles1()")

	for _, p := range constants.Products {
		utils.Printfln("Product %v of category %v with price %.2f", p.Name, p.Category, p.Price)
	}
}

func WorkWithFiles2() {
	utils.Printfln("WorkWithFiles2()")

	if data, err := os.ReadFile("config.json"); err == nil {
		utils.Printfln("Config file content: %v", string(data))
	} else {
		utils.Printfln("File read error: %v", err.Error())
	}
}

func WorkWithFiles3() {
	utils.Printfln("WorkWithFiles3()")

	if data, err := os.ReadFile("config.json"); err == nil {
		var config models.Config
		var decoder = json.NewDecoder(strings.NewReader(string(data)))
		decoder.Decode(&config)
		utils.Printfln("Config file content: %v", config)
	} else {
		utils.Printfln("File read error: %v", err.Error())
	}
}

func WorkWithFiles4() {
	utils.Printfln("WorkWithFiles4()")

	if data, err := os.ReadFile("config.json"); err == nil {
		var config models.Config
		var decoder = json.NewDecoder(strings.NewReader(string(data)))
		decoder.Decode(&config)
		constants.Products = append(constants.Products, config.AdditionalProducts...)
		utils.Printfln("Config file content: %v", constants.Products)
	} else {
		utils.Printfln("File read error: %v", err.Error())
	}
}

func WorkWithFiles5() {
	utils.Printfln("WorkWithFiles5()")

	if file, err := os.Open("config.json"); err == nil {
		defer file.Close()

		var config models.Config
		decoder := json.NewDecoder(file)
		decoder.Decode(&config)
		utils.Printfln("Config file content: %v", config)
	} else {
		utils.Printfln("File open error: %v", err.Error())
	}
}

func WorkWithFiles6() {
	utils.Printfln("WorkWithFiles6()")

	if file, err := os.Open("config.json"); err == nil {
		defer file.Close()

		usernameAsBytes := make([]byte, 5)
		file.ReadAt(usernameAsBytes, 17)
		file.Seek(49, 0)

		var config models.Config
		decoder := json.NewDecoder(file)
		decoder.Decode(&config.AdditionalProducts)

		utils.Printfln("Username: %v", string(usernameAsBytes))
		utils.Printfln("Config file content: %v", config.AdditionalProducts)
	} else {
		utils.Printfln("File open error: %v", err.Error())
	}
}

func WorkWithFiles7() {
	utils.Printfln("WorkWithFiles7()")

	totalPrice := 0.00
	for _, p := range constants.Products {
		totalPrice += p.Price
	}

	textToWrite := fmt.Sprintf("Time: %v, Total price: %.2f\n", time.Now().Format("Mon 15:04:05"), totalPrice)

	if err := os.WriteFile("totalPriceReport1.txt", []byte(textToWrite), 0666); err == nil {
		utils.Printfln("File write success")
	} else {
		utils.Printfln("File write error: %v", err.Error())
	}
}

func WorkWithFiles8() {
	utils.Printfln("WorkWithFiles8()")

	totalPrice := 0.00
	for _, p := range constants.Products {
		totalPrice += p.Price
	}

	textToWrite := fmt.Sprintf("Time: %v, Total price: %.2f\n", time.Now().Format("Mon 15:04:05"), totalPrice)

	if file, err := os.OpenFile("totalPriceReport2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666); err == nil {
		defer file.Close()
		file.WriteString(textToWrite)
		utils.Printfln("File write success")
	} else {
		utils.Printfln("File open error: %v", err.Error())
	}
}

func WorkWithFiles9() {
	utils.Printfln("WorkWithFiles9()")

	if file, err := os.OpenFile("totalPriceReport3.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666); err == nil {
		defer file.Close()

		jsonEncoder := json.NewEncoder(file)
		jsonEncoder.Encode(constants.Products)
		utils.Printfln("File write success")
	} else {
		utils.Printfln("File open error: %v", err.Error())
	}
}

func WorkWithFiles10() {
	utils.Printfln("WorkWithFiles10()")

	if file, err := os.CreateTemp(".", "temp-*.json"); err == nil {
		defer file.Close()

		jsonEncoder := json.NewEncoder(file)
		jsonEncoder.Encode(constants.Products)
		utils.Printfln("File write success")
	} else {
		utils.Printfln("File open error: %v", err.Error())
	}
}

func WorkWithFiles11() {
	utils.Printfln("WorkWithFiles11()")

	if path, err := os.UserHomeDir(); err == nil {
		path = filepath.Join(path, "WorkingWithFiles", "totalPriceReport4.txt")

		utils.Printfln("Full path: %v", path)
		utils.Printfln("Volume name: %v", filepath.VolumeName(path))
		utils.Printfln("Directory: %v", filepath.Dir(path))
		utils.Printfln("Filename with extension: %v", filepath.Base(path))
		utils.Printfln("File extension: %v", filepath.Ext(path))
	}
}

func WorkWithFiles12() {
	utils.Printfln("WorkWithFiles12()")

	path, _ := os.Getwd()
	path = filepath.Join(path, "Folder1", "Folder2", "totalPriceReport4.txt")

	if err := os.MkdirAll(filepath.Dir(path), 0766); err == nil {
		if file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666); err == nil {
			defer file.Close()

			jsonEncoder := json.NewEncoder(file)
			jsonEncoder.Encode(constants.Products)
			utils.Printfln("File write success")
		} else {
			utils.Printfln("File open error: %v", err.Error())
		}
	} else {
		utils.Printfln("Directory creation error: %v", err.Error())
	}
}

func WorkWithFiles13() {
	utils.Printfln("WorkWithFiles13()")

	path, _ := os.Getwd()
	dirEntries, _ := os.ReadDir(path)

	utils.Printfln("Current working directory: %v", path)
	for _, dirEntry := range dirEntries {
		utils.Printfln("Path entry: %v, is directory: %v", dirEntry.Name(), dirEntry.IsDir())
	}
}

func WorkWithFiles14() {
	utils.Printfln("WorkWithFiles14()")

	fileNames := []string{"not-exist-file.txt", "config.json"}

	for _, fileName := range fileNames {
		stat, err := os.Stat(fileName)

		if os.IsNotExist(err) {
			utils.Printfln("File %v exists %v", fileName, !os.IsNotExist(err))
		} else {
			utils.Printfln("File %v exists %v with size %v", fileName, !os.IsNotExist(err), stat.Size())
		}
	}
}

func WorkWithFiles15() {
	utils.Printfln("WorkWithFiles15()")

	path, _ := os.Getwd()
	foundFiles, _ := filepath.Glob(filepath.Join(path, "total*.txt"))

	for _, foundFile := range foundFiles {
		utils.Printfln("Found file: %v", foundFile)
	}
}

func workWithFiles16_1(path string, dirEntry os.DirEntry, dirErr error) (err error) {
	info, _ := dirEntry.Info()
	utils.Printfln("Path: %v, size: %v", path, info.Size())
	return
}

func WorkWithFiles16() {
	utils.Printfln("WorkWithFiles16()")
	path, _ := os.Getwd()
	filepath.WalkDir(path, workWithFiles16_1)
}
