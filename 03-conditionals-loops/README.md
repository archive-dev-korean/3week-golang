# 03. Conditionals And Loops

## 목표

- `if`, `else if`, `else`로 조건을 분기합니다.
- `switch`로 값에 따라 분기합니다.
- Go의 반복문인 `for`를 여러 형태로 사용합니다.

## 실행

```bash
go run ./03-conditionals-loops
```

## 핵심 설명

- Go에는 `while` 키워드가 없습니다. `for 조건` 형태로 while처럼 사용합니다.
- `[]string{"apple", "banana", "cherry"}`는 문자열 여러 개를 담는 슬라이스입니다. `string` 앞의 `[]`는 값이 여러 개 들어가는 목록 형태라는 뜻입니다.
    - 슬라이스는 배열과 비슷하지만 길이를 나중에 늘릴 수 있어 Go에서 목록 데이터를 다룰 때 자주 사용합니다.
- `range`는 슬라이스, 배열, 맵 등을 순회할 때 사용합니다.
- `for index, value := range fruits`에서 `index`는 위치 번호이고 `value`는 해당 위치의 실제 값입니다.
- `switch`의 `case`는 기본적으로 자동으로 빠져나옵니다. 다른 언어처럼 `break`를 매번 쓰지 않아도 됩니다.

## `[]string` 예제

```go
fruits := []string{"apple", "banana", "cherry"}
```

위 코드는 문자열 슬라이스를 만들고 `fruits` 변수에 저장합니다.

- `[]string`: 문자열을 여러 개 담는 슬라이스 타입
- `{...}`: 처음 넣을 값 목록
- `"apple"`, `"banana"`, `"cherry"`: 슬라이스에 들어가는 문자열 값

이 슬라이스는 `range`와 함께 순회할 수 있습니다.

```go
for index, value := range fruits {
	fmt.Printf("[%d] %s\n", index, value)
}
```

## 실습

1부터 100까지 숫자 중 짝수만 출력하는 반복문을 추가해 보세요.
