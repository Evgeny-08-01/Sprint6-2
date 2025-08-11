package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleGet(res http.ResponseWriter, req *http.Request, logger *log.Logger) {

	if _, err := os.Stat("index.html"); os.IsNotExist(err) {
		logger.Fatal("Ошибка при чтении данных:", err)
	}
	http.ServeFile(res, req, "index.html")
}

func HandlePost(res http.ResponseWriter, req *http.Request, logger *log.Logger) {
	req.ParseMultipartForm(10 << 20) // 10 MB
	if req.Method != "POST" {
		logger.Println("Метод не является методом POST")
		http.Error(res, "Метод не является методом POST", http.StatusInternalServerError)
		return
	}
	file, header, err := req.FormFile("myFile")
	if err != nil {
		logger.Println("Не удалось получить файл")
		http.Error(res, "Не удалось получить файл", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		logger.Println("Ошибка при чтении файла")
		http.Error(res, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}
	convString := service.TexttoMorseAndRevers(string(data))

	currentTime := time.Now().UTC()
	moscowTime := currentTime.Add(3 * time.Hour)

	fileName := fmt.Sprintf("%s_%s",
		moscowTime.Format("2006.01.02 15.04.05"), header.Filename)

	file2, err := os.Create(fileName)
	if err != nil {
		logger.Println("Ошибка при создании файла")
		http.Error(res, "Ошибка при создании файла", http.StatusInternalServerError)
		return
	}

	_, err = file2.WriteString(convString)
	if err != nil {
		logger.Println("Ошибка при записи результата конвертации строки в файл")
		http.Error(res, "Ошибка при записи результата конвертации строки в файл", http.StatusInternalServerError)
		return
	}
	defer file2.Close()

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err = res.Write([]byte(convString))
	if err != nil {
		logger.Println("Ошибка при передаче файла")
		http.Error(res, "Ошибка при передаче файла", http.StatusInternalServerError)
		return
	}
}
