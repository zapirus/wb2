package patterns

import "fmt"

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

// "Наследники" (нет)
type colt struct {
	gun
}

func colts() iGun {
	return &colt{
		gun: gun{
			name:  "Colt gun",
			power: 4,
		},
	}
}

type m16 struct {
	gun
}

func m16s() iGun {
	return &m16{
		gun: gun{
			name:  "colt gun",
			power: 1,
		},
	}
}

// Фабрика
func getGun(gunType string) (iGun, error) {
	switch gunType {
	case "colt":
		return colts(), nil
	case "m16":
		return m16s(), nil
	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}

func ExampleFactory() {
	fmt.Println("Factory example")
	fmt.Println()

	s, _ := getGun("colt")
	ms, _ := getGun("m16")

	printDetails(s)
	printDetails(ms)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
