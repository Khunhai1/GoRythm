// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

package audio

import (
	_ "embed"
	"encoding/json"
)

// A Beat struct contains the time and beat number of a beat
type Beat struct {
	Time    float64 `json:"time"`    // The time of the beat
	BeatNum int     `json:"beatNum"` // The beat number
}

//go:embed assets/beatmap/beatmap.json
var beatmapData []byte

// Load the beatmap and select only every 2nd beat
func LoadBeatmap() ([]Beat, error) {
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
