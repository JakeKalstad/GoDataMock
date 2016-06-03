package generate

import (
	"math/rand"
	"reflect"
	"time"
)

type Generatorer interface {
	Get(data interface{}, num int) []interface{}
}

type Generator struct {
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func rString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func rInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(36)
}

func rFloat64() float64 {
	return float64(rInt())
}

//Get(myDataArray, number of items)
func (g *Generator) Get(data interface{}, num int) interface{} {
	original := reflect.ValueOf(data)
	copy := reflect.New(original.Type()).Elem()
	translateRecursive(copy, original, num)
	return copy.Interface()
}

func translateRecursive(copy, original reflect.Value, num int) {
	switch original.Kind() {
	case reflect.Ptr:
		originalValue := original.Elem()
		if !originalValue.IsValid() {
			return
		}
		copy.Set(reflect.New(originalValue.Type()))
		translateRecursive(copy.Elem(), originalValue, num)

	case reflect.Interface:
		originalValue := original.Elem()
		if !originalValue.IsValid() {
			return
		}
		copyValue := reflect.New(originalValue.Type()).Elem()
		translateRecursive(copyValue, originalValue, num)
		copy.Set(copyValue)
	case reflect.Struct:
		for i := 0; i < original.NumField(); i++ {
			translateRecursive(copy.Field(i), original.Field(i), num)
		}

	case reflect.Slice:
		copy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < num; i++ {
			translateRecursive(copy.Index(0), original.Index(0), num)
		}
	case reflect.Map:
		copy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			translateRecursive(copyValue, originalValue, num)
			copy.SetMapIndex(key, copyValue)
		}
	case reflect.String:
		copy.SetString(rString(rInt()))
	case reflect.Int:
		copy.SetInt(int64(rInt()))
	case reflect.Bool:
		copy.SetBool(rInt()%2 == 0)
	case reflect.Float64:
		copy.SetFloat((rFloat64() + .03) * (rFloat64() + .12))
	default:
		copy.Set(original)
	}
}
