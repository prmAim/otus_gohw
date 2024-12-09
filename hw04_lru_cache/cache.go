package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int               // ёмкость (количество сохраняемых в кэше элементов)
	queue    List              // очередь \[последних используемых элементов\] на основе двусвязного списка
	items    map[Key]*ListItem // словарь, отображающий ключ (строка) на элемент очереди
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

// Set Добавить значение в кэш по ключу.
// Алгоритм работы кэша:
// - при добавлении элемента:
// - если элемент присутствует в словаре, то обновить его значение и переместить элемент в начало очереди;
// - если элемента нет в словаре, то добавить в словарь и в начало очереди
// (при этом, если размер очереди больше ёмкости кэша,
// то необходимо удалить последний элемент из очереди и его значение из словаря);
// - возвращаемое значение - флаг, присутствовал ли элемент в кэше.
func (cache *lruCache) Set(key Key, value interface{}) bool {
	if cache.capacity == 0 || key == "" {
		return false
	}

	// если элемент в словаре (cache.items[key]) есть, то ..., иначе ...
	if item, isOk := cache.items[key]; isOk {
		item.Value = value
		cache.queue.MoveToFront(item)
		cache.items[key] = cache.queue.Front()
		return isOk
	}

	// Если превысили кол-во элементов, чем размер очереди, то ...
	if cache.capacity == cache.queue.Len() {
		item := cache.queue.Back()

		for k, val := range cache.items {
			if val == item {
				delete(cache.items, k)
				break
			}
		}

		cache.queue.Remove(item)
	}

	// Добавили новый элемент в начало
	newItem := cache.queue.PushFront(value)
	cache.items[key] = newItem

	return false
}

// Get Получить значение из кэша по ключу.
func (cache *lruCache) Get(key Key) (interface{}, bool) {
	if key == "" || cache.capacity == 0 {
		return nil, false
	}

	// если элемент в словаре (cache.items[key]) есть, то ..., иначе ...
	if item, isOk := cache.items[key]; isOk {
		cache.queue.MoveToFront(item)
		cache.items[key] = cache.queue.Front()
		return item.Value, isOk
	}

	return nil, false
}

// Clear Очистить кэш.
func (cache *lruCache) Clear() {
	cache.items = make(map[Key]*ListItem, cache.capacity)
	cache.queue = NewList()
}
