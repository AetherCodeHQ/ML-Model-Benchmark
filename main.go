package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	models := []struct {
		Name     string
		Accuracy float64
		Latency  float64
		Memory   float64
	}{
		{"logistic-regression", 0.82, 1.2, 128},
		{"random-forest", 0.89, 5.4, 512},
		{"neural-network", 0.91, 12.8, 2048},
		{"gradient-boosting", 0.90, 3.1, 768},
	}
	fmt.Println("ML Model Benchmark")
	fmt.Println("==================")
	fmt.Printf("  %-22s %-8s %-10s %-10s Score\n", "Model", "Acc", "Latency", "Memory")
	for _, m := range models {
		score := m.Accuracy * 100 / (1 + math.Log(m.Latency+1)) * (1024 / m.Memory)
		fmt.Printf("  %-22s %.1f%%  %.1fms   %.0fMB   %.2f\n", m.Name, m.Accuracy*100, m.Latency, m.Memory, score)
	}
}