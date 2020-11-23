package fsm

import "fmt"

type EventStatus int32

const (
	MessageEvent EventStatus = 1
	ButtonEvent EventStatus = 2
)

type Event struct {
	Massage string
	Status EventStatus
}

type Desc struct {
	Event      Event
	DescNodeId int32
}

type NodeInfo struct {
	NodeId int32
	NodeName string
	IsStop bool
}

type GraphNode struct {
	NodeInfo NodeInfo
	desc map[Event]*GraphNode
}

type FSM struct {
	nodes []*GraphNode
	head *GraphNode
}

func NewFSM(NodeLen int) *FSM {
	fsm := &FSM{}
	for i := 0; i < NodeLen; i++ {
		node := &GraphNode{
			NodeInfo: NodeInfo{},
			desc:     make(map[Event]*GraphNode),
		}
		fsm.nodes = append(fsm.nodes, node)
	}
	fsm.head = fsm.nodes[0]
	return fsm
}

func (f *FSM) SetGraphNode(nodeInfo NodeInfo, descs... Desc) error {
	f.nodes[nodeInfo.NodeId].NodeInfo = nodeInfo
	for _, desc := range descs {
		f.nodes[nodeInfo.NodeId].desc[desc.Event] = f.nodes[desc.DescNodeId]
	}
	return nil
}

func (f *FSM) GetNextStatus(nodeId int32, event Event) (int32, error) {
	nextNode := f.nodes[nodeId].desc[event]
	return nextNode.NodeInfo.NodeId, nil
}

func (f *FSM) GetStartStatusId() int32 {
	return f.head.NodeInfo.NodeId
}

func (f *FSM) GetStatusButtonEvents(statusId int32) ([]Event, error) {
	var res []Event
	for event, _ := range f.nodes[statusId].desc {
		if event.Status == ButtonEvent {
			res = append(res, event)
		}
	}
	return res, nil
}

func (f *FSM) GetStatusMessageEvent(statusId int32) (Event, bool, error){
	for event, _ := range f.nodes[statusId].desc {
		if event.Status == MessageEvent {
			return event, true, nil
		}
	}
	return Event{}, false, fmt.Errorf("not find message event")
}


func (f *FSM) IsStopNode(statusId int32) (bool, error) {
	return f.nodes[statusId].NodeInfo.IsStop, nil
}