package main

import (
	"net"
	"fmt"
)
func (s *store) subscribe(channel string, conn net.Conn) string {
	client := client{conn: conn}
	s.subs[channel] = append(s.subs[channel], client)
	return "OK"
}

func (s *store) publish(channel string, message string) {
	subs, ok := s.subs[channel]
	if ok {
		for _, subscriber := range subs {
			fmt.Fprintf(subscriber.conn, "+%s\n", message)
		}
	}
}