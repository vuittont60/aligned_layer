package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/Layr-Labs/eigensdk-go/signer"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
	"github.com/Layr-Labs/incredible-squaring-avs/common"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
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
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func taskGeneratorMain(ctx *cli.Context) error {

	log.Println("Initializing Task Generator")
	config, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}
	fmt.Println("Config:", string(configJson))

	// avsReader, err := chainio.NewAvsReaderFromConfig(c)
	// if err != nil {
	// 	panic("hola")
	// }

	chainId, err := config.EthHttpClient.ChainID(context.Background())
	if err != nil {
		panic("Hola")
	}

	privateKeySigner, err := signer.NewPrivateKeySigner(config.EcdsaPrivateKey, chainId)
	if err != nil {
		panic("Hola")
	}
	config.Signer = privateKeySigner

	avsWriter, err := chainio.NewAvsWriterFromConfig(config)
	if err != nil {
		panic("hola")
	}

	proof := make([]byte, 32)

	_, taskIndex, err := avsWriter.SendNewTaskVerifyProof(context.Background(), proof, common.GnarkPlonkBls12_381, types.QUORUM_THRESHOLD_NUMERATOR, types.QUORUM_NUMBERS)

	if err != nil {
		panic("Couldn't send the task")
	}

	fmt.Println("Created a new task with index: ", taskIndex)

	return nil
}
