package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:           PortID,
		PostList:         []Post{},
		SendPostList:     []SendPost{},
		TimedoutPostList: []TimedoutPost{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated ID in post
	postIdMap := make(map[uint64]bool)
	postCount := gs.GetPostCount()
	for _, elem := range gs.PostList {
		if _, ok := postIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for post")
		}
		if elem.Id >= postCount {
			return fmt.Errorf("post id should be lower or equal than the last id")
		}
		postIdMap[elem.Id] = true
	}
	// Check for duplicated ID in sendPost
	sendPostIdMap := make(map[uint64]bool)
	sendPostCount := gs.GetSendPostCount()
	for _, elem := range gs.SendPostList {
		if _, ok := sendPostIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for sendPost")
		}
		if elem.Id >= sendPostCount {
			return fmt.Errorf("sendPost id should be lower or equal than the last id")
		}
		sendPostIdMap[elem.Id] = true
	}
	// Check for duplicated ID in timedoutPost
	timedoutPostIdMap := make(map[uint64]bool)
	timedoutPostCount := gs.GetTimedoutPostCount()
	for _, elem := range gs.TimedoutPostList {
		if _, ok := timedoutPostIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for timedoutPost")
		}
		if elem.Id >= timedoutPostCount {
			return fmt.Errorf("timedoutPost id should be lower or equal than the last id")
		}
		timedoutPostIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
