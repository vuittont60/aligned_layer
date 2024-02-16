package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/Layr-Labs/incredible-squaring-avs/aggregator"
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



	agg, err := aggregator.NewAggregator(config)
	if err != nil {
		return err
	}

	err = agg.Start(context.Background())
	if err != nil {
		return err
	}

	avsReader, err := chainio.NewAvsReaderFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create EthReader", "err", err)
		return nil, err
	}
	/*

	chainId, err := c.EthHttpClient.ChainID(context.Background())
	if err != nil {
		c.Logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	privateKeySigner, err := signer.NewPrivateKeySigner(c.EcdsaPrivateKey, chainId)
	if err != nil {
		c.Logger.Error("Cannot create signer", "err", err)
		return nil, err
	}
	c.Signer = privateKeySigner

	*/
	avsWriter, err := chainio.NewAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create EthSubscriber", "err", err)
		ret

	avsWriter.SendNewTaskVerifyProof(context.Background(), proof, verifierId, types.QUORUM_THRESHOLD_NUMERATOR, types.QUORUM_NUMBERS)

	return nil
}

