package game

import (
	"testing"
)

func TestCalculateHandScore(t *testing.T) {
	tests := []struct {
		name     string
		hand     []int
		expected int
	}{
		{
			name:     "Basit Seri",
			hand:     []int{21, 22, 23, 5, 6, 7},
			expected: 6, // 3 + 3
		},
		{
			name:     "Farklı Renklerde Grup",
			hand:     []int{9, 22, 35, 48}, // 9 (Kırmızı), 22 (Mavi), 35 (Yeşil)
			expected: 3,                    // 3'lü grup
		},
		{
			name:     "Okey ile Tamamlanmış Seri",
			hand:     []int{21, 23, OK}, // 21-OK-23 → 21-22-23
			expected: 3,
		},
		{
			name:     "Karışık Kombinasyonlar",
			hand:     []int{1, 2, 3, 14, 15, 16, 9, 22, 35}, // 3 seri (1-2-3, 14-15-16) + 1 grup (9-22-35)
			expected: 9,                                     // 3 + 3 + 3
		},
		{
			name:     "Maksimum Skor Sınırı",
			hand:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			expected: 12, // 4x3=12
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateHandScore(tt.hand)
			if got != tt.expected {
				t.Errorf("Test %s başarısız: Beklenen %d, Hesaplanan %d", tt.name, tt.expected, got)
			}
		})
	}
}

func TestCalculatePairScore(t *testing.T) {
	tests := []struct {
		name     string
		hand     []int
		expected int
	}{
		{
			name:     "Tek Çift",
			hand:     []int{5, 5, 10, 11},
			expected: 2,
		},
		{
			name:     "Çoklu Çiftler",
			hand:     []int{5, 5, 7, 7, 9, 9},
			expected: 6,
		},
		{
			name:     "Okey ile Çift",
			hand:     []int{5, OK, OK}, // 5-OK ve OK-OK
			expected: 4,
		},
		{
			name:     "Maksimum Çift Sınırı",
			hand:     []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7},
			expected: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculatePairScore(tt.hand)
			if got != tt.expected {
				t.Errorf("Test %s başarısız: Beklenen %d, Hesaplanan %d", tt.name, tt.expected, got)
			}
		})
	}
}

func TestEdgeCases(t *testing.T) {
	t.Run("Okey ile İki Farklı Kombinasyon", func(t *testing.T) {
		hand := []int{21, 23, OK, 5, 5} // Seri (21-OK-23) + Çift (5-5)
		expected := 5
		if got := CalculateHandScore(hand); got != expected {
			t.Errorf("Beklenen %d, Hesaplanan %d", expected, got)
		}
	})
}
