# 04. Functions And Errors

## 목표

- 함수를 선언하고 호출합니다.
- Go의 다중 반환값을 사용합니다.
- `error` 값을 검사하는 기본 패턴을 익힙니다.

## 실행

```bash
go run ./04-functions-errors
```

## 핵심 설명

- `func greet(name string) string`은 문자열을 받아 문자열을 반환합니다.
- `divide`는 결과값과 에러를 함께 반환합니다.
- Go에서는 예외 처리 대신 `if err != nil` 패턴으로 실패를 직접 확인하는 코드가 흔합니다.

## 실습

두 정수를 받아 더한 값을 반환하는 `add` 함수를 만들고 `main`에서 호출해 보세요.
