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
	<H1>Hey There</H1>
	<p>I want to condess someting. this things has been eating me up from year and now i can't stop but tell you about it.
	I want to telle  you that I hate Soan Papdi,  it tastes like medicine. And momos,  they taste like </p>

	<p> I don't know why this is not working</p>
	`

	to := []string{"shubhammahar1306@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	assert.NoError(err)
}
