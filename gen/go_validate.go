package gen

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/gocode"
)

func GenerateGoValidators() error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd: %w", err)
	}

	apiValidatorsPath := path.Join(wd, "api")
	cueBuildCtx := cuecontext.New()
	instances := load.Instances([]string{"./api"}, &load.Config{})
	cVals, err := cueBuildCtx.BuildInstances(instances)
	if err != nil {
		return fmt.Errorf("processing cue files in 'api': %w", err)
	}
	for i, cv := range cVals {
		genconf := gocode.Config{
			Prefix:       "cuegen",
			ValidateName: "CUEValidate",
			CompleteName: "CUEComplete",
			RuntimeVar:   "",
		}

		generatedCode, err := gocode.Generate(apiValidatorsPath, cv, &genconf)
		if err != nil {
			jsn, jsonErr := cv.MarshalJSON()
			return fmt.Errorf("gocode.Generate: in %q %s (jsonErr=%v): %w", apiValidatorsPath, jsn, jsonErr, err)
		}

		filename := fmt.Sprintf("zz_cue_validators.%d.go", i)
		if len(cVals) == 1 {
			filename = "zz_cue_validators.go"
		}
		outputFileName := path.Join(apiValidatorsPath, filename)
		if err := os.WriteFile(outputFileName, generatedCode, fs.FileMode(0644)); err != nil {
			return fmt.Errorf("writing generated code to file %q: %w", outputFileName, err)
		}
	}

	// TODO fmt file
	// TODO clean old files

	return nil
}
