/*

  Copyright 2019 ZeroEx Intl.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

pragma solidity 0.5.15;


// solhint-disable no-empty-blocks
contract ICoordinatorRegistryCore
{
    /// @dev Emitted when a Coordinator endpoint is set.
    event CoordinatorEndpointSet(
        address coordinatorOperator,
        string coordinatorEndpoint
    );

    /// @dev Called by a Coordinator operator to set the endpoint of their Coordinator.
    /// @param coordinatorEndpoint Endpoint of the Coordinator as a string.
    function setCoordinatorEndpoint(string calldata coordinatorEndpoint) external;

    /// @dev Gets the endpoint for a Coordinator.
    /// @param coordinatorOperator Operator of the Coordinator endpoint.
    /// @return coordinatorEndpoint Endpoint of the Coordinator as a string.
    function getCoordinatorEndpoint(address coordinatorOperator)
        external
        view
        returns (string memory coordinatorEndpoint);
}
