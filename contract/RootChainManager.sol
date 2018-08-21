pragma solidity ^0.4.20;

contract RootChainManager {
    struct UserInfo {
        address from;
        uint256 timestamp;
    }

    mapping (address => UserInfo) public contractUser;
    event SubmitContract(
        address addr,
        address from
    );

    function submitAddress(address _addr) public {
        require(_addr != 0x0);
        uint32 size;
        assembly {
            size := extcodesize(_addr)
        }
        require(size>0);
        contractUser[_addr] = UserInfo({
           from: msg.sender,
           timestamp: block.timestamp
        });

        emit SubmitContract(_addr, msg.sender);
    }
}
