package util

import (
	"fmt"
	"testing"
)

var (
	StartNode = NodeInfo{NodeId: 0, NodeName: "start node", IsStop: false}
	GetMessageNode = NodeInfo{NodeId: 1, NodeName: "get message node", IsStop: false}
	DoubleCheckNode = NodeInfo{NodeId: 2, NodeName: "double check node", IsStop: false}
	SuccessNode = NodeInfo{NodeId: 3, NodeName: "success node", IsStop: true}
	FailNode = NodeInfo{NodeId: 4, NodeName: "fail node", IsStop: true}
)

var (
	SetMessage Event = Event{Massage: "请输入内容", Status: MessageEvent}
	Confirm Event = Event{Massage: "确认", Status: ButtonEvent}
	Cancel Event = Event{Massage: "取消", Status: ButtonEvent}
)

func TestFSM_SetGraphNode(t *testing.T) {
	fsm := NewFSM(5)

	fsm.SetGraphNode(StartNode,
		Desc{Event: SetMessage, DescNodeId: 1},
		Desc{Event: Cancel, DescNodeId: 4})

	fsm.SetGraphNode(GetMessageNode,
		Desc{Event: Confirm, DescNodeId: 3},
		Desc{Event: Cancel, DescNodeId: 4})
	fsm.SetGraphNode(SuccessNode)
	fsm.SetGraphNode(FailNode)

	statusId := fsm.GetStartStatusId()

	for {
		if fsm.IsStopNode(statusId) {
			return
		}
		events := fsm.GetStatusEvent(statusId)
		var id int32
		for _, event := range events {
			fmt.Println(event.Status, event.Massage)
		}
		fmt.Scanf("%d", id)
		statusId = id
	}
}