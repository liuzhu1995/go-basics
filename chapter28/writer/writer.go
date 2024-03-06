package main

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}

	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbs(name string) error {
	// 创建一个名为name的文件，如果文件已存在会重置为空文件
	// 如果成功返回的File文件对象可用于I/O
	f, err := os.Create(name)

	if err != nil {
		return err
	}

	defer f.Close()

	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don’t just check errors, handle them gracefully.")
	sw.writeln("Don't panic.")
	sw.writeln("Make the zero value useful.")
	sw.writeln("The bigger the interface, the weaker the abstraction.")
	sw.writeln("interface{} says nothing.")
	sw.writeln("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.")
	sw.writeln("Documentation is for users.")
	sw.writeln("A little copying is better than a little dependency.")
	sw.writeln("Clear is better than clever.")
	sw.writeln("Concurrency is not parallelism.")
	sw.writeln("Don’t communicate by sharing memory, share memory by communicating.")
	sw.writeln("Channels orchestrate; mutexes serialize.")
	fmt.Println(sw.w, "sw")

	return sw.err
}

func main() {
	err := proverbs("/writer.txt")

	fmt.Println(err)

}
