# test-selectel-linter


Приватный плагин для golangci-lint и `go vet` (анализатор для логов).

Коротко:
- Проверяет capitalisation сообщений (должны начинаться со строчной буквы).
- Проверяет, что сообщение содержит только английские буквы и пробелы.
- Детектирует возможные конфиденциальные данные (password, token, api_key и т.п.).

Поддерживаемые функции логирования: `Info`, `Error`, `Debug`, `Warn` в `log/slog`.

---

Клонирование репозитория (если ещё не клонировано):

```bash
git clone https://github.com/wywyy3cee/test-selectel-linter.git
cd test-selectel-linter
```

## 1) Быстрая сборка плагина (.so) для golangci-lint 

```bash
# требуется CGO_ENABLED=1
bash build.sh
# результат: ./bin/selectellinter.so
```

Чтобы golangci-lint использовал этот плагин как приватный, нужно:
- собрать `golangci-lint` с поддержкой CGO (см. ниже), либо использовать абсолютный путь к .so в конфиге и запускать кастомный бинарь.

Пример `.golangci.yml` (в корне проекта, указывайте реальный путь к .so):

```yaml
version: "2"
linters:
	disable-all: true
	custom:
		selectellinter:
			path: /полный/путь/к/selectellinter.so
			description: Selectel linter for log messages
	enable:
		- selectellinter
```

Важно: golangci-lint и плагин должны быть собраны для одной и той же платформы и использовать совпадающие версии зависимостей; для приватного плагина лучше собрать golangci-lint локально с `CGO_ENABLED=1`:

```bash
git clone https://github.com/golangci/golangci-lint.git
cd golangci-lint
CGO_ENABLED=1 make build
# затем использовать ./bin/golangci-lint с указанным .golangci.yml
```

## 2) Локальное использование через `go vet` (рекомендуемый простой вариант)

Собрать исполняемый `vet`-инструмент, который использует ваш анализатор:

```bash
mkdir -p ./bin
go build -o ./bin/vet-selectellinter ./cmd/test-selectel-linter
```

Запустить анализ на пакетах (пример на тестовых данных):

```bash
go vet -vettool=./bin/vet-selectellinter ./pkg/linter/testdata
```

## 3) Тесты

```bash
go test ./pkg/linter/ -v
```

## 4) Примеры проблем, которые обнаружит линтер

Первая строчная буква:
```go
slog.Info("Uppercase message") // отчет: сообщение должно начинаться со строчной
```

Неанглийский текст и спец. символы / эмоджи:
```go
slog.Info("текст на русском") // отчет: только английские буквы и пробелы
```

Чувствительные данные:
```go
password := "secret"
slog.Error("auth failed", "password", password) // отчет: потенциально чувствительные данные
```

---

Если нужно — короткая демонстрация сборки и локального запуска `go vet` находится в `build.sh` и сборка `./bin/vet-selectellinter` показывает, как запускать анализатор без golangci-lint.

Все команды минимальные и воспроизводимые.
