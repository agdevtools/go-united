package middleware

import (
    "database/sql"
    "encoding/json" // package to encode and decode the json into struct and vice versa
    "fmt"
    "go-united/models" // models package where User schema is defined
    "log"
    "net/http" // used to access the request and response object of the api
    "os"       // used to read the environment variable
    "strconv"  // package used to covert string into int type
    "github.com/gorilla/mux" // used to get the params from the route
    "github.com/joho/godotenv" // package used to read the .env file
    _ "github.com/lib/pq"      // postgres golang driver
)

// response format
type response struct {
    ID      int64  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}

type team_response struct {
    ID      int64  `json:"player_id,omitempty"`
    Message string `json:"message,omitempty"`
}

type GoService interface {
     GetAllPlayers() (http.ResponseWriter, *http.Request)
}

type GoServiceImp struct{}

type DBService interface {
     getAllPlayers() ([]models.Team, error)
}

type DBServiceImp struct{}

// GetAllUPlayers will return all the players from the team.
func  GetAllPlayers(w http.ResponseWriter, r *http.Request) {
    fmt.Println("****  In GetAllPlayers implement my intereface test ******* can you believe it??")
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    // get all the users in the db
    var s DBService
    s = DBServiceImp{}
    players, err := s.getAllPlayers()

    if err != nil {
        log.Fatalf("Unable to get all user. %v", err)
    }

    // send all the users as response
    json.NewEncoder(w).Encode(players)
}

// create connection with postgres db
func createConnection() *sql.DB {
    // load .env file
    err := godotenv.Load(".env")

    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Open the connection
    db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

    if err != nil {
        panic(err)
    }

    // check the connection
    err = db.Ping()

    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!")
    // return the connection
    return db
}

func TestTest() string {
 return "this is the real method"
}



// CreatePlayer create a player in the postgres db
func CreatePlayer(w http.ResponseWriter, r *http.Request) {
    // set the header to content type x-www-form-urlencoded
    // Allow all origin to handle cors issue
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // create an empty user of type models.User
    var team models.Team

    // decode the json request to user
    err := json.NewDecoder(r.Body).Decode(&team)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call insert user function and pass the user
    insertID := insertPlayer(team)

    // format a response object
    res := team_response{
        ID:      insertID,
        Message: "Player created successfully",
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}


// GetPlayer will return a single player by its id
func GetPlayer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    // get the userid from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id type from string to int
    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    // call the getPlayer function with player id to retrieve a single player
    player, err := getPlayer(int64(id))

    if err != nil {
        log.Fatalf("Unable to get player. %v", err)
    }

    // send the response
    json.NewEncoder(w).Encode(player)
}

// UpdatePlayer update player's detail in the postgres db
func UpdatePlayer(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // get the userid from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id type from string to int
    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    // create an empty user of type models.User
    var team models.Team

    // decode the json request to user
    err = json.NewDecoder(r.Body).Decode(&team)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call update user to update the user
    updatedRows := updatePlayer(int64(id), team)

    // format the message string
    msg := fmt.Sprintf("Player updated successfully. Total rows/record affected %v", updatedRows)

    // format the response message
    res := team_response{
        ID:      int64(id),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// DeletePlayer delete player's detail in the postgres db
func DeletePlayer(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // get the userid from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id in string to int
    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    // call the deleteUser, convert the int to int64
    deletedRows := deletePlayer(int64(id))

    // format the message string
    msg := fmt.Sprintf("User deleted successfully. Total rows/record affected %v", deletedRows)

    // format the reponse message
    res := team_response{
        ID:      int64(id),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

//------------------------- handler functions ----------------
// insert one player in the Team table in the DB
func insertPlayer(team models.Team) int64 {

    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the insert sql query
    // returning userid will return the id of the inserted user
    sqlStatement := `INSERT INTO team (player_id, player_name, player_position) VALUES ($1, $2, $3) RETURNING player_id`

    // the inserted id will store in this id
    var id int64

    // execute the sql statement
    // Scan function will save the insert id in the id
    err := db.QueryRow(sqlStatement, team.PLAYER_ID, team.Player_name, team.Player_position).Scan(&id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single player record %v", id)

    // return the inserted id
    return id
}

// get one player from the DB by its userid
func getPlayer(id int64) (models.Team, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create a user of models.User type
    var team models.Team

    // create the select sql query
    sqlStatement := `SELECT * FROM team WHERE player_id=$1`

    // execute the sql statement
    row := db.QueryRow(sqlStatement, id)

    // unmarshal the row object to user
    err := row.Scan(&team.PLAYER_ID, &team.Player_name, &team.Player_position)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return team, nil
    case nil:
        return team, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    // return empty user on error
    return team, err
}


// update player in the DB
func updatePlayer(id int64, team models.Team) int64 {

    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the update sql query
    sqlStatement := `UPDATE team SET player_name=$2, player_position=$3 WHERE player_id=$1`

    // execute the sql statement
    res, err := db.Exec(sqlStatement, id, team.Player_name, team.Player_position)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected, err := res.RowsAffected()

    if err != nil {
        log.Fatalf("Error while checking the affected rows. %v", err)
    }

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}


// delete player in the DB
func deletePlayer(id int64) int64 {

    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the delete sql query
    sqlStatement := `DELETE FROM team WHERE player_id=$1`

    // execute the sql statement
    res, err := db.Exec(sqlStatement, id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected, err := res.RowsAffected()

    if err != nil {
        log.Fatalf("Error while checking the affected rows. %v", err)
    }

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}

// get one user from the DB by its userid
func (s DBServiceImp) getAllPlayers() ([]models.Team, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    var teams []models.Team

    // create the select sql query
    sqlStatement := `SELECT * FROM team`

    // execute the sql statement
    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // close the statement
    defer rows.Close()

    // iterate over the rows
    for rows.Next() {
        var team models.Team

        // unmarshal the row object to user
        err = rows.Scan(&team.PLAYER_ID, &team.Player_name, &team.Player_position)

        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }

        // append the user in the users slice
        teams = append(teams, team)

    }

    // return empty user on error
    return teams, err
}