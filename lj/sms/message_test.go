package sms_test

import (
	"testing"

	"github.com/qiniu/api.v7/sms"
	"encoding/json"
	"fmt"
)

func TestMessage(t *testing.T) {

	// SendMessage
	args := sms.MessagesRequest{
		SignatureID: "1131017766455742464",
		TemplateID:  "1131211147840589824",
		Mobiles:     []string{"18801732070"},
		Parameters: map[string]interface{}{
			"code": 123456,
		},
	}

	msg, err := json.Marshal(args)
	fmt.Println("encoded data : ")
	         fmt.Println(msg)
	         fmt.Println(string(msg))

	ret, err := manager.SendMessage(args)

	if err != nil {
		t.Fatalf("SendMessage() error: %v\n", err)
	}

	if len(ret.JobID) == 0 {
		t.Fatal("SendMessage() error: The job id cannot be empty")
	}
}
