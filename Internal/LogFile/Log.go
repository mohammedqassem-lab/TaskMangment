package logfile

import (
	"fmt"
	"log/slog"
	"os"
)

func InitLogger() func() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	}

	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: slog.LevelError,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return func() {
		file.Close()
	}
}

func LogErr(err error) {
	if err == nil {
		return
	}
	fmt.Println("Error from logger:", err.Error())
	slog.Error(err.Error())
}
