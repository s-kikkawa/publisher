package message

import (
    "log"
    "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
    "github.com/aws/aws-sdk-go/aws/session"
    "encoding/json"
)

const (
    DO_NOTHING = false // メッセージ送信しない場合はtrueにしてください
    AWS_REGION = "ap-northeast-1"
    QUEUE_URL  = "https://sqs.ap-northeast-1.amazonaws.com/123456789012/TestQueue"
)

type Message struct {
    OperationType string `json:"operationType"`
    ID string `json:"id"`
    ItemCode string `json:"itemCode"`
    Text   string   `json:"text"`
}

// SQSにメッセージを送信します
func SendMessage(operationType string, idStr string, itemCode string, text string){
    if DO_NOTHING{
        log.Print("DO_NOTHING = true のためメッセージは送信しません")
        msgBody := createMessage(operationType, idStr, itemCode, text)
        log.Print(msgBody)
        return
    }
    sess := session.Must(session.NewSession())
    svc := sqs.New(sess, aws.NewConfig().WithRegion(AWS_REGION))
    // 送信内容を作成
    msgBody := createMessage(operationType, idStr, itemCode, text)
    params := &sqs.SendMessageInput{
        MessageBody:  aws.String(msgBody),
        QueueUrl:     aws.String(QUEUE_URL),
        DelaySeconds: aws.Int64(1),
    }
    sqsRes, err := svc.SendMessage(params)
    if err != nil {
        log.Fatal(err)
    }
    log.Print("SQSMessageID", *sqsRes.MessageId)
    log.Print("SQSMessageBody", msgBody)
}

// json形式でメッセージ本体を作成します
func createMessage(operationType string, idStr string, itemCode string, text string) string {
    message := new(Message)
    message.OperationType = operationType
    message.ID = idStr
    message.ItemCode = itemCode
    message.Text = text
    message_json, _ := json.Marshal(message)
    return string(message_json)
}
