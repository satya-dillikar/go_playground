package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

var (
	ErrFileNotFound  = fmt.Errorf("file not found")
	ErrJsonUnmarshal = fmt.Errorf("failed to unmarshal")
)

type errorOne struct{}

func (e errorOne) Error() string {
	return "Error One happended"
}

func main() {
	main1()
	//main2()
}
func main2() {
	var err1 errorOne
	err2 := do10()
	if err1 == err2 {
		fmt.Println("Equality Operator: Both errors are equal")
	}
	if errors.Is(err1, err2) {
		fmt.Println("Is function: Both errors are equal")
	}
}
func do10() error {
	return errorOne{}
}

func main1() {
	err1 := errorOne{}

	err2 := do2()

	err3 := do3()
	fmt.Println("err1: ", err1)
	fmt.Println("err2: ", err2)
	fmt.Println("err3: ", err3)
	if err1 == err2 {
		fmt.Println("Equality Operator1: Both errors are equal")
	} else {
		fmt.Println("Equality Operator1: Both errors are not equal")
	}

	if errors.Is(err2, ErrFileNotFound) {
		fmt.Println("Equality Operator2: Both errors are equal")
	} else {
		fmt.Println("Equality Operator2: Both errors are not equal")
	}
	if errors.Is(err3, ErrJsonUnmarshal) {
		fmt.Println("Equality Operator3: Both errors are equal")
	} else {
		fmt.Println("Equality Operator3: Both errors are not equal")
	}
}

func do() error {
	return fmt.Errorf("E2: %w", errorOne{})
}

func do2() error {
	pluginConfigPath := ""
	_, err := ioutil.ReadFile(pluginConfigPath)
	if err != nil {
		fmt.Println("read failed", err)
		return fmt.Errorf("plugin-config-path: %s, error: %w", pluginConfigPath, ErrFileNotFound)
	}
	return nil
}

func do3() error {
	var config string
	var pluginConfig []byte

	err := json.Unmarshal([]byte(pluginConfig), &config)
	if err != nil {
		fmt.Println("unmarshal failed", err)
		return fmt.Errorf("plugin config : %s, error: %w", string(pluginConfig), ErrJsonUnmarshal)
	}
	return nil
}
