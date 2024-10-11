package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type store struct {
	data map[string]string 
	// list map[string][]string
	sets map[string]map[string]bool
}

func (s *store) set(key string, value string) string {
	s.data[key] = value
	return "OK"
}

func (s *store) get(key string) string {
	value, exists := s.data[key]
	if !exists {
		return "NULL" // Returning NULL for non-existent keys
	}
	return value
}

func (s *store) del(key string) string {
	_, exists := s.data[key]
	if exists {
		delete(s.data, key)
		return "OK"
	}
	return "NULL" // Returning NULL if key doesn't exist
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
	
    case "INCR":
        return fmt.Sprintf("%v", s.incr(args[0]))
    case "INCRBY":
        return fmt.Sprintf("%v", s.incrBy(args[0], args[1]))
    case "DECR":
        return fmt.Sprintf("%v", s.decr(args[0]))
    case "DECRBY":
        return fmt.Sprintf("%v", s.decrBy(args[0], args[1]))
	case "SADD":
        return fmt.Sprintf("%v", s.sadd(args[0], args[1]))
     case "SREM":
        return fmt.Sprintf("%v", s.srem(args[0], args[1]))
     case "SMEMBERS":
        members := s.smembers(args[0])
        result := ""
      
        for _, member := range members {
           result += fmt.Sprintf("%v ", member)
        }
      
        return strings.TrimSpace(result)
     case "SISMEMBER":
        return fmt.Sprintf("%v", s.sismember(args[0], args[1]))    
    default:
        return "ERROR : unkown command"
    }
}


func handleConnection(conn net.Conn, s *store) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.Fields(input) // Use Fields to split on spaces, tabs, etc.

		if len(parts) == 0 {
			fmt.Fprintln(conn, "ERROR: Empty command")
			continue
		}

		command := strings.ToUpper(parts[0]) // Make command case-insensitive
		args := parts[1:]

		response := s.handleCommand(command, args)

		fmt.Fprintln(conn, response) // Respond to client
	}
}

func main() {
	s := &store{data: make(map[string]string)}

	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server listening on port 6379...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn, s)
	}
}
