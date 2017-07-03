package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

type Post struct {
	Title   string
	Content int
}

func main() {
	// f, err := os.OpenFile("test.file", os.O_RDWR|os.O_CREATE, 0755)
	postMap := make(map[int]*Post)

	timeStart := time.Now()
	// f, err := os.Open("test.file")
	f, err := os.Create("test.file")
	if err != nil {
		fmt.Println(err)
	}
	// defer f.Close()
	for i := 0; i < 10; i++ {
		postMap[i] = &Post{"title", i}

	}
	// timeStart := time.Now()
	// data, _ := msgpack.Marshal(&postMap)

	// //JSON Encode
	// jenc := json.NewEncoder(f)
	// err = jenc.Encode(&postMap)

	// //MSGPACK Encode
	// menc := msgpack.NewEncoder(f)
	// err = menc.Encode(&postMap)

	// GOB Encode
	genc := gob.NewEncoder(f)
	err = genc.Encode(&postMap)

	fmt.Println(time.Since(timeStart))
	if err != nil {
		fmt.Println(err)
	}
	f.Sync()
	fi, _ := f.Stat()
	fmt.Println(fi.Size())
	f.Close()

	rf, err := os.Open("test.file")
	defer rf.Close()

	// //JSON
	// jdec := json.NewDecoder(rf)
	// err = jdec.Decode(&postMap)

	// //MSG
	// mdec := msgpack.NewDecoder(rf)
	// err = mdec.Decode(&postMap)

	//GOB
	gdec := gob.NewDecoder(rf)
	err = gdec.Decode(&postMap)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(time.Since(timeStart))
}

// github.com/calmh/xdr

// func generateXDR() []*XDRA {
//         a := make([]*XDRA, 0, 1000)
//         for i := 0; i < 1000; i++ {
//                 a = append(a, &XDRA{
//                         Name:     randString(16),
//                         BirthDay: time.Now().UnixNano(),
//                         Phone:    randString(10),
//                         Siblings: rand.Int31n(5),
//                         Spouse:   rand.Intn(2) == 1,
//                         Money:    math.Float64bits(rand.Float64()),
//                 })
//         }
//         return a
// }

// func BenchmarkXDR2Marshal(b *testing.B) {
//         b.StopTimer()
//         data := generateXDR()
//         b.ReportAllocs()
//         b.StartTimer()
//         for i := 0; i < b.N; i++ {
//                 data[rand.Intn(len(data))].MarshalXDR()
//         }
// }

// func BenchmarkXDR2Unmarshal(b *testing.B) {
//         b.StopTimer()
//         data := generateXDR()
//         ser := make([][]byte, len(data))
//         for i, d := range data {
//                 ser[i] = d.MustMarshalXDR()
//         }
//         b.ReportAllocs()
//         b.StartTimer()
//         for i := 0; i < b.N; i++ {
//                 n := rand.Intn(len(ser))
//                 o := XDRA{}
//                 err := o.UnmarshalXDR(ser[n])
//                 if err != nil {
//                         b.Fatalf("xdr failed to unmarshal: %s (%s)", err, ser[n])
//                 }
//                 // Validate unmarshalled data.
//                 if validate != "" {
//                         i := data[n]
//                         correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay
//                         if !correct {
//                                 b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
//                         }
//                 }
//         }
// }
