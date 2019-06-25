// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package nodeprotocol

import (
        "errors"
        "encoding/json"
        "io/ioutil"

        "github.com/ethereum/go-ethereum/log"
        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/node"
        "github.com/ethereum/go-ethereum/params"
)

var StateAccessClient map[string]*NodeDataClient
var SyncStartingBlock uint64

type NodeData struct {
        Hash   common.Hash  `json:"hash"`
        Id     string       `json:"id"`
        Number uint64       `json:"number"`
        Ip     string       `json:"ip"`
}

type NodeDatabaseHandlers struct {
        Client *NodeDataClient
}

// client for submitting jobs and providing a repository-like interface
type NodeDataClient struct {
         Updates chan Update
}

type Update interface {
        ExitChan() chan error
        Run(nodeDatabase map[uint64]NodeData) (map[uint64]NodeData, error)
}

// Read all node protocol data from the database
type ReadNodeData struct {
    nodeDatabase  chan map[uint64]NodeData
    exitChan chan error
}

func InitiateNodeProtocolStateAccess(nodeType string) {
        // Make sure client mapping is initiated
        if len(StateAccessClient) == 0 {
                StateAccessClient = make(map[string]*NodeDataClient)
        }

        db := node.DefaultDataDir() + "/geth/chaindata/" + nodeType

        // initialize empty-object json file if not found
	if _, err := ioutil.ReadFile(db); err != nil {
		str := "{}"
		if err = ioutil.WriteFile(db, []byte(str), 0644); err != nil {
			log.Error("Error Creating Node Protocol State DB")
		}
	}

	// create channel to communicate over
	updates := make(chan Update)

	// start watching updates channel new updates
	go ProcessUpdates(updates, db)

	// create dependencies
	client := &NodeDataClient{Updates: updates}

        // Save client to mapping for future use
        StateAccessClient[nodeType] = client

        // Initiate sync from highest known protocol data entry
        SyncStartingBlock = uint64(0)
        for _, nodeType := range params.NodeTypes {
                number, _, err := GetNodeDataStateLatest(nodeType.Name)
                if err == nil && number > SyncStartingBlock {
                        SyncStartingBlock = number
                }
        }

}

func SaveNodeDataStateGroup(nodeType string, data map[uint64]NodeData) {
        if client, ok := StateAccessClient[nodeType]; ok {
                client.SaveNodeDataGroup(data)
        }
}
func SaveNodeDataState(nodeType string, hash common.Hash, id string, blockNumber uint64, ip string) {
        if client, ok := StateAccessClient[nodeType]; ok {
                client.SaveNodeData(NodeData{Id: id, Hash: hash, Number: blockNumber, Ip: ip}, blockNumber)
        }
}
func DeleteNodeDataState(nodeType string, hash common.Hash, blockNumber uint64) {
        if client, ok := StateAccessClient[nodeType]; ok {
                client.DeleteNodeData(blockNumber)
        }
}
func GetNodeDataState(nodeType string, blockNumber uint64) (NodeData, error) {
        if client, ok := StateAccessClient[nodeType]; ok {
                return client.GetNodeData(blockNumber)
        }
        return NodeData{}, errors.New("Node Protocol Client Not Ready")
}
func GetFullNodeDataState(nodeType string) (map[uint64]NodeData, error) {
        if client, ok := StateAccessClient[nodeType]; ok {
                return client.GetNodeDatabase()
        }
        return nil, errors.New("Node Protocol Client Not Ready")
}
func GetNodeDataStateLatest(nodeType string) (uint64, NodeData, error) {
        if client, ok := StateAccessClient[nodeType]; ok {
                number, data, err := client.GetLatestNodeData()
                log.Info("Protocol Head Data State", "Type", nodeType, "Number", number, "Hash", data.Hash, "ID", data.Id, "IP", data.Ip)
                return number, data, err
         }
         return 0, NodeData{}, errors.New("Node Protocol Client Not Ready")
}
func ReplaceNodeDataState(nodeType string, hash common.Hash, id string, blockNumber uint64, ip string) {
        _, err := GetNodeDataState(nodeType, blockNumber)
        if err == nil {
                DeleteNodeDataState(nodeType, hash, blockNumber)
        }
        SaveNodeDataState(nodeType, hash, id, blockNumber, ip)
}

// GetNodeData returns specific id based on hash
func (c *NodeDataClient) GetNodeData(blockNumber uint64) (NodeData, error) {
	nodeDatabase, err := c.getNodeDataHash()
	if err != nil {
		return NodeData{}, err
	}
        if data, ok := nodeDatabase[blockNumber]; ok {
	        return data, nil
        }
        return NodeData{}, errors.New("Node Protocol Entry Does Not Exist")
}

// DeleteNodeData deletes specific id based on hash
func (c *NodeDataClient) DeleteNodeData(blockNumber uint64) error {
	update := NewDeleteNodeDataUpdate(blockNumber)
	c.Updates <- update

	if err := <-update.ExitChan(); err != nil {
		return err
	}
	return nil
}

// GetNodeData returns all node data as map
func (c *NodeDataClient) GetNodeDatabase() (map[uint64]NodeData, error) {
        var nodeDatabase map[uint64]NodeData

	nodeDatabase, err := c.getNodeDataHash()
	if err != nil {
		return nodeDatabase, err
	}

	return nodeDatabase, nil
}

// GetLatestNodeData returns latest NodeData
func (c *NodeDataClient) GetLatestNodeData() (uint64, NodeData, error) {
        maxKey := uint64(0)
        var maxData NodeData

	nodeDatabase, err := c.getNodeDataHash()
	if err != nil {
		return maxKey, maxData, err
	}

	for key, value := range nodeDatabase {
                if key > maxKey {
                        maxKey = key
                        maxData = value
                }
	}
	return maxKey, maxData, nil
}

// SaveNodeDataGroup saves a group of NodeData
func (c *NodeDataClient) SaveNodeDataGroup(data map[uint64]NodeData) (map[uint64]NodeData, error) {
	update := NewSaveNodeDataGroupUpdate(data)

	c.Updates <- update

	if err := <-update.ExitChan(); err != nil {
		return data, err
	}
	return <-update.saved, nil
}

// SaveNodeData saves a new NodeData entry
func (c *NodeDataClient) SaveNodeData(data NodeData, blockNumber uint64) (NodeData, error) {
	update := NewSaveNodeDataUpdate(data, blockNumber)

	c.Updates <- update

	if err := <-update.ExitChan(); err != nil {
		return NodeData{}, err
	}
	return <-update.saved, nil
}

// Get all NodeData as a map
func (h *NodeDatabaseHandlers) GetNodeDatabase() {
        nodeDatabase, err := h.Client.GetNodeDatabase()
        if err != nil {
                log.Error("Error Accessing Node Protocol State Data", "Error", err)
                return
        }
        log.Trace("Node Database Accessed", "Data", nodeDatabase)
}

func ProcessUpdates(updates chan Update, db string) {
    for {
        j := <-updates

        // Read the database
        nodeDatabase := make(map[uint64]NodeData, 0)
        content, err := ioutil.ReadFile(db)
        if err == nil {
            if err = json.Unmarshal(content, &nodeDatabase); err == nil {

                // Run the job
                nodeDatabaseMod, err := j.Run(nodeDatabase)

                // If there were modifications, write them back to the database
                if err == nil && nodeDatabaseMod != nil {
                    b, err := json.Marshal(nodeDatabaseMod)
                    if err == nil {
                        err = ioutil.WriteFile(db, b, 0644)
                    }
                }
            }
        }

        j.ExitChan() <- err
    }
}

func NewReadNodeData() *ReadNodeData {
    return &ReadNodeData{
        nodeDatabase: make(chan map[uint64]NodeData, 1),
        exitChan:     make(chan error, 1),
    }
}

func (j ReadNodeData) ExitChan() chan error {
    return j.exitChan
}

func (j ReadNodeData) Run(nodeDatabase map[uint64]NodeData) (map[uint64]NodeData, error) {
    j.nodeDatabase <- nodeDatabase

    return nil, nil
}

func (c *NodeDataClient) getNodeDataHash() (map[uint64]NodeData, error) {
	update := NewReadNodeData()
	c.Updates <- update

	if err := <-update.ExitChan(); err != nil {
		return make(map[uint64]NodeData, 0), err
	}
	return <-update.nodeDatabase, nil
}

// Save NodeData Group to state
type SaveNodeDataGroupUpdate struct {
	toSave   map[uint64]NodeData
	saved    chan map[uint64]NodeData
	exitChan chan error
}
func NewSaveNodeDataGroupUpdate(data map[uint64]NodeData) *SaveNodeDataGroupUpdate {
	return &SaveNodeDataGroupUpdate{
		toSave:      data,
		saved:    make(chan map[uint64]NodeData, 1),
		exitChan: make(chan error, 1),
	}
}
func (j SaveNodeDataGroupUpdate) ExitChan() chan error {
	return j.exitChan
}
func (j SaveNodeDataGroupUpdate) Run(nodeDatabase map[uint64]NodeData) (map[uint64]NodeData, error) {
	var data map[uint64]NodeData
        data = j.toSave
        for key, val := range data {
	        nodeDatabase[key] = val
        }
	j.saved <- data
	return nodeDatabase, nil
}

// Save NodeData to state
type SaveNodeDataUpdate struct {
	toSave   NodeData
	saved    chan NodeData
	exitChan chan error
}
func NewSaveNodeDataUpdate(data NodeData, blockNumber uint64) *SaveNodeDataUpdate {
	return &SaveNodeDataUpdate{
		toSave:      data,
		saved:    make(chan NodeData, 1),
		exitChan: make(chan error, 1),
	}
}
func (j SaveNodeDataUpdate) ExitChan() chan error {
	return j.exitChan
}
func (j SaveNodeDataUpdate) Run(nodeDatabase map[uint64]NodeData) (map[uint64]NodeData, error) {
	var data NodeData
        data = j.toSave
	nodeDatabase[data.Number] = data

	j.saved <- data
	return nodeDatabase, nil
}

// Read NodeData from state
type ReadNodeDatabaseUpdate struct {
	nodeDatabase chan map[uint64]NodeData
	exitChan     chan error
}

func NewReadNodeDatabaseUpdate() *ReadNodeDatabaseUpdate {
	return &ReadNodeDatabaseUpdate{
		nodeDatabase: make(chan map[uint64]NodeData, 1),
		exitChan:     make(chan error, 1),
	}
}
func (j ReadNodeDatabaseUpdate) ExitChan() chan error {
	return j.exitChan
}
func (j ReadNodeDatabaseUpdate) Run(nodeDatabase map[uint64]NodeData) (map[uint64]NodeData, error) {
	j.nodeDatabase <- nodeDatabase

	return nil, nil
}

// Delete NodeData from state
type DeleteNodeDataUpdate struct {
	toDelete uint64
	exitChan chan error
}
func NewDeleteNodeDataUpdate(blockNumber uint64) *DeleteNodeDataUpdate {
	return &DeleteNodeDataUpdate{
		toDelete: blockNumber,
		exitChan: make(chan error, 1),
	}
}
func (j DeleteNodeDataUpdate) ExitChan() chan error {
	return j.exitChan
}
func (j DeleteNodeDataUpdate) Run(nodeDatabase map[uint64]NodeData) (map[uint64]NodeData, error) {
	delete(nodeDatabase, j.toDelete)
	return nodeDatabase, nil
}
