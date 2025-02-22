package game

import (
	"log"
	"os"
)

var (
	Logger  *log.Logger
	logFile *os.File
)

/*
func init() {
	// Log dosyası için klasör oluştur
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal("Log dizini oluşturulamadı:", err)
	}

	// Timestamp ile log dosyası oluştur
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	logPath := filepath.Join(logDir, fmt.Sprintf("game_log_%s.txt", timestamp))

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Log dosyası açılamadı:", err)
	}

	logFile = file
	Logger = log.New(file, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	Logger.Println("=== Okey Oyunu Log Başlangıcı ===")
}*/

func CloseLogger() {
	if logFile != nil {
		Logger.Println("=== Okey Oyunu Log Sonu ===")
		logFile.Close()
	}
}

func LogInfo(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Printf(format, v...)
	}
}

func LogNewHand(playerNum int, hand []int) {
	LogInfo("\n=== Oyuncu %d için Yeni El Başlıyor ===\n", playerNum)
	LogInfo("El: %v", hand)
}

func LogScore(playerNum int, score int, details string) {
	LogInfo("\n=== Oyuncu %d Skor Hesaplaması ===\n", playerNum)
	LogInfo("Skor: %d", score)
	LogInfo("Detaylar: %s", details)
}

func LogCombinations(combinations [][]int, combinationType string) {
	LogInfo("\n=== Bulunan %s ===\n", combinationType)
	for i, combo := range combinations {
		LogInfo("%s %d: %v", combinationType, i+1, combo)
	}
}
