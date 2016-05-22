package main

import (
	"log"
	"math/rand"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var conn *websocket.Conn

func handler(w http.ResponseWriter, r *http.Request) {
	var err error
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	go func() error {
		for {
			messageType, _, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
			}
			if err = conn.WriteMessage(messageType, []byte(getRandomName())); err != nil {
				log.Println("write message failed. ", err)
				return err
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Wait()

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9998", nil)
}

func getRandomName() string {
	n := len(peopleNames)
	k := rand.Intn(n)
	return peopleNames[k]
}

var peopleNames = []string{
	"Michael",
	"Jennifer",
	"Christopher",
	"Amanda",
	"Jason",
	"Jessica",
	"David",
	"Melissa",
	"James",
	"Sarah",
	"Matthew",
	"Heather",
	"Joshua",
	"Nicole",
	"John",
	"Amy",
	"Robert",
	"Elizabeth",
	"Joseph",
	"Michelle",
	"Daniel",
	"Kimberly",
	"Brian",
	"Angela",
	"Justin",
	"Stephanie",
	"William",
	"Tiffany",
	"Ryan",
	"Christina",
	"Eric",
	"Lisa",
	"Nicholas",
	"Rebecca",
	"Jeremy",
	"Crystal",
	"Andrew",
	"Kelly",
	"Timothy",
	"Erin",
	"Jonathan",
	"Laura",
	"Adam",
	"Amber",
	"Kevin",
	"Rachel",
	"Anthony",
	"Jamie",
	"Thomas",
	"April",
	"Richard",
	"Mary",
	"Jeffrey",
	"Sara",
	"Steven",
	"Andrea",
	"Charles",
	"Shannon",
	"Brandon",
	"Megan",
	"Mark",
	"Emily",
	"Benjamin",
	"Julie",
	"Scott",
	"Danielle",
	"Aaron",
	"Erica",
	"Paul",
	"Katherine",
	"Nathan",
	"Maria",
	"Travis",
	"Kristin",
	"Patrick",
	"Lauren",
	"Chad",
	"Kristen",
	"Stephen",
	"Ashley",
	"Kenneth",
	"Christine",
	"Gregory",
	"Brandy",
	"Jacob",
	"Tara",
	"Dustin",
	"Katie",
	"Jesse",
	"Monica",
	"Jose",
	"Carrie",
	"Shawn",
	"Alicia",
	"Sean",
	"Courtney",
	"Bryan",
	"Misty",
	"Derek",
	"Kathryn",
	"Bradley",
	"Patricia",
	"Edward",
	"Holly",
	"Donald",
	"Stacy",
	"Samuel",
	"Karen",
	"Peter",
	"Anna",
	"Keith",
	"Tracy",
	"Kyle",
	"Brooke",
	"Ronald",
	"Samantha",
	"Juan",
	"Allison",
	"George",
	"Melanie",
	"Jared",
	"Leslie",
	"Douglas",
	"Brandi",
	"Gary",
	"Susan",
	"Erik",
	"Cynthia",
	"Phillip",
	"Natalie",
	"Raymond",
	"Jill",
	"Joel",
	"Dawn",
	"Corey",
	"Dana",
	"Shane",
	"Vanessa",
	"Larry",
	"Veronica",
	"Marcus",
	"Lindsay",
	"Zachary",
	"Tina",
	"Craig",
	"Kristina",
	"Derrick",
	"Stacey",
	"Todd",
	"Wendy",
	"Jeremiah",
	"Lori",
	"Antonio",
	"Catherine",
	"Carlos",
	"Kristy",
	"Shaun",
	"Heidi",
	"Dennis",
	"Sandra",
	"Frank",
	"Jacqueline",
	"Philip",
	"Kathleen",
	"Cory",
	"Christy",
	"Brent",
	"Leah",
	"Gabriel",
	"Valerie",
	"Nathaniel",
	"Pamela",
	"Randy",
	"Erika",
	"Luis",
	"Tanya",
	"Curtis",
	"Natasha",
	"Jeffery",
	"Katrina",
	"Alexander",
	"Lindsey",
	"Russell",
	"Melinda",
	"Casey",
	"Monique",
	"Jerry",
	"Teresa",
	"Wesley",
	"Denise",
	"Brett",
	"Tammy",
	"Luke",
	"Tonya",
	"Lucas",
	"Julia",
	"Seth",
	"Candice",
	"Billy",
	"Gina"}
