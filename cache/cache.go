package cache

type Cache struct {
	imageMap map[string][]byte
}

func InitCache() *Cache {
	imgMap := make(map[string][]byte)
	return &Cache{
		imageMap : imgMap,
	}
}

func (c *Cache) IsPresent(key string) bool {
	return c.imageMap[key] != nil 
}

func (c *Cache) Get(key string) []byte {
	return c.imageMap[key]
}

func (c *Cache) Add(key string, value []byte) {
	c.imageMap[key] = value
}