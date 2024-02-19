// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "./IAlignedLayerTaskManager.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";

/**
 * @title Primary entrypoint for procuring services from AlignedLayer.
 * @author Layr Labs, Inc.
 */
contract AlignedLayerServiceManager is ServiceManagerBase {
    using BytesLib for bytes;

    IAlignedLayerTaskManager
        public immutable alignedLayerTaskManager;

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlyAlignedLayerTaskManager() {
        require(
            msg.sender == address(alignedLayerTaskManager),
            "onlyAlignedLayerTaskManager: not from credible aligned layer task manager"
        );
        _;
    }

    constructor(
        IBLSRegistryCoordinatorWithIndices _registryCoordinator,
        ISlasher _slasher,
        IAlignedLayerTaskManager _alignedLayerTaskManager
    ) ServiceManagerBase(_registryCoordinator, _slasher) {
        alignedLayerTaskManager = _alignedLayerTaskManager;
    }

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    /// @dev The Slasher contract is under active development and its interface expected to change.
    ///      We recommend writing slashing logic without integrating with the Slasher at this point in time.
    function freezeOperator(
        address operatorAddr
    ) external override onlyAlignedLayerTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }
}
