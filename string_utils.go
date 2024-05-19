package goutils

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// RemoveDiacritics remove acentos de uma string
func RemoveDiacritics(value string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, value)
	return result
}

// IsNilOrWhiteSpace verifica se uma string é nil ou vazia
func IsNilOrWhiteSpace(value *string) bool {
	return value == nil || strings.TrimSpace(*value) == ""
}

// NotIsNilOrWhiteSpace verifica se uma string não é nil ou vazia
func NotIsNilOrWhiteSpace(value *string) bool {
	return !IsNilOrWhiteSpace(value)
}

// ExtractNumbers extrai números de uma string
func ExtractNumbers(value string) string {
	if IsNilOrWhiteSpace(&value) {
		return ""
	}

	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(value, -1)
	return strings.Join(matches, "")
}

// IsOnlyNumber verifica se uma string contém apenas números
func IsOnlyNumber(value string) bool {
	if IsNilOrWhiteSpace(&value) {
		return false
	}
	return len(ExtractNumbers(value)) == len(value)
}

// ConvertToFloat64 converte uma string para float64
func ConvertToFloat64(value string) (float64, error) {
	cleaned := strings.ReplaceAll(value, ".", "")
	cleaned = strings.ReplaceAll(cleaned, ",", ".")
	return strconv.ParseFloat(cleaned, 64)
}

// Truncate trunca uma string para um tamanho máximo, se necessário
func Truncate(value string, maxLength int) string {
	if len(value) <= maxLength {
		return value
	}
	return value[:maxLength]
}

// FormatCpfCnpj formata um CPF ou CNPJ
func FormatCpfCnpj(value string) string {
	cpfPattern := `^(\d{3})(\d{3})(\d{3})(\d{2})$`
	cnpjPattern := `^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})$`

	reCpf := regexp.MustCompile(cpfPattern)
	reCnpj := regexp.MustCompile(cnpjPattern)

	if reCpf.MatchString(value) {
		return reCpf.ReplaceAllString(value, "$1.$2.$3-$4")
	} else if reCnpj.MatchString(value) {
		return reCnpj.ReplaceAllString(value, "$1.$2.$3/$4-$5")
	}
	return value
}
