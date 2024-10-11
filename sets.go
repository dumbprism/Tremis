package main

func (s *store) sadd(key string, value string) bool {
	if s.sets[key] == nil {
		s.sets[key] = make(map[string]bool)
	}

	if s.sets[key][value] {
		return false
	}

	s.sets[key][value] = true

	return true
}

func (s *store) srem(key string, value string) bool {
	if s.sets[key] == nil {
		return false
	}

	if !s.sets[key][value] {
		return false
	}

	delete(s.sets[key], value)

	return true
}

func (s *store) smembers(key string) []string {
	if s.sets[key] == nil {
		return []string{}
	}

	members := make([]string, 0, len(s.sets[key]))

	for member := range s.sets[key] {
		members = append(members, member)
	}

	return members
}

func (s *store) sismember(key string, value string) bool {
	if s.sets[key] == nil {
		return false
	}

	return s.sets[key][value]
}
