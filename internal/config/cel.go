package config

import (
	"errors"

	"github.com/google/cel-go/cel"
)

func NewCelAstMap(c FilterStruct) (*map[string]cel.Program, error) {
	celAstMap := make(map[string]cel.Program)

	var errRet error

	for t, v := range c.Tables {
		for _, vv := range v {
			switch d := vv.(type) {
			case string:
				continue

			case map[string]string:
				if cond, ok := d["condition"]; ok {
					env, err := cel.NewEnv(
						cel.Variable("row", cel.MapType(cel.StringType, cel.DynType)),
					)
					if err != nil {
						errRet = err
						break
					}

					ast, iss := env.Compile(cond)
					if iss.Err() != nil {
						errRet = iss.Err()
						break
					}

					if ast.IsChecked() && ast.OutputType() != cel.BoolType {
						errRet = errors.New("output must be boolean")
						break
					}

					prg, err := env.Program(ast)
					if err != nil {
						errRet = err
						break
					}

					op, _ := d["operation"]

					celAstMap[t+"."+op] = prg
				}
			}
		}
	}

	return &celAstMap, errRet
}
