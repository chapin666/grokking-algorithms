package chapter6

import (
	"strings"
)

var graph = map[string][]string{
	"you":    []string{"alice", "bob", "claire"},
	"bob":    []string{"anuj", "peggy"},
	"alice":  []string{"peggy"},
	"claire": []string{"thom", "jonny"},
	"anuj":   []string{},
	"peggy":  []string{},
	"thom":   []string{},
	"jonny":  []string{},
}

type deque []string

func search(name string) *string {

	// 广度遍历依赖队列，深度遍历依赖栈。这里选择队列
	var users deque

	// 先添加自己的好友到队列
	users = append(users, graph[name]...)

	for len(users) > 0 {
		user := users[0]
		users = append(users[:0], users[1:]...)
		if person_is_seller(user) {
			return &user
		} else {
			users = append(users, graph[user]...)
		}
	}

	return nil
}

func person_is_seller(name string) bool {
	return strings.HasSuffix(name, "m")
}
