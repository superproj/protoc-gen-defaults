// Copyright 2021 Linka Cloud  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package module

import (
	"fmt"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	"go.linka.cloud/protoc-gen-defaults/defaults"
)

func Defaults() *Module {
	return &Module{
		ModuleBase: &pgs.ModuleBase{},
		imports:    make(map[string]struct{}),
		oneOfs:     make(map[string]struct{}),
	}
}

type Module struct {
	*pgs.ModuleBase
	ctx     pgsgo.Context
	tpl     *template.Template
	imports map[string]struct{}
	oneOfs  map[string]struct{}
}

func (m *Module) Name() string {
	return "defaults"
}

func (m *Module) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("fields").Funcs(map[string]interface{}{
		"package": m.ctx.PackageName,
		"name":    m.ctx.Name,
		"comment": func(s string) string {
			var out string
			parts := strings.Split(s, "\n")
			for i, v := range parts {
				if i == len(parts)-1 && v == "" {
					return out
				}
				out += "// " + v + "\n"
			}
			return out
		},
		"imports": func() string {
			var imports string
			for v := range m.imports {
				imports += fmt.Sprintf("\"%s\"\n", v)
			}
			return imports
		},
		"gen": func(m pgs.Message) bool {
			var ignored bool
			ok, err := m.Extension(defaults.E_Ignored, &ignored)
			if err != nil || !ok {
				return true
			}
			return !ignored
		},
		"enabled": func(m pgs.Message) bool {
			var disabled bool
			ok, err := m.Extension(defaults.E_Disabled, &disabled)
			if err != nil || !ok {
				return true
			}
			return !disabled
		},
		"defaults": func(f pgs.Field) string {
			v, _ := m.genFieldDefaults(f)
			return v
		},
	})
	m.tpl = template.Must(tpl.Parse(defaultsTpl))
}

func (m *Module) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
	for _, f := range targets {
		m.generate(f)
	}
	return m.Artifacts()
}

func (m *Module) generate(f pgs.File) {
	if len(f.Messages()) == 0 {
		return
	}
	for _, msg := range f.Messages() {
		m.Check(msg)
	}
	name := m.ctx.OutputPath(f).SetExt(".defaults.go")
	m.AddGeneratorTemplateFile(name.String(), m.tpl, f)
}

func (m *Module) isOneOfDone(oneOf pgs.OneOf) bool {
	_, done := m.oneOfs[oneOf.FullyQualifiedName()]
	return done
}

func (m *Module) setOneOfDone(oneOf pgs.OneOf) {
	if oneOf == nil {
		return
	}
	m.oneOfs[oneOf.FullyQualifiedName()] = struct{}{}
}

const defaultsTpl = `{{ comment .SyntaxSourceCodeInfo.LeadingComments }}
{{ range .SyntaxSourceCodeInfo.LeadingDetachedComments }}
{{ comment . }}
{{ end }}
// Code generated by protoc-gen-defaults. DO NOT EDIT.

package {{ package . }}

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	{{ imports }}
)

var (
	_ *timestamppb.Timestamp
	_ *durationpb.Duration
	_ *wrapperspb.BoolValue
)

{{ range .AllMessages }}

{{ if gen . }}
func (x *{{ name . }}) Default() {
	{{- if enabled . }}
		{{- range .Fields }}
			{{- defaults . }}
		{{- end }}
	{{- end }} 
}
{{- end }}
{{ end }}
`
