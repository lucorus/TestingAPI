package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
  PathToConfigFile = "config.json"
  PathToJsFile = "js/script.js"
  PathToMainPageFile = "templates/main_page.html"
  PathToJsDir = "/js/"
  PathToCssDir = "/css/"
)


// Функция для обработки запросов на главную страницу
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


// Функция для обработки запросов на /send_request/
func sendRequestHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    return
  }

  body, err := ioutil.ReadAll(r.Body)
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


// посылает запрос с указанными данными
func SendRequest(url, method string, data map[string]interface{}, headers map[string]string) (int, string, error) {
  jsonData, err := json.Marshal(data)
  if err != nil {
    return 0, "nil", fmt.Errorf("error marshalling data: %v", err)
  }

  req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
  if err != nil {
    return 0, "nil", fmt.Errorf("error creating request: %v", err)
  }

  for key, value := range headers {
    req.Header.Set(key, value)
  }

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return 0, "nil", fmt.Errorf("error sending request: %v", err)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return resp.StatusCode, "nil", fmt.Errorf("error reading response body: %v", err)
  }

  return resp.StatusCode, string(body), nil
}


// получает порт и домен из конфига
func GetData(filename string) (string, string, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	data := struct{Port string `json:"port"`; Domain string `json:"domain"`}{}
	json.Unmarshal(byteValue, &data)

	return data.Port, data.Domain, nil
}


type Config struct {
	Settings map[string]string
}

// Получает данные конфигурации из JSON файла
func GetConfigData(filename string) (Config, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return Config{}, err
	}

	var settings map[string]string
	if err := json.Unmarshal(byteValue, &settings); err != nil {
		return Config{}, err
	}

	return Config{Settings: settings}, nil
}


// Создает строки конфигурации для js файла
func CreateJSContent(config Config) string {
	jsContent := ""

	for key, value := range config.Settings {
		jsContent += fmt.Sprintf("const %s = '%s';\n", key, value)
	}

	return jsContent
}


// Заменяет предыдущую конфигурацию js файла
func ReplaceConfigData(filename string, newLines string, n int) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")

	if len(lines) > n {
		lines = append([]string{newLines}, lines[n:]...)
	} else {
		lines = append([]string{newLines}, lines...)
	}

	// Записываем обновленное содержимое обратно в файл
	err = ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		return err
	}

	return nil
}


func main() {
  http.HandleFunc("/", mainPageHandler)
  http.HandleFunc("/send_request/", sendRequestHandler)
	http.Handle("/js/", http.StripPrefix(PathToJsDir, http.FileServer(http.Dir("./js"))))
	http.Handle("/css/", http.StripPrefix(PathToCssDir, http.FileServer(http.Dir("./css"))))

  port, domain, err := GetData(PathToConfigFile)
  if err != nil {
    fmt.Println("Error in getting config info: ", err.Error())
    return
  }

  config, err := GetConfigData("config.json")
	if err != nil {
		fmt.Println("Ошибка при получении данных конфигурации:", err)
		return
	}

	// Создаем новую настройку для JS файла
	newContent := CreateJSContent(config)

	// Удаляем прошлые константы
	err = ReplaceConfigData(PathToJsFile, newContent, 25)
	if err != nil {
		fmt.Println("Ошибка при обновлении конфигурации", err)
		return
	}

	fmt.Println("Конфигурация успешно обновлена!")

  fmt.Println("Server started at " + domain + ":" + port)
  http.ListenAndServe(domain + ":" + port, nil)
}
