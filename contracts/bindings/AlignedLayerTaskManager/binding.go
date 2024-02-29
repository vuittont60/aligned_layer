// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAlignedLayerTaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BLSOperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type BLSOperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// BLSOperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type BLSOperatorStateRetrieverOperator struct {
	OperatorId [32]byte
	Stake      *big.Int
}

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IAlignedLayerTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IAlignedLayerTaskManagerTask struct {
	Proof                     []byte
	PubInput                  []byte
	VerifierId                uint16
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IAlignedLayerTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IAlignedLayerTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	ProofIsCorrect     bool
}

// IAlignedLayerTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IAlignedLayerTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// ContractAlignedLayerTaskManagerMetaData contains all meta data concerning the ContractAlignedLayerTaskManager contract.
var ContractAlignedLayerTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsPubkeyRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSPubkeyRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubInput\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBLSOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structBLSOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structBLSOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIAlignedLayerTaskManager.Task\",\"components\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubInput\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIAlignedLayerTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ProofIsCorrect\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIAlignedLayerTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIAlignedLayerTaskManager.Task\",\"components\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubInput\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIAlignedLayerTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ProofIsCorrect\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAlignedLayerTaskManager.Task\",\"components\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubInput\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAlignedLayerTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ProofIsCorrect\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAlignedLayerTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162004dfb38038062004dfb833981016040819052620000359162000169565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b59190620001b0565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316633561deb16040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001339190620001b0565b6001600160a01b031660c0525063ffffffff1660e05250620001d7565b6001600160a01b03811681146200016657600080fd5b50565b600080604083850312156200017d57600080fd5b82516200018a8162000150565b602084015190925063ffffffff81168114620001a557600080fd5b809150509250929050565b600060208284031215620001c357600080fd5b8151620001d08162000150565b9392505050565b60805160a05160c05160e051614bc06200023b60003960008181610251015281816105090152610e4a015260008181610306015261185a0152600081816103e601528181611e68015261200901526000818161040d0152611cbc0152614bc06000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c80636d14a9871161010f5780638f310c24116100a2578063f5c9899d11610071578063f5c9899d14610507578063f63c5bab1461052d578063f8c8765e14610535578063fabc1cbc1461054857600080fd5b80638f310c24146104ad57806399c4a6b1146104c0578063cefdc1d4146104d3578063f2fde38b146104f457600080fd5b80637afa1eed116100de5780637afa1eed14610466578063886f1195146104795780638b00ce7c1461048c5780638da5cb5b1461049c57600080fd5b80636d14a987146104085780636efb46361461042f578063715018a61461045057806372d18e8d1461045857600080fd5b80633563b0d1116101875780635ac86ab7116101565780635ac86ab7146103835780635c975abb146103b65780635decc3f5146103be57806368304835146103e157600080fd5b80633563b0d11461032857806346bb7bae146103485780634f739f741461035b578063595c6a671461037b57600080fd5b8063245a7bfc116101c3578063245a7bfc146102885780632cb223d5146102b35780632d89f6fc146102e15780633561deb11461030157600080fd5b806310d67a2f146101f5578063136439dd1461020a578063171f1d5b1461021d5780631ad431891461024c575b600080fd5b6102086102033660046136dc565b61055b565b005b610208610218366004613700565b610617565b61023061022b36600461387e565b610756565b6040805192151583529015156020830152015b60405180910390f35b6102737f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610243565b609b5461029b906001600160a01b031681565b6040516001600160a01b039091168152602001610243565b6102d36102c13660046138ec565b60996020526000908152604090205481565b604051908152602001610243565b6102d36102ef3660046138ec565b60986020526000908152604090205481565b61029b7f000000000000000000000000000000000000000000000000000000000000000081565b61033b610336366004613909565b6108e0565b6040516102439190613a50565b610208610356366004613d26565b610c59565b61036e610369366004613de2565b6110d8565b6040516102439190613ee6565b61020861175c565b6103a6610391366004613fa1565b606654600160ff9092169190911b9081161490565b6040519015158152602001610243565b6066546102d3565b6103a66103cc3660046138ec565b609a6020526000908152604090205460ff1681565b61029b7f000000000000000000000000000000000000000000000000000000000000000081565b61029b7f000000000000000000000000000000000000000000000000000000000000000081565b61044261043d366004613fc4565b611823565b604051610243929190614084565b61020861229c565b60975463ffffffff16610273565b609c5461029b906001600160a01b031681565b60655461029b906001600160a01b031681565b6097546102739063ffffffff1681565b6033546001600160a01b031661029b565b6102086104bb3660046140cd565b6122b0565b6102086104ce366004614165565b612311565b6104e66104e1366004614226565b612528565b604051610243929190614268565b6102086105023660046136dc565b6126ba565b7f0000000000000000000000000000000000000000000000000000000000000000610273565b610273606481565b610208610543366004614289565b612730565b610208610556366004613700565b612881565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105d291906142e5565b6001600160a01b0316336001600160a01b03161461060b5760405162461bcd60e51b815260040161060290614302565b60405180910390fd5b610614816129dd565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561065f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610683919061435a565b61069f5760405162461bcd60e51b815260040161060290614377565b606654818116146107185760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610602565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061079e5761079e6143bf565b60200201518951600160200201518a602001516000600281106107c3576107c36143bf565b60200201518b602001516001600281106107df576107df6143bf565b602090810291909101518c518d83015160405161083c9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c61085f91906143d5565b90506108d26108786108718884612ad4565b8690612b6b565b610880612bff565b6108c86108b9856108b3604080518082018252600080825260209182015281518083019092526001825260029082015290565b90612ad4565b6108c28c612cbf565b90612b6b565b886201d4c0612d4f565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610922573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061094691906142e5565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610988573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ac91906142e5565b9050600085516001600160401b038111156109c9576109c9613719565b6040519080825280602002602001820160405280156109fc57816020015b60608152602001906001900390816109e75790505b50905060005b8651811015610c4e576000878281518110610a1f57610a1f6143bf565b016020015160405163889ae3e560e01b815260f89190911c6004820181905263ffffffff8916602483015291506000906001600160a01b0386169063889ae3e590604401600060405180830381865afa158015610a80573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610aa891908101906143f7565b905080516001600160401b03811115610ac357610ac3613719565b604051908082528060200260200182016040528015610b0857816020015b6040805180820190915260008082526020820152815260200190600190039081610ae15790505b50848481518110610b1b57610b1b6143bf565b602002602001018190525060005b8151811015610c38576000828281518110610b4657610b466143bf565b6020908102919091018101516040805180820182528281529051631b32722560e01b81526004810183905260ff8816602482015263ffffffff8e166044820152919350918201906001600160a01b038b1690631b32722590606401602060405180830381865afa158015610bbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610be29190614487565b6001600160601b0316815250868681518110610c0057610c006143bf565b60200260200101518381518110610c1957610c196143bf565b6020026020010181905250508080610c30906144c6565b915050610b29565b5050508080610c46906144c6565b915050610a02565b509695505050505050565b609b546001600160a01b03163314610cb35760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610602565b6000610cc560808501606086016138ec565b9050366000610cd760808701876144e1565b90925090506000610cee60c0880160a089016138ec565b905060986000610d0160208901896138ec565b63ffffffff1663ffffffff1681526020019081526020016000205487604051602001610d2d9190614595565b6040516020818303038152906040528051906020012014610db65760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610602565b6000609981610dc860208a018a6138ec565b63ffffffff1663ffffffff1681526020019081526020016000205414610e455760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610602565b610e6f7f000000000000000000000000000000000000000000000000000000000000000085614662565b63ffffffff164363ffffffff161115610ee05760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610602565b600086604051602001610ef391906146b8565b604051602081830303815290604052805190602001209050600080610f1b8387878a8c611823565b9150915060005b8581101561101a578460ff1683602001518281518110610f4457610f446143bf565b6020026020010151610f5691906146c6565b6001600160601b0316606484600001518381518110610f7757610f776143bf565b60200260200101516001600160601b0316610f9291906146f5565b1015611008576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610602565b80611012816144c6565b915050610f22565b5060408051808201825263ffffffff43168152602080820184905291519091611047918c91849101614714565b60405160208183030381529060405280519060200120609960008c600001602081019061107491906138ec565b63ffffffff1663ffffffff168152602001908152602001600020819055507f46c382fc6ec03c126042e96e5be20ec429b355ffe4b56d4614d3a6b627da3a7e8a826040516110c3929190614714565b60405180910390a15050505050505050505050565b6111036040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611143573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061116791906142e5565b90506111946040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516385020d4960e01b81526001600160a01b038a16906385020d49906111c4908b9089908990600401614740565b600060405180830381865afa1580156111e1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611209919081019061478a565b815260405163e192e9ad60e01b81526001600160a01b0383169063e192e9ad9061123b908b908b908b90600401614818565b600060405180830381865afa158015611258573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611280919081019061478a565b6040820152856001600160401b0381111561129d5761129d613719565b6040519080825280602002602001820160405280156112d057816020015b60608152602001906001900390816112bb5790505b50606082015260005b60ff811687111561166d576000856001600160401b038111156112fe576112fe613719565b604051908082528060200260200182016040528015611327578160200160208202803683370190505b5083606001518360ff1681518110611341576113416143bf565b602002602001018190525060005b8681101561156d5760008c6001600160a01b0316633064620d8a8a8581811061137a5761137a6143bf565b905060200201358e88600001518681518110611398576113986143bf565b60200260200101516040518463ffffffff1660e01b81526004016113d59392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156113f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114169190614841565b90508a8a8560ff1681811061142d5761142d6143bf565b6001600160c01b03841692013560f81c9190911c60019081161415905061155a57856001600160a01b031663480858668a8a8581811061146f5761146f6143bf565b905060200201358d8d8860ff1681811061148b5761148b6143bf565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa1580156114e1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611505919061486a565b85606001518560ff168151811061151e5761151e6143bf565b60200260200101518481518110611537576115376143bf565b63ffffffff9092166020928302919091019091015282611556816144c6565b9350505b5080611565816144c6565b91505061134f565b506000816001600160401b0381111561158857611588613719565b6040519080825280602002602001820160405280156115b1578160200160208202803683370190505b50905060005b828110156116325784606001518460ff16815181106115d8576115d86143bf565b602002602001015181815181106115f1576115f16143bf565b602002602001015182828151811061160b5761160b6143bf565b63ffffffff909216602092830291909101909101528061162a816144c6565b9150506115b7565b508084606001518460ff168151811061164d5761164d6143bf565b60200260200101819052505050808061166590614887565b9150506112d9565b506000896001600160a01b0316633561deb16040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116d291906142e5565b60405163eda1076360e01b81529091506001600160a01b0382169063eda1076390611705908b908b908e906004016148a7565b600060405180830381865afa158015611722573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261174a919081019061478a565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156117a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117c8919061435a565b6117e45760405162461bcd60e51b815260040161060290614377565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60408051808201909152606080825260208201526040805180820190915260008082526020820181905290815b86811015611a4d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c1af6b24898984818110611899576118996143bf565b9050013560f81c60f81b60f81c888860a0015185815181106118bd576118bd6143bf565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015611919573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061193d91906148d1565b6001600160401b03191661196d86604001518381518110611960576119606143bf565b6020026020010151612f73565b67ffffffffffffffff191614611a095760405162461bcd60e51b81526020600482015260616024820152600080516020614b6b83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610602565b611a3985604001518281518110611a2257611a226143bf565b602002602001015183612b6b90919063ffffffff16565b915080611a45816144c6565b915050611850565b506040805180820190915260608082526020820152866001600160401b03811115611a7a57611a7a613719565b604051908082528060200260200182016040528015611aa3578160200160208202803683370190505b506020820152866001600160401b03811115611ac157611ac1613719565b604051908082528060200260200182016040528015611aea578160200160208202803683370190505b5081526020850151516000906001600160401b03811115611b0d57611b0d613719565b604051908082528060200260200182016040528015611b36578160200160208202803683370190505b50905060008660200151516001600160401b03811115611b5857611b58613719565b604051908082528060200260200182016040528015611b81578160200160208202803683370190505b5090506000611bc58b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612fb692505050565b905060005b886020015151811015611e3057611bf089602001518281518110611960576119606143bf565b848281518110611c0257611c026143bf565b60209081029190910101528015611cba5783611c1f6001836148fc565b81518110611c2f57611c2f6143bf565b602002602001015160001c848281518110611c4c57611c4c6143bf565b602002602001015160001c11611cba576040805162461bcd60e51b8152602060048201526024810191909152600080516020614b6b83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610602565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633064620d858381518110611cfb57611cfb6143bf565b60200260200101518c8c600001518581518110611d1a57611d1a6143bf565b60200260200101516040518463ffffffff1660e01b8152600401611d579392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611d74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d989190614841565b6001600160c01b0316838281518110611db357611db36143bf565b602002602001018181525050611e1c611e15611de984868581518110611ddb57611ddb6143bf565b60200260200101511661311f565b611e0f8c602001518581518110611e0257611e026143bf565b6020026020010151613150565b906131eb565b8790612b6b565b955080611e28816144c6565b915050611bca565b505060005b60ff81168a11156121705760008b8b8360ff16818110611e5757611e576143bf565b9050013560f81c60f81b60f81c90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c56828c8c60c001518660ff1681518110611eb057611eb06143bf565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015611f0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f309190614487565b85602001518360ff1681518110611f4957611f496143bf565b6001600160601b03909216602092830291909101820152850151805160ff8416908110611f7857611f786143bf565b602002602001015185600001518360ff1681518110611f9957611f996143bf565b60200260200101906001600160601b031690816001600160601b03168152505060005b8960200151518163ffffffff161015612166576000612002858363ffffffff1681518110611fec57611fec6143bf565b60200260200101518460ff161c60019081161490565b15612153577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a43cde89848e898663ffffffff1681518110612050576120506143bf565b60200260200101518f60e001518960ff1681518110612071576120716143bf565b60200260200101518663ffffffff1681518110612090576120906143bf565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa1580156120f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121189190614487565b8751805160ff871690811061212f5761212f6143bf565b602002602001018181516121439190614913565b6001600160601b03169052506001015b508061215e8161493b565b915050611fbc565b5050600101611e35565b50506000806121898c868a606001518b60800151610756565b91509150816121fa5760405162461bcd60e51b81526020600482015260436024820152600080516020614b6b83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610602565b8061225b5760405162461bcd60e51b81526020600482015260396024820152600080516020614b6b83398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610602565b50506000878260405160200161227292919061495f565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6122a46132d0565b6122ae600061332a565b565b60006122bf60208501856138ec565b63ffffffff81166000818152609a6020526040808220805460ff191660011790555192935033927ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb059190a35050505050565b609c546001600160a01b031633146123755760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610602565b6040805160c0810182526060808252602082018190526000928201839052808201839052608082015260a081019190915288888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f890181900481028201810190925287815290889088908190840183828082843760009201919091525050505060208083019190915261ffff861660408084019190915263ffffffff4381166060850152861660a08401528051601f85018390048302810183019091528381529084908490819084018382808284376000920191909152505050506080820152604051612484908290602001614a03565b60408051601f1981840301815282825280516020918201206097805463ffffffff9081166000908152609890945293909220555416907fc76bc85127c84882e4e28847aeef6dac7f160a43a5e7f22a550d32d06e4efce3906124e7908490614a03565b60405180910390a26097546125039063ffffffff166001614662565b6097805463ffffffff191663ffffffff92909216919091179055505050505050505050565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612563576125636143bf565b60209081029190910101526040516385020d4960e01b81526000906001600160a01b038816906385020d499061259f9088908690600401614a97565b600060405180830381865afa1580156125bc573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526125e4919081019061478a565b6000815181106125f6576125f66143bf565b6020908102919091010151604051633064620d60e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b03891690633064620d90606401602060405180830381865afa158015612662573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126869190614841565b6001600160c01b03169050600061269c8261337c565b9050816126aa8a838a6108e0565b9550955050505050935093915050565b6126c26132d0565b6001600160a01b0381166127275760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610602565b6106148161332a565b600054610100900460ff16158080156127505750600054600160ff909116105b8061276a5750303b15801561276a575060005460ff166001145b6127cd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610602565b6000805460ff1916600117905580156127f0576000805461ff0019166101001790555b6127fb8560006133d9565b6128048461332a565b609b80546001600160a01b038086166001600160a01b031992831617909255609c805492851692909116919091179055801561287a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128f891906142e5565b6001600160a01b0316336001600160a01b0316146129285760405162461bcd60e51b815260040161060290614302565b6066541981196066541916146129a65760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610602565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200161074b565b6001600160a01b038116612a6b5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610602565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152612af06135ed565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015612b2357612b25565bfe5b5080612b635760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610602565b505092915050565b6040805180820190915260008082526020820152612b8761360b565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015612b23575080612b635760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610602565b612c07613629565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080612cef600080516020614b4b833981519152866143d5565b90505b612cfb816134c3565b9093509150600080516020614b4b833981519152828309831415612d35576040805180820190915290815260208101919091529392505050565b600080516020614b4b833981519152600182089050612cf2565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190612d8161364e565b60005b6002811015612f46576000612d9a8260066146f5565b9050848260028110612dae57612dae6143bf565b60200201515183612dc0836000614aeb565b600c8110612dd057612dd06143bf565b6020020152848260028110612de757612de76143bf565b60200201516020015183826001612dfe9190614aeb565b600c8110612e0e57612e0e6143bf565b6020020152838260028110612e2557612e256143bf565b6020020151515183612e38836002614aeb565b600c8110612e4857612e486143bf565b6020020152838260028110612e5f57612e5f6143bf565b6020020151516001602002015183612e78836003614aeb565b600c8110612e8857612e886143bf565b6020020152838260028110612e9f57612e9f6143bf565b602002015160200151600060028110612eba57612eba6143bf565b602002015183612ecb836004614aeb565b600c8110612edb57612edb6143bf565b6020020152838260028110612ef257612ef26143bf565b602002015160200151600160028110612f0d57612f0d6143bf565b602002015183612f1e836005614aeb565b600c8110612f2e57612f2e6143bf565b60200201525080612f3e816144c6565b915050612d84565b50612f4f61366d565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600081600001518260200151604051602001612f99929190918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b60006101008251111561302a5760405162461bcd60e51b815260206004820152603660248201527f4269746d61705574696c732e62797465734172726179546f4269746d61703a206044820152756279746573417272617920697320746f6f206c6f6e6760501b6064820152608401610602565b815161303857506000919050565b6000808360008151811061304e5761304e6143bf565b0160200151600160f89190911c81901b92505b84518110156131165784818151811061307c5761307c6143bf565b0160200151600160f89190911c1b9150828216156131025760405162461bcd60e51b815260206004820152603a60248201527f4269746d61705574696c732e62797465734172726179546f4269746d61703a2060448201527f72657065617420656e74727920696e20627974657341727261790000000000006064820152608401610602565b9181179161310f816144c6565b9050613061565b50909392505050565b6000805b821561314a576131346001846148fc565b909216918061314281614b03565b915050613123565b92915050565b6040805180820190915260008082526020820152815115801561317557506020820151155b15613193575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020614b4b83398151915284602001516131c691906143d5565b6131de90600080516020614b4b8339815191526148fc565b905292915050565b919050565b60408051808201909152600080825260208201526102008261ffff16106132475760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610602565b8161ffff166001141561325b57508161314a565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1611156132c557600161ffff871660ff83161c811614156132a8576132a58484612b6b565b93505b6132b28384612b6b565b92506201fffe600192831b169101613277565b509195945050505050565b6033546001600160a01b031633146122ae5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610602565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606000805b6101008110156133d2576001811b9150838216156133c257828160f81b6040516020016133b0929190614b1b565b60405160208183030381529060405292505b6133cb816144c6565b9050613382565b5050919050565b6065546001600160a01b03161580156133fa57506001600160a01b03821615155b61347c5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610602565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26134bf826129dd565b5050565b60008080600080516020614b4b8339815191526003600080516020614b4b83398151915286600080516020614b4b833981519152888909090890506000613539827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020614b4b833981519152613545565b91959194509092505050565b60008061355061366d565b61355861368b565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015612b235750826135e25760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610602565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061363c6136a9565b81526020016136496136a9565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461061457600080fd5b6000602082840312156136ee57600080fd5b81356136f9816136c7565b9392505050565b60006020828403121561371257600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561375157613751613719565b60405290565b60405161010081016001600160401b038111828210171561375157613751613719565b604051601f8201601f191681016001600160401b03811182821017156137a2576137a2613719565b604052919050565b6000604082840312156137bc57600080fd5b6137c461372f565b9050813581526020820135602082015292915050565b600082601f8301126137eb57600080fd5b604051604081018181106001600160401b038211171561380d5761380d613719565b806040525080604084018581111561382457600080fd5b845b818110156132c5578035835260209283019201613826565b60006080828403121561385057600080fd5b61385861372f565b905061386483836137da565b815261387383604084016137da565b602082015292915050565b600080600080610120858703121561389557600080fd5b843593506138a686602087016137aa565b92506138b5866060870161383e565b91506138c48660e087016137aa565b905092959194509250565b63ffffffff8116811461061457600080fd5b80356131e6816138cf565b6000602082840312156138fe57600080fd5b81356136f9816138cf565b60008060006060848603121561391e57600080fd5b8335613929816136c7565b92506020848101356001600160401b038082111561394657600080fd5b818701915087601f83011261395a57600080fd5b81358181111561396c5761396c613719565b61397e601f8201601f1916850161377a565b9150808252888482850101111561399457600080fd5b80848401858401376000848284010152508094505050506139b7604085016138e1565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015613a42578385038a52825180518087529087019087870190845b81811015613a2d578351805184528a01516001600160601b03168a840152928901926040909201916001016139fd565b50509a87019a955050918501916001016139df565b509298975050505050505050565b6020815260006136f960208301846139c0565b600060c08284031215613a7557600080fd5b50919050565b600060408284031215613a7557600080fd5b60006001600160401b03821115613aa657613aa6613719565b5060051b60200190565b600082601f830112613ac157600080fd5b81356020613ad6613ad183613a8d565b61377a565b82815260059290921b84018101918181019086841115613af557600080fd5b8286015b84811015610c4e578035613b0c816138cf565b8352918301918301613af9565b600082601f830112613b2a57600080fd5b81356020613b3a613ad183613a8d565b82815260069290921b84018101918181019086841115613b5957600080fd5b8286015b84811015610c4e57613b6f88826137aa565b835291830191604001613b5d565b600082601f830112613b8e57600080fd5b81356020613b9e613ad183613a8d565b82815260059290921b84018101918181019086841115613bbd57600080fd5b8286015b84811015610c4e5780356001600160401b03811115613be05760008081fd5b613bee8986838b0101613ab0565b845250918301918301613bc1565b60006101808284031215613c0f57600080fd5b613c17613757565b905081356001600160401b0380821115613c3057600080fd5b613c3c85838601613ab0565b83526020840135915080821115613c5257600080fd5b613c5e85838601613b19565b60208401526040840135915080821115613c7757600080fd5b613c8385838601613b19565b6040840152613c95856060860161383e565b6060840152613ca78560e086016137aa565b6080840152610120840135915080821115613cc157600080fd5b613ccd85838601613ab0565b60a0840152610140840135915080821115613ce757600080fd5b613cf385838601613ab0565b60c0840152610160840135915080821115613d0d57600080fd5b50613d1a84828501613b7d565b60e08301525092915050565b600080600060808486031215613d3b57600080fd5b83356001600160401b0380821115613d5257600080fd5b613d5e87838801613a63565b9450613d6d8760208801613a7b565b93506060860135915080821115613d8357600080fd5b50613d9086828701613bfc565b9150509250925092565b60008083601f840112613dac57600080fd5b5081356001600160401b03811115613dc357600080fd5b602083019150836020828501011115613ddb57600080fd5b9250929050565b60008060008060008060808789031215613dfb57600080fd5b8635613e06816136c7565b95506020870135613e16816138cf565b945060408701356001600160401b0380821115613e3257600080fd5b613e3e8a838b01613d9a565b90965094506060890135915080821115613e5757600080fd5b818901915089601f830112613e6b57600080fd5b813581811115613e7a57600080fd5b8a60208260051b8501011115613e8f57600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015613edb57815163ffffffff1687529582019590820190600101613eb9565b509495945050505050565b600060208083528351608082850152613f0260a0850182613ea5565b905081850151601f1980868403016040870152613f1f8383613ea5565b92506040870151915080868403016060870152613f3c8383613ea5565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015613f935784878303018452613f81828751613ea5565b95880195938801939150600101613f67565b509998505050505050505050565b600060208284031215613fb357600080fd5b813560ff811681146136f957600080fd5b600080600080600060808688031215613fdc57600080fd5b8535945060208601356001600160401b0380821115613ffa57600080fd5b61400689838a01613d9a565b90965094506040880135915061401b826138cf565b9092506060870135908082111561403157600080fd5b5061403e88828901613bfc565b9150509295509295909350565b600081518084526020808501945080840160005b83811015613edb5781516001600160601b03168752958201959082019060010161405f565b604081526000835160408084015261409f608084018261404b565b90506020850151603f198483030160608501526140bc828261404b565b925050508260208301529392505050565b60008060008060c085870312156140e357600080fd5b84356001600160401b03808211156140fa57600080fd5b61410688838901613a63565b95506141158860208901613a7b565b94506141248860608901613a7b565b935060a087013591508082111561413a57600080fd5b5061414787828801613b19565b91505092959194509250565b803561ffff811681146131e657600080fd5b60008060008060008060008060a0898b03121561418157600080fd5b88356001600160401b038082111561419857600080fd5b6141a48c838d01613d9a565b909a50985060208b01359150808211156141bd57600080fd5b6141c98c838d01613d9a565b90985096508691506141dd60408c01614153565b955060608b013591506141ef826138cf565b90935060808a0135908082111561420557600080fd5b506142128b828c01613d9a565b999c989b5096995094979396929594505050565b60008060006060848603121561423b57600080fd5b8335614246816136c7565b925060208401359150604084013561425d816138cf565b809150509250925092565b82815260406020820152600061428160408301846139c0565b949350505050565b6000806000806080858703121561429f57600080fd5b84356142aa816136c7565b935060208501356142ba816136c7565b925060408501356142ca816136c7565b915060608501356142da816136c7565b939692955090935050565b6000602082840312156142f757600080fd5b81516136f9816136c7565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b801515811461061457600080fd5b60006020828403121561436c57600080fd5b81516136f98161434c565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000826143f257634e487b7160e01b600052601260045260246000fd5b500690565b6000602080838503121561440a57600080fd5b82516001600160401b0381111561442057600080fd5b8301601f8101851361443157600080fd5b805161443f613ad182613a8d565b81815260059190911b8201830190838101908783111561445e57600080fd5b928401925b8284101561447c57835182529284019290840190614463565b979650505050505050565b60006020828403121561449957600080fd5b81516001600160601b03811681146136f957600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156144da576144da6144b0565b5060010190565b6000808335601e198436030181126144f857600080fd5b8301803591506001600160401b0382111561451257600080fd5b602001915036819003821315613ddb57600080fd5b6000808335601e1984360301811261453e57600080fd5b83016020810192503590506001600160401b0381111561455d57600080fd5b803603831315613ddb57600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6020815260006145a58384614527565b60c060208501526145ba60e08501828461456c565b9150506145ca6020850185614527565b601f19808685030160408701526145e284838561456c565b935061ffff6145f360408901614153565b16606087015260608701359250614609836138cf565b63ffffffff831660808701526146226080880188614527565b93509150808685030160a08701525061463c83838361456c565b9250505060a084013561464e816138cf565b63ffffffff811660c0850152509392505050565b600063ffffffff808316818516808303821115614681576146816144b0565b01949350505050565b8035614695816138cf565b63ffffffff16825260208101356146ab8161434c565b8015156020840152505050565b6040810161314a828461468a565b60006001600160601b03808316818516818304811182151516156146ec576146ec6144b0565b02949350505050565b600081600019048311821515161561470f5761470f6144b0565b500290565b60808101614722828561468a565b63ffffffff8351166040830152602083015160608301529392505050565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561476d57600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561479d57600080fd5b82516001600160401b038111156147b357600080fd5b8301601f810185136147c457600080fd5b80516147d2613ad182613a8d565b81815260059190911b820183019083810190878311156147f157600080fd5b928401925b8284101561447c578351614809816138cf565b825292840192908401906147f6565b63ffffffff8416815260406020820152600061483860408301848661456c565b95945050505050565b60006020828403121561485357600080fd5b81516001600160c01b03811681146136f957600080fd5b60006020828403121561487c57600080fd5b81516136f9816138cf565b600060ff821660ff81141561489e5761489e6144b0565b60010192915050565b6040815260006148bb60408301858761456c565b905063ffffffff83166020830152949350505050565b6000602082840312156148e357600080fd5b815167ffffffffffffffff19811681146136f957600080fd5b60008282101561490e5761490e6144b0565b500390565b60006001600160601b0383811690831681811015614933576149336144b0565b039392505050565b600063ffffffff80831681811415614955576149556144b0565b6001019392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561499a5781518552938201939082019060010161497e565b5092979650505050505050565b60005b838110156149c25781810151838201526020016149aa565b838111156149d1576000848401525b50505050565b600081518084526149ef8160208601602086016149a7565b601f01601f19169290920160200192915050565b602081526000825160c06020840152614a1f60e08401826149d7565b90506020840151601f1980858403016040860152614a3d83836149d7565b925061ffff60408701511660608601526060860151915063ffffffff808316608087015260808701519250818685030160a0870152614a7c84846149d7565b93508060a08801511660c08701525050508091505092915050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b81811015614ade57845183529383019391830191600101614ac2565b5090979650505050505050565b60008219821115614afe57614afe6144b0565b500190565b600061ffff80831681811415614955576149556144b0565b60008351614b2d8184602088016149a7565b6001600160f81b031993909316919092019081526001019291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220197c5bfbc1f18eaa434b1c484c1365cc3a5b401cacb6ce7fef7bbdeaf610859464736f6c634300080c0033",
}

// ContractAlignedLayerTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAlignedLayerTaskManagerMetaData.ABI instead.
var ContractAlignedLayerTaskManagerABI = ContractAlignedLayerTaskManagerMetaData.ABI

// ContractAlignedLayerTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAlignedLayerTaskManagerMetaData.Bin instead.
var ContractAlignedLayerTaskManagerBin = ContractAlignedLayerTaskManagerMetaData.Bin

// DeployContractAlignedLayerTaskManager deploys a new Ethereum contract, binding an instance of ContractAlignedLayerTaskManager to it.
func DeployContractAlignedLayerTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractAlignedLayerTaskManager, error) {
	parsed, err := ContractAlignedLayerTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAlignedLayerTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAlignedLayerTaskManager{ContractAlignedLayerTaskManagerCaller: ContractAlignedLayerTaskManagerCaller{contract: contract}, ContractAlignedLayerTaskManagerTransactor: ContractAlignedLayerTaskManagerTransactor{contract: contract}, ContractAlignedLayerTaskManagerFilterer: ContractAlignedLayerTaskManagerFilterer{contract: contract}}, nil
}

// ContractAlignedLayerTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractAlignedLayerTaskManager struct {
	ContractAlignedLayerTaskManagerCaller     // Read-only binding to the contract
	ContractAlignedLayerTaskManagerTransactor // Write-only binding to the contract
	ContractAlignedLayerTaskManagerFilterer   // Log filterer for contract events
}

// ContractAlignedLayerTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAlignedLayerTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAlignedLayerTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAlignedLayerTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAlignedLayerTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAlignedLayerTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAlignedLayerTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAlignedLayerTaskManagerSession struct {
	Contract     *ContractAlignedLayerTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractAlignedLayerTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAlignedLayerTaskManagerCallerSession struct {
	Contract *ContractAlignedLayerTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// ContractAlignedLayerTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAlignedLayerTaskManagerTransactorSession struct {
	Contract     *ContractAlignedLayerTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// ContractAlignedLayerTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAlignedLayerTaskManagerRaw struct {
	Contract *ContractAlignedLayerTaskManager // Generic contract binding to access the raw methods on
}

// ContractAlignedLayerTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAlignedLayerTaskManagerCallerRaw struct {
	Contract *ContractAlignedLayerTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAlignedLayerTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAlignedLayerTaskManagerTransactorRaw struct {
	Contract *ContractAlignedLayerTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAlignedLayerTaskManager creates a new instance of ContractAlignedLayerTaskManager, bound to a specific deployed contract.
func NewContractAlignedLayerTaskManager(address common.Address, backend bind.ContractBackend) (*ContractAlignedLayerTaskManager, error) {
	contract, err := bindContractAlignedLayerTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManager{ContractAlignedLayerTaskManagerCaller: ContractAlignedLayerTaskManagerCaller{contract: contract}, ContractAlignedLayerTaskManagerTransactor: ContractAlignedLayerTaskManagerTransactor{contract: contract}, ContractAlignedLayerTaskManagerFilterer: ContractAlignedLayerTaskManagerFilterer{contract: contract}}, nil
}

// NewContractAlignedLayerTaskManagerCaller creates a new read-only instance of ContractAlignedLayerTaskManager, bound to a specific deployed contract.
func NewContractAlignedLayerTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractAlignedLayerTaskManagerCaller, error) {
	contract, err := bindContractAlignedLayerTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManagerCaller{contract: contract}, nil
}

// NewContractAlignedLayerTaskManagerTransactor creates a new write-only instance of ContractAlignedLayerTaskManager, bound to a specific deployed contract.
func NewContractAlignedLayerTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAlignedLayerTaskManagerTransactor, error) {
	contract, err := bindContractAlignedLayerTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManagerTransactor{contract: contract}, nil
}

// NewContractAlignedLayerTaskManagerFilterer creates a new log filterer instance of ContractAlignedLayerTaskManager, bound to a specific deployed contract.
func NewContractAlignedLayerTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAlignedLayerTaskManagerFilterer, error) {
	contract, err := bindContractAlignedLayerTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManagerFilterer{contract: contract}, nil
}

// bindContractAlignedLayerTaskManager binds a generic wrapper to an already deployed contract.
func bindContractAlignedLayerTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAlignedLayerTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAlignedLayerTaskManager.Contract.ContractAlignedLayerTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.ContractAlignedLayerTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.ContractAlignedLayerTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAlignedLayerTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractAlignedLayerTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractAlignedLayerTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractAlignedLayerTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractAlignedLayerTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractAlignedLayerTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractAlignedLayerTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractAlignedLayerTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractAlignedLayerTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.Aggregator(&_ContractAlignedLayerTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.Aggregator(&_ContractAlignedLayerTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractAlignedLayerTaskManager.Contract.AllTaskHashes(&_ContractAlignedLayerTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractAlignedLayerTaskManager.Contract.AllTaskHashes(&_ContractAlignedLayerTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractAlignedLayerTaskManager.Contract.AllTaskResponses(&_ContractAlignedLayerTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractAlignedLayerTaskManager.Contract.AllTaskResponses(&_ContractAlignedLayerTaskManager.CallOpts, arg0)
}

// BlsPubkeyRegistry is a free data retrieval call binding the contract method 0x3561deb1.
//
// Solidity: function blsPubkeyRegistry() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) BlsPubkeyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "blsPubkeyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsPubkeyRegistry is a free data retrieval call binding the contract method 0x3561deb1.
//
// Solidity: function blsPubkeyRegistry() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) BlsPubkeyRegistry() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.BlsPubkeyRegistry(&_ContractAlignedLayerTaskManager.CallOpts)
}

// BlsPubkeyRegistry is a free data retrieval call binding the contract method 0x3561deb1.
//
// Solidity: function blsPubkeyRegistry() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) BlsPubkeyRegistry() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.BlsPubkeyRegistry(&_ContractAlignedLayerTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns((uint96[],uint96[]), bytes32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns((uint96[],uint96[]), bytes32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractAlignedLayerTaskManager.Contract.CheckSignatures(&_ContractAlignedLayerTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns((uint96[],uint96[]), bytes32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractAlignedLayerTaskManager.Contract.CheckSignatures(&_ContractAlignedLayerTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) Generator() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.Generator(&_ContractAlignedLayerTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.Generator(&_ContractAlignedLayerTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(BLSOperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(BLSOperatorStateRetrieverCheckSignaturesIndices)).(*BLSOperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractAlignedLayerTaskManager.Contract.GetCheckSignaturesIndices(&_ContractAlignedLayerTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractAlignedLayerTaskManager.Contract.GetCheckSignaturesIndices(&_ContractAlignedLayerTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]BLSOperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]BLSOperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]BLSOperatorStateRetrieverOperator)).(*[][]BLSOperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]BLSOperatorStateRetrieverOperator, error) {
	return _ContractAlignedLayerTaskManager.Contract.GetOperatorState(&_ContractAlignedLayerTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]BLSOperatorStateRetrieverOperator, error) {
	return _ContractAlignedLayerTaskManager.Contract.GetOperatorState(&_ContractAlignedLayerTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]BLSOperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]BLSOperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]BLSOperatorStateRetrieverOperator)).(*[][]BLSOperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]BLSOperatorStateRetrieverOperator, error) {
	return _ContractAlignedLayerTaskManager.Contract.GetOperatorState0(&_ContractAlignedLayerTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]BLSOperatorStateRetrieverOperator, error) {
	return _ContractAlignedLayerTaskManager.Contract.GetOperatorState0(&_ContractAlignedLayerTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractAlignedLayerTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractAlignedLayerTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractAlignedLayerTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractAlignedLayerTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractAlignedLayerTaskManager.Contract.LatestTaskNum(&_ContractAlignedLayerTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractAlignedLayerTaskManager.Contract.LatestTaskNum(&_ContractAlignedLayerTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) Owner() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.Owner(&_ContractAlignedLayerTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.Owner(&_ContractAlignedLayerTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractAlignedLayerTaskManager.Contract.Paused(&_ContractAlignedLayerTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractAlignedLayerTaskManager.Contract.Paused(&_ContractAlignedLayerTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractAlignedLayerTaskManager.Contract.Paused0(&_ContractAlignedLayerTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractAlignedLayerTaskManager.Contract.Paused0(&_ContractAlignedLayerTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.PauserRegistry(&_ContractAlignedLayerTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.PauserRegistry(&_ContractAlignedLayerTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.RegistryCoordinator(&_ContractAlignedLayerTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.RegistryCoordinator(&_ContractAlignedLayerTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.StakeRegistry(&_ContractAlignedLayerTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractAlignedLayerTaskManager.Contract.StakeRegistry(&_ContractAlignedLayerTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractAlignedLayerTaskManager.Contract.TaskNumber(&_ContractAlignedLayerTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractAlignedLayerTaskManager.Contract.TaskNumber(&_ContractAlignedLayerTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractAlignedLayerTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractAlignedLayerTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractAlignedLayerTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractAlignedLayerTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractAlignedLayerTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractAlignedLayerTaskManager.Contract.TrySignatureAndApkVerification(&_ContractAlignedLayerTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractAlignedLayerTaskManager.Contract.TrySignatureAndApkVerification(&_ContractAlignedLayerTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x99c4a6b1.
//
// Solidity: function createNewTask(bytes proof, bytes pubInput, uint16 verifierId, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, proof []byte, pubInput []byte, verifierId uint16, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.contract.Transact(opts, "createNewTask", proof, pubInput, verifierId, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x99c4a6b1.
//
// Solidity: function createNewTask(bytes proof, bytes pubInput, uint16 verifierId, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) CreateNewTask(proof []byte, pubInput []byte, verifierId uint16, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.CreateNewTask(&_ContractAlignedLayerTaskManager.TransactOpts, proof, pubInput, verifierId, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x99c4a6b1.
//
// Solidity: function createNewTask(bytes proof, bytes pubInput, uint16 verifierId, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactorSession) CreateNewTask(proof []byte, pubInput []byte, verifierId uint16, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.CreateNewTask(&_ContractAlignedLayerTaskManager.TransactOpts, proof, pubInput, verifierId, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.Initialize(&_ContractAlignedLayerTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.Initialize(&_ContractAlignedLayerTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.Pause(&_ContractAlignedLayerTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.Pause(&_ContractAlignedLayerTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.PauseAll(&_ContractAlignedLayerTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.PauseAll(&_ContractAlignedLayerTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x8f310c24.
//
// Solidity: function raiseAndResolveChallenge((bytes,bytes,uint16,uint32,bytes,uint32) task, (uint32,bool) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IAlignedLayerTaskManagerTask, taskResponse IAlignedLayerTaskManagerTaskResponse, taskResponseMetadata IAlignedLayerTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x8f310c24.
//
// Solidity: function raiseAndResolveChallenge((bytes,bytes,uint16,uint32,bytes,uint32) task, (uint32,bool) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) RaiseAndResolveChallenge(task IAlignedLayerTaskManagerTask, taskResponse IAlignedLayerTaskManagerTaskResponse, taskResponseMetadata IAlignedLayerTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.RaiseAndResolveChallenge(&_ContractAlignedLayerTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x8f310c24.
//
// Solidity: function raiseAndResolveChallenge((bytes,bytes,uint16,uint32,bytes,uint32) task, (uint32,bool) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactorSession) RaiseAndResolveChallenge(task IAlignedLayerTaskManagerTask, taskResponse IAlignedLayerTaskManagerTaskResponse, taskResponseMetadata IAlignedLayerTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.RaiseAndResolveChallenge(&_ContractAlignedLayerTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.RenounceOwnership(&_ContractAlignedLayerTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.RenounceOwnership(&_ContractAlignedLayerTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x46bb7bae.
//
// Solidity: function respondToTask((bytes,bytes,uint16,uint32,bytes,uint32) task, (uint32,bool) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IAlignedLayerTaskManagerTask, taskResponse IAlignedLayerTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x46bb7bae.
//
// Solidity: function respondToTask((bytes,bytes,uint16,uint32,bytes,uint32) task, (uint32,bool) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) RespondToTask(task IAlignedLayerTaskManagerTask, taskResponse IAlignedLayerTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.RespondToTask(&_ContractAlignedLayerTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x46bb7bae.
//
// Solidity: function respondToTask((bytes,bytes,uint16,uint32,bytes,uint32) task, (uint32,bool) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactorSession) RespondToTask(task IAlignedLayerTaskManagerTask, taskResponse IAlignedLayerTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.RespondToTask(&_ContractAlignedLayerTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.SetPauserRegistry(&_ContractAlignedLayerTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.SetPauserRegistry(&_ContractAlignedLayerTaskManager.TransactOpts, newPauserRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.TransferOwnership(&_ContractAlignedLayerTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.TransferOwnership(&_ContractAlignedLayerTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.Unpause(&_ContractAlignedLayerTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAlignedLayerTaskManager.Contract.Unpause(&_ContractAlignedLayerTaskManager.TransactOpts, newPausedStatus)
}

// ContractAlignedLayerTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerInitializedIterator struct {
	Event *ContractAlignedLayerTaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAlignedLayerTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAlignedLayerTaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAlignedLayerTaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAlignedLayerTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAlignedLayerTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAlignedLayerTaskManagerInitialized represents a Initialized event raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractAlignedLayerTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManagerInitializedIterator{contract: _ContractAlignedLayerTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAlignedLayerTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAlignedLayerTaskManagerInitialized)
				if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractAlignedLayerTaskManagerInitialized, error) {
	event := new(ContractAlignedLayerTaskManagerInitialized)
	if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAlignedLayerTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerNewTaskCreatedIterator struct {
	Event *ContractAlignedLayerTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAlignedLayerTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAlignedLayerTaskManagerNewTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAlignedLayerTaskManagerNewTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAlignedLayerTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAlignedLayerTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAlignedLayerTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IAlignedLayerTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0xc76bc85127c84882e4e28847aeef6dac7f160a43a5e7f22a550d32d06e4efce3.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (bytes,bytes,uint16,uint32,bytes,uint32) task)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractAlignedLayerTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManagerNewTaskCreatedIterator{contract: _ContractAlignedLayerTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0xc76bc85127c84882e4e28847aeef6dac7f160a43a5e7f22a550d32d06e4efce3.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (bytes,bytes,uint16,uint32,bytes,uint32) task)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractAlignedLayerTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAlignedLayerTaskManagerNewTaskCreated)
				if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTaskCreated is a log parse operation binding the contract event 0xc76bc85127c84882e4e28847aeef6dac7f160a43a5e7f22a550d32d06e4efce3.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (bytes,bytes,uint16,uint32,bytes,uint32) task)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractAlignedLayerTaskManagerNewTaskCreated, error) {
	event := new(ContractAlignedLayerTaskManagerNewTaskCreated)
	if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAlignedLayerTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerOwnershipTransferredIterator struct {
	Event *ContractAlignedLayerTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAlignedLayerTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAlignedLayerTaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAlignedLayerTaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAlignedLayerTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAlignedLayerTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAlignedLayerTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAlignedLayerTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManagerOwnershipTransferredIterator{contract: _ContractAlignedLayerTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAlignedLayerTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAlignedLayerTaskManagerOwnershipTransferred)
				if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractAlignedLayerTaskManagerOwnershipTransferred, error) {
	event := new(ContractAlignedLayerTaskManagerOwnershipTransferred)
	if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAlignedLayerTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerPausedIterator struct {
	Event *ContractAlignedLayerTaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAlignedLayerTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAlignedLayerTaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAlignedLayerTaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAlignedLayerTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAlignedLayerTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAlignedLayerTaskManagerPaused represents a Paused event raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractAlignedLayerTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManagerPausedIterator{contract: _ContractAlignedLayerTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractAlignedLayerTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAlignedLayerTaskManagerPaused)
				if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) ParsePaused(log types.Log) (*ContractAlignedLayerTaskManagerPaused, error) {
	event := new(ContractAlignedLayerTaskManagerPaused)
	if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAlignedLayerTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerPauserRegistrySetIterator struct {
	Event *ContractAlignedLayerTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAlignedLayerTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAlignedLayerTaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAlignedLayerTaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAlignedLayerTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAlignedLayerTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAlignedLayerTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractAlignedLayerTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManagerPauserRegistrySetIterator{contract: _ContractAlignedLayerTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractAlignedLayerTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAlignedLayerTaskManagerPauserRegistrySet)
				if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractAlignedLayerTaskManagerPauserRegistrySet, error) {
	event := new(ContractAlignedLayerTaskManagerPauserRegistrySet)
	if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAlignedLayerTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractAlignedLayerTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAlignedLayerTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAlignedLayerTaskManagerTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAlignedLayerTaskManagerTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAlignedLayerTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAlignedLayerTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAlignedLayerTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractAlignedLayerTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractAlignedLayerTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractAlignedLayerTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAlignedLayerTaskManagerTaskChallengedSuccessfully)
				if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractAlignedLayerTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractAlignedLayerTaskManagerTaskChallengedSuccessfully)
	if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractAlignedLayerTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractAlignedLayerTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAlignedLayerTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerTaskCompletedIterator struct {
	Event *ContractAlignedLayerTaskManagerTaskCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAlignedLayerTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAlignedLayerTaskManagerTaskCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAlignedLayerTaskManagerTaskCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAlignedLayerTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAlignedLayerTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAlignedLayerTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractAlignedLayerTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManagerTaskCompletedIterator{contract: _ContractAlignedLayerTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractAlignedLayerTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAlignedLayerTaskManagerTaskCompleted)
				if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractAlignedLayerTaskManagerTaskCompleted, error) {
	event := new(ContractAlignedLayerTaskManagerTaskCompleted)
	if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAlignedLayerTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerTaskRespondedIterator struct {
	Event *ContractAlignedLayerTaskManagerTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAlignedLayerTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAlignedLayerTaskManagerTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAlignedLayerTaskManagerTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAlignedLayerTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAlignedLayerTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAlignedLayerTaskManagerTaskResponded represents a TaskResponded event raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerTaskResponded struct {
	TaskResponse         IAlignedLayerTaskManagerTaskResponse
	TaskResponseMetadata IAlignedLayerTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x46c382fc6ec03c126042e96e5be20ec429b355ffe4b56d4614d3a6b627da3a7e.
//
// Solidity: event TaskResponded((uint32,bool) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractAlignedLayerTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManagerTaskRespondedIterator{contract: _ContractAlignedLayerTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x46c382fc6ec03c126042e96e5be20ec429b355ffe4b56d4614d3a6b627da3a7e.
//
// Solidity: event TaskResponded((uint32,bool) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractAlignedLayerTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAlignedLayerTaskManagerTaskResponded)
				if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskResponded is a log parse operation binding the contract event 0x46c382fc6ec03c126042e96e5be20ec429b355ffe4b56d4614d3a6b627da3a7e.
//
// Solidity: event TaskResponded((uint32,bool) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractAlignedLayerTaskManagerTaskResponded, error) {
	event := new(ContractAlignedLayerTaskManagerTaskResponded)
	if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAlignedLayerTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerUnpausedIterator struct {
	Event *ContractAlignedLayerTaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAlignedLayerTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAlignedLayerTaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAlignedLayerTaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAlignedLayerTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAlignedLayerTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAlignedLayerTaskManagerUnpaused represents a Unpaused event raised by the ContractAlignedLayerTaskManager contract.
type ContractAlignedLayerTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractAlignedLayerTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAlignedLayerTaskManagerUnpausedIterator{contract: _ContractAlignedLayerTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAlignedLayerTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAlignedLayerTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAlignedLayerTaskManagerUnpaused)
				if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAlignedLayerTaskManager *ContractAlignedLayerTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractAlignedLayerTaskManagerUnpaused, error) {
	event := new(ContractAlignedLayerTaskManagerUnpaused)
	if err := _ContractAlignedLayerTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
