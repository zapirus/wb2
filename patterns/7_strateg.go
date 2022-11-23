package patterns

import "fmt"

type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

// Сеттер, обеспечивающий выбор стратегии
func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

// Инетерфейс сратегий
type evictionAlgo interface {
	evict(c *cache)
}

// Стратегии с заглушками вместо реализации

// Стратегия 1
type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

// Стратегия 2
type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

// Стратегия 3
type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}

func ExampleStrategy() {
	fmt.Println("Strategy example")
	fmt.Println()

	// Инициализация и заполнение кэша со стратегией lfu
	lfu := &lfu{}
	cache := initCache(lfu)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	// Смена стратегии
	lru := &lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	// Смена стратегии
	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")

}
