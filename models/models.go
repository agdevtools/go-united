package models
// User schema of the user table
type User struct {
    ID       int64  `json:"id"`
    Name     string `json:"name"`
    Location string `json:"location"`
    Age      int64  `json:"age"`
}

// User schema of the team table
type Team struct {
    PLAYER_ID      int64  `json:"player_id"`
    player_name     string `json:"player_name"`
    player_position string `json:"player_position"`
}