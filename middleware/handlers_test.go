package middleware
import (
"testing"
 "fmt"
//"github.com/gorilla/mux"
)

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
// type Testy interface {
//      TestTest() string
// }
//
// type TestyImp struct{}
//
// func (e TestyImp) TestTest() string {
//         return "Hello TesttEST FAKER"
// }
//
// type DBService interface {
//      getAllPlayers() ([]models.Team, error)
// }
//
// type DBServiceImp struct{}
//
// type GoService interface {
//      GetAllPlayers() (http.ResponseWriter, *http.Request)
// }
//
// type GoServiceImp struct{}
//
// //**** implement a consumer that uses the testy interface ****
// func someConsumer(g Testy) {
//   fmt.Println("implement my consumer ")
//   fmt.Println(g.TestTest())
//   fmt.Println("FINISH my consumer ")
// }
//
// func Test_English(t *testing.T) {
//             var voice Speak
//             voice = English{}
//             fmt.Println(voice.SayHello())
// }
//
// type MiddleWareInt interface {
//   getAllPlayers(w http.ResponseWriter, r *http.Request)
// }
//
// type MiddleWareIntImpl struct {}

// type GoService interface {
//      getAllPlayers() ([]models.Team, error)
// }
//
// type GoServiceImp struct{}
//
// func (s GoServiceImp) getAllPlayers() ([]models.Team, error) {
//         var teams []models.Team
//         var err error
//         return teams, err
// }

// func  Test_GetAllPlayers_int(t *testing.T) {
//       fmt.Println("********implement my go service intereface test")
//       var w http.ResponseWriter
//               var r *http.Request
//               var mine GoService
//               mine = GoServiceImp{}
//               mine.GetAllPlayers()
// }
//
//
// //This is an interesting one router calls the interface passed to function
// func (m GoServiceImp) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
//      router.HandleFunc("/api/team", m.GetAllPlayers).Methods("GET", "OPTIONS")
// //         var w http.ResponseWriter
// //         var r *http.Request
//         var mine GoService
//         mine = GoServiceImp{}
//         mine.GetAllPlayers(w,r)
//      fmt.Println("implement my intereface 22 ")
// }

// func (m GoServiceImp) GetAllPlayers() {
//      router.HandleFunc("/api/team", m.GetAllPlayers).Methods("GET", "OPTIONS")
//      fmt.Println("implement my intereface 66666666 ")
// }

func  Test_GetAllPlayers(t *testing.T) {
      fmt.Println("****  Test_GetAllPlayers implement my intereface test")
//       router := mux.NewRouter()
//       router.HandleFunc("/api/team", GetAllPlayers).Methods("GET", "OPTIONS")
//       var mine DBService
//       mine = DBServiceImp{}
//       mine.getAllPlayers()
}

// func  Test_TestTest(t *testing.T) {
//   fmt.Println("testing Test Test123")
//     fmt.Println(TestTest())
// }
//
// func  Test_TestTest2(t *testing.T) {
//   fmt.Println("testing Test Test fake")
//   var testy2 Testy
//   testy2 = TestyImp{}
//     fmt.Println(testy2.TestTest())
//     someConsumer(testy2)
// }


