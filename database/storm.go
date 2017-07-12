package main

import (
	"fmt"
	"time"

	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/gob"
	"github.com/asdine/storm/codec/json"
	"github.com/asdine/storm/codec/msgpack"
	"github.com/asdine/storm/codec/protobuf"
	"github.com/asdine/storm/codec/sereal"
)

var db, _ = storm.Open("normal.db")
var gobDb, _ = storm.Open("gob.db", storm.Codec(gob.Codec))
var jsonDb, _ = storm.Open("json.db", storm.Codec(json.Codec))
var serealDb, _ = storm.Open("sereal.db", storm.Codec(sereal.Codec))
var protobufDb, _ = storm.Open("protobuf.db", storm.Codec(protobuf.Codec))
var msgpackDb, _ = storm.Open("msgpack.db", storm.Codec(msgpack.Codec))

// Data type
type Data struct {
	ID        int    `storm:"id"`
	Title     string `storm:"index"`
	Content   string
	CreatedAt time.Time `storm:"index"`
}

func main() {
	timeStart := time.Now()
	for i := 0; i < 1; i++ {
		data := &Data{i + 1, "Title", "Content", time.Now()}
		err := gobDb.Save(data)
		// err = jsonDb.Save(data)
		// err = serealDb.Save(data)
		// err = protobufDb.Save(data)
		// err := msgpackDb.Save(data)
		if err != nil {
			fmt.Print("fuck")
		}
	}
	fmt.Println(time.Since(timeStart))

	timeStart = time.Now()
	for i := 0; i < 1; i++ {
		data := &Data{i + 1, "Title", "Content", time.Now()}
		// err := gobDb.Save(data)
		err := jsonDb.Save(data)
		// err = serealDb.Save(data)
		// err = protobufDb.Save(data)
		// err := msgpackDb.Save(data)
		if err != nil {
			fmt.Print("fuck")
		}
	}
	fmt.Println(time.Since(timeStart))

	timeStart = time.Now()
	for i := 0; i < 1; i++ {
		data := &Data{i + 1, "Title", "Content", time.Now()}
		// err := gobDb.Save(data)
		// err = jsonDb.Save(data)
		err := serealDb.Save(data)
		// err = protobufDb.Save(data)
		// err := msgpackDb.Save(data)
		if err != nil {
			fmt.Print("fuck")
		}
	}
	fmt.Println(time.Since(timeStart))

	timeStart = time.Now()
	for i := 0; i < 1; i++ {
		data := &Data{i + 1, "Title", "Content", time.Now()}
		// err := gobDb.Save(data)
		// err = jsonDb.Save(data)
		// err = serealDb.Save(data)
		err := protobufDb.Save(data)
		// err := msgpackDb.Save(data)
		if err != nil {
			fmt.Print("fuck")
		}
	}
	fmt.Println(time.Since(timeStart))

	timeStart = time.Now()
	for i := 0; i < 1; i++ {
		data := &Data{i + 1, "Title", "Content", time.Now()}
		// err := gobDb.Save(data)
		// err = jsonDb.Save(data)
		// err = serealDb.Save(data)
		// err = protobufDb.Save(data)
		err := msgpackDb.Save(data)
		if err != nil {
			fmt.Print("fuck")
		}
	}
	fmt.Println(time.Since(timeStart))

	timeStart = time.Now()
	for i := 0; i < 1; i++ {
		data := &Data{i + 1, "Title", "Content", time.Now()}
		// err := gobDb.Save(data)
		// err = jsonDb.Save(data)
		// err = serealDb.Save(data)
		// err = protobufDb.Save(data)
		err := db.Save(data)
		if err != nil {
			fmt.Print("fuck")
		}
	}
	fmt.Println(time.Since(timeStart))
}
