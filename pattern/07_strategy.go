package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Стратегия предлагает определить семейство схожих алгоритмов, которые часто изменяются или расширяются, и вынести их в собственные структуры, называемые стратегиями.

Применение:
    Когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
    Когда есть множество похожих структур, отличающихся только некоторым поведением.
    Когда нет желания обнажать детали реализации алгоритмов для других структур.

Плюсы:
    Замена алгоритмов на лету.
    Изоляция кода и данных алгоритмов от остальных структур.
    Уход от наследования к делегированию.

Минусы:
    Усложняет программу за счёт дополнительных структур.
    Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

Примеры:
	Выбор алгоритма кеширования в зависимости от нагрузки.
*/

type EvictionAlgo interface {
	evict(c *Cache)
}

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting cache using FIFO strtegy")
}

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting cache using LRU strtegy")
}

func main() {
	lfu := &Lru{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("d", "4")
}
