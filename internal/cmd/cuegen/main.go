package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"cuelang.org/go/cue/errors"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/gocode"
	_ "cuelang.org/go/pkg"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Join(cwd, "graph")
	pkg := "./model"
	inst := cue.Build(load.Instances([]string{pkg}, &load.Config{
		Dir:        dir,
		ModuleRoot: dir,
		Module:     "github.com/bketelsen/gqlops/graph",
	}))[0]
	if err := inst.Err; err != nil {
		log.Fatal("build:", err)
	}

	goPkg := "./graph/model"
	b, err := gocode.Generate(goPkg, inst, nil)
	if err != nil {
		log.Fatal(errStr(err))
	}

	goFile := filepath.Join("graph", "model", "cue_gen.go")
	_ = ioutil.WriteFile(goFile, b, 0644)

	/*	want, err := ioutil.ReadFile(goFile)
		if err != nil {
			log.Fatal(err)
		}

		if d := diff.Diff(string(want), string(b)); d != "" {
			log.Errorf("files differ:\n%v", d)
		}
	*/

}

func errStr(err error) string {
	if err == nil {
		return "nil"
	}
	buf := &bytes.Buffer{}
	errors.Print(buf, err, nil)
	r := regexp.MustCompile(`.cue:\d+:\d+`)
	return r.ReplaceAllString(buf.String(), ".cue:x:x")
}
