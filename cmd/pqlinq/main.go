package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"text/template"

	_ "github.com/lib/pq"
	"golang.org/x/tools/imports"

	"github.com/lebenasa/pqlinq"
	"github.com/lebenasa/pqprobe"
)

func main() {
	connectionString, tableName, packageName, outdir := parseFlags()
	if connectionString == "" || tableName == "" {
		return
	}
	data := genData(connectionString, tableName, packageName)

	// Generating linq code:
	funcs := pqlinq.GenFuncMap()
	funcs["packageName"] = func() string { return data.PackageName }
	funcs["toLower"] = func(in string) string { return strings.ToLower(in) }
	tableBuf := processTemplate("template/table.go.template", "", funcs, data)
	linqBuf := processTemplate("template/linq.go.template", "", funcs, data)
	fieldsBufs := processTemplate("template/fields.go.template", "", funcs, data)
	//filterBufs := map[string][]byte{}
	//for _, field := range data.Table.Fields {
	//	filterBufs[field.Name] = processTemplate("template/filter.go.template", "", funcs, field)
	//}

	// Write output files:
	outdir += "/" + packageName
	createPathIfNotExists(outdir)
	tableOutName := outdir + "/" + packageName + ".go"
	linqOutName := outdir + "/" + packageName + "_linq.go"
	fieldsOutName := outdir + "/" + packageName + "_fields.go"
	writeOutput(tableOutName, tableBuf)
	writeOutput(linqOutName, linqBuf)
	writeOutput(fieldsOutName, fieldsBufs)
	//for name, buf := range filterBufs {
	//	outname := outdir + "/" + name + "_filter.go"
	//	writeOutput(outname, buf)
	//}
}

func genData(connectionString, tableName, packageName string) pqlinq.Data {
	prober, err := pqprobe.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Connecting to database %v: %v", connectionString, err)
	}

	table, err := prober.QueryTable(tableName)
	if err != nil {
		log.Fatalf("Querying table fields: %v", err)
	}

	return pqlinq.NewData(packageName, table)
}

func processTemplate(assetName, outname string, funcs template.FuncMap, data interface{}) (out []byte) {
	code, err := pqlinq.Asset(assetName)
	if err != nil {
		log.Fatalf("%v: Retrieving table template asset: %v", assetName, err)
	}

	templ := template.Must(template.New("table").Funcs(funcs).Parse(string(code)))
	buff := bytes.NewBuffer([]byte{})
	err = templ.Execute(buff, data)
	if err != nil {
		log.Fatalf("%v: Executing table template: %v", assetName, err)
	}

	out, err = imports.Process(outname, buff.Bytes(), nil)
	if err != nil {
		log.Fatalf("%v: Executing goimports at table template: %v", assetName, err)
	}
	return out
}

func createPathIfNotExists(path string) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				log.Fatalf("Creating path %v: %v", path, err)
			}
		} else {
			log.Fatalf("Checking path %v: %v", path, err)
		}
	}
}

func writeOutput(filename string, data []byte) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Creating file %v: %v", filename, err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		log.Fatalf("Writing to file %v: %v", filename, err)
	}
}
