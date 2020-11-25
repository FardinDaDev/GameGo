package logger

import (
	"log"
	"os"
	"sync"
)

type Logger struct {
	fileName string
	*log.Logger
}

var (
	logg *Logger
	once sync.Once
)

func NewLogger(fileName string) *Logger {
	once.Do(func() {
		logg = func(filename string) *Logger {
			file, _ := os.OpenFile(fileName, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0777)

			return &Logger{
				fileName: fileName,
				Logger: log.New(file, "SDL2 >> ", log.Lshortfile|log.Ltime),
			}
		}(fileName)
	})
	return logg
}

var GameDir = NewLogger("./game.log")