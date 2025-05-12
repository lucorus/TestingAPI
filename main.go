package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

const (
	PathToConfigFile   = "config.json"
	PathToJsFile       = "js/script.js"
	PathToMainPageFile = "templates/main_page.html"
	PathToJsDir        = "/js/"
	PathToCssDir       = "/css/"
)

// обрабатывает запросы на главную страницу.
func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles(PathToMainPageFile)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// обрабатывает запросы на /send_request/
func sendRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var payload RequestPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	statusCode, responseBody, err := SendRequest(payload.URL, payload.Method, payload.Data, payload.Headers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"status": statusCode,
		"body":   responseBody,
		"error":  err,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type RequestPayload struct {
	URL     string                 `json:"url"`
	Method  string                 `json:"method"`
	Data    map[string]interface{} `json:"data"`
	Headers map[string]string      `json:"headers"`
}

// отправляет запрос
func SendRequest(url, method string, data map[string]interface{}, headers map[string]string) (int, string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, "", fmt.Errorf("error marshalling data: %v", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, "", fmt.Errorf("error creating request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", fmt.Errorf("error reading response body: %v", err)
	}

	return resp.StatusCode, string(body), nil
}

// получает порт и домен из конфига
func GetData(filename string) (string, string, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return "", "", fmt.Errorf("error opening config file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return "", "", fmt.Errorf("error reading config file: %v", err)
	}

	data := struct {
		Port   string `json:"port"`
		Domain string `json:"domain"`
	}{}
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return "", "", fmt.Errorf("error unmarshalling config data: %v", err)
	}

	return data.Port, data.Domain, nil
}

type Config struct {
	Settings map[string]string
}

// получает данные цветов из конфига
func GetConfigData(filename string) (Config, error) {
	byteValue, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %v", err)
	}

	var settings map[string]string
	if err := json.Unmarshal(byteValue, &settings); err != nil {
		return Config{}, fmt.Errorf("error unmarshalling config data: %v", err)
	}

	return Config{Settings: settings}, nil
}

// создает строки конфигурации для JS файла
func CreateJSContent(config Config) string {
	var jsContent strings.Builder

	for key, value := range config.Settings {
		jsContent.WriteString(fmt.Sprintf("const %s = '%s';\n", key, value))
	}

	return jsContent.String()
}

// заменяет старую конфигурацию в JS файле
func ReplaceConfigData(filename string, newLines string, n int) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading JS file: %v", err)
	}

	lines := strings.Split(string(content), "\n")

	if len(lines) > n {
		lines = append([]string{newLines}, lines[n:]...)
	} else {
		lines = append([]string{newLines}, lines...)
	}

	// Записываем обновленное содержимое в файл
	err = os.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		return fmt.Errorf("error writing JS file: %v", err)
	}

	return nil
}

func main() {
  mux := http.NewServeMux()
    
  mux.HandleFunc("/", mainPageHandler)
  mux.HandleFunc("/send_request/", sendRequestHandler)
  mux.Handle(PathToJsDir, http.StripPrefix(PathToJsDir, http.FileServer(http.Dir("./js"))))
  mux.Handle(PathToCssDir, http.StripPrefix(PathToCssDir, http.FileServer(http.Dir("./css"))))

  // Настраиваем CORS
  c := cors.New(cors.Options{
    AllowedOrigins:   []string{"*"}, // Разрешаем все домены
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Content-Type", "Authorization"},
    AllowCredentials: true,
    Debug: true, // Выводит отладочную информацию в консоль
  })

  handler := c.Handler(mux)

  // Получаем параметры сервера из конфига
  port, domain, err := GetData(PathToConfigFile)
  if err != nil {
      fmt.Println("Ошибка при получении данных конфигурации: port, host. Будут использован localhost:8081")
      port = "8081"
      domain = "localhost"
  }

  // Обновляем JS конфиг
  config, err := GetConfigData(PathToConfigFile)
  if err != nil {
    fmt.Println("Ошибка при получении данных конфигурации:", err)
    return
  }

  newContent := CreateJSContent(config)
  err = ReplaceConfigData(PathToJsFile, newContent, 25)
  if err != nil {
    fmt.Println("Ошибка при обновлении конфигурации:", err)
    return
  }

  fmt.Println("Конфигурация успешно обновлена!")
  fmt.Println("Server started at " + domain + ":" + port)
    
  // Запускаем сервер с CORS middleware
  http.ListenAndServe(domain+":"+port, handler)
}