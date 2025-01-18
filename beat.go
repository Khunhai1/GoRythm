package main

import (
	_ "embed"
	"encoding/json"
)

type Beat struct {
	Time    float64 `json:"time"`
	BeatNum int     `json:"beatNum"`
}

var (
	//go:embed assets/beatmap/beatmap.json
	beatmapData []byte
)

// Load the beatmap and select only every 2nd beat
func loadBeatmap() ([]Beat, error) {
	var beatmap []Beat
	err := json.Unmarshal(beatmapData, &beatmap)
	if err != nil {
		return nil, err
	}

	// Select every 2nd beat
	var filteredBeatmap []Beat
	for i, beat := range beatmap {
		if i%2 == 0 {
			filteredBeatmap = append(filteredBeatmap, beat)
		}
	}
	return filteredBeatmap, nil
}
