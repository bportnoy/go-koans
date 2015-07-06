package go_koans

import (
	"bytes"
	"io"
	"fmt"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		b := make([]byte, 8)

		for {
			n, err := in.Read(b)
			fmt.Printf("b = %v\n", b)
			fmt.Printf("string = %q\n", b[:n])
			out.Write(b[:n])
			fmt.Printf("out = %v\n", out.String())
			if err == io.EOF {
				break
			}
		}


		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		b := make([]byte, 5)

		_, err := in.Read(b)

		fmt.Printf("err = %v", err)

		out.Write(b)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
