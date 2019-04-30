package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fraenky8/tables-to-go/internal/cli"
	"github.com/fraenky8/tables-to-go/pkg/config"
)

// CmdArgs represents the supported command line args
type CmdArgs struct {
	Help bool
	*config.Settings
}

// NewCmdArgs creates and prepares the command line arguments with default values
func NewCmdArgs() (args *CmdArgs) {

	innerargs = &CmdArgs{
		Settings: config.NewSettings(),
	}

	flag.BoolVar(&innerargs.Help, "?", false, "shows help and usage")
	flag.BoolVar(&innerargs.Help, "help", false, "shows help and usage")
	flag.BoolVar(&innerargs.Verbose, "v", innerargs.Verbose, "verbose output")
	flag.BoolVar(&innerargs.VVerbose, "vv", innerargs.VVerbose, "more verbose output")

	flag.StringVar(&innerargs.DbType, "t", innerargs.DbType, fmt.Sprintf("type of database to use, currently supported: %v", innerargs.SupportedDbTypes()))
	flag.StringVar(&innerargs.User, "u", innerargs.User, "user to connect to the database")
	flag.StringVar(&innerargs.Pswd, "p", innerargs.Pswd, "password of user")
	flag.StringVar(&innerargs.DbName, "d", innerargs.DbName, "database name")
	flag.StringVar(&innerargs.Schema, "s", innerargs.Schema, "schema name")
	flag.StringVar(&innerargs.Host, "h", innerargs.Host, "host of database")
	flag.StringVar(&innerargs.Port, "port", innerargs.Port, "port of database host, if not specified, it will be the default ports for the supported databases")

	flag.StringVar(&innerargs.OutputFilePath, "of", innerargs.OutputFilePath, "output file path, default is current working directory")
	flag.StringVar(&innerargs.OutputFormat, "format", innerargs.OutputFormat, "format of struct fields (columns): camelCase (c) or original (o)")
	flag.StringVar(&innerargs.Prefix, "pre", innerargs.Prefix, "prefix for file- and struct names")
	flag.StringVar(&innerargs.Suffix, "suf", innerargs.Suffix, "suffix for file- and struct names")
	flag.StringVar(&innerargs.PackageName, "pn", innerargs.PackageName, "package name")
	flag.StringVar(&innerargs.Null, "null", innerargs.Null, "representation of NULL columns: sql.Null* (sql) or primitive pointers (native|primitive)")

	flag.BoolVar(&innerargs.NoInitialism, "no-initialism", innerargs.NoInitialism, "disable the conversion to upper-case words in column names")

	flag.BoolVar(&innerargs.TagsNoDb, "tags-no-db", innerargs.TagsNoDb, "do not create db-tags")

	flag.BoolVar(&innerargs.TagsMastermindStructable, "tags-structable", innerargs.TagsMastermindStructable, "generate struct with tags for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	flag.BoolVar(&innerargs.TagsMastermindStructableOnly, "tags-structable-only", innerargs.TagsMastermindStructableOnly, "generate struct with tags ONLY for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	flag.BoolVar(&innerargs.IsMastermindStructableRecorder, "structable-recorder", innerargs.IsMastermindStructableRecorder, "generate a structable.Recorder field")

	flag.BoolVar(&innerargs.TagsSQL, "experimental-tags-sql", innerargs.TagsSQL, "generate struct with sql-tags")
	flag.BoolVar(&innerargs.TagsSQLOnly, "experimental-tags-sql-only", innerargs.TagsSQLOnly, "generate struct with ONLY sql-tags")

	flag.Parse()

	return innerargs
}

// main function to run the transformations
func main() {

	cmdArgs := NewCmdArgs()

	if cmdArgs.Help {
		flag.Usage()
		os.Exit(0)
	}

	if err := cmdArgs.Verify(); err != nil {
		fmt.Printf("settings verification error: %v", err)
		os.Exit(1)
	}

	if err := cli.Run(cmdArgs.Settings); err != nil {
		fmt.Printf("run error: %v", err)
		os.Exit(1)
	}
}
