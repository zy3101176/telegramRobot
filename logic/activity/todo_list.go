package activity

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"telegram_robot/util/fsm"
)

var (
	StartNode = fsm.NodeInfo{NodeId: 0, NodeName: "start node", IsStop: false}
	GetMessageNode = fsm.NodeInfo{NodeId: 1, NodeName: "get message node", IsStop: false}
	DoubleCheckNode = fsm.NodeInfo{NodeId: 2, NodeName: "double check node", IsStop: false}
	SuccessNode = fsm.NodeInfo{NodeId: 3, NodeName: "success node", IsStop: true}
	FailNode = fsm.NodeInfo{NodeId: 4, NodeName: "fail node", IsStop: true}
)

var (
	SetMessage = fsm.Event{Massage: "请输入内容", Status: fsm.MessageEvent}
	Confirm = fsm.Event{Massage: "确认", Status: fsm.ButtonEvent}
	Cancel = fsm.Event{Massage: "取消", Status: fsm.ButtonEvent}
)

type TodoListActivity struct {
	activityFsm *fsm.FSM
}

func (t *TodoListActivity) Init() error {
	activityFsm := fsm.NewFSM(5)

	activityFsm.SetGraphNode(StartNode,
		fsm.Desc{Event: SetMessage, DescNodeId: 1},
		fsm.Desc{Event: Cancel, DescNodeId: 4})

	activityFsm.SetGraphNode(GetMessageNode,
		fsm.Desc{Event: Confirm, DescNodeId: 3},
		fsm.Desc{Event: Cancel, DescNodeId: 4})
	activityFsm.SetGraphNode(SuccessNode)
	activityFsm.SetGraphNode(FailNode)
	t.activityFsm = activityFsm
	return nil
}

func (t *TodoListActivity) DataProcess(statusId int32, update *tgbotapi.Update) (int32, *tgbotapi.MessageConfig, error) {
	fun := "TodoListActivity.DataProcess -->"
	var messageChatId int64
	var nextStatusId int32
	if update.CallbackQuery != nil {
		buttonEvents, err := t.activityFsm.GetStatusButtonEvents(statusId)
		if err != nil {
			return 0, nil, errors.Wrap(err, fun)
		}
		messageChatId = update.CallbackQuery.Message.Chat.ID
		for _, event := range buttonEvents {
			if event.Massage == update.CallbackQuery.Data {
				//event.CallBackFunc(messageChatId, update.Message.Text)
				nextStatusId, err = t.activityFsm.GetNextStatus(statusId, event)
				if err != nil {
					return 0, nil, errors.Wrap(err, fun)
				}
				break
			}
		}
	} else if update.Message != nil {
		event, ok, err := t.activityFsm.GetStatusMessageEvent(statusId)
		if err != nil {
			return 0, nil, errors.Wrap(err, fun)
		}
		if !ok {
			return 0, nil, errors.New("not find message event")
		}
		messageChatId = update.Message.Chat.ID
		//event.CallBackFunc(messageChatId, update.Message.Text)
		nextStatusId, err = t.activityFsm.GetNextStatus(statusId, event)
	}
	message, err := t.GetMessageResp(messageChatId, nextStatusId)
	if err != nil {
		return 0, nil, errors.Wrap(err, fun)
	}
	return nextStatusId, message, nil
}

func (t *TodoListActivity) GetMessageResp(chatId int64, statusId int32) (*tgbotapi.MessageConfig, error)  {
	fun := "TodoListActivity.GetMessageResp -->"
	buttonEvents, err := t.activityFsm.GetStatusButtonEvents(statusId)
	if err != nil {
		return nil, errors.Wrap(err, fun)
	}
	messageEvent, ok, err := t.activityFsm.GetStatusMessageEvent(statusId)
	if err != nil {
		return nil, errors.Wrap(err, fun)
	}
	var msg string
	if ok {
		msg = messageEvent.Massage
	} else {
		msg = ""
	}
	resMessage := tgbotapi.NewMessage(chatId, msg)
	var keyboardMarkup tgbotapi.InlineKeyboardMarkup
	var button []tgbotapi.InlineKeyboardButton
	for _, event := range buttonEvents {
		item := tgbotapi.NewInlineKeyboardButtonData(event.Massage,event.Massage)
		button = append(button, item)
	}
	if len(button) != 0 {
		row := tgbotapi.NewInlineKeyboardRow(button...)
		keyboardMarkup = tgbotapi.NewInlineKeyboardMarkup(row)
		resMessage.ReplyMarkup = keyboardMarkup
	}
	return &resMessage, nil
}

func (t *TodoListActivity) IsFinish(status int32) (bool, error) {
	return t.activityFsm.IsStopNode(status)
}