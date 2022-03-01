package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//SetConfigVars is a...
func SetConfigVars() {

	// filepath = "development.json"
	absfilepath, err := filepath.Abs(os.Getenv("HOME") + "/go/src/oms/config/development.json")
	fmt.Println("absfilepath", absfilepath)
	// dir, err := os.Getwd()
	fmt.Println("err", err)
	if os.Getenv("GOENV") == "production" {
		// filepath = "production.json"
		absfilepath, _ = filepath.Abs(os.Getenv("HOME") + "/go/src/oms/config/production.json")
	}

	cfile, err := ioutil.ReadFile(absfilepath)
	if err != nil {
		fmt.Printf("// error while reading file %s\n", absfilepath)
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	configvals := make(map[string]string)
	err1 := json.Unmarshal(cfile, &configvals)
	if err1 != nil {
		fmt.Println("error:", err1)
	}

	for key := range configvals {
		os.Setenv(key, configvals[key])
	}
}
