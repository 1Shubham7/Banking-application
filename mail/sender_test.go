package mail

import (
	"testing"

	"github.com/1shubham7/bank/util"
	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	// to skip the test
	if testing.Short(){
		t.Skip()
	}

	assert := assert.New(t)
	config, err := util.LoadConfig("..")
	assert.NoError(err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "I hate Soan Papdi and Momos"

	content := `
	<h1>Hey There</h1>
	<p>I want to confess something. This has been eating me up for years, and now I can't help but tell you about it.
	I want to tell you that I hate Soan Papdi — it tastes like medicine. And momos... I hate the taste of momos.</p>
	`

	to := []string{"shubhammahar1306@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	assert.NoError(err)
}
