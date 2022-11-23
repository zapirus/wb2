package patterns

import "fmt"

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// Интерфейс команды
type command interface {
	execute()
}

// Команда 1
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

// Команда 2
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

// Интерфейс получателя
type device interface {
	on()
	off()
}

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
}

func (t *tv) off() {
	t.isRunning = false
}

func ExampleCommand() {
	fmt.Println("Command example")
	fmt.Println()

	// Полчатель
	tv := &tv{}

	// Команды
	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	// Отправители с разными командами
	onButton := &button{
		command: onCommand,
	}
	onButton.press()
	fmt.Println(tv.isRunning)

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
	fmt.Println(tv.isRunning)
}
