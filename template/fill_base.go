package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"io"
	"os"
	"text/template"

	"github.com/imdario/mergo"
)

var (
	inputOption = flag.String("option", "", "option")
	outputFile  = flag.String("o", "", "output file. stdout is default.")
)

type optInner struct {
	ParentImageName string
	ParentImageTag  string
	UbuntuTag       string
	UseRust         bool
	RustVersion     string
	UseGo           bool
	GoVersion       string
	UseNode         bool
	NodeVersion     string
	UseDeno         bool
	DenoVersion     string
	Dood            bool
	EntryPoint      string
}

type Option struct {
	ParentImageName *string
	ParentImageTag  *string
	UbuntuTag       *string
	UseRust         *bool
	RustVersion     *string
	UseGo           *bool
	GoVersion       *string
	UseNode         *bool
	NodeVersion     *string
	UseDeno         *bool
	DenoVersion     *string
	Dood            *bool
	EntryPoint      *string
}

func (o Option) intoInner() optInner {
	inner := optInner{}

	if o.ParentImageName != nil {
		inner.ParentImageName = *o.ParentImageName
	}
	if o.ParentImageTag != nil {
		inner.ParentImageTag = *o.ParentImageTag
	}
	if o.UbuntuTag != nil {
		inner.UbuntuTag = *o.UbuntuTag
	}
	if o.UseRust != nil {
		inner.UseRust = *o.UseRust
	}
	if o.RustVersion != nil {
		inner.RustVersion = *o.RustVersion
	}
	if o.UseGo != nil {
		inner.UseGo = *o.UseGo
	}
	if o.GoVersion != nil {
		inner.GoVersion = *o.GoVersion
	}
	if o.UseNode != nil {
		inner.UseNode = *o.UseNode
	}
	if o.NodeVersion != nil {
		inner.NodeVersion = *o.NodeVersion
	}
	if o.UseDeno != nil {
		inner.UseDeno = *o.UseDeno
	}
	if o.DenoVersion != nil {
		inner.DenoVersion = *o.DenoVersion
	}
	if o.Dood != nil {
		inner.Dood = *o.Dood
	}
	if o.EntryPoint != nil {
		inner.EntryPoint = *o.EntryPoint
	}

	return inner
}

//go:embed base.template
var baseTemplate string

var defaultInner optInner = optInner{
	UbuntuTag:   "jammy-20230624",
	UseRust:     false,
	RustVersion: "1.71.1",
	UseGo:       false,
	GoVersion:   "1.21.0",
	UseNode:     false,
	NodeVersion: "20.5.0",
	UseDeno:     false,
	DenoVersion: "1.36.0",
	Dood:        false,
}

var defaultOption Option = Option{
	UbuntuTag:   &defaultInner.UbuntuTag,
	UseRust:     &defaultInner.UseRust,
	RustVersion: &defaultInner.RustVersion,
	UseGo:       &defaultInner.UseGo,
	GoVersion:   &defaultInner.GoVersion,
	UseNode:     &defaultInner.UseNode,
	NodeVersion: &defaultInner.NodeVersion,
	UseDeno:     &defaultInner.UseDeno,
	DenoVersion: &defaultInner.DenoVersion,
	Dood:        &defaultInner.Dood,
}

func main() {
	flag.Parse()

	var opt Option
	if *inputOption != "" {
		bin := must(io.ReadAll(must(os.Open(*inputOption))))
		var override Option
		err := json.Unmarshal(bin, &override)
		if err != nil {
			panic(err)
		}
		err = mergo.Merge(&override, defaultOption)
		if err != nil {
			panic(err)
		}
		opt = override
	} else {
		opt = defaultOption
	}

	var output io.WriteCloser
	if *outputFile != "" {
		output = must(os.Create(*outputFile))
	} else {
		output = os.Stdout
	}

	t := template.Must(template.New("foo").Parse(baseTemplate))
	if err := t.Execute(output, opt.intoInner()); err != nil {
		panic(err)
	}
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
