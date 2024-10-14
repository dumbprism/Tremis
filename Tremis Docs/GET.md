
# GET command

Now that we have understood the `SET` commands functionality, there must be a way to fetch these values and showcase them to the user so that it is successfully  getting stored. Well, that is when we use the `GET` commands.

---
The `GET` commands primary purpose is to retrieve the values that are stored in the database. The only thing required here is to specify the key name and you'll have your value associated to the key. Just remind yourself of how "objects" work in a programming language.

To retrieve a value from an object we specify the objects name and the key associated to it

```JavaScript

const Objects = {
	firstName : "John",
	lastName  : "Doe"
}
console.log(Objects.firstName)

```

To better understand `GET` you can have a look at the above code example where `Object.firstName` gives you the first name.

In tremis, as mentioned already we just have to specify the proper key value else you'll be returned with a `NULL` value

```Tremis
	GET <key>
```

example  : 

`GET name`
output
`John Doe`

`GET age`
output
`10`

`GET country`
output
`NULL`

---

```Go
func (s *store) get(key string) string {
	value, exists := s.data[key]
	if !exists {
		return "NULL" // Returning NULL for non-existent keys
	}
	return value
}
```

The above Golang code explains how `GET` works and it is pretty much easy to understand.
