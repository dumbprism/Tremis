package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type store struct {
	data map[string]string
	list map[string][]string
	sets map[string]map[string]bool
	subs map[string][]client
	disk *diskStore
}

func (s *store) set(key string, value string) string {
	s.data[key] = value
	s.disk.save(s.data)
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
		s.disk.save(s.data)
		return "OK"
	}
	return "NULL" // Returning NULL if key doesn't exist
}

func (s *store) help() string {
	return `
Available commands and their syntax:
1. SET key value             - Set a value for a key
2. GET key                   - Get the value of a key
3. DEL key                   - Delete a key
4. INCR key                  - Increment the value of a key (must be an integer)
5. INCRBY key increment      - Increment the value of a key by a specific integer
6. DECR key                  - Decrement the value of a key (must be an integer)
7. DECRBY key decrement      - Decrement the value of a key by a specific integer
8. LPUSH list value          - Insert a value at the beginning of a list
9. RPUSH list value          - Insert a value at the end of a list
10. LPOP list                - Remove and return the first element of a list
11. RPOP list                - Remove and return the last element of a list
12. LLEN list                - Get the length of a list
13. LINDEX list index        - Get the element at the specified index in a list
14. SADD set value           - Add a value to a set
15. SREM set value           - Remove a value from a set
16. SMEMBERS set             - Get all members of a set
17. SISMEMBER set value      - Check if a value is a member of a set
18. SUBSCRIBE channel        - Subscribe to a channel
19. PUBLISH channel message  - Publish a message to a channel
20. --help                    - Display this help message
`
}

func (s *store) handleCommand(command string, args []string, conn net.Conn) string {
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
	case "LPUSH":
		return fmt.Sprintf("%v", s.lPush(args[0], args[1]))
	case "RPUSH":
		return fmt.Sprintf("%v", s.rPush(args[0], args[1]))
	case "LPOP":
		return s.lPop(args[0])
	case "RPOP":
		return s.rPop(args[0])
	case "LLEN":
		return fmt.Sprintf("%v", s.lLen(args[0]))
	case "LINDEX":
		index, _ := strconv.Atoi(args[1])
		return s.lIndex(args[0], index)
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

	case "SUBSCRIBE":
		return s.subscribe(args[0], conn)
	case "PUBLISH":
		s.publish(args[0], strings.Join(args[1:], " "))
		return "OK"

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

		response := s.handleCommand(command, args, conn)

		fmt.Fprintln(conn, response) // Respond to client
	}
}

func main() {
	// Initialize both data and sets maps
	s := &store{
		data: make(map[string]string),
		list: make(map[string][]string),
		sets: make(map[string]map[string]bool),
		subs: make(map[string][]client),
		disk: &diskStore{},
		
	}

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
