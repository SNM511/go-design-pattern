package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotification_Notify(t *testing.T) {
	assert.Nil(t, NewErrorNotification(NewEmailMsgSender([]string{"test@test.com"})).Notify("panic happened!"))
	assert.Nil(t, NewWarnNotification(NewPhoneMsgSender([]string{"123456"})).Notify("something went wrong."))
}
