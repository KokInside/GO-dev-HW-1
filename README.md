# GO development HW1

## Uniq

### Перед началом работы:

```bash
go mod tidy
```

### Запуск тестов:

```bash
go test -v ./uniq/
```

### Запуск программы из файла:

```bash
go run main.go [-c] [-d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]
```

### Создание исполняемого файла:

```bash
go build -o ./bin/
```

#### Запуск исполняемого файла:

```bash
./bin/uniq [-c] [-d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]

## Calc

### Запуск программы из файла:

```bash
go run main.go "[your expression]"
```

### Создание исполняемого файла:

```bash
go build -o ./bin/
```

### Запуск исполнямого файла:

```bash
./bin/calc "[your expression]"
```

### Запуск тестов:

```bash
go test -v ./calc/
```