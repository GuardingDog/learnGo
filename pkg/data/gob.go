package data

import (
	"encoding/gob"
	"fmt"
	"os"
)

// Person 示例DemoStrcut
type Person struct {
	Name string
	Age  int
}

// GobDemo Gob 使用示例
func GobDemo() {
	p := Person{
		Name: "Alice",
		Age:  30,
	}

	err := encodePerson(p, "person.gob")
	if err != nil {
		fmt.Println("编码失败：", err)
		return
	}

	fmt.Println("编码成功！")

	decodedPerson, err := decodePerson("person.gob")
	if err != nil {
		fmt.Println("解码失败：", err)
		return
	}

	fmt.Println("解码成功！")
	fmt.Println("解码后的Person对象：", decodedPerson)
}

// 编码Person对象为Gob格式，并写入文件
func encodePerson(p Person, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(p)
	if err != nil {
		return err
	}

	return nil
}

// 从文件中读取Gob数据，并解码为Person对象
func decodePerson(filename string) (Person, error) {
	var decodedPerson Person

	file, err := os.Open(filename)
	if err != nil {
		return decodedPerson, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&decodedPerson)
	if err != nil {
		return decodedPerson, err
	}

	return decodedPerson, nil
}
