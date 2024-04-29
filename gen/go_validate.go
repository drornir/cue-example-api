package gen

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/gocode"
)

func GenerateGoValidators() error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd: %w", err)
	}

	apiValidatorsPath := path.Join(wd, "api")

	genconf := gocode.Config{
		Prefix:       "",
		ValidateName: "",
		CompleteName: "",
		RuntimeVar:   "",
	}

	var instance cue.InstanceOrValue

	generatedCode, err := gocode.Generate(apiValidatorsPath, instance, &genconf)
	if err != nil {
		return fmt.Errorf("gocode.Generate: %w", err)
	}

	outputFileName := path.Join(apiValidatorsPath, "zz_api_validators.go")
	if err := os.WriteFile(outputFileName, generatedCode, fs.FileMode(0644)); err != nil {
		return fmt.Errorf("writing generated code to file %q: %w", outputFileName, err)
	}

	// TODO fmt file

	return nil
}
