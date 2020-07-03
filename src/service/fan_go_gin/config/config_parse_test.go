package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	conf := &Config{
		HttpPort: "9000",
		MySQLHost: "127.0.0.1:3006",
	}
	fmt.Println("----- vv -------------")
	value := reflect.ValueOf(conf)

	valueElum := value.Elem() //conf 的值
	valueFieldQuantity := valueElum.NumField() //conf 的值的个数
	valueFieldType := valueElum.Type() //conf 的值的类型

	fmt.Println(value)
	fmt.Println("vv.Elem:", valueElum)
	fmt.Println("vft:", valueFieldType)
	fmt.Println("vv.Elem.NumField:", valueFieldQuantity)

	for i := 0;i<valueFieldQuantity;i++ {
		tempElumField := valueElum.Field(i)
		tempElumFieldType := valueFieldType.Field(i)
		filedTag := tempElumFieldType.Tag
		fmt.Println("\t ", tempElumField, tempElumFieldType, ", tag:", filedTag)
		fmt.Println("string:", tempElumField.String())
	}





	fmt.Println("------------------")
	tt := reflect.TypeOf(conf)
	fmt.Println(tt)




}
