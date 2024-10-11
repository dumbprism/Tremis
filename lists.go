package main

func (s *store) lPush(key string, value string) int {
	if s.list[key] == nil {
	   s.list[key] = make([]string, 0)
	}
   
	s.list[key] = append([]string{value}, s.list[key]...)
   
	return len(s.list[key])
 }
 
 func (s *store) rPush(key string, value string) int {
	if s.list[key] == nil {
	   s.list[key] = make([]string, 0)
	}
   
	s.list[key] = append(s.list[key], value)
   
	return len(s.list[key])
 }
 
 func (s *store) lPop(key string) string {
	if s.list[key] == nil {
	   return ""
	}
   
	value := s.list[key][0]
	s.list[key] = s.list[key][1:]
   
	return value
 }
 
 func (s *store) rPop(key string) string {
	if s.list[key] == nil {
	   return ""
	}
   
	value := s.list[key][len(s.list[key])-1]
	s.list[key] = s.list[key][:len(s.list[key])-1]
   
	return value
 }
 
 func (s *store) lLen(key string) int {
	if s.list[key] == nil {
	   return 0
	}
   
	return len(s.list[key])
 }
 
 func (s *store) lIndex(key string, index int) string {
	if s.list[key] == nil || index >= len(s.list[key]) || index < 0 {
	   return ""
	}
   
	return s.list[key][index]
 }