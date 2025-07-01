package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
)

type CLI struct {
	Create CreateCmd `cmd:"" help:"create this"`
	Delete DeleteCmd `cmd:"" help:"delete this"`
}

type CreateCmd struct {
	Names []string `short:"n" long:"name" help:"names to create"`
}

type DeleteCmd struct {
	Names []string `short:"n" long:"name" help:"names to delete"`
}


func (c *CreateCmd) Run() error {
	if len(c.Names) == 0 {
		return fmt.Errorf("at least one name must be provided")
	}

	for _, name := range c.Names {
		fmt.Printf("created %s\n", name)
	}

	return nil
}


func (d *DeleteCmd) Run() error {
	if len(d.Names) == 0 {
		return fmt.Errorf("at least one name must be provided")
	}

	for _, name := range d.Names {
		fmt.Printf("deleted %s\n", name)
	}

	return nil
}

func main() {
	var cli CLI
	
	ctx := kong.Parse(&cli,
		kong.Name("shinytool"),
		kong.Description("shinytool description"),
	)
	
	err := ctx.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
