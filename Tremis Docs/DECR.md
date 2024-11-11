
### Decrement by command

The `DECR` command works just as the name suggests. Just few changes to the increment functionality and we are good to go with decrement as well.

---

The `DECR` command decrements a value by 1. 

Once the value has been specified for a key, if an integer is detected then it makes sure to decrement the value by 1. It is just like the normal decrement which we notice in our programming languages. 


```Tremis

SET age 10
DECR age
GET age
```

We get the value of the `age` key as 9, instead of 10 since we used the `DECR` command. The commands are quite easier to pick up and understand since they were created with the sole purpose of easy understanding.

---

```Go
func (s *store) decr(key string) int {

    value, _ := strconv.Atoi(s.get(key))

    s.set(key, strconv.Itoa(value-1))

    return value - 1

 }

```

this is quite a straight forward code to decrement a value. Again the value is first converted to integer and then the value is decremented by 1 and the value is returned. 


