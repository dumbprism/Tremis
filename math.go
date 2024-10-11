package main 

import  (
	"strconv"
)

func (s *store) incr(key string) int {
	value, _ := strconv.Atoi(s.get(key))
   
	s.set(key, strconv.Itoa(value+1))
   
	return value + 1
 }
 
 func (s *store) incrBy(key string, incr string) int {
	number, _ := strconv.Atoi(incr)
   
	value, _ := strconv.Atoi(s.get(key))
   
	s.set(key, strconv.Itoa(value+number))
   
	return value + number
 }
 
 func (s *store) decr(key string) int {
	value, _ := strconv.Atoi(s.get(key))
   
	s.set(key, strconv.Itoa(value-1))
   
	return value - 1
 }
 
 func (s *store) decrBy(key string, decr string) int {
	number, _ := strconv.Atoi(decr)
   
	value, _ := strconv.Atoi(s.get(key))
   
	s.set(key, strconv.Itoa(value-number))
   
	return value - number
 }

