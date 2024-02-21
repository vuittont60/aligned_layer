package task_generator

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signer"
	"github.com/yetanotherco/aligned_layer/aggregator/types"
	"github.com/yetanotherco/aligned_layer/common"
	"github.com/yetanotherco/aligned_layer/core/chainio"
	"github.com/yetanotherco/aligned_layer/core/config"
)

type TaskGenerator struct {
	logger    logging.Logger
	avsWriter chainio.AvsWriterer
}

func NewTaskGenerator(c *config.Config) (*TaskGenerator, error) {
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

	avsWriter, err := chainio.NewAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create AVS writer", "err", err)
		return nil, err
	}

	return &TaskGenerator{
		logger:    c.Logger,
		avsWriter: avsWriter,
	}, nil
}

func (tg *TaskGenerator) Start(ctx context.Context) error {
	tg.logger.Infof("Starting task generator.")

	ticker := time.NewTicker(10 * time.Second)
	tg.logger.Infof("Task generator set to send new task every 10 seconds...")
	defer ticker.Stop()

	taskNum := int64(0)
	// ticker doesn't tick immediately, so we send the first task here
	// see https://github.com/golang/go/issues/17601

	// We are randomizing bytes for bad proofs, all should fail
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var proof []byte
	badProof := make([]byte, 32)
	r.Read(badProof)
	proof = badProof

	_ = tg.SendNewTask(proof, common.LambdaworksCairo)
	taskNum++

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			taskNum++

			// If taskNum is even, a verify Cairo task is sent
			// if taskNum is odd, a verify Gnark Plonk task is sent
			// This should be an additional configuration parameter

			if taskNum%2 == 0 {
				// Randomly creates tasks to verify correct and incorrect cairo proofs
				if r.Intn(3) != 0 {
					var err error
					proof, err = os.ReadFile("tests/testing_data/fibo_5.proof")
					if err != nil {
						panic("Could not read Cairo proof file")
					}
				} else {
					badProof := make([]byte, 32)
					r.Read(badProof)
					proof = badProof
				}
				err := tg.SendNewTask(proof, common.LambdaworksCairo)
				if err != nil {
					// we log the errors inside sendNewTask() so here we just continue to the next task
					continue
				}
			} else {
				if r.Intn(3) != 0 {
					var err error
					proof, err = os.ReadFile("tests/testing_data/plonk_cubic_circuit.proof")
					if err != nil {
						panic("Could not read PLONK proof file")
					}
				} else {
					badProof := make([]byte, 32)
					r.Read(badProof)
					proof = badProof
				}
				err := tg.SendNewTask(proof, common.GnarkPlonkBls12_381)
				if err != nil {
					// we log the errors inside sendNewTask() so here we just continue to the next task
					continue
				}
			}

		}
	}
}

// sendNewTask sends a new task to the task manager contract
func (tg *TaskGenerator) SendNewTask(proof []byte, verifierId common.VerifierId) error {
	_, taskIndex, err := tg.avsWriter.SendNewTaskVerifyProof(context.Background(), proof, verifierId, types.QUORUM_THRESHOLD_NUMERATOR, types.QUORUM_NUMBERS)
	if err != nil {
		tg.logger.Error("Task generator failed to send proof", "err", err)
		return err
	}

	tg.logger.Infof("Generated new task with index %d", taskIndex)

	return nil
}
