package main

import (
	"fmt"
	"io/ioutil"
	"os"

	goflags "github.com/jessevdk/go-flags"

	"golang.org/x/mod/modfile"
)

type Flags struct {
	Help bool `short:"h" long:"help" description:"display this help"`
}

func main() {
	progName := "gomodconflict"
	if len(os.Args) > 0 {
		progName = os.Args[0]
	}
	usage := fmt.Sprintf("%s [options] <go.mod> [<go.mod> [...]]", progName)

	flags := Flags{}
	p := goflags.NewNamedParser("", goflags.PrintErrors|goflags.PassDoubleDash|goflags.PassAfterNonOption)
	p.AddGroup(usage, "", &flags)
	goPaths, err := p.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to parse flags: %s\n", err)
		os.Exit(1)
	}
	if flags.Help {
		p.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	numPaths := len(goPaths)
	if numPaths < 1 {
		fmt.Fprintf(os.Stderr, "%s\n", usage)
		os.Exit(1)
	}

	deps := map[string][]string{}
	for i, goModPath := range goPaths {
		requires, err := goModParse(goModPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to parse %s: %s\n", goModPath, err.Error())
			os.Exit(1)
		}
		for k, v := range requires {
			if _, ok := deps[k]; !ok {
				deps[k] = make([]string, numPaths)
			}
			deps[k][i] = v
		}
	}

	exitCode := 0
	for depPath, vals := range deps {
		ok := true
		expectedVal := ""
		for _, val := range vals {
			if val == "" {
				continue
			}
			if expectedVal == "" {
				expectedVal = val
				continue
			}
			if expectedVal != val {
				ok = false
				exitCode = 1
			}
		}
		if !ok {
			fmt.Printf("%s\n", depPath)
			for i, val := range vals {
				if val != "" {
					fmt.Printf("  %s: %s\n", goPaths[i], val)
				}
			}
		}
	}

	os.Exit(exitCode)
}

func goModParse(path string) (map[string]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	gm, err := modfile.Parse(path, data, nil)
	if err != nil {
		return nil, err
	}

	results := map[string]string{}
	for _, require := range gm.Require {
		results[require.Mod.Path] = require.Mod.Version
	}
	for _, replace := range gm.Replace {
		results[replace.Old.Path] = fmt.Sprintf("%s=>%s", replace.New.Path, replace.New.Version)
	}
	return results, nil
}
