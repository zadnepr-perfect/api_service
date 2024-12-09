package logging

import (
	"log"
	"os"
)

// NewLogger создает новый экземпляр логгера
func NewLogger() *log.Logger {
	return log.New(os.Stdout, "INFO: ", log.LstdFlags)
}
