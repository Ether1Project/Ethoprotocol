pragma solidity ^0.4.18;

contract PinStorage {

    function deleteContract() public onlyOwner {
        selfdestruct(owner);
    }

    uint32 public ReplicationFactor;
    uint32 public PinCount;
    address internal owner; // owner of the contract
    address internal operator; // operator or control contract
    mapping (string => uint32) PinData;
    string[] public Pins;

    function PinStorage () public {
        owner = msg.sender;
        operator = msg.sender;    
        PinCount = 0;
        ReplicationFactor = 0;
    }

    function SetReplicationFactor(uint32 set) public onlyOwnerOrOperator {
        ReplicationFactor = set;
    }

    function FindArrayKey(string pin) private constant returns (uint) {
        uint i = 0;
        while (keccak256(Pins[i]) != keccak256(pin)) {
            i++;
        }
        return i;
    }

    function RemoveByValue(string pin) private {
        uint i = FindArrayKey(pin);
        RemoveByIndex(i);
    }

    function RemoveByIndex(uint i) private {
        while (i < (Pins.length - 1)) {
            Pins[i] = Pins[i+1];
            i++;
        }
        Pins.length--;
    }

    function PinAdd(string pin, uint32 size) public onlyOwnerOrOperator {
        Pins.push(pin);
        PinData[pin] = size;
        PinCount++;
    }

    function PinRemove(string pin) public onlyOwnerOrOperator {
        RemoveByValue(pin);
        delete PinData[pin];
        PinCount--;
    }
    
    function GetPinSize(string pin) public constant returns(uint32) {
         return PinData[pin];
    }

    modifier onlyOwner() 
    {
       if (msg.sender != owner) revert();
       _;
    }

    modifier onlyOwnerOrOperator() 
    {
       if (msg.sender != owner && msg.sender != operator) revert();
       _;
    }

    function changeOperator(address newOperator) public onlyOwner
    {
       operator = newOperator;
    }
}
