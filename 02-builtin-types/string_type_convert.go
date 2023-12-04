package main

import "fmt"

func testConversion() {
	var s string = "字符串"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Printf("%v\n%v\n%v\n", s, bs, rs)

	for _, r := range s {
		fmt.Println(r)
	}
}

// https://go.dev/src/runtime/string.go
// https://go.dev/src/runtime/utf8.go

// go build -gcflags="-S" .\string_type_convert.go
//        rel 3+0 t=23 type:string+0
//        rel 3+0 t=23 type:[]uint8+0
//        rel 3+0 t=23 type:[]int32+0
//        rel 3+0 t=23 type:*os.File+0
//        rel 3+0 t=23 type:int32+0
//        rel 3+0 t=23 type:*os.File+0
//        rel 31+4 t=14 go:string."字符串"+0
//        rel 41+4 t=7 runtime.stringtoslicebyte+0
//        rel 65+4 t=14 go:string."字符串"+0
//        rel 75+4 t=7 runtime.stringtoslicerune+0
//        rel 124+4 t=14 go:string."字符串"+0
//        rel 134+4 t=7 runtime.convTstring+0
//        rel 141+4 t=14 type:string+0
//        rel 177+4 t=7 runtime.convTslice+0
//        rel 184+4 t=14 type:[]uint8+0
//        rel 225+4 t=7 runtime.convTslice+0
//        rel 232+4 t=14 type:[]int32+0
//        rel 255+4 t=14 os.Stdout+0
//        rel 262+4 t=14 go:itab.*os.File,io.Writer+0
//        rel 269+4 t=14 go:string."%v\n%v\n%v\n"+0
//        rel 296+4 t=7 fmt.Fprintf+0
//        rel 321+4 t=7 runtime.convT32+0
//        rel 328+4 t=14 type:int32+0
//        rel 345+4 t=14 os.Stdout+0
//        rel 352+4 t=14 go:itab.*os.File,io.Writer+0
//        rel 370+4 t=7 fmt.Fprintln+0
//        rel 393+4 t=14 go:string."字符串"+0
//        rel 427+4 t=7 runtime.decoderune+0
//        rel 451+4 t=7 runtime.morestack_noctxt+0
