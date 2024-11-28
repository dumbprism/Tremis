package main

func (s *store) help() string {
	return `
		Available commands and their syntax:
		1. SET key value             - Set a value for a key : SET <key> <value>
		2. GET key                   - Get the value of a key : GET <key>
		3. DEL key                   - Delete a key : DEL <key>
		4. INCR key                  - Increment the value of a key (must be an integer) : INCR <key>
		5. INCRBY key increment      - Increment the value of a key by a specific integer : INCRBY <key> <increment>
		6. DECR key                  - Decrement the value of a key (must be an integer) : DECR <key>
		7. DECRBY key decrement      - Decrement the value of a key by a specific integer : DECRBY <key> <decrement>
		8. LPUSH list value          - Insert a value at the beginning of a list : LPUSH <list> <value>
		9. RPUSH list value          - Insert a value at the end of a list : RPUSH <list> <value>
		10. LPOP list                - Remove and return the first element of a list : LPOP <list>
		11. RPOP list                - Remove and return the last element of a list  : RPOP <list>
		12. LLEN list                - Get the length of a list : LLEN <list>
		13. LINDEX list index        - Get the element at the specified index in a list : LINDEX <list> <index>
		14. SADD set value           - Add a value to a set : SADD <set> <value>
		15. SREM set value           - Remove a value from a set : SREM <set> <value>
		16. SMEMBERS set             - Get all members of a set : SMEMBERS <set>
		17. SISMEMBER set value      - Check if a value is a member of a set : SISMEMBER <set> <value>
		18. SUBSCRIBE channel        - Subscribe to a channel : SUBSCRIBE <channel>
		19. PUBLISH channel message  - Publish a message to a channel : PUBLISH <channel> <message>
		20. --help                    - Display this help message : --HELP
	`

}
