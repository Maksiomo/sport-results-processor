pragma solidity ^0.8.0;

// SPDX-License-Identifier: MIT

contract CompetitionTeamsRegistry {
    event CompetitionTeamsRecorded(uint256 indexed id, bytes32 recordHash);
    mapping(uint256 => bool) public recorded;
    mapping(uint256 => bytes32) public hashes;

    function recordCompetitionTeams(uint256 id, bytes32 recordHash) external {
        require(!recorded[id], "CompetitionTeams already recorded");
        recorded[id] = true;
        hashes[id] = recordHash;
        emit CompetitionTeamsRecorded(id, recordHash);
    }

    function validateCompetitionTeams(uint256 id, bytes32 recordHash) external view returns (bool) {
        return recorded[id] && hashes[id] == recordHash;
    }
}