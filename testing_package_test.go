package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestTestingPackage(t *testing.T) {
	t.Skip("skipping~") // 실패일 경우에 넣으면 안됨
	t.Run("A=1", func(t *testing.T) { t.Log("it is A") })
	t.Run("A=2", func(t *testing.T) { t.Log("it is A") })
	t.Run("B=1", func(t *testing.T) { t.Log("it is B") })
}

/*
	go test -run ''        # 모든 테스트를 실행합니다.
	go test -run TestingPackage       # TestTestingPackage 에서의 [TestingPackage]와 일치하는 테스트 실행
	go test -run TestingPackage/A=    # [TestingPackage] "A ="와 일치하는 하위 테스트 실행
	go test -run /A=1      # 모든 최상위 테스트중에서 [A = 1]하고 일치하는 하위 테스트 실행
	go test -fuzz FuzzSilheom  # [FuzzSilheom]과 일치하는 Fuzz 대상
*/

// 루프가 다 돌때까지 종료되지 않음
func BenchmarkLoop(b *testing.B) {
	// 병렬 등의 성능 테스트 시 cpu 샤용량 등의 기능
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
		}
	})
}

// output에 정의된 출력으로 나와야 정상
func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// 안녕하세요.
	// goodbye
}

func FuzzHex(f *testing.F) {
	for _, seed := range [][]byte{{}, {0}, {9}, {0xa}, {0xf}, {1, 2, 3, 4}} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, in []byte) {
		enc := hex.EncodeToString(in)
		out, err := hex.DecodeString(enc)
		if err != nil {
			t.Fatalf("%v: decode: %v", in, err)
		}
		if !bytes.Equal(in, out) {
			t.Fatalf("%v: not equal after round trip: %v", in, out)
		}
	})
}

func TestGroupedParallel(t *testing.T) {
	for _, tc := range []struct{ Name string }{} {
		tc := tc // 범위 변수 캡처
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			// ...
		})
	}

	// Run은 하위 병렬 테스트가 완료 될 때까지 종료되지 않음!
	t.Run("group", func(t *testing.T) {
		t.Run("Test1", 나는병렬테스트)
		t.Run("Test2", 나는병렬테스트)
		t.Run("Test3", 나는병렬테스트)
	})
	// 해체 해야됨
}

func 나는병렬테스트(t *testing.T) {}
