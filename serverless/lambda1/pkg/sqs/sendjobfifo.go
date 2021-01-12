// Go Api server
// @jeffotoni
// 2019-06-01

package sqs

import (
	"log"
	"time"

	apicoreSqs "github.com/jeffotoni/organizando.seu.projeto/serverless/pkg/sqs"
)

type JobSqsFifo struct {
	Resource string    `json:"resource"`
	UserID   int64     `json:"user_id"`
	Topic    string    `json:"topic"`
	Received time.Time `json:"received"`
	Sent     time.Time `json:"sent"`
}

// Send Job FIFO
func SendJobFifo(jsonJob string) (bool, string) {

	// publica na fila o json FIFO
	ok, errsqs := apicoreSqs.SendFifo(apicoreSqs.Melinotificationfifo, jsonJob)
	if !ok {
		msgError := "Erro ao fazer sqs.Send in FIFO: " + errsqs
		log.Println(msgError)
		return false, msgError
	}

	return true, ""
}
