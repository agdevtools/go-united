package models

// User schema of the team table
type Team struct {
    PLAYER_ID      int64  `json:"player_id"`
    Player_name     string `json:"player_name"`
    Player_position string `json:"player_position"`
    Team_id         int64  `json: "team_id"`
}

type Teams struct {
    Team_id      int64  `json:"team_id"`
    Team_name     string `json:"team_name"`
}