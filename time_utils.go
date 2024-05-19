package goutils

import (
	"strconv"
	"time"
)

// UtcOffSet ajusta o fuso horário de uma data para UTC de acordo com o valor passado
func UtcOffSet(value *time.Time, utc any) {
	switch utc.(type) {
	case int:
		location := time.FixedZone("UTC", utc.(int)*60*60)
		*value = value.In(location)
	case string:
		utcInt, _ := strconv.Atoi(utc.(string))
		location := time.FixedZone("UTC", int(utcInt)*60*60)
		*value = value.In(location)
	}
}
