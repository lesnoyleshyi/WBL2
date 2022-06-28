/*
Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс
к сложной системе/фреймворку.
Применимость:
	- есть необходимость пользоваться лишь частью стороннего пакета/библиотеки.
Плюсы:
	- простое использование сложной системы без необходимости её детального изучения;
	- возможность использовать шаблоны, если предполагается, что
		конфигурация сложной системы не будет меняться.
Минусы:
	- если это не внешняя библиотека, то, возможно, сложность системы - это
		признак нарушения принципа единственной ответственности;
	- риск превращения фасада в объект бога;
	- невозможность тонкой конфигурации, доступной при использовании всего функционала.
*/

/*
Например, в компании принято пользоваться пакетом sarama, но его API избыточен
для большинства наших задач. Поэтому мы создаём свою библиотеку-оболочку и
пользуемся ей. Например, чтобы не было головной боли с синхронизацией различных
версий клиентов Kafka, мы можем "зашить" стандартный конфиг, в котором укажем
версию Kafka, оставив клиентам возможность изменять только список брокеров
и ID клиента.
*/

package main

import (
	"gopkg.in/Shopify/sarama.v1"
	"log"
	"time"
)

var brokers []string = []string{"host:port1", "host:port2"}

func main() {
	kafkaProducer, err := NewSyncProducer(brokers, "myServiceID")
	if err != nil {
		log.Fatalln("bad thing happened: ", err)
	}

	msg := sarama.ProducerMessage{
		Topic:     "some_topic",
		Key:       sarama.ByteEncoder("hello"),
		Value:     sarama.ByteEncoder("world"),
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: 1,
		Timestamp: time.Now(),
	}

	if _, _, err := kafkaProducer.SendMessage(&msg); err != nil {
		log.Println(err)
	}

	if err := kafkaProducer.Close(); err != nil {
		log.Fatalln(err)
	}
}

func NewSyncProducer(brokers []string, clientID string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()

	config.Version = sarama.V2_0_0_0
	config.Net.MaxOpenRequests = 100
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.ClientID = clientID

	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}
