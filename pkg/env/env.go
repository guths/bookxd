package env

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func readEnv(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	defer file.Close()

	stats, err := file.Stat()

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	fmt.Printf("Env name: %s\n", stats.Name())

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, " ")

		if len(words) > 1 {
			return fmt.Errorf("env have invalid format, possible blank spaces after declaration")
		}

		input := words[0]

		re := regexp.MustCompile(`^([^=]+)=(.*)$`)
		match := re.FindStringSubmatch(input)

		if len(match) == 3 {
			variableName := match[1]
			variableValue := match[2]

			err := validateEnvVariable(variableName, variableValue)

			if err != nil {
				return fmt.Errorf(err.Error())
			}

			setEnvVariable(variableName, variableValue)
		} else {
			fmt.Println("Invalid input format.")
		}

	}

	return nil
}

func setEnvVariable(name, value string) error {
	err := os.Setenv(name, value)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func validateEnvVariable(name, value string) error {
	osEnvRule := "[a-zA-Z_]{1,}[a-zA-Z0-9_]"

	re := regexp.MustCompile(osEnvRule)

	if !re.MatchString(name) {
		return fmt.Errorf("%s has invalid format", name)
	}

	if !re.MatchString(value) {
		return fmt.Errorf("value: %s of %s variable, has invalid format", value, name)
	}

	return nil
}

func SetEnvs() {
	filename := ".env.development"
	err := readEnv(filename)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
