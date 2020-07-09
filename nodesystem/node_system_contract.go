// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nodesystem

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// NodeSystemABI is the input ABI used to generate the binding from.
const NodeSystemABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"nodeIdMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"checkExistence\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatureBlockHash\",\"type\":\"bytes\"}],\"name\":\"nodeCheckIn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"addNodeType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"nodeIpMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeAddressMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeTypeArray\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nodeTypesCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nodeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"mainNodeMap\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeTypeCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeTypeMap\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastNodeActivity\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"signatureCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatureBlockHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"newNodeType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getSignatures\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// NodeSystemBin is the compiled bytecode used for deploying new contracts.
const NodeSystemBin = `0x608060405234801561001057600080fd5b50600980546001600160a01b031990811633908117909255600a8054909116909117905560006007819055600855611eb68061004d6000396000f3fe6080604052600436106100f35760003560e01c80636da49b831161008a578063c7cfc9ea11610059578063c7cfc9ea146102a3578063d0d3f5ba146102d4578063d425bd81146102e9578063ed5b5015146102fc576100f3565b80636da49b831461021f5780638d80313414610234578063b3ab15fb14610263578063b97454ba14610283576100f3565b80631887f13a116100c65780631887f13a1461019d5780631f901c4f146101bd5780634cff1950146101dd5780634d382c701461020a576100f3565b806302d1a17f146100f857806302f2b5bd1461012e57806306e5b5461461015b578063142685bf1461017d575b600080fd5b34801561010457600080fd5b50610118610113366004611951565b610329565b6040516101259190611d2c565b60405180910390f35b34801561013a57600080fd5b5061014e61014936600461189c565b610346565b6040516101259190611ce0565b34801561016757600080fd5b5061017b610176366004611919565b6103c5565b005b34801561018957600080fd5b5061017b610198366004611a46565b6105fc565b3480156101a957600080fd5b506101186101b8366004611951565b610757565b3480156101c957600080fd5b506101186101d836600461187e565b610774565b3480156101e957600080fd5b506101fd6101f8366004611a97565b610786565b6040516101259190611cee565b34801561021657600080fd5b5061011861082c565b34801561022b57600080fd5b50610118610832565b34801561024057600080fd5b5061025461024f366004611951565b610838565b60405161012593929190611cff565b34801561026f57600080fd5b5061017b61027e36600461187e565b6108f1565b34801561028f57600080fd5b506101fd61029e36600461187e565b610943565b3480156102af57600080fd5b506102c36102be36600461187e565b6109ab565b604051610125959493929190611c75565b3480156102e057600080fd5b5061014e610af5565b61014e6102f7366004611986565b61117b565b34801561030857600080fd5b5061031c61031736600461187e565b61151c565b6040516101259190611ccf565b805160208183018101805160008252928201919093012091525481565b600080836040516103579190611c52565b9081526020016040518091039020546104571415801561039157506001600160a01b03841660009081526001602052604090205461045714155b80156103bd57506002826040516103a89190611c52565b90815260200160405180910390205461045714155b949350505050565b3360009081526005602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084526060939283018282801561045b5780601f106104305761010080835404028352916020019161045b565b820191906000526020600020905b81548152906001019060200180831161043e57829003601f168201915b50505050509050806040516020016104739190611c52565b604051602081830303815290604052805190602001206004826040516104999190611c52565b908152604051602091819003820181206104b592909101611c5e565b6040516020818303038152906040528051906020012014801561052257506040516020016104e290611c6a565b60405160208183030381529060405280519060200120816040516020016105099190611c52565b6040516020818303038152906040528051906020012014155b61052857fe5b61053061160d565b506040805160c081018252328152602080820187815287518385015260608301869052608083018790524360a08401523360009081526003835293909320825181546001600160a01b0319166001600160a01b03909116178155925180519293849390926105a592600185019291019061164c565b5060408201516002820155606082015180516105cb9160038401916020909101906116a9565b50608082015180516105e79160048401916020909101906116a9565b5060a082015181600501559050505050505050565b6009546001600160a01b031633148015906106225750600a546001600160a01b03163314155b1561062c57600080fd5b8160405160200161063d9190611c52565b604051602081830303815290604052805190602001206004836040516106639190611c52565b9081526040516020918190038201812061067f92909101611c5e565b60405160208183030381529060405280519060200120146106d257816006600854815481106106aa57fe5b9060005260206000200190805190602001906106c79291906116a9565b506008805460010190555b816004836040516106e39190611c52565b908152602001604051809103902060000190805190602001906107079291906116a9565b50806004836040516107199190611c52565b90815260200160405180910390206001018190555060006004836040516107409190611c52565b908152604051908190036020019020600401555050565b805160208183018101805160028252928201919093012091525481565b60016020526000908152604090205481565b6006818154811061079357fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152935090918301828280156108245780601f106107f957610100808354040283529160200191610824565b820191906000526020600020905b81548152906001019060200180831161080757829003601f168201915b505050505081565b60085481565b60075481565b805160208183018101805160048252928201938201939093209190925280546040805160026001841615610100026000190190931692909204601f81018590048502830185019091528082529192909183918301828280156108db5780601f106108b0576101008083540402835291602001916108db565b820191906000526020600020905b8154815290600101906020018083116108be57829003601f168201915b5050505050908060010154908060040154905083565b6009546001600160a01b031633148015906109175750600a546001600160a01b03163314155b1561092157600080fd5b600a80546001600160a01b0319166001600160a01b0392909216919091179055565b60056020908152600091825260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156108245780601f106107f957610100808354040283529160200191610824565b6003602081815260009283526040928390208054600280830154948301805487516101006001831615026000190190911692909204601f81018690048602830186019097528682526001600160a01b03909216959293909290830182828015610a555780601f10610a2a57610100808354040283529160200191610a55565b820191906000526020600020905b815481529060010190602001808311610a3857829003601f168201915b5050505060048301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015610ae55780601f10610aba57610100808354040283529160200191610ae5565b820191906000526020600020905b815481529060010190602001808311610ac857829003601f168201915b5050505050908060050154905085565b336000908152600560205260408082209051600491610b1391611c5e565b908152604080519182900360209081018320600101543360009081526005909252919020670de0b6b3a764000090910291600491610b5091611c5e565b90815260200160405180910390206002016000336001600160a01b03166001600160a01b0316815260200190815260200160002060050154148015610c16575033600090815260056020526040908190209051600491610baf91611c5e565b90815260405160209181900382018120610bcb92909101611c5e565b60408051601f19818403018152828252805160209182012033600090815260058352929092209192610bfe929101611c5e565b60405160208183030381529060405280519060200120145b610c1c57fe5b336000818152600560205260409081902090516108fc91600491610c409190611c5e565b90815260408051602092819003830181203360009081526002909101909352908220600501548015939093029291818181858888f19350505050158015610c8b573d6000803e3d6000fd5b50336000908152600560205260408082209051600491610caa91611c5e565b9081526040805191829003602090810183203360009081526002918201909252919020610cd992910190611c5e565b90815260408051918290036020908101832060009081905533815260018252828120819055600590915220600291600491610d1391611c5e565b908152604080519182900360209081018320336000908152600290910190915220610d4391600390910190611c5e565b908152604080519182900360209081018320600090819055338152600590915220600491610d719190611c5e565b90815260408051602092819003830181203360009081526002909101845282812054600590945291909120600491610da99190611c5e565b90815260200160405180910390206002016000600460056000336001600160a01b03166001600160a01b03168152602001908152602001600020604051610df09190611c5e565b908152602001604051809103902060030160006001600460056000336001600160a01b03166001600160a01b03168152602001908152602001600020604051610e399190611c5e565b9081526040805160209281900383019020600490810154939093038452838201949094529183016000908120546001600160a01b031685528483019590955292820184209490945533835260059093529082902091519091610e9a91611c5e565b908152602001604051809103902060030160006001600460056000336001600160a01b03166001600160a01b03168152602001908152602001600020604051610ee39190611c5e565b90815260200160405180910390206004015403815260200190815260200160002060009054906101000a90046001600160a01b0316600460056000336001600160a01b03166001600160a01b03168152602001908152602001600020604051610f4c9190611c5e565b90815260200160405180910390206003016000600460056000336001600160a01b03166001600160a01b03168152602001908152602001600020604051610f939190611c5e565b90815260408051602092819003830190203360008181526002909201845282822054855284840195909552928101832080546001600160a01b0319166001600160a01b03969096169590951790945591815260059091528190209051600491610ffb91611c5e565b908152602001604051809103902060030160006001600460056000336001600160a01b03166001600160a01b031681526020019081526020016000206040516110449190611c5e565b90815260200160405180910390206004015403815260200190815260200160002060006101000a8154906001600160a01b030219169055600460056000336001600160a01b03166001600160a01b031681526020019081526020016000206040516110af9190611c5e565b908152604080516020928190038301902033600090815260029182019093529082208281556001810180546001600160a01b031916905591906110f490830182611723565b611102600383016000611723565b611110600483016000611723565b5060006005918201819055338152602091909152604090819020905160049161113891611c5e565b908152604080516020928190038301902060040180546000190190553360009081526005909252812061116a91611723565b506007805460001901905560015b90565b600060048560405161118d9190611c52565b908152602001604051809103902060010154670de0b6b3a764000002341480156111bd57506111bd338585610346565b801561122f5750846040516020016111d59190611c52565b604051602081830303815290604052805190602001206004866040516111fb9190611c52565b9081526040516020918190038201812061121792909101611c5e565b60405160208183030381529060405280519060200120145b61123557fe5b3360009081526005602090815260409091208651611255928801906116a9565b50336004866040516112679190611c52565b9081526020016040518091039020600301600060048860405161128a9190611c52565b908152602001604051809103902060040154815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b03160217905550336004866040516112df9190611c52565b90815260408051602092819003830181203360009081526002909101909352912060010180546001600160a01b0319166001600160a01b0393909316929092179091553490600490611332908890611c52565b90815260408051918290036020908101832033600090815260029091019091522060050191909155849060049061136a908890611c52565b90815260408051602092819003830190203360009081526002918201845291909120835161139e94919092019201906116a9565b50826004866040516113b09190611c52565b9081526040805160209281900383019020336000908152600290910183522082516113e59360039092019291909101906116a9565b50816004866040516113f79190611c52565b90815260408051602092819003830190203360009081526002909101835220825161142c9360049092019291909101906116a9565b5060048560405161143d9190611c52565b90815260200160405180910390206004015460048660405161145f9190611c52565b908152604080516020928190038301812033600090815260029091019093529082209290925561045791611494908790611c52565b9081526040805160209281900383018120939093553360009081526001909252902061045790819055906002906114cc908690611c52565b9081526020016040518091039020819055506004856040516114ee9190611c52565b9081526040519081900360200190206004018054600190810190915560078054820190559050949350505050565b6001600160a01b0381166000908152600360209081526040808320600101805482518185028101850190935280835260609492939192909184015b828210156116025760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156115ee5780601f106115c3576101008083540402835291602001916115ee565b820191906000526020600020905b8154815290600101906020018083116115d157829003601f168201915b505050505081526020019060010190611557565b505050509050919050565b6040518060c0016040528060006001600160a01b0316815260200160608152602001600081526020016060815260200160608152602001600081525090565b828054828255906000526020600020908101928215611699579160200282015b8281111561169957825180516116899184916020909101906116a9565b509160200191906001019061166c565b506116a592915061176a565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106116ea57805160ff1916838001178555611717565b82800160010185558215611717579182015b828111156117175782518255916020019190600101906116fc565b506116a592915061178d565b50805460018160011615610100020316600290046000825580601f106117495750611767565b601f016020900490600052602060002090810190611767919061178d565b50565b61117891905b808211156116a55760006117848282611723565b50600101611770565b61117891905b808211156116a55760008155600101611793565b80356117b281611e30565b92915050565b600082601f8301126117c957600080fd5b81356117dc6117d782611d61565b611d3a565b81815260209384019390925082018360005b8381101561181a57813586016118048882611824565b84525060209283019291909101906001016117ee565b5050505092915050565b600082601f83011261183557600080fd5b81356118436117d782611d82565b9150808252602083016020830185838301111561185f57600080fd5b61186a838284611dea565b50505092915050565b80356117b281611e44565b60006020828403121561189057600080fd5b60006103bd84846117a7565b6000806000606084860312156118b157600080fd5b60006118bd86866117a7565b935050602084013567ffffffffffffffff8111156118da57600080fd5b6118e686828701611824565b925050604084013567ffffffffffffffff81111561190357600080fd5b61190f86828701611824565b9150509250925092565b60008060006060848603121561192e57600080fd5b833567ffffffffffffffff81111561194557600080fd5b6118bd868287016117b8565b60006020828403121561196357600080fd5b813567ffffffffffffffff81111561197a57600080fd5b6103bd84828501611824565b6000806000806080858703121561199c57600080fd5b843567ffffffffffffffff8111156119b357600080fd5b6119bf87828801611824565b945050602085013567ffffffffffffffff8111156119dc57600080fd5b6119e887828801611824565b935050604085013567ffffffffffffffff811115611a0557600080fd5b611a1187828801611824565b925050606085013567ffffffffffffffff811115611a2e57600080fd5b611a3a87828801611824565b91505092959194509250565b60008060408385031215611a5957600080fd5b823567ffffffffffffffff811115611a7057600080fd5b611a7c85828601611824565b9250506020611a8d85828601611873565b9150509250929050565b600060208284031215611aa957600080fd5b60006103bd8484611873565b6000611ac18383611b4e565b9392505050565b611ad181611dce565b82525050565b6000611ae282611dbc565b611aec8185611dc0565b935083602082028501611afe85611daa565b8060005b85811015611b385784840389528151611b1b8582611ab5565b9450611b2683611daa565b60209a909a0199925050600101611b02565b5091979650505050505050565b611ad181611dd9565b6000611b5982611dbc565b611b638185611dc0565b9350611b73818560208601611df6565b611b7c81611e26565b9093019392505050565b6000611b9182611dbc565b611b9b8185611dc9565b9350611bab818560208601611df6565b9290920192915050565b600081546001811660008114611bd25760018114611bf557611c34565b607f6002830416611be38187611dc9565b60ff1984168152955085019250611c34565b60028204611c038187611dc9565b9550611c0e85611db0565b60005b82811015611c2d57815488820152600190910190602001611c11565b5050850192505b505092915050565b60006117b2600083611dc9565b611ad181611178565b6000611ac18284611b86565b6000611ac18284611bb5565b60006117b282611c3c565b60a08101611c838288611ac8565b611c906020830187611c49565b8181036040830152611ca28186611b4e565b90508181036060830152611cb68185611b4e565b9050611cc56080830184611c49565b9695505050505050565b60208082528101611ac18184611ad7565b602081016117b28284611b45565b60208082528101611ac18184611b4e565b60608082528101611d108186611b4e565b9050611d1f6020830185611c49565b6103bd6040830184611c49565b602081016117b28284611c49565b60405181810167ffffffffffffffff81118282101715611d5957600080fd5b604052919050565b600067ffffffffffffffff821115611d7857600080fd5b5060209081020190565b600067ffffffffffffffff821115611d9957600080fd5b506020601f91909101601f19160190565b60200190565b60009081526020902090565b5190565b90815260200190565b919050565b60006117b282611dde565b151590565b6001600160a01b031690565b82818337506000910152565b60005b83811015611e11578181015183820152602001611df9565b83811115611e20576000848401525b50505050565b601f01601f191690565b611e3981611dce565b811461176757600080fd5b611e398161117856fea365627a7a723158207fafafb8542901605e57428df232726754e7c2c53f655bb5a22bcc4bc47167d06c6578706572696d656e74616cf564736f6c637828302e352e31322d646576656c6f702e323031392e382e31332b636f6d6d69742e65393163366163620066`

// DeployNodeSystem deploys a new Ethereum contract, binding an instance of NodeSystem to it.
func DeployNodeSystem(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeSystem, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeSystemABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeSystemBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeSystem{NodeSystemCaller: NodeSystemCaller{contract: contract}, NodeSystemTransactor: NodeSystemTransactor{contract: contract}, NodeSystemFilterer: NodeSystemFilterer{contract: contract}}, nil
}

// NodeSystem is an auto generated Go binding around an Ethereum contract.
type NodeSystem struct {
	NodeSystemCaller     // Read-only binding to the contract
	NodeSystemTransactor // Write-only binding to the contract
	NodeSystemFilterer   // Log filterer for contract events
}

// NodeSystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeSystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeSystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeSystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeSystemSession struct {
	Contract     *NodeSystem       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeSystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeSystemCallerSession struct {
	Contract *NodeSystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NodeSystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeSystemTransactorSession struct {
	Contract     *NodeSystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NodeSystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeSystemRaw struct {
	Contract *NodeSystem // Generic contract binding to access the raw methods on
}

// NodeSystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeSystemCallerRaw struct {
	Contract *NodeSystemCaller // Generic read-only contract binding to access the raw methods on
}

// NodeSystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeSystemTransactorRaw struct {
	Contract *NodeSystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeSystem creates a new instance of NodeSystem, bound to a specific deployed contract.
func NewNodeSystem(address common.Address, backend bind.ContractBackend) (*NodeSystem, error) {
	contract, err := bindNodeSystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeSystem{NodeSystemCaller: NodeSystemCaller{contract: contract}, NodeSystemTransactor: NodeSystemTransactor{contract: contract}, NodeSystemFilterer: NodeSystemFilterer{contract: contract}}, nil
}

// NewNodeSystemCaller creates a new read-only instance of NodeSystem, bound to a specific deployed contract.
func NewNodeSystemCaller(address common.Address, caller bind.ContractCaller) (*NodeSystemCaller, error) {
	contract, err := bindNodeSystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeSystemCaller{contract: contract}, nil
}

// NewNodeSystemTransactor creates a new write-only instance of NodeSystem, bound to a specific deployed contract.
func NewNodeSystemTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeSystemTransactor, error) {
	contract, err := bindNodeSystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeSystemTransactor{contract: contract}, nil
}

// NewNodeSystemFilterer creates a new log filterer instance of NodeSystem, bound to a specific deployed contract.
func NewNodeSystemFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeSystemFilterer, error) {
	contract, err := bindNodeSystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeSystemFilterer{contract: contract}, nil
}

// bindNodeSystem binds a generic wrapper to an already deployed contract.
func bindNodeSystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeSystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeSystem *NodeSystemRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeSystem.Contract.NodeSystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeSystem *NodeSystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeSystem.Contract.NodeSystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeSystem *NodeSystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeSystem.Contract.NodeSystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeSystem *NodeSystemCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeSystem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeSystem *NodeSystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeSystem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeSystem *NodeSystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeSystem.Contract.contract.Transact(opts, method, params...)
}

// GetSignatures is a free data retrieval call binding the contract method 0xed5b5015.
//
// Solidity: function getSignatures(nodeAddress address) constant returns(bytes[])
func (_NodeSystem *NodeSystemCaller) GetSignatures(opts *bind.CallOpts, nodeAddress common.Address) ([][]byte, error) {
	var (
		ret0 = new([][]byte)
	)
	out := ret0
	err := _NodeSystem.contract.Call(opts, out, "getSignatures", nodeAddress)
	return *ret0, err
}

// GetSignatures is a free data retrieval call binding the contract method 0xed5b5015.
//
// Solidity: function getSignatures(nodeAddress address) constant returns(bytes[])
func (_NodeSystem *NodeSystemSession) GetSignatures(nodeAddress common.Address) ([][]byte, error) {
	return _NodeSystem.Contract.GetSignatures(&_NodeSystem.CallOpts, nodeAddress)
}

// GetSignatures is a free data retrieval call binding the contract method 0xed5b5015.
//
// Solidity: function getSignatures(nodeAddress address) constant returns(bytes[])
func (_NodeSystem *NodeSystemCallerSession) GetSignatures(nodeAddress common.Address) ([][]byte, error) {
	return _NodeSystem.Contract.GetSignatures(&_NodeSystem.CallOpts, nodeAddress)
}

// LastNodeActivity is a free data retrieval call binding the contract method 0xc7cfc9ea.
//
// Solidity: function lastNodeActivity( address) constant returns(nodeAddress address, signatureCount uint256, signatureBlockHash bytes, publicKey bytes, blockHeight uint256)
func (_NodeSystem *NodeSystemCaller) LastNodeActivity(opts *bind.CallOpts, arg0 common.Address) (struct {
	NodeAddress        common.Address
	SignatureCount     *big.Int
	SignatureBlockHash []byte
	PublicKey          []byte
	BlockHeight        *big.Int
}, error) {
	ret := new(struct {
		NodeAddress        common.Address
		SignatureCount     *big.Int
		SignatureBlockHash []byte
		PublicKey          []byte
		BlockHeight        *big.Int
	})
	out := ret
	err := _NodeSystem.contract.Call(opts, out, "lastNodeActivity", arg0)
	return *ret, err
}

// LastNodeActivity is a free data retrieval call binding the contract method 0xc7cfc9ea.
//
// Solidity: function lastNodeActivity( address) constant returns(nodeAddress address, signatureCount uint256, signatureBlockHash bytes, publicKey bytes, blockHeight uint256)
func (_NodeSystem *NodeSystemSession) LastNodeActivity(arg0 common.Address) (struct {
	NodeAddress        common.Address
	SignatureCount     *big.Int
	SignatureBlockHash []byte
	PublicKey          []byte
	BlockHeight        *big.Int
}, error) {
	return _NodeSystem.Contract.LastNodeActivity(&_NodeSystem.CallOpts, arg0)
}

// LastNodeActivity is a free data retrieval call binding the contract method 0xc7cfc9ea.
//
// Solidity: function lastNodeActivity( address) constant returns(nodeAddress address, signatureCount uint256, signatureBlockHash bytes, publicKey bytes, blockHeight uint256)
func (_NodeSystem *NodeSystemCallerSession) LastNodeActivity(arg0 common.Address) (struct {
	NodeAddress        common.Address
	SignatureCount     *big.Int
	SignatureBlockHash []byte
	PublicKey          []byte
	BlockHeight        *big.Int
}, error) {
	return _NodeSystem.Contract.LastNodeActivity(&_NodeSystem.CallOpts, arg0)
}

// MainNodeMap is a free data retrieval call binding the contract method 0x8d803134.
//
// Solidity: function mainNodeMap( string) constant returns(name string, collateral uint256, nodeTypeCount uint256)
func (_NodeSystem *NodeSystemCaller) MainNodeMap(opts *bind.CallOpts, arg0 string) (struct {
	Name          string
	Collateral    *big.Int
	NodeTypeCount *big.Int
}, error) {
	ret := new(struct {
		Name          string
		Collateral    *big.Int
		NodeTypeCount *big.Int
	})
	out := ret
	err := _NodeSystem.contract.Call(opts, out, "mainNodeMap", arg0)
	return *ret, err
}

// MainNodeMap is a free data retrieval call binding the contract method 0x8d803134.
//
// Solidity: function mainNodeMap( string) constant returns(name string, collateral uint256, nodeTypeCount uint256)
func (_NodeSystem *NodeSystemSession) MainNodeMap(arg0 string) (struct {
	Name          string
	Collateral    *big.Int
	NodeTypeCount *big.Int
}, error) {
	return _NodeSystem.Contract.MainNodeMap(&_NodeSystem.CallOpts, arg0)
}

// MainNodeMap is a free data retrieval call binding the contract method 0x8d803134.
//
// Solidity: function mainNodeMap( string) constant returns(name string, collateral uint256, nodeTypeCount uint256)
func (_NodeSystem *NodeSystemCallerSession) MainNodeMap(arg0 string) (struct {
	Name          string
	Collateral    *big.Int
	NodeTypeCount *big.Int
}, error) {
	return _NodeSystem.Contract.MainNodeMap(&_NodeSystem.CallOpts, arg0)
}

// NodeAddressMap is a free data retrieval call binding the contract method 0x1f901c4f.
//
// Solidity: function nodeAddressMap( address) constant returns(uint256)
func (_NodeSystem *NodeSystemCaller) NodeAddressMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeSystem.contract.Call(opts, out, "nodeAddressMap", arg0)
	return *ret0, err
}

// NodeAddressMap is a free data retrieval call binding the contract method 0x1f901c4f.
//
// Solidity: function nodeAddressMap( address) constant returns(uint256)
func (_NodeSystem *NodeSystemSession) NodeAddressMap(arg0 common.Address) (*big.Int, error) {
	return _NodeSystem.Contract.NodeAddressMap(&_NodeSystem.CallOpts, arg0)
}

// NodeAddressMap is a free data retrieval call binding the contract method 0x1f901c4f.
//
// Solidity: function nodeAddressMap( address) constant returns(uint256)
func (_NodeSystem *NodeSystemCallerSession) NodeAddressMap(arg0 common.Address) (*big.Int, error) {
	return _NodeSystem.Contract.NodeAddressMap(&_NodeSystem.CallOpts, arg0)
}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() constant returns(uint256)
func (_NodeSystem *NodeSystemCaller) NodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeSystem.contract.Call(opts, out, "nodeCount")
	return *ret0, err
}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() constant returns(uint256)
func (_NodeSystem *NodeSystemSession) NodeCount() (*big.Int, error) {
	return _NodeSystem.Contract.NodeCount(&_NodeSystem.CallOpts)
}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() constant returns(uint256)
func (_NodeSystem *NodeSystemCallerSession) NodeCount() (*big.Int, error) {
	return _NodeSystem.Contract.NodeCount(&_NodeSystem.CallOpts)
}

// NodeIdMap is a free data retrieval call binding the contract method 0x02d1a17f.
//
// Solidity: function nodeIdMap( string) constant returns(uint256)
func (_NodeSystem *NodeSystemCaller) NodeIdMap(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeSystem.contract.Call(opts, out, "nodeIdMap", arg0)
	return *ret0, err
}

// NodeIdMap is a free data retrieval call binding the contract method 0x02d1a17f.
//
// Solidity: function nodeIdMap( string) constant returns(uint256)
func (_NodeSystem *NodeSystemSession) NodeIdMap(arg0 string) (*big.Int, error) {
	return _NodeSystem.Contract.NodeIdMap(&_NodeSystem.CallOpts, arg0)
}

// NodeIdMap is a free data retrieval call binding the contract method 0x02d1a17f.
//
// Solidity: function nodeIdMap( string) constant returns(uint256)
func (_NodeSystem *NodeSystemCallerSession) NodeIdMap(arg0 string) (*big.Int, error) {
	return _NodeSystem.Contract.NodeIdMap(&_NodeSystem.CallOpts, arg0)
}

// NodeIpMap is a free data retrieval call binding the contract method 0x1887f13a.
//
// Solidity: function nodeIpMap( string) constant returns(uint256)
func (_NodeSystem *NodeSystemCaller) NodeIpMap(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeSystem.contract.Call(opts, out, "nodeIpMap", arg0)
	return *ret0, err
}

// NodeIpMap is a free data retrieval call binding the contract method 0x1887f13a.
//
// Solidity: function nodeIpMap( string) constant returns(uint256)
func (_NodeSystem *NodeSystemSession) NodeIpMap(arg0 string) (*big.Int, error) {
	return _NodeSystem.Contract.NodeIpMap(&_NodeSystem.CallOpts, arg0)
}

// NodeIpMap is a free data retrieval call binding the contract method 0x1887f13a.
//
// Solidity: function nodeIpMap( string) constant returns(uint256)
func (_NodeSystem *NodeSystemCallerSession) NodeIpMap(arg0 string) (*big.Int, error) {
	return _NodeSystem.Contract.NodeIpMap(&_NodeSystem.CallOpts, arg0)
}

// NodeTypeArray is a free data retrieval call binding the contract method 0x4cff1950.
//
// Solidity: function nodeTypeArray( uint256) constant returns(string)
func (_NodeSystem *NodeSystemCaller) NodeTypeArray(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NodeSystem.contract.Call(opts, out, "nodeTypeArray", arg0)
	return *ret0, err
}

// NodeTypeArray is a free data retrieval call binding the contract method 0x4cff1950.
//
// Solidity: function nodeTypeArray( uint256) constant returns(string)
func (_NodeSystem *NodeSystemSession) NodeTypeArray(arg0 *big.Int) (string, error) {
	return _NodeSystem.Contract.NodeTypeArray(&_NodeSystem.CallOpts, arg0)
}

// NodeTypeArray is a free data retrieval call binding the contract method 0x4cff1950.
//
// Solidity: function nodeTypeArray( uint256) constant returns(string)
func (_NodeSystem *NodeSystemCallerSession) NodeTypeArray(arg0 *big.Int) (string, error) {
	return _NodeSystem.Contract.NodeTypeArray(&_NodeSystem.CallOpts, arg0)
}

// NodeTypeMap is a free data retrieval call binding the contract method 0xb97454ba.
//
// Solidity: function nodeTypeMap( address) constant returns(string)
func (_NodeSystem *NodeSystemCaller) NodeTypeMap(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NodeSystem.contract.Call(opts, out, "nodeTypeMap", arg0)
	return *ret0, err
}

// NodeTypeMap is a free data retrieval call binding the contract method 0xb97454ba.
//
// Solidity: function nodeTypeMap( address) constant returns(string)
func (_NodeSystem *NodeSystemSession) NodeTypeMap(arg0 common.Address) (string, error) {
	return _NodeSystem.Contract.NodeTypeMap(&_NodeSystem.CallOpts, arg0)
}

// NodeTypeMap is a free data retrieval call binding the contract method 0xb97454ba.
//
// Solidity: function nodeTypeMap( address) constant returns(string)
func (_NodeSystem *NodeSystemCallerSession) NodeTypeMap(arg0 common.Address) (string, error) {
	return _NodeSystem.Contract.NodeTypeMap(&_NodeSystem.CallOpts, arg0)
}

// NodeTypesCounter is a free data retrieval call binding the contract method 0x4d382c70.
//
// Solidity: function nodeTypesCounter() constant returns(uint256)
func (_NodeSystem *NodeSystemCaller) NodeTypesCounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeSystem.contract.Call(opts, out, "nodeTypesCounter")
	return *ret0, err
}

// NodeTypesCounter is a free data retrieval call binding the contract method 0x4d382c70.
//
// Solidity: function nodeTypesCounter() constant returns(uint256)
func (_NodeSystem *NodeSystemSession) NodeTypesCounter() (*big.Int, error) {
	return _NodeSystem.Contract.NodeTypesCounter(&_NodeSystem.CallOpts)
}

// NodeTypesCounter is a free data retrieval call binding the contract method 0x4d382c70.
//
// Solidity: function nodeTypesCounter() constant returns(uint256)
func (_NodeSystem *NodeSystemCallerSession) NodeTypesCounter() (*big.Int, error) {
	return _NodeSystem.Contract.NodeTypesCounter(&_NodeSystem.CallOpts)
}

// AddNode is a paid mutator transaction binding the contract method 0xd425bd81.
//
// Solidity: function addNode(newNodeType string, id string, ip string, port string) returns(bool)
func (_NodeSystem *NodeSystemTransactor) AddNode(opts *bind.TransactOpts, newNodeType string, id string, ip string, port string) (*types.Transaction, error) {
	return _NodeSystem.contract.Transact(opts, "addNode", newNodeType, id, ip, port)
}

// AddNode is a paid mutator transaction binding the contract method 0xd425bd81.
//
// Solidity: function addNode(newNodeType string, id string, ip string, port string) returns(bool)
func (_NodeSystem *NodeSystemSession) AddNode(newNodeType string, id string, ip string, port string) (*types.Transaction, error) {
	return _NodeSystem.Contract.AddNode(&_NodeSystem.TransactOpts, newNodeType, id, ip, port)
}

// AddNode is a paid mutator transaction binding the contract method 0xd425bd81.
//
// Solidity: function addNode(newNodeType string, id string, ip string, port string) returns(bool)
func (_NodeSystem *NodeSystemTransactorSession) AddNode(newNodeType string, id string, ip string, port string) (*types.Transaction, error) {
	return _NodeSystem.Contract.AddNode(&_NodeSystem.TransactOpts, newNodeType, id, ip, port)
}

// AddNodeType is a paid mutator transaction binding the contract method 0x142685bf.
//
// Solidity: function addNodeType(name string, collateral uint256) returns()
func (_NodeSystem *NodeSystemTransactor) AddNodeType(opts *bind.TransactOpts, name string, collateral *big.Int) (*types.Transaction, error) {
	return _NodeSystem.contract.Transact(opts, "addNodeType", name, collateral)
}

// AddNodeType is a paid mutator transaction binding the contract method 0x142685bf.
//
// Solidity: function addNodeType(name string, collateral uint256) returns()
func (_NodeSystem *NodeSystemSession) AddNodeType(name string, collateral *big.Int) (*types.Transaction, error) {
	return _NodeSystem.Contract.AddNodeType(&_NodeSystem.TransactOpts, name, collateral)
}

// AddNodeType is a paid mutator transaction binding the contract method 0x142685bf.
//
// Solidity: function addNodeType(name string, collateral uint256) returns()
func (_NodeSystem *NodeSystemTransactorSession) AddNodeType(name string, collateral *big.Int) (*types.Transaction, error) {
	return _NodeSystem.Contract.AddNodeType(&_NodeSystem.TransactOpts, name, collateral)
}

// CheckExistence is a paid mutator transaction binding the contract method 0x02f2b5bd.
//
// Solidity: function checkExistence(nodeAddress address, id string, ip string) returns(bool)
func (_NodeSystem *NodeSystemTransactor) CheckExistence(opts *bind.TransactOpts, nodeAddress common.Address, id string, ip string) (*types.Transaction, error) {
	return _NodeSystem.contract.Transact(opts, "checkExistence", nodeAddress, id, ip)
}

// CheckExistence is a paid mutator transaction binding the contract method 0x02f2b5bd.
//
// Solidity: function checkExistence(nodeAddress address, id string, ip string) returns(bool)
func (_NodeSystem *NodeSystemSession) CheckExistence(nodeAddress common.Address, id string, ip string) (*types.Transaction, error) {
	return _NodeSystem.Contract.CheckExistence(&_NodeSystem.TransactOpts, nodeAddress, id, ip)
}

// CheckExistence is a paid mutator transaction binding the contract method 0x02f2b5bd.
//
// Solidity: function checkExistence(nodeAddress address, id string, ip string) returns(bool)
func (_NodeSystem *NodeSystemTransactorSession) CheckExistence(nodeAddress common.Address, id string, ip string) (*types.Transaction, error) {
	return _NodeSystem.Contract.CheckExistence(&_NodeSystem.TransactOpts, nodeAddress, id, ip)
}

// NodeCheckIn is a paid mutator transaction binding the contract method 0x06e5b546.
//
// Solidity: function nodeCheckIn(signatures bytes[], publicKey bytes, signatureBlockHash bytes) returns()
func (_NodeSystem *NodeSystemTransactor) NodeCheckIn(opts *bind.TransactOpts, signatures [][]byte, publicKey []byte, signatureBlockHash []byte) (*types.Transaction, error) {
	return _NodeSystem.contract.Transact(opts, "nodeCheckIn", signatures, publicKey, signatureBlockHash)
}

// NodeCheckIn is a paid mutator transaction binding the contract method 0x06e5b546.
//
// Solidity: function nodeCheckIn(signatures bytes[], publicKey bytes, signatureBlockHash bytes) returns()
func (_NodeSystem *NodeSystemSession) NodeCheckIn(signatures [][]byte, publicKey []byte, signatureBlockHash []byte) (*types.Transaction, error) {
	return _NodeSystem.Contract.NodeCheckIn(&_NodeSystem.TransactOpts, signatures, publicKey, signatureBlockHash)
}

// NodeCheckIn is a paid mutator transaction binding the contract method 0x06e5b546.
//
// Solidity: function nodeCheckIn(signatures bytes[], publicKey bytes, signatureBlockHash bytes) returns()
func (_NodeSystem *NodeSystemTransactorSession) NodeCheckIn(signatures [][]byte, publicKey []byte, signatureBlockHash []byte) (*types.Transaction, error) {
	return _NodeSystem.Contract.NodeCheckIn(&_NodeSystem.TransactOpts, signatures, publicKey, signatureBlockHash)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xd0d3f5ba.
//
// Solidity: function removeNode() returns(bool)
func (_NodeSystem *NodeSystemTransactor) RemoveNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeSystem.contract.Transact(opts, "removeNode")
}

// RemoveNode is a paid mutator transaction binding the contract method 0xd0d3f5ba.
//
// Solidity: function removeNode() returns(bool)
func (_NodeSystem *NodeSystemSession) RemoveNode() (*types.Transaction, error) {
	return _NodeSystem.Contract.RemoveNode(&_NodeSystem.TransactOpts)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xd0d3f5ba.
//
// Solidity: function removeNode() returns(bool)
func (_NodeSystem *NodeSystemTransactorSession) RemoveNode() (*types.Transaction, error) {
	return _NodeSystem.Contract.RemoveNode(&_NodeSystem.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(newOperator address) returns()
func (_NodeSystem *NodeSystemTransactor) SetOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _NodeSystem.contract.Transact(opts, "setOperator", newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(newOperator address) returns()
func (_NodeSystem *NodeSystemSession) SetOperator(newOperator common.Address) (*types.Transaction, error) {
	return _NodeSystem.Contract.SetOperator(&_NodeSystem.TransactOpts, newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(newOperator address) returns()
func (_NodeSystem *NodeSystemTransactorSession) SetOperator(newOperator common.Address) (*types.Transaction, error) {
	return _NodeSystem.Contract.SetOperator(&_NodeSystem.TransactOpts, newOperator)
}
