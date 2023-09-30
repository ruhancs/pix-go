package kafka

import (
	"context"
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ruhanc/pix-go/internal/application/dto"
	"github.com/ruhanc/pix-go/internal/application/factory"
	"github.com/ruhanc/pix-go/internal/domain/entity"
)

type KafkaProccessor struct {
	Database *sql.DB
	Producer *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

func NewKafkaProcessor(database *sql.DB, producer *ckafka.Producer, deliveryChan chan ckafka.Event) *KafkaProccessor {
	return &KafkaProccessor{
		Database: database,
		Producer: producer,
		DeliveryChan: deliveryChan,
	}
}

func (k *KafkaProccessor) Consume(ctx context.Context) {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":"localhost:9092",
		"group.id": "consumergroup",
		//ler sempre da mais nova msg no topico
		"auto.offset.reset":"earliest",
	}

	consumer,err := ckafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	//topicos que serao consumidos
	topics := []string{"test"}
	consumer.SubscribeTopics(topics, nil)

	fmt.Println("kafks consumer has been started")

	//manter lendo as msg
	for {
		msg,err := consumer.ReadMessage(-1)
		if err == nil {
			k.proccessMessage(ctx,msg)
		}
	}
}

func (k *KafkaProccessor) proccessMessage(ctx context.Context, msg *ckafka.Message) {
	transactionsTopic := "transactions"
	transactionConfirmationTopic := "transaction_confirmation"

	switch topic := *msg.TopicPartition.Topic; topic {
	//realizar as operacoes referentes a cada topico
	case transactionsTopic:
		k.proccessTransaction(ctx,msg)
	case transactionConfirmationTopic:
	default:
		fmt.Println("not valid topic", string(msg.Value))
	}
}

//processamento de msg no topico transactions
func (k *KafkaProccessor) proccessTransaction(ctx context.Context, msg *ckafka.Message) error {
	transaction := dto.NewTransaction()
	//transaforma o valor da msg de json em transaction para validar os campos da msg
	err := transaction.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	transactionUseCase := factory.TransactionUseCaseFactory(k.Database)

	createdTransaction,err := transactionUseCase.RegisterTransaction(
		ctx,
		transaction.AccountID,
		transaction.PixKeyTo,
		transaction.Description,
		transaction.Amount,
	)
	if err != nil {
		fmt.Println("error to register transaction: ", err)
		return err
	}

	//enviar para o banco de destino
	topic := "bank"+createdTransaction.PixKeyTo.Account.Bank.Code
	transaction.ID = createdTransaction.ID
	transaction.Status = entity.TransactionPending
	//converter a transacao para json para publicar no kafka
	transactionJson,err := transaction.ToJson()
	if err != nil {
		return err
	}

	//publicar a msg no kafka para enviar para o banco da conta de destino
	err = Publish(string(transactionJson), topic, k.Producer, k.DeliveryChan)
	if err != nil {
		return err
	}

	return nil
}