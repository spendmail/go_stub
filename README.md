
---
1. Coding:

 - fix all <stub> occurrences
 - go mod init github.com/spendmail/stub
 - go mod tidy
 - implement new features

---
2. Linting & Testing:
```
golangci-lint run .
golangci-lint run ./...
go test -v -count=1 -race -timeout=1m .
go test -v -race -count=100 .
go test -v -count=1 -timeout=30s -tags bench .
```

---
3. Module Availability
```
cd /tmp/
go get -v -d github.com/spendmail/stub@develop
```

---
4. Build сервиса:
```
make build
```

---
5. Запуск сервиса:
```
make launch
```

---
6. Запуск сервиса (docker-compose):
```
cd /tmp
git clone --branch develop git@github.com:spendmail/stub.git stub
cd stub
make run
```

---
7. Проверка работы:
```
wget http://localhost:8888/path/1024/param1/param2.jpg

