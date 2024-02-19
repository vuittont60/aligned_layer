package aggregator

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/rpc"

	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

var (
	TaskNotFoundError400                     = errors.New("400. Task not found")
	OperatorNotPartOfTaskQuorum400           = errors.New("400. Operator not part of quorum")
	TaskResponseDigestNotFoundError500       = errors.New("500. Failed to get task response digest")
	UnknownErrorWhileVerifyingSignature400   = errors.New("400. Failed to verify signature")
	SignatureVerificationFailed400           = errors.New("400. Signature verification failed")
	CallToGetCheckSignaturesIndicesFailed500 = errors.New("500. Failed to get check signatures indices")
)

func (agg *Aggregator) startServer(ctx context.Context) error {

	err := rpc.Register(agg)
	if err != nil {
		agg.logger.Fatal("Format of service TaskManager isn't correct. ", "err", err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(agg.serverIpPortAddr, nil)
	if err != nil {
		agg.logger.Fatal("ListenAndServe", "err", err)
	}

	return nil
}

type SignedTaskResponse struct {
	TaskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
	BlsSignature bls.Signature
	OperatorId   bls.OperatorId
}

// rpc endpoint which is called by operator
// reply doesn't need to be checked. If there are no errors, the task response is accepted
// rpc framework forces a reply type to exist, so we put bool as a placeholder
func (agg *Aggregator) ProcessSignedTaskResponse(signedTaskResponse *SignedTaskResponse, reply *bool) error {
	agg.logger.Infof("Received signed task response: %#v", signedTaskResponse)
	taskIndex := signedTaskResponse.TaskResponse.ReferenceTaskIndex
	taskResponseDigest, err := core.GetTaskResponseDigest(&signedTaskResponse.TaskResponse)

	agg.tasksMu.Lock()
	agg.tasks[taskIndex]
	agg.tasksMu.Unlock()

	/*
		agg.tasksMu.Lock()
		agg.tasks[taskIndex] = newTask
		agg.tasksMu.Unlock()

		quorumThresholdPercentages := make([]uint32, len(newTask.QuorumNumbers))
		for i, _ := range newTask.QuorumNumbers {
			quorumThresholdPercentages[i] = newTask.QuorumThresholdPercentage
		}
		// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
		// and it should monitor the chain and only expire the task aggregation once the chain has reached that block number.
		taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds

		agg.blsAggregationService.InitializeNewTask(taskIndex, newTask.TaskCreatedBlock, newTask.QuorumNumbers, quorumThresholdPercentages, taskTimeToExpiry)
	*/

	if err != nil {
		agg.logger.Error("Failed to get task response digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}
	agg.taskResponsesMu.Lock()
	if _, ok := agg.taskResponses[taskIndex]; !ok {
		agg.taskResponses[taskIndex] = make(map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse)
	}
	if _, ok := agg.taskResponses[taskIndex][taskResponseDigest]; !ok {
		agg.taskResponses[taskIndex][taskResponseDigest] = signedTaskResponse.TaskResponse
	}
	agg.taskResponsesMu.Unlock()

	/*
		agg.tasksMu.Lock()
		agg.tasks[taskIndex] = newTask
		agg.tasksMu.Unlock()

		quorumThresholdPercentages := make([]uint32, len(newTask.QuorumNumbers))
		for i, _ := range newTask.QuorumNumbers {
			quorumThresholdPercentages[i] = newTask.QuorumThresholdPercentage
		}
		// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
		// and it should monitor the chain and only expire the task aggregation once the chain has reached that block number.
		taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds

		agg.blsAggregationService.InitializeNewTask(taskIndex, newTask.TaskCreatedBlock, newTask.QuorumNumbers, quorumThresholdPercentages, taskTimeToExpiry)
	*/

	err = agg.blsAggregationService.ProcessNewSignature(
		context.Background(), taskIndex, taskResponseDigest,
		&signedTaskResponse.BlsSignature, signedTaskResponse.OperatorId,
	)

	fmt.Println("Error status: ", err)

	return err
}
