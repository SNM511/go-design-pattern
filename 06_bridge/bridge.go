package bridge

/*
	一个类存在多种变化时，
	例如如下的实现MsgSender，可以有Send2Email、Send2Phone等；实现MsgNotify，可以有Error级别、Warn级别等；
	使用桥接模式，抽象接口，各自实现后通过组合的形式来解耦；
	（这里可以使用工厂模式生成各自实现实例）
*/

// IMsgSender IMsgSender
type IMsgSender interface {
	Send(msg string) error
}

// EmailMsgSender 发送邮件
// 可能还有 电话、短信等各种实现
type EmailMsgSender struct {
	emails []string
}

// NewEmailMsgSender NewEmailMsgSender
func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

func (s *EmailMsgSender) Send(msg string) error {
	// 这里去发送消息
	for _, email := range s.emails {
		println("Send email to "+email+" with msg: ", msg)
	}
	return nil
}

type PhoneMsgSender struct {
	phoneNumbers []string
}

func NewPhoneMsgSender(numbers []string) *PhoneMsgSender {
	return &PhoneMsgSender{phoneNumbers: numbers}
}

func (s *PhoneMsgSender) Send(msg string) error {
	// 这里去发送消息
	for _, phoneNumber := range s.phoneNumbers {
		println("Send msg to phone "+phoneNumber+": ", msg)
	}
	return nil
}

// INotification 通知接口
type INotification interface {
	Notify(msg string) error
}

// ErrorNotification 错误通知
// 后面可能还有 warning 各种级别
type ErrorNotification struct {
	sender IMsgSender
}

func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send("Error: " + msg)
}

type WarnNotification struct {
	sender IMsgSender
}

func NewWarnNotification(sender IMsgSender) *WarnNotification {
	return &WarnNotification{sender: sender}
}

func (n *WarnNotification) Notify(msg string) error {
	return n.sender.Send("Warn: " + msg)
}
