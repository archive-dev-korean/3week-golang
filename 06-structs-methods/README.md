# 06. Structs And Methods

## 목표

- 구조체로 관련 데이터를 하나의 타입으로 묶습니다.
- 값 리시버와 포인터 리시버 메서드를 구분합니다.
- 구조체 필드에 접근하고 값을 변경합니다.

## 실행

```bash
go run ./06-structs-methods
```

## 핵심 설명

- `type Person struct`는 새로운 구조체 타입을 정의합니다.
    - 마찬가지로 변수 선언시 순서 주의하기 !
- `func (p Person) Introduce()`는 `Person` 값에 연결된 메서드입니다.
- `func (p *Person) HaveBirthday()`는 포인터 리시버를 사용하므로 원본 값을 변경할 수 있습니다.

## 실습

`Person`에 `Email` 필드를 추가하고 `Introduce` 출력에 포함해 보세요.
