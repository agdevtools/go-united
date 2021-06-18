package models

// User schema of the team table
type Team struct {
    PLAYER_ID      int64  `json:"player_id"`
    Player_name     string `json:"player_name"`
    Player_position string `json:"player_position"`
}