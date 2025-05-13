pragma solidity ^0.8.0;

// SPDX-License-Identifier: MIT

contract TeamPersonRegistry {
    event TeamPersonRecorded(uint256 indexed id, string recordHash);
    mapping(uint256 => bool) public recorded;
    mapping(uint256 => string) public hashes;

    function recordTeamPerson(uint256 id, string calldata recordHash) external {
        require(!recorded[id], "TeamPerson already recorded");
        recorded[id] = true;
        hashes[id] = recordHash;
        emit TeamPersonRecorded(id, recordHash);
    }

    function validateTeamPerson(uint256 id, string calldata recordHash) external view returns (bool) {
        return recorded[id] && keccak256(bytes(hashes[id])) == keccak256(bytes(recordHash));
    }
}

contract CompetitionRegistry {
    event CompetitionRecorded(uint256 indexed id, string recordHash);
    mapping(uint256 => bool) public recorded;
    mapping(uint256 => string) public hashes;

    function recordCompetition(uint256 id, string calldata recordHash) external {
        require(!recorded[id], "Competition already recorded");
        recorded[id] = true;
        hashes[id] = recordHash;
        emit CompetitionRecorded(id, recordHash);
    }

    function validateCompetition(uint256 id, string calldata recordHash) external view returns (bool) {
        return recorded[id] && keccak256(bytes(hashes[id])) == keccak256(bytes(recordHash));
    }
}

contract CompetitionTeamsRegistry {
    event CompetitionTeamsRecorded(uint256 indexed id, string recordHash);
    mapping(uint256 => bool) public recorded;
    mapping(uint256 => string) public hashes;

    function recordCompetitionTeams(uint256 id, string calldata recordHash) external {
        require(!recorded[id], "CompetitionTeams already recorded");
        recorded[id] = true;
        hashes[id] = recordHash;
        emit CompetitionTeamsRecorded(id, recordHash);
    }

    function validateCompetitionTeams(uint256 id, string calldata recordHash) external view returns (bool) {
        return recorded[id] && keccak256(bytes(hashes[id])) == keccak256(bytes(recordHash));
    }
}

contract MatchParticipantRegistry {
    event MatchParticipantRecorded(uint256 indexed id, string recordHash);
    mapping(uint256 => bool) public recorded;
    mapping(uint256 => string) public hashes;

    function recordMatchParticipant(uint256 id, string calldata recordHash) external {
        require(!recorded[id], "MatchParticipant already recorded");
        recorded[id] = true;
        hashes[id] = recordHash;
        emit MatchParticipantRecorded(id, recordHash);
    }

    function validateMatchParticipant(uint256 id, string calldata recordHash) external view returns (bool) {
        return recorded[id] && keccak256(bytes(hashes[id])) == keccak256(bytes(recordHash));
    }
}

contract MatchRegistry {
    event MatchRecorded(uint256 indexed id, string recordHash);
    mapping(uint256 => bool) public recorded;
    mapping(uint256 => string) public hashes;

    function recordMatch(uint256 id, string calldata recordHash) external {
        require(!recorded[id], "Match already recorded");
        recorded[id] = true;
        hashes[id] = recordHash;
        emit MatchRecorded(id, recordHash);
    }

    function validateMatch(uint256 id, string calldata recordHash) external view returns (bool) {
        return recorded[id] && keccak256(bytes(hashes[id])) == keccak256(bytes(recordHash));
    }
}

contract PersonRegistry {
    event PersonRecorded(uint256 indexed id, string recordHash);
    mapping(uint256 => bool) public recorded;
    mapping(uint256 => string) public hashes;

    function recordPerson(uint256 id, string calldata recordHash) external {
        require(!recorded[id], "Person already recorded");
        recorded[id] = true;
        hashes[id] = recordHash;
        emit PersonRecorded(id, recordHash);
    }

    function validatePerson(uint256 id, string calldata recordHash) external view returns (bool) {
        return recorded[id] && keccak256(bytes(hashes[id])) == keccak256(bytes(recordHash));
    }
}

contract PrizeRegistry {
    event PrizeRecorded(uint256 indexed id, string recordHash);
    mapping(uint256 => bool) public recorded;
    mapping(uint256 => string) public hashes;

    function recordPrize(uint256 id, string calldata recordHash) external {
        require(!recorded[id], "Prize already recorded");
        recorded[id] = true;
        hashes[id] = recordHash;
        emit PrizeRecorded(id, recordHash);
    }

    function validatePrize(uint256 id, string calldata recordHash) external view returns (bool) {
        return recorded[id] && keccak256(bytes(hashes[id])) == keccak256(bytes(recordHash));
    }
}

contract SportRegistry {
    event SportRecorded(uint256 indexed id, string recordHash);
    mapping(uint256 => bool) public recorded;
    mapping(uint256 => string) public hashes;

    function recordSport(uint256 id, string calldata recordHash) external {
        require(!recorded[id], "Sport already recorded");
        recorded[id] = true;
        hashes[id] = recordHash;
        emit SportRecorded(id, recordHash);
    }

    function validateSport(uint256 id, string calldata recordHash) external view returns (bool) {
        return recorded[id] && keccak256(bytes(hashes[id])) == keccak256(bytes(recordHash));
    }
}

contract TeamAchievementsRegistry {
    event TeamAchievementsRecorded(uint256 indexed id, string recordHash);
    mapping(uint256 => bool) public recorded;
    mapping(uint256 => string) public hashes;

    function recordTeamAchievements(uint256 id, string calldata recordHash) external {
        require(!recorded[id], "TeamAchievements already recorded");
        recorded[id] = true;
        hashes[id] = recordHash;
        emit TeamAchievementsRecorded(id, recordHash);
    }

    function validateTeamAchievements(uint256 id, string calldata recordHash) external view returns (bool) {
        return recorded[id] && keccak256(bytes(hashes[id])) == keccak256(bytes(recordHash));
    }
}

contract TeamRegistry {
    event TeamRecorded(uint256 indexed id, string recordHash);
    mapping(uint256 => bool) public recorded;
    mapping(uint256 => string) public hashes;

    function recordTeam(uint256 id, string calldata recordHash) external {
        require(!recorded[id], "Team already recorded");
        recorded[id] = true;
        hashes[id] = recordHash;
        emit TeamRecorded(id, recordHash);
    }

    function validateTeam(uint256 id, string calldata recordHash) external view returns (bool) {
        return recorded[id] && keccak256(bytes(hashes[id])) == keccak256(bytes(recordHash));
    }
}