package main

import (
  "fmt"
  "net"
  "log"
  "bufio"
)

func main()  {
  li, err := net.Listen("tcp", ":50080")
  if err != nil {
    log.Fatalln(err)
  }
  defer li.Close()

  for {
    conn, err := li.Accept()
    if err != nil {
      log.Fatalln(err)
    }

    go handle(conn)
  }
}

func handle(conn net.Conn)  {
  scanner := bufio.NewScanner(conn)
  for scanner.Scan() {
    ln := scanner.Text()
    fmt.Println(ln)
    fmt.Fprintf(conn, "I heard you say %s\n", ln)
  }
  defer conn.Close()

  fmt.Println("Code got here.")
}