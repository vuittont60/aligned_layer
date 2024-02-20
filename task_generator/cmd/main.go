package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/Layr-Labs/incredible-squaring-avs/task_generator"
)

var (
	// Version is the version of the binary.
	Version   string
	GitCommit string
	GitDate   string
)

func main() {
	app := cli.NewApp()
	app.Flags = config.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "Aligned Layer Task Generator"
	app.Usage = "Aligned Layer Task Generator"
	app.Description = "Service that sends number to be credibly squared by operator nodes."

	app.Action = taskGeneratorMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Task generator application failed.", "Message:", err)
	}
}

func taskGeneratorMain(ctx *cli.Context) error {
	log.Println("Initializing Task Generator")
	config, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}

	taskGen, err := task_generator.NewTaskGenerator(config)
	if err != nil {
		return err
	}

	err = taskGen.Start(context.Background())
	if err != nil {
		return err
	}

	return nil
}
