package another

// import (
// 	"fmt"
// 	"log"
// 	"sync"
// )

// var (
// 	rwMutex sync.RWMutex
// 	data    = make(map[string]string)
// )

// Задача по использованию sync.RWMutex
// для доступа к карте (map) - это классический способ обеспечить
// конкурентную безопасность при работе с картой в Go.
// Напишите методы read() и write(),
// обратите внимание на выбор правильных методов мьютекса в функции read
// func read(key string) string {
// 	rwMutex.RLock()
// 	defer rwMutex.RUnlock()
// 	return data[key]
// }

// func write(key, value string) {
// 	rwMutex.Lock()
// 	defer rwMutex.Unlock()
// 	data[key] = value
// }

// func main() {
// 	var key string
// 	var value string
// 	_, err := fmt.Scan(&key, &value)
// 	if err != nil {
// 		log.Println("Error", err)
// 	}
// 	write(key, value)
// 	fmt.Println(read(key))
// }
