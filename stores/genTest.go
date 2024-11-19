package stores

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	featureFile := "features/create_store.feature" // Replace with your .feature file name

	file, err := os.Open(featureFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	stepDefs := map[string]string{
		"Given": "func %s() error {\n    return godog.ErrPending\n}\n",
		"When":  "func %s() error {\n    return godog.ErrPending\n}\n",
		"Then":  "func %s() error {\n    return godog.ErrPending\n}\n",
	}

	re := regexp.MustCompile(`^(Given|When|Then) (.+)$`)

	output := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		match := re.FindStringSubmatch(line)
		if len(match) > 0 {
			stepType := match[1]
			stepText := match[2]
			functionName := strings.Title(strings.ReplaceAll(stepText, " ", ""))
			output = append(output, fmt.Sprintf(stepDefs[stepType], functionName))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Generated Step Definitions:")
	for _, def := range output {
		fmt.Println(def)
	}
}
