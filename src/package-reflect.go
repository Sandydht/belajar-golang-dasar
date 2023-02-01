package src

import (
	"fmt"
	"reflect"
)

/**
Package Relfect
- Dalam bahasa pemrograman, biasanya ada fitur reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan
- Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect
- Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
- https://golang.org/pkg/reflect/
*/

type Sample struct {
	Name        string `required:"true" max:"10"`
	Description string `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func PackageReflect() {
	sample := Sample{"Sandy", "Fullstack Software Engineer"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	required := structField.Tag.Get("required")
	max := structField.Tag.Get("max")
	fmt.Println("structField:", structField.Name)
	fmt.Println("required:", required)
	fmt.Println("max:", max)

	sample.Name = ""
	fmt.Println("IsValid:", IsValid(sample))
}
