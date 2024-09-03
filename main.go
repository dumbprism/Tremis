package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

type store struct {
    data map[string]string
}

func (s *store) set(key string, value string) string {
    s.data[key] = value
  
    return value
}

func (s *store) get(key string) string {
    return s.data[key]
}

func (s *store) del(key string) {
    delete(s.data, key)
}

func (s *store) handleCommand(command string, args []string) string {
    switch command {
     case "SET":
        // Using Join to save the rest of the received data
        return s.set(args[0], strings.Join(args[1:], " "))
     case "GET":
        return s.get(args[0])
     case "DEL":
        s.del(args[0])
        return "DELETED"
     default:
        return "ERROR: Unknown command"
    }
}

func handleConnection(conn net.Conn, s *store) {
    defer conn.Close()

    scanner := bufio.NewScanner(conn)

    for scanner.Scan() {
        input := scanner.Text()
        parts := strings.Split(input, " ")

        if len(parts) < 2 {
            fmt.Fprintln(conn, "ERROR: Unknown command")
            continue
        }

        command := parts[0]
        args := parts[1:]

        response := s.handleCommand(command, args)

        fmt.Fprintln(conn, response)
    }
}

func main() {
    s := &store{data: make(map[string]string)}

    listener, err := net.Listen("tcp", ":6379")
    if err != nil {
        panic(err)
    }

    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            panic(err)
        }

        go handleConnection(conn, s)
    }
}