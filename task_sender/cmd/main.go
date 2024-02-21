package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/yetanotherco/aligned_layer/core/config"
	"github.com/yetanotherco/aligned_layer/task_generator"
)

var (
	// Version is the version of the binary.
	Version   string
	GitCommit string
	GitDate   string
)

var (
	ProofFileFlag = cli.StringFlag{
		Name:     "proof",
		Required: true,
		Usage:    "Load proof from `FILE`",
	}

	VerifierIdFlag = cli.StringFlag{
		Name:     "verifier-id",
		Required: true,
		Usage:    "Set verifier ID",
	}
)

var flags = []cli.Flag{
	config.ConfigFileFlag,
	config.AlignedLayerDeploymentFileFlag,
	config.SharedAvsContractsDeploymentFileFlag,
	config.EcdsaPrivateKeyFlag,
	ProofFileFlag,
	VerifierIdFlag,
}

func main() {
	app := cli.NewApp()
	app.Flags = config.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "Aligned Layer Task Sender"
	app.Usage = "Aligned Layer Task Sender"
	app.Description = "Service that sends proofs to verify by operator nodes."

	app.Action = taskSenderMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Task sender application failed.", "Message:", err)
	}
}

func taskSenderMain(ctx *cli.Context) error {
	log.Println("Initializing Task Sender")
	config, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}

	taskGen, err := task_generator.NewTaskGenerator(config)
	if err != nil {
		return err
	}

	err = taskGen.SendNewTask(proof, VerifierId)
	if err != nil {
		return err
	}

	return nil
}
