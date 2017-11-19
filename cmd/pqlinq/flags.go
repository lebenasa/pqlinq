package main

import (
	"flag"
	"fmt"
)

func usage() {
	fmt.Println("Generate linq code to access table in specific database.")
	fmt.Println("Will generate package folder inside output directory.")
	fmt.Println("Usage: pqlinq [database connection string] [table name]")
	fmt.Println("Example: pqlinq postgres://user:pass@host/database musics")
	flag.PrintDefaults()
}

func parseFlags() (connectionString, tableName, packageName, outputDir string) {
	flag.Usage = usage
	flagPackage := flag.String("package", "", "package name for generated code, defaults to table name")
	flagPackageShort := flag.String("p", "", "package name for generated code, defaults to table name")
	flagOutput := flag.String("output", "", "output directory, defaults to current directory")
	flagOutputShort := flag.String("o", "", "output directory, defaults to current directory")
	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		return
	}

	connectionString = flag.Arg(0)
	tableName = flag.Arg(1)
	packageName = tableName

	if flagPackage != nil && *flagPackage != "" {
		packageName = *flagPackage
	}
	if flagPackageShort != nil && *flagPackageShort != "" {
		packageName = *flagPackageShort
	}

	outputDir = "."
	if flagOutput != nil && *flagOutput != "" {
		outputDir = *flagOutput
	}
	if flagOutputShort != nil && *flagOutputShort != "" {
		outputDir = *flagOutputShort
	}

	return
}
