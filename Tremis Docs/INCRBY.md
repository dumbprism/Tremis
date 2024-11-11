
# Increment by command

The `INCRBY` command works similar to increment but with just a small update

---

The `INCRBY` command let's you decide how much you want to increment a value with. 

Once you have specified the key and give it a value you can then increment the value by how much ever you need. 

Consider it to be a simple addition operation where you already know a value and the other value is also set by you. 


```Tremis

SET age 10
INCRBY(age 20)
GET age
```

Here we notice that once the value of age is set to 10, we are incrementing the age by the value of 2. The output thus would be 12 instead of 10.

---

```Go

 func (s *store) incrBy(key string, incr string) int {
	 number, _ := strconv.Atoi(incr)
    value, _ := strconv.Atoi(s.get(key))
    s.set(key, strconv.Itoa(value+number))
    return value + number
}

```

The above explains the simple increment by functionality which is coded in Golang.  Since we are considering the value to be in String format we are first converting it in the form of integer and thus moving forward with the incrementation process.