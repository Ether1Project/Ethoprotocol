//pragma solidity 0.4.25;
//pragma solidity 0.5.12-develop.2019.8.13+commit.e91c6acb.Linux.g++;
pragma experimental ABIEncoderV2;

contract NodeSystem {

    mapping (string => uint) public nodeIdMap; 
    mapping (address => uint) public nodeAddressMap; 
    mapping (string => uint) public nodeIpMap; 
    mapping (address => signedNodeValidation) public lastNodeActivity;
    mapping (string => nodeType) public mainNodeMap;
    mapping (address => string) public nodeTypeMap; 

    string[] public nodeTypeArray;

    uint public nodeCount;
    uint public nodeTypesCounter;

    address internal owner;
    address internal operator;
    
    struct Node {
        uint index;
        address nodeAddress;
        string nodeId;
        string nodeIp;
        string nodePort;
        uint collateral;
    }

    struct signedNodeValidation {
        address nodeAddress;
        bytes[] signatures;
        uint signatureCount;
	bytes signatureBlockHash;
        bytes publicKey;
        uint blockHeight;
    }

    struct nodeType {
        string name;
        uint collateral;
        mapping (address => Node) nodeMap; 
        mapping (uint => address) nodeIndexMap;
        uint nodeTypeCount;
    }
    
    constructor() public {
        owner = msg.sender;
        operator = msg.sender;
        nodeCount = 0;
        nodeTypesCounter = 0;
    }    

    function addNodeType(string memory name, uint collateral) ownerOrOperator public {

        if(keccak256(abi.encodePacked((mainNodeMap[name].name))) != keccak256(abi.encodePacked((name)))) {
	    nodeTypeArray[nodeTypesCounter] = name;
            nodeTypesCounter++;
        }

        mainNodeMap[name].name = name;
        mainNodeMap[name].collateral = collateral;
        mainNodeMap[name].nodeTypeCount = 0;
    }

    function nodeCheckIn(bytes[] memory signatures, bytes memory publicKey, bytes memory signatureBlockHash) public {

        string memory nodeCheckInType = nodeTypeMap[msg.sender];
 
        assert(keccak256(abi.encodePacked((mainNodeMap[nodeCheckInType].name))) == keccak256(abi.encodePacked((nodeCheckInType))) && keccak256(abi.encodePacked((nodeCheckInType))) != keccak256(abi.encodePacked((""))));

        signedNodeValidation memory validation = signedNodeValidation({nodeAddress:tx.origin, signatures:signatures, signatureCount:signatures.length, signatureBlockHash:signatureBlockHash, publicKey:publicKey, blockHeight:block.number});
        lastNodeActivity[msg.sender] = validation;
    }
    
    function getSignatures(address nodeAddress) public view returns (bytes[] memory) {
        return lastNodeActivity[nodeAddress].signatures;
    }

    function addNode(string memory newNodeType, string memory id, string memory ip, string memory port) public payable returns (bool) {
        assert(msg.value == mainNodeMap[newNodeType].collateral * 1 ether && checkExistence(msg.sender, id, ip) && keccak256(abi.encodePacked((mainNodeMap[newNodeType].name))) == keccak256(abi.encodePacked((newNodeType))));

        nodeTypeMap[msg.sender] = newNodeType;

        mainNodeMap[newNodeType].nodeIndexMap[mainNodeMap[newNodeType].nodeTypeCount] = msg.sender;

        mainNodeMap[newNodeType].nodeMap[msg.sender].nodeAddress = msg.sender;
        mainNodeMap[newNodeType].nodeMap[msg.sender].collateral = msg.value;
        mainNodeMap[newNodeType].nodeMap[msg.sender].nodeId = id;
        mainNodeMap[newNodeType].nodeMap[msg.sender].nodeIp = ip;
        mainNodeMap[newNodeType].nodeMap[msg.sender].nodePort = port;
        mainNodeMap[newNodeType].nodeMap[msg.sender].index = mainNodeMap[newNodeType].nodeTypeCount;

        nodeIdMap[id] = 1111;
        nodeAddressMap[msg.sender] = 1111;
        nodeIpMap[ip] = 1111;

        mainNodeMap[newNodeType].nodeTypeCount++;
	nodeCount++;
        return true;
    }

    function removeNode() public returns (bool) {
        assert(mainNodeMap[nodeTypeMap[msg.sender]].nodeMap[msg.sender].collateral == mainNodeMap[nodeTypeMap[msg.sender]].collateral * 1 ether && keccak256(abi.encodePacked((nodeTypeMap[msg.sender]))) == keccak256(abi.encodePacked((mainNodeMap[nodeTypeMap[msg.sender]].name))));

        msg.sender.transfer(mainNodeMap[nodeTypeMap[msg.sender]].nodeMap[msg.sender].collateral);

        delete nodeIdMap[mainNodeMap[nodeTypeMap[msg.sender]].nodeMap[msg.sender].nodeId];
        delete nodeAddressMap[msg.sender];
        delete nodeIpMap[mainNodeMap[nodeTypeMap[msg.sender]].nodeMap[msg.sender].nodeIp];

        mainNodeMap[nodeTypeMap[msg.sender]].nodeMap[mainNodeMap[nodeTypeMap[msg.sender]].nodeIndexMap[mainNodeMap[nodeTypeMap[msg.sender]].nodeTypeCount-1]].index = mainNodeMap[nodeTypeMap[msg.sender]].nodeMap[msg.sender].index;
        mainNodeMap[nodeTypeMap[msg.sender]].nodeIndexMap[mainNodeMap[nodeTypeMap[msg.sender]].nodeMap[msg.sender].index] = mainNodeMap[nodeTypeMap[msg.sender]].nodeIndexMap[mainNodeMap[nodeTypeMap[msg.sender]].nodeTypeCount-1];

        delete mainNodeMap[nodeTypeMap[msg.sender]].nodeIndexMap[mainNodeMap[nodeTypeMap[msg.sender]].nodeTypeCount-1];
        delete mainNodeMap[nodeTypeMap[msg.sender]].nodeMap[msg.sender];

        mainNodeMap[nodeTypeMap[msg.sender]].nodeTypeCount--;

        delete nodeTypeMap[msg.sender];

        nodeCount--;

        return true;
    }

    function checkExistence(address nodeAddress, string memory id, string memory ip) public returns (bool) {
        return (nodeIdMap[id] != 1111 && nodeAddressMap[nodeAddress] != 1111 && nodeIpMap[ip] != 1111);
    }

    function setOperator(address newOperator) ownerOrOperator public {
        operator = newOperator;
    }

    modifier ownerOrOperator() {         
        if (msg.sender != owner && msg.sender != operator) revert();
        _;
    }
}
