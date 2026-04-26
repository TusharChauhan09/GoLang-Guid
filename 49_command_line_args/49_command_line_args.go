// 49_command_line_args.go
// Topic: Command-line args — os.Args and flag package
//
// RAW ARGS
//   os.Args[0] = program path
//   os.Args[1:] = remaining args
//
// flag PACKAGE — typed flags
//   var name = flag.String("name", "default", "usage")
//   var n    = flag.Int("n", 1, "count")
//   var ok   = flag.Bool("ok", false, "switch")
//   var d    = flag.Duration("d", time.Second, "duration")
//
//   flag.Parse()                  // call after defining all flags
//   flag.Args()                   // remaining positional
//   flag.NArg(), flag.NFlag()
//
//   Bind to existing var:
//   var port int
//   flag.IntVar(&port, "port", 8080, "port")
//
// FLAG SYNTAX (accepted)
//   -name=value
//   -name value      (NOT for bool)
//   --name=value
//   bools:  -ok   -ok=true   -ok=false
//
// CUSTOM FLAG SETS
//   fs := flag.NewFlagSet("sub", flag.ExitOnError)
//   fs.String(...); fs.Parse(args)
//
// THIRD-PARTY (subcommands etc): cobra, urfave/cli
//
// Run examples:
//   go run 49_command_line_args.go -name Ada -n 3 extra1 extra2
//   go run 49_command_line_args.go -h

package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// Raw
	fmt.Println("os.Args:", os.Args)

	name := flag.String("name", "world", "name to greet")
	count := flag.Int("n", 1, "number of greetings")
	loud := flag.Bool("loud", false, "shout")
	wait := flag.Duration("wait", 0, "delay before greeting")

	var port int
	flag.IntVar(&port, "port", 8080, "server port")

	flag.Parse()

	if *wait > 0 {
		time.Sleep(*wait)
	}

	greet := fmt.Sprintf("hello, %s", *name)
	if *loud {
		greet = fmt.Sprintf("HELLO, %s!", *name)
	}
	for i := 0; i < *count; i++ {
		fmt.Println(greet)
	}

	fmt.Println("port:", port)
	fmt.Println("positional args:", flag.Args())
}
