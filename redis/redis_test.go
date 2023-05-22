package redis

import (
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	// Redis 클라이언트 생성
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Redis 인증 비밀번호 (없는 경우 빈 문자열 사용)
		DB:       0,  // Redis 데이터베이스 번호
	})

	// Redis 분산 락 구현을 위한 Redisson 인스턴스 생성
	redisson := redissync.New(client)

	// 분산 락 객체 생성
	lock := redisson.NewLock("my_lock")

	// 분산 락 획득
	err := lock.Acquire()
	if err != nil {
		panic(err)
	}

	// 작업 수행
	fmt.Println("분산 락을 획득하여 작업을 수행합니다.")

	// 분산 락 해제
	err = lock.Release()
	if err != nil {
		panic(err)
	}

	// Redis 클라이언트 연결 종료
	err = client.Close()
	if err != nil {
		panic(err)
	}
}
