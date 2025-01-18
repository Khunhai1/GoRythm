package main

import (
	"encoding/json"
	"log"
	"os"
)

type Beat struct {
	Time    float64 `json:"time"`
	BeatNum int     `json:"beatNum"`
}

var beatmap []Beat
var modifiedBeatmap []Beat

// Load and process the beatmap to keep only every 4th beat
func loadBeatmap() {
	data, err := os.ReadFile("beatmap.json")
	if err != nil {
		log.Fatalf("failed to read beatmap file: %v", err)
	}
	err = json.Unmarshal(data, &beatmap)
	if err != nil {
		log.Fatalf("failed to unmarshal beatmap data: %v", err)
	}
	log.Printf("loaded %d beats", len(beatmap))
	log.Printf("first beat time: %f", beatmap[0].Time)

	// Process to select every 2nd beat
	for i, beat := range beatmap {
		if i%2 == 0 {
			modifiedBeatmap = append(modifiedBeatmap, beat)
		}
	}

	log.Printf("modified beatmap with every 4th beat contains %d beats", len(modifiedBeatmap))
	log.Printf("first modified beat time: %f", modifiedBeatmap[0].Time)
	log.Printf("second modified beat time: %f", modifiedBeatmap[1].Time)
	log.Printf("third modified beat time: %f", modifiedBeatmap[2].Time)
}
