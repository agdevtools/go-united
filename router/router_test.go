package router

import (
"net/http"
"testing"
"github.com/gorilla/mux"
"net/http/httptest"
"fmt"
)

var m *mux.Router

var err error

var req *http.Request

var respRec *httptest.ResponseRecorder

func setup() {
    //mux router with added question routes
    m = mux.NewRouter()
    Router(m)

    //The response recorder used to record HTTP responses
    respRec = httptest.NewRecorder()
}

func TestGet400(t *testing.T) {
    setup()
    fmt.Println("****  Test_router implement my intereface test")
    //Testing get of non existent question type
    req, err = http.NewRequest("GET", "/api/team", nil)
    if err != nil {
        t.Fatal("Creating 'GET /api/team' request failed!")
    }

    m.ServeHTTP(respRec, req)

    if respRec.Code != http.StatusBadRequest {
        t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
    }
}

//
// type Speak interface {
//         SayHello() string
// }
//
// type English struct{}
//
// func (e English) SayHello() string {
//         return "Hello"
// }
//
// func Test_English(t *testing.T) {
//             var voice Speak
//             voice = English{}
//             fmt.Println(voice.SayHello())
// }
//
// func  Test_Router(t *testing.T) {
//       fmt.Println("****  Test_Router implement my intereface test")
//       router := mux.NewRouter()
//       router.HandleFunc("/api/team", middleware.GetAllPlayers).Methods("GET", "OPTIONS")
// //       var mine DBService
// //       mine = DBServiceImp{}
// //       mine.getAllPlayers()
// }
//
// // type MiddleWareInt interface {
// //   GetAllPlayers(w http.ResponseWriter, r *http.Request)
// // }
// //
// // type MiddleWareIntImpl struct {}
//
//
// // func (m MiddleWareIntImpl) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
// //      fmt.Println("implement my intereface ")
// // }
// //
// // func  Test_GetAllPlayers(t *testing.T) {
// //       fmt.Println("implement my intereface test")
// //       var w http.ResponseWriter
// //       var r *http.Request
// //       var mine MiddleWareInt
// //       mine = MiddleWareIntImpl{}
// //       mine.GetAllPlayers(w,r)
// // }
//
