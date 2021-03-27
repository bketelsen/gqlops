package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/bketelsen/gqlops/graph/model"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

func load() []*model.Profile {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)
	files, err := os.ReadDir("data/profiles/")
	if err != nil {
		panic(err)
	}
	var profiles []*model.Profile
	for _, file := range files {
		slug := strings.Replace(file.Name(), ".yaml", "", -1)
		bb, err := os.ReadFile(path.Join("data", "profiles", file.Name()))
		if err != nil {
			panic(err)
		}
		var buf bytes.Buffer
		context := parser.NewContext()
		if err := markdown.Convert([]byte(bb), &buf, parser.WithContext(context)); err != nil {
			panic(err)
		}
		metaData := meta.Get(context)
		title := metaData["name"]
		fmt.Println(slug, title)
		age := metaData["age"]

		profile := &model.Profile{ID: slug, Age: age.(int), Name: title.(string)}
		fmt.Println("validating")
		err = profile.Validate()
		fmt.Println(profile)

		if err != nil {
			panic(err)
		}
		profiles = append(profiles, profile)
	}

	return profiles
}
