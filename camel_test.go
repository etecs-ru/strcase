/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Ian Coleman
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package strcase

import (
	"testing"
)

func toCamel(tb testing.TB) {
	cases := [][]string{
		{"test_case", "TestCase"},
		{"test.case", "TestCase"},
		{"test", "Test"},
		{"TestCase", "TestCase"},
		{" test  case ", "TestCase"},
		{"", ""},
		{"many_many_words", "ManyManyWords"},
		{"AnyKind of_string", "AnyKindOfString"},
		{"odd-fix", "OddFix"},
		{"numbers2And55with000", "Numbers2And55With000"},
		{"ID", "Id"},
		{"айди", "Айди"},
		{"нечётный-фикс", "НечётныйФикс"},
		{"Любая например_строка", "ЛюбаяНапримерСтрока"},
		{"много_много_слов", "МногоМногоСлов"},
		{" тестовый  кейс ", "ТестовыйКейс"},
		{"ТестовыйКейс", "ТестовыйКейс"},
		{"тестовый_кейс", "ТестовыйКейс"},
		{"тестовый.кейс", "ТестовыйКейс"},
		{"тест", "Тест"},
		{"@смешная-собачка", "СмешнаяСобачка"},
		{"плохо.распарсил[0].джейсон", "ПлохоРаспарсил0Джейсон"},
		{"!~сломанное-сообщение.один", "СломанноеСообщениеОдин"},
		{"!~сломанное-сообщение.!~&один", "СломанноеСообщениеОдин"},
		{"!~сломанное-*^$(*4сообщение.!~&один", "Сломанное4СообщениеОдин"},
		{"3!@#$%^&*f(&o^%o$#5b@#$%^&a*&^%$r", "3Foo5Bar"},
		{"3!@#$%^&*К(&о^%ш$ка#5м@#$%^&ы*&^%$шка", "3Кошка5Мышка"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToCamel(in)
		if result != out {
			tb.Errorf("%q (%q != %q)", in, result, out)
		}
	}
}

func TestToCamel(t *testing.T) {
	toCamel(t)
}

func BenchmarkToCamel(b *testing.B) {
	benchmarkCamelTest(b, toCamel)
}

func toLowerCamel(tb testing.TB) {
	cases := [][]string{
		{"foo-bar", "fooBar"},
		{"TestCase", "testCase"},
		{"", ""},
		{"AnyKind of_string", "anyKindOfString"},
		{"AnyKind.of-string", "anyKindOfString"},
		{"ID", "id"},
		{"some string", "someString"},
		{" some string", "someString"},
		{"любая строка", "любаяСтрока"},
		{" некая Строка", "некаяСтрока"},
		{"Идентификатор", "идентификатор"},
		{"ЛюбойТип любой_строки", "любойТипЛюбойСтроки"},
		{"ЛюбойТип.любой-строки", "любойТипЛюбойСтроки"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToLowerCamel(in)
		if result != out {
			tb.Errorf("%q (%q != %q)", in, result, out)
		}
	}
}

func TestToLowerCamel(t *testing.T) {
	toLowerCamel(t)
}

func TestCustomAcronymsToCamel(t *testing.T) {
	tests := []struct {
		name         string
		acronymKey   string
		acronymValue string
		expected     string
	}{
		{
			name:         "API Custom Acronym",
			acronymKey:   "API",
			acronymValue: "api",
			expected:     "Api",
		},
		{
			name:         "ABCDACME Custom Acroynm",
			acronymKey:   "ABCDACME",
			acronymValue: "AbcdAcme",
			expected:     "AbcdAcme",
		},
		{
			name:         "PostgreSQL Custom Acronym",
			acronymKey:   "PostgreSQL",
			acronymValue: "PostgreSQL",
			expected:     "PostgreSQL",
		},
		{
			name:         "Кириллица Custom Acronym",
			acronymKey:   "Кириллица",
			acronymValue: "Кириллица",
			expected:     "Кириллица",
		},
		{
			name:         "Кириллица-крлц Custom Acronym",
			acronymKey:   "Кириллица",
			acronymValue: "крлц",
			expected:     "Крлц",
		},
	}
	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			ConfigureAcronym(test.acronymKey, test.acronymValue)
			if result := ToCamel(test.acronymKey); result != test.expected {
				t.Errorf("expected custom acronym result %s, got %s", test.expected, result)
			}
		})
	}
}

func TestCustomAcronymsToLowerCamel(t *testing.T) {
	tests := []struct {
		name         string
		acronymKey   string
		acronymValue string
		expected     string
	}{
		{
			name:         "API Custom Acronym",
			acronymKey:   "API",
			acronymValue: "api",
			expected:     "api",
		},
		{
			name:         "ABCDACME Custom Acroynm",
			acronymKey:   "ABCDACME",
			acronymValue: "AbcdAcme",
			expected:     "abcdAcme",
		},
		{
			name:         "PostgreSQL Custom Acronym",
			acronymKey:   "PostgreSQL",
			acronymValue: "PostgreSQL",
			expected:     "postgreSQL",
		},
	}
	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			ConfigureAcronym(test.acronymKey, test.acronymValue)
			if result := ToLowerCamel(test.acronymKey); result != test.expected {
				t.Errorf("expected custom acronym result %s, got %s", test.expected, result)
			}
		})
	}
}

func BenchmarkToLowerCamel(b *testing.B) {
	benchmarkCamelTest(b, toLowerCamel)
}

func benchmarkCamelTest(b *testing.B, fn func(testing.TB)) {
	for n := 0; n < b.N; n++ {
		fn(b)
	}
}
