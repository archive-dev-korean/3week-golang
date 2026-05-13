# 3주차 Go 실습 자료

Go 기초 문법부터 표준 라이브러리만 사용하는 HTTP CRUD API까지 이어지는 강의용 예제입니다.
각 폴더는 `main.go` 하나와 `README.md` 하나로 구성되어 있으며, 폴더 단위로 바로 실행할 수 있습니다.

## 준비

```bash
cd /home/ubuntu/data/3week-golang
go version
```

Go 1.22 이상을 기준으로 작성했습니다.

## 실습 순서

| 순서 | 폴더 | 주제 |
| --- | --- | --- |
| 1 | `01-hello` | Go 프로그램 구조, `package main`, `fmt.Println` |
| 2 | `02-variables-types` | 변수 선언, 타입, 포맷 출력 |
| 3 | `03-conditionals-loops` | 조건문, `switch`, `for`, `range` |
| 4 | `04-functions-errors` | 함수, 다중 반환값, 에러 처리 |
| 5 | `05-collections` | 배열, 슬라이스, 맵 |
| 6 | `06-structs-methods` | 구조체, 메서드, 포인터 리시버 |
| 7 | `07-http-crud` | HTTP 서버, JSON, REST 스타일 CRUD |

## 실행 방법

각 실습 폴더를 지정해서 실행합니다.

```bash
go run ./01-hello
go run ./02-variables-types
go run ./03-conditionals-loops
go run ./04-functions-errors
go run ./05-collections
go run ./06-structs-methods
go run ./07-http-crud
```

전체 예제가 빌드되는지 확인하려면 다음 명령을 사용합니다.

```bash
go test ./...
```

## 과제 제출 안내
1. 3week-golang Repository를 Fork 합니다.
2. 각자 Fork된 Repository 내에서 실README.md를 읽고 실습을 수행합니다(각 README에 과제 존재)
3. 실습 완료 후 PR로 제출합니다.
