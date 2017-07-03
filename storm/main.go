package main

import (
	"fmt"
	"time"

	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/msgpack"
)

func main() {

	var db, _ = storm.Open("msgp.db", storm.Codec(msgpack.Codec))
	// var db, _ = storm.Open("json.db")

	type User struct {
		ID        int    `storm:"id,increment"` // primary key
		Group     string //`storm:"index"`        // this field will be indexed
		Email     string // this field will be indexed with a unique constraint
		Name      string // this field will not be indexed
		Age       int    `storm:"index"`
		CreatedAt time.Time
	}
	timeStart := time.Now()
	var user User
	for i := 1; i < 1000; i++ {

		// user := User{
		// 	Group:     "staff",
		// 	Email:     "john@provider.com",
		// 	Name:      "John",
		// 	Age:       rand.Intn(100),
		// 	CreatedAt: time.Now(),
		// }
		// err := db.Save(&user)
		// if err != nil {
		// 	fmt.Println("ERr0r:", err)
		// }
		err := db.One("ID", i, &user)
		// err == nil

		// user.ID++
		// err = db.Save(&user)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println(&user)
		// err == storm.ErrAlreadyExists

		// var users []User
		// timeStart := time.Now()
		// err = db.All(&users)
		// fmt.Println(time.Since(timeStart))
		// // fmt.Println(&users)

		// fi, e := os.Stat("json.db")
		// if e != nil {
		// 	fmt.Print(e)
		// }
		// // get the size
		// fmt.Println(fi.Size())

		// time.Sleep(50 * time.Millisecond)

	}

	// var users []User
	// timeStart := time.Now()
	// err := db.All(&users)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(&user)
	fmt.Println(time.Since(timeStart))
	fmt.Println(1000 / 15.38)

	// var users []User
	// // db.AllByIndex("Age", &users)
	// db.Range("Age", 10, 21, &users)
	// timeStart := time.Now()
	// for i := 0; i < len(users); i++ {
	// 	fmt.Println(users[i])
	// }
	// fmt.Println(time.Since(timeStart))
	// // fmt.Println(&users)

}
