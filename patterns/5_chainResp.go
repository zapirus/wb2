package patterns

import "fmt"

type department interface {
	execute(*patient)
	setNext(department)
}

// Обработчик 1
type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.hasInsurnce {
		fmt.Printf("Reception registering %s\n", p.name)
		r.next.execute(p)
		return
	}
	fmt.Printf("%s has no insurance\n", p.name)
}

func (r *reception) setNext(next department) {
	r.next = next
}

// Обработчик 2
type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.heavyDesease {
		fmt.Printf("Doctor prescribes treatment for %s\n", p.name)
		d.next.execute(p)
		return
	}
	fmt.Printf("Doctor cured %s\n", p.name)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

// Обработчик 3
type hospital struct {
	next department
}

func (h *hospital) execute(p *patient) {
	fmt.Printf("%s admitted to hospital\n", p.name)
}

func (h *hospital) setNext(next department) {
	h.next = next
}

// Объект-запрос, проходящий обработку
type patient struct {
	name         string
	hasInsurnce  bool
	heavyDesease bool
}

func ExampleChainOfResp() {
	fmt.Println("Chain of responsibility example")
	fmt.Println()

	// Последний обработчик в цепочке
	hospital := &hospital{}

	doctor := &doctor{}
	doctor.setNext(hospital)

	// Первый обработчик
	reception := &reception{}
	reception.setNext(doctor)

	patient1 := &patient{name: "abc"}
	patient2 := &patient{name: "def", hasInsurnce: true}
	patient3 := &patient{name: "def", hasInsurnce: true, heavyDesease: true}

	// Старт цепочки с первого обработчика
	reception.execute(patient1)
	fmt.Println()
	reception.execute(patient2)
	fmt.Println()
	reception.execute(patient3)
}
