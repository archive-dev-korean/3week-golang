# Go 언어를 시작하기에 앞서

이 문서는 `3week-golang` 강의를 시작하기 전에 필요한 Go 설치, 개발 환경 확인, Go Modules 초기화, Go 언어의 기본 특성을 정리한 사전 안내 문서입니다.

강의 저장소는 다음 주소를 기준으로 합니다.

```text
https://github.com/LitmusChaos-Korea-Mentorship/3week-golang.git
```

Go 관련 공식 설치 문서는 아래에서 확인할 수 있습니다.

```text
https://go.dev/doc/install
```

## 1. 전체 진행 순서

Go 실습은 다음 순서로 준비합니다.

1. Go 설치
2. 터미널에서 `go version`으로 설치 확인
3. 강의 저장소 클론
4. 저장소 루트로 이동
5. `go.mod` 확인 또는 생성
6. 예제 실행

`go mod init`, `go run`, `go test` 같은 명령은 모두 Go가 설치되어 있어야 사용할 수 있습니다. 따라서 Go 설치 확인이 가장 먼저입니다.

## 2. Go 설치

### Windows

Windows에서는 공식 설치 페이지에서 `.msi` 설치 파일을 내려받아 설치하는 방법이 가장 단순합니다.

1. 브라우저에서 `https://go.dev/doc/install` 또는 `https://go.dev/dl/`에 접속합니다.
2. Windows용 installer를 다운로드합니다.
3. `.msi` 파일을 실행하고 안내에 따라 설치합니다.
4. 새 PowerShell 또는 터미널을 열고 설치를 확인합니다.

```powershell
go version
```

정상 설치되었다면 다음과 비슷한 결과가 나옵니다.

```text
go version go1.xx.x windows/amd64
```

### macOS

macOS도 공식 설치 파일을 사용하는 방식이 가장 일반적입니다.

1. `https://go.dev/doc/install` 또는 `https://go.dev/dl/`에 접속합니다.
2. macOS용 package installer를 다운로드합니다.
3. 설치 후 새 터미널을 열어 확인합니다.

```bash
go version
```

Homebrew를 사용하는 환경이라면 다음 방식도 가능합니다.

```bash
brew install go
go version
```

### Linux

Linux에서는 배포판 패키지 매니저를 쓰거나 공식 tarball을 직접 설치할 수 있습니다.

Ubuntu 계열에서 패키지 매니저를 사용하는 예시는 다음과 같습니다.

```bash
sudo apt update
sudo apt install golang-go
go version
```

다만 배포판 패키지 매니저의 Go 버전은 최신 안정 버전보다 낮을 수 있습니다. 강의에서 특정 버전 이상이 필요하다면 공식 설치 문서의 tarball 설치 방법을 따르는 편이 좋습니다.

공식 설치 문서:

```text
https://go.dev/doc/install
```

### WSL을 사용하는 경우

Windows에서 WSL로 실습한다면 중요한 점이 있습니다.

Windows PowerShell에 Go가 설치되어 있어도 WSL 안에서는 별도의 Linux 환경입니다. 즉, WSL 터미널에서 `go version`이 동작해야 WSL에서 Go 실습을 진행할 수 있습니다.

WSL 안에서 확인합니다.

```bash
go version
```

동작하지 않는다면 WSL 내부에 Go를 설치해야 합니다.

```bash
sudo apt update
sudo apt install golang-go
go version
```

## 3. 설치 확인

Go 설치가 끝났다면 다음 명령을 실행합니다.

```bash
go version
```

정상 출력 예시는 다음과 같습니다.

```text
go version go1.xx.x linux/amd64
```

또는 Windows에서는 다음처럼 보일 수 있습니다.

```text
go version go1.xx.x windows/amd64
```

`go version`이 동작하지 않으면 아직 Go가 설치되지 않았거나, `PATH` 환경 변수에 Go 실행 파일 경로가 잡히지 않은 상태입니다.

## 4. 강의 저장소 준비

Go 설치가 확인되면 강의 저장소를 클론합니다.

```bash
git clone https://github.com/LitmusChaos-Korea-Mentorship/3week-golang.git
cd 3week-golang
```

이미 저장소를 받은 상태라면 저장소 루트 디렉터리로 이동하면 됩니다.

예를 들어 WSL에서 다음 위치에 있다면:

```bash
/mnt/d/Litmus/3week/3week-golang
```

아래처럼 이동합니다.

```bash
cd /mnt/d/Litmus/3week/3week-golang
```

## 5. Go Modules와 go.mod

Go Modules는 Go 프로젝트의 의존성과 모듈 이름을 관리하는 방식입니다.

`go.mod` 파일은 다음 정보를 담습니다.

- 이 프로젝트의 모듈 경로
- 사용하는 Go 버전
- 외부 라이브러리 의존성

예시:

```go
module github.com/LitmusChaos-Korea-Mentorship/3week-golang

go 1.22
```

이 강의 저장소의 모듈 경로는 저장소 URL에서 `.git`을 제외한 값입니다.

```text
github.com/LitmusChaos-Korea-Mentorship/3week-golang
```

## 6. go mod init은 언제 실행하는가?

`go mod init`은 Go 모듈을 처음 만들 때 실행합니다.

저장소 루트에 이미 `go.mod`가 있다면 다시 실행하지 않습니다.

먼저 확인합니다.

```bash
ls
```

또는 Windows PowerShell에서는:

```powershell
dir
```

`go.mod`가 없다면 저장소 루트에서 다음 명령을 실행합니다.

```bash
go mod init github.com/LitmusChaos-Korea-Mentorship/3week-golang
```

그 결과 루트 디렉터리에 `go.mod`가 생성됩니다.

```text
3week-golang/
  go.mod
  01-hello/
  02-variables-types/
  03-conditionals-loops/
  ...
```

## 7. go mod init 오류가 나는 이유

다음처럼 모듈 경로 없이 실행하면 오류가 날 수 있습니다.

```bash
go mod init
```

오류 예시는 다음과 같습니다.

```text
go: cannot determine module path for source directory ... (outside GOPATH, module path must be specified)
```

이 오류는 Go가 현재 프로젝트의 모듈 경로를 자동으로 결정할 수 없다는 뜻입니다.

해결 방법은 모듈 경로를 직접 지정하는 것입니다.

```bash
go mod init github.com/LitmusChaos-Korea-Mentorship/3week-golang
```

정리하면 `go mod init` 자체가 문제가 아니라, 모듈 경로를 생략한 것이 문제입니다.

## 8. 예제 실행

저장소 루트에서 각 예제를 실행합니다.

```bash
go run ./01-hello
go run ./02-variables-types
go run ./03-conditionals-loops
```

특정 폴더 안으로 들어가서 실행할 수도 있습니다.

```bash
cd 01-hello
go run .
```

다시 루트로 돌아가려면:

```bash
cd ..
```

## 9. 전체 빌드 확인

전체 예제가 컴파일되는지 확인하려면 저장소 루트에서 다음 명령을 실행합니다.

```bash
go test ./...
```

테스트 파일이 없더라도 각 패키지가 정상적으로 빌드되는지 확인할 수 있습니다.

코드 포맷은 다음 명령으로 정리합니다.

```bash
go fmt ./...
```

## 10. 자주 쓰는 Go 명령어

| 명령어 | 의미 |
| --- | --- |
| `go version` | 설치된 Go 버전 확인 |
| `go mod init <module>` | 새 Go 모듈 생성 |
| `go mod tidy` | 의존성 정리 |
| `go run <path>` | Go 프로그램 실행 |
| `go build <path>` | 실행 파일 빌드 |
| `go test ./...` | 현재 모듈의 모든 패키지 테스트 및 빌드 확인 |
| `go fmt ./...` | Go 코드 포맷 정리 |
| `go env` | Go 환경 변수 확인 |

## 11. VS Code 설정

VS Code를 사용한다면 Go 확장을 설치하는 것이 좋습니다.

- 확장 이름: `Go`
- 확장 ID: `golang.go`

Go 파일을 열면 VS Code가 추가 도구 설치를 안내할 수 있습니다. 대표적으로 `gopls`는 Go 언어 서버로, 자동 완성, 정의로 이동, 오류 표시 등에 사용됩니다.

## 12. Go 언어의 주요 특성

### 단순한 문법

Go는 문법이 비교적 작고 단순합니다. 처음 배우는 입장에서도 프로그램의 흐름을 읽기 쉽도록 설계되어 있습니다.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go")
}
```

Go 프로그램은 보통 `package main`과 `func main()`에서 시작합니다.

### 컴파일 언어

Go는 실행 전에 기계가 실행할 수 있는 형태로 컴파일됩니다.

```bash
go build ./01-hello
```

컴파일 언어이기 때문에 실행 속도가 빠르고, 서버나 CLI 도구를 하나의 실행 파일로 배포하기 좋습니다.

### 정적 타입 언어

Go는 변수의 타입이 명확한 언어입니다.

```go
var name string = "gopher"
var age int = 10
```

타입 추론도 지원합니다.

```go
name := "gopher"
age := 10
```

`:=`는 Go에서 자주 사용하는 짧은 변수 선언 방식입니다.

### 표준 라이브러리가 강력함

Go는 외부 프레임워크 없이도 많은 기능을 표준 라이브러리로 제공합니다.

- 문자열 처리
- 파일 입출력
- JSON 처리
- HTTP 서버와 클라이언트
- 테스트
- 시간 처리
- 동시성 처리

예를 들어 HTTP 서버도 표준 라이브러리만으로 만들 수 있습니다.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, HTTP")
    })

    http.ListenAndServe(":8080", nil)
}
```

### 동시성 지원

Go는 동시성 처리를 언어 차원에서 지원합니다. 대표적인 기능은 `goroutine`과 `channel`입니다.

```go
go doSomething()
```

`go` 키워드를 붙이면 함수를 별도의 goroutine으로 실행할 수 있습니다.

채널은 goroutine 사이에서 값을 주고받을 때 사용합니다.

```go
ch := make(chan string)
```

서버 개발, 네트워크 프로그램, 백그라운드 작업 처리에서 Go가 자주 사용되는 이유 중 하나입니다.

### 명시적인 에러 처리

Go는 예외를 던지고 잡는 방식보다 에러 값을 직접 반환하고 확인하는 방식을 선호합니다.

```go
result, err := doSomething()
if err != nil {
    return err
}
```

처음에는 반복적으로 보일 수 있지만, 실패 가능성을 코드에서 분명하게 드러낸다는 장점이 있습니다.

## 13. Go 프로젝트의 기본 구조

작은 학습 프로젝트는 다음처럼 단순하게 시작할 수 있습니다.

```text
3week-golang/
  go.mod
  01-hello/
    main.go
  02-variables-types/
    main.go
```

각 실습 폴더가 실행 가능한 프로그램이라면 각 폴더의 `main.go`는 보통 다음 구조를 가집니다.

```go
package main

func main() {
    // 실행 시작 지점
}
```

여러 파일로 나누더라도 같은 폴더 안에서는 같은 package 이름을 사용해야 합니다.

## 14. GOPATH와 Go Modules

예전 Go는 `GOPATH`라는 특정 작업 공간 안에서 프로젝트를 관리하는 방식이 중심이었습니다.

현재는 대부분 Go Modules를 사용합니다. 그래서 프로젝트 폴더가 어디에 있든 `go.mod`를 기준으로 개발할 수 있습니다.

예를 들어 다음 위치처럼 `GOPATH` 밖에 있어도 괜찮습니다.

```bash
/mnt/d/Litmus/3week/3week-golang
```

대신 `go mod init`을 처음 실행할 때는 모듈 경로를 직접 적어야 합니다.

```bash
go mod init github.com/LitmusChaos-Korea-Mentorship/3week-golang
```

## 15. 자주 만나는 문제

### `go: command not found`

Go가 설치되어 있지 않거나, 설치는 되었지만 `PATH`에 등록되지 않은 상태입니다.

먼저 설치 여부를 확인하고, 새 터미널을 열어 다시 실행합니다.

```bash
go version
```

그래도 동작하지 않으면 공식 설치 문서를 기준으로 Go를 다시 설치합니다.

```text
https://go.dev/doc/install
```

### `go.mod file not found`

현재 디렉터리 또는 상위 디렉터리에 `go.mod`가 없다는 뜻입니다.

저장소 루트로 이동한 뒤 `go.mod`가 있는지 확인합니다.

```bash
cd 3week-golang
ls
```

없다면 처음 한 번만 생성합니다.

```bash
go mod init github.com/LitmusChaos-Korea-Mentorship/3week-golang
```

### `package ... is not in std`

실행 경로나 import 경로가 잘못되었을 가능성이 큽니다.

저장소 루트에서 다음처럼 실행합니다.

```bash
go run ./01-hello
```

또는 해당 폴더 안에서 실행합니다.

```bash
cd 01-hello
go run .
```

## 16. 학습 순서

처음에는 다음 순서로 익히는 것이 좋습니다.

1. `package main`, `func main()` 구조
2. 변수와 타입
3. 조건문과 반복문
4. 함수와 에러 처리
5. 배열, 슬라이스, 맵
6. 구조체와 메서드
7. HTTP 서버와 JSON
8. 모듈, 패키지, 테스트

Go를 배울 때 중요한 것은 문법을 많이 외우는 것보다 작은 프로그램을 직접 실행하면서 `go run`, `go test`, `go fmt` 흐름에 익숙해지는 것입니다.

## 17. 빠른 시작 명령어

Go가 이미 설치되어 있고 저장소를 받은 상태라면 다음 순서로 시작합니다.

```bash
cd 3week-golang
go version
go mod init github.com/LitmusChaos-Korea-Mentorship/3week-golang
go run ./01-hello
go test ./...
```

단, `go.mod`가 이미 있다면 `go mod init`은 생략합니다.
