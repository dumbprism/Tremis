
The sole purpose of creating set as the name suggests was to set values to your database and then add an optional disk storage to it. Storing it in the disk is quite upto the user to decide but in general it would work as an in-memory database set functionality. 

---


the functionality is pretty straightforward and it stores the values in the RAM 

```Tremis
SET <key> <value>
```

Example : 
``SET name John Doe
`SET age 10`
`SET state true`

Once you have set the value for assurance you get a `OK` message.

The setting of these values works quite easily through key-value pairs. 

let us take the name example :  

`name` here is the key whose value would be `John Doe`. Now the sole reason for choosing key-value pair was for faster retrieval of values and keeping no lag time. 

Here, if a key already consists an existing value, then to store another value, it over writes the older values and keeps the ew updates values as the present value. It is like Objects in your programming languages (key-value pairs) . 

---

```Go
func (s *store) set(key string, value string) string {
s.data[key] = value
    return "OK"
}
```

With just a simple function written in Golang, the Set function was created. As already mentioned it is stored in key-value pairs which is conspicuous in the above code snippet. 

