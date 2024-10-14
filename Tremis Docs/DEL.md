
# DELETE command

The `DEL` command is the last key-value pair command that is indeed used to delete a key value pair from the table.

---

the `DEL` command comes handy in situations when you want to delete something from the database. the functionality is quite simple and even with the simple syntax it makes it easy for the users to use it as well. 

```Tremis
DEL <key>
```

This removes the key from the database and also the value associated to the key

Example 
`
`DEL name`
output 
`DELETED`


---

```Go
func (s *store) del(key string) string {
	_, exists := s.data[key]
	if exists {
	 delete(s.data, key)
	 return "OK"
	}
	return "NULL" // Returning NULL if key doesn't exist
}
```

Well the above Golang code makes the work easier by taking in a key value and removing the needful key from the database and if there does not exists any key then it just returns `NULL` in return. 