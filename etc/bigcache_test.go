package etc

import (
	"fmt"
	"github.com/allegro/bigcache"
	"log"
	"testing"
	"time"
	"unsafe"
)

func TestBigCache(t *testing.T) {
	forever := make(chan struct{})
	config := bigcache.Config{
		//  파편 수(2의 거듭제곱)
		Shards: 1024,

		//엔트리가 퇴출될 때까지의 시간
		//LifeWindow: 10 * time.Minute,
		LifeWindow: 3 * time.Second,

		//기한이 지난 엔트리를 삭제하는 간격(정리).  <= 0으로 설정하면 작업이 수행되지 않습니다.  1초 미만으로 설정하면 역효과가 납니다.큰 캐시의 해상도는 1초입니다.
		CleanWindow: 2 * time.Second,

		// rps * life Window, 초기 메모리 할당에서만 사용
		MaxEntriesInWindow: 1000 * 10 * 60,

		MaxEntrySize: 500, //최대 엔트리 크기(바이트), 초기 메모리 할당에만 사용됩니다.

		// 추가 메모리 할당에 대한 정보를 인쇄합니다.
		Verbose: true,

		// 캐시는 이 제한(MB)보다 많은 메모리를 할당할 수 없습니다.  값에 도달하면 가장 오래된 엔트리는 새로운 엔트리에 대해 덮어쓸 수 있습니다. 0 값은 크기 제한이 없음을 의미합니다.
		HardMaxCacheSize: 8192,

		//가장 오래된 엔트리가 유효기간 또는 남은 공간이 없기 때문에 삭제되었을 때 호출되는 콜백. 또는 삭제가 호출되었기 때문입니다. 이유를 나타내는 비트마스크가 반환됩니다. 기본값은 0 입니다.이것은 콜백이 없음을 의미하며 가장 오래된 엔트리의 래핑이 해제되지 않습니다.
		OnRemove: func(key string, entry []byte) {
			fmt.Println(key, "가 사라짐", string(entry))
		},

		// OnRemoveWithReason은 가장 오래된 엔트리가 만료되거나 남은 공간이 없기 때문에 삭제되었을 때 호출되는 콜백입니다.
		// 또는 삭제가 호출되었기 때문입니다. 이유를 나타내는 상수가 통과됩니다.
		// 기본값은 0 입니다.이것은 콜백이 없음을 의미하며 가장 오래된 엔트리의 래핑이 해제되지 않습니다.
		// OnRemove가 지정된 경우 무시됩니다.
		OnRemoveWithReason: nil,
	}

	cache, initErr := bigcache.NewBigCache(config)
	if initErr != nil {
		log.Fatal(initErr)
	}

	cache.Set("my-unique-key", []byte("메롱"))
	iter := cache.Iterator()

	cache.Set("my-unique-key1", []byte("메롱1"))
	cache.Set("my-unique-key2", []byte("메롱2"))
	cache.Set("key", []byte("키"))
	for iter.SetNext() {

		fmt.Println(iter.Value())
	}

	tick := time.Tick(time.Second)
	for {
		select {
		case <-tick:
			if entry, err := cache.Get("my-unique-key"); err == nil {
				fmt.Println("히히", string(entry))
			} else {
				fmt.Println(err)
				break
			}
		}
	}

	<-forever

}

func TestSize(t *testing.T) {
	a := int(123)
	b := int64(123)
	c := "foasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdofoasfasdfasfafasdfasdfasdfasdo" +
		""
	d := struct {
		FieldA float32
		FieldB string
	}{0, "bar"}

	fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("b: %T, %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("c: %T, %d\n", c, unsafe.Sizeof(c))
	fmt.Printf("d: %T, %d\n", d, unsafe.Sizeof(d))
}
