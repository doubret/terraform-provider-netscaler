package main

import (
	"github.com/doubret/citrix-netscaler-nitro-go-specs/specs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func tf_schema_type(fieldType string) string {
	fieldType = strings.TrimSuffix(fieldType, "[]")

	isEnum := nitro.IsEnum(fieldType)

	fieldType = strings.TrimPrefix(fieldType, "(")
	fieldType = strings.TrimSuffix(fieldType, ")")

	base := "TypeString"

	switch fieldType {
	case "ip", "ip_mask", "string":
		base = "TypeString"
	case "double", "int":
		base = "TypeInt"
	case "bool":
		base = "TypeBool"
	}

	if isEnum {
		base = "TypeString"
	}

	return base
}

func is_in(value string, array []string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}

	return false
}

func main() {
	spec, err := nitro.ReadSpec("C:/Users/ce.breteche/go/src/github.com/doubret/citrix-netscaler-nitro-yaml-specs/yml")

	if err != nil {
		log.Println("Failed to read spec : ", err)
	} else {
		funcMap := template.FuncMap{
			"title":          strings.Title,
			"is_array":       nitro.IsArray,
			"go_type":        nitro.GoType,
			"name":           nitro.Name,
			"tf_schema_type": tf_schema_type,
			"is_in":          is_in,
		}

		templates := template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/resource.tmpl", "templates/binding.tmpl", "templates/provider.tmpl"))

		for key, value := range spec.Resources {
			context := struct {
				Name   string
				Schema *nitro.Resource
			}{
				key,
				value,
			}

			writer, err := os.Create(filepath.Join("netscaler", "resources", key+".go"))

			if err != nil {
				log.Println("Failed to create file : ", err)
			}

			err = templates.ExecuteTemplate(writer, "resource.tmpl", context)

			if err != nil {
				log.Println("Failed to execute template : ", err)
			}
		}

		for key, value := range spec.Bindings {
			context := struct {
				Name   string
				Schema *nitro.Binding
			}{
				key,
				value,
			}

			writer, err := os.Create(filepath.Join("netscaler", "bindings", key+".go"))

			if err != nil {
				log.Println("Failed to create file : ", err)
			}

			err = templates.ExecuteTemplate(writer, "binding.tmpl", context)

			if err != nil {
				log.Println("Failed to execute template : ", err)
			}
		}

		writer, err := os.Create(filepath.Join("netscaler", "provider.go"))
		err = templates.ExecuteTemplate(writer, "provider.tmpl", spec)
		if err != nil {
			log.Println("Failed to execute provider template : ", err)
		}

		return
	}
}
