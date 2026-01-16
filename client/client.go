package client

import "fmt"

type client struct {
	Name string
}

func NewClient(name string) *client {
	return &client{Name: name}
}

func (c *client) GetName() string {
	return fmt.Sprintf("Hello World! My name is %s", c.Name)
}
