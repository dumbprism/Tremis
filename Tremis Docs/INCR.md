
# INCREMENT command

The `incr` command is used for incrementing numbers when required.

---

As the command suggests, Increment is used to increment a key value by 1 and if there does not exist any key then the value is automatically set to 0.

```Tremis
INCR age 
```

Let us suppose there already exists a key with the variable name a `age` whose value is set to 10.

After I implement the INCR function , the output gives 11.

It is simple to understand and easy to implement. 

---

```Go
func (s *store) incr(key string) int {
	value, _ := strconv.Atoi(s.get(key))
	s.set(key, strconv.Itoa(value+1))
	return value + 1
}
```

The above code increments the value by 1 and thus gives the desired output, Simple to execute and implement in Golang.



