pragma solidity ^0.8.0;

// SPDX-License-Identifier: MIT

contract TeamRegistry {
    event TeamRecorded(uint256 indexed id, bytes32 recordHash);
    mapping(uint256 => bool) public recorded;
    mapping(uint256 => bytes32) public hashes;

    function recordTeam(uint256 id, bytes32 recordHash) external {
        require(!recorded[id], "Team already recorded");
        recorded[id] = true;
        hashes[id] = recordHash;
        emit TeamRecorded(id, recordHash);
    }

    function validateTeam(uint256 id, bytes32 recordHash) external view returns (bool) {
        return recorded[id] && hashes[id] == recordHash;
    }
}