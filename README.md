## Установка:

1.  **Скачайте Go**

    [Официальный сайт Go](https://go.dev/dl/)

2.  **Клонируйте репозиторий**

    ```bash
    git clone https://github.com/lucorus/TestingAPI
    ```

3.  **Смените активную директорию**

    ```bash
    cd TestingAPI
    ```

4.  **Запустите проект**

    ```bash
    go run main.go
    ```

Готово! Теперь вы можете открыть проект, набрав в адресную строку текст: `localhost:8081/`

## Редактирование:

1. Если порт, использующийся по умолчанию (8081), уже занят, вы можете сменить его на другой, изменив значение `port` на номер свободного порта в файле `config.json`

2. Если вы хотите использовать свой шаблон для вывода информации, измените значение `PathToMainPageFile` в файле `main.go` на путь к вашему файлу (помните, что используемый вами js или css код должен находиться в папках js и css соответсвенно)
