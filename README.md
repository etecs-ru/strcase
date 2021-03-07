# strcase
[![Go Reference](https://pkg.go.dev/badge/github.com/etecs-ru/strcase.svg)](https://pkg.go.dev/github.com/etecs-ru/strcase)
[![golangci-lint](https://github.com/etecs-ru/strcase/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/etecs-ru/strcase/actions/workflows/golangci-lint.yml)
[![Coverage](https://codecov.io/gh/etecs-ru/strcase/branch/master/graph/badge.svg?token=INN6PQXII4)](https://codecov.io/gh/etecs-ru/strcase)
[![Go Report Card](https://goreportcard.com/badge/github.com/etecs-ru/strcase)](https://goreportcard.com/report/github.com/etecs-ru/strcase)

#### This is the fork of [iancoleman/strcase](https://github.com/iancoleman/strcase) adding support for cyrillic letters in CamelCase. Поддерживает трансляцию в CamelCase, например `Ключ.Значение`. API пакета не менялось, примеры использования можно найти в `camel_test.go`.


strcase is a go package for converting string case to various cases (e.g. [snake case](https://en.wikipedia.org/wiki/Snake_case) or [camel case](https://en.wikipedia.org/wiki/CamelCase)) to see the full conversion table below.

## Example

```go
s := "AnyKind of_string"
```

| Function                                  | Result               |
|-------------------------------------------|----------------------|
| `ToSnake(s)`                              | `any_kind_of_string` |
| `ToSnakeWithIgnore(s, '.')`               | `any_kind.of_string` |
| `ToScreamingSnake(s)`                     | `ANY_KIND_OF_STRING` |
| `ToKebab(s)`                              | `any-kind-of-string` |
| `ToScreamingKebab(s)`                     | `ANY-KIND-OF-STRING` |
| `ToDelimited(s, '.')`                     | `any.kind.of.string` |
| `ToScreamingDelimited(s, '.', '', true)`  | `ANY.KIND.OF.STRING` |
| `ToScreamingDelimited(s, '.', ' ', true)` | `ANY.KIND OF.STRING` |
| `ToCamel(s)`                              | `AnyKindOfString`    |
| `ToLowerCamel(s)`                         | `anyKindOfString`    |


## Install

```bash
go get -u github.com/iancoleman/strcase
```

## Custom Acronyms for ToCamel && ToLowerCamel

Often times text can contain specific acronyms which you need to be handled a certain way.
Out of the box `strcase` treats the string "ID" as "Id" or "id" but there is no way to cater
for every case in the wild.

To configure your custom acronym globally you can use the following before running any conversion

```go
package foo

import (
    "github.com/iancoleman/strcase"
)

func init() {
    // results in "Api" using ToCamel("API")
    // results in "api" using ToLowerCamel("API")
    strcase.ConfigureAcronym("API", "api")
    
    // results in "PostgreSQL" using ToCamel("PostgreSQL")
    // results in "postgreSQL" using ToLowerCamel("PostgreSQL")
    strcase.ConfigureAcronym("PostgreSQL", "PostgreSQL")

}

```
