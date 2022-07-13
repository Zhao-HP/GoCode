package test

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"testing"
	"time"
)

func TestAssign(t *testing.T) {
	m1 := map[string]int{
		"K1": 1,
		"K2": 2,
		"K3": 3,
	}

	m2 := map[string]int{
		"K4": 4,
		"K5": 5,
		"K6": 6,
	}
	m3 := lo.Assign(m1, m2)
	for k := range m3 {
		fmt.Println(k, " --- ", m3[k])
	}
}

func TestAsync(t *testing.T) {
	c := lo.Async(func() string {
		time.Sleep(time.Second * 5)
		return "a"
	})

	select {
	case r := <-c:
		fmt.Println(r, "lo.Async.result")
	}
}

func TestAttempt(t *testing.T) {
	_, _ = lo.Attempt(10, func(i int) error {
		if i > 3 {
			return nil
		}
		fmt.Println("i = ", i)
		return errors.New("attempt error")
	})
}

func TestAttemptWithDelay(t *testing.T) {
	i1, i2, _ := lo.AttemptWithDelay(10, time.Second*2, func(i int, duration time.Duration) error {
		if i > 3 {
			return nil
		}
		fmt.Println("i = ", i)
		return errors.New("attemptWithDelay error")
	})

	fmt.Println("i1 = ", i1)
	fmt.Println("i2 = ", i2)
}

func TestChunk(t *testing.T) {
	ia := []int{1, 2, 3, 4, 5, 6, 7}

	iaa := lo.Chunk(ia, 2)
	for i := range iaa {
		for i2 := range iaa[i] {
			fmt.Println(iaa[i][i2])
		}
	}
}

func TestClamp(t *testing.T) {
	fmt.Println(lo.Clamp(1, 2, -1))
}

func TestCoalesce(t *testing.T) {

	r, _ := lo.Coalesce(0, 0, 1, 2, 3)
	fmt.Println(r)

}

func TestContainsBy(t *testing.T) {
	r := lo.ContainsBy([]int{1, 2, 3, 4, 5, 6, 7}, func(i int) bool {
		return i == 5
	})
	fmt.Println(r)
}

func TestDifference(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 5, 6, 7}
	a2 := []int{4, 5, 6, 7, 8, 9, 10}
	r1, r2 := lo.Difference(a1, a2)
	fmt.Println(r1)
	fmt.Println(r2)
}

func TestDrop(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r1 := lo.Drop(a1, 3)
	fmt.Println(r1)
}

func TestDropWhile(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r1 := lo.DropWhile(a1, func(i int) bool {
		return i >= 5
	})
	fmt.Println(r1)
}
