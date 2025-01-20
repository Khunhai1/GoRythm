package audio

import (
	"testing"
)

// Mock beatmap JSON data
var mockBeatmapData = []byte(`[
    {"time": 0.5, "beatNum": 1},
    {"time": 1.0, "beatNum": 2},
    {"time": 1.5, "beatNum": 3},
    {"time": 2.0, "beatNum": 4},
    {"time": 2.5, "beatNum": 5}
]`)

// Mock the embedded beatmap data
func init() {
	beatmapData = mockBeatmapData
}

// TestLoadBeatmap tests the loadBeatmap function
func TestLoadBeatmap(t *testing.T) {
	// Expected beats after filtering
	expected := []Beat{
		{Time: 0.5, BeatNum: 1},
		{Time: 1.5, BeatNum: 3},
		{Time: 2.5, BeatNum: 5},
	}

	// Load the beatmap
	beats, err := LoadBeatmap()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the number of beats matches the expected number
	if len(beats) != len(expected) {
		t.Fatalf("Expected %d beats, got %d", len(expected), len(beats))
	}

	// Compare each beat with the expected beat
	for i, beat := range beats {
		if beat != expected[i] {
			t.Errorf("Expected beat %v, got %v", expected[i], beat)
		}
	}
}
