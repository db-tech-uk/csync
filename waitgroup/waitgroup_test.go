package waitgroup

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	wg := New()

	wg.Add(1)

	go func() {
		doWork()
		wg.Done()
	}()

	wg.Wait()
}

func TestChain(t *testing.T) {
	wg := New()

	wg.Add(1)

	go func() {
		doWork()
		wg.Done()
	}()

	wg.Wait()
	wg.Add(1)

	go func() {
		doWork()
		wg.Done()
	}()

	wg.Wait()
}

func TestChainAlt(t *testing.T) {
	wg := New()

	wg.Add(1)

	go func() {
		doWork()
		wg.Done()
	}()

	wg.Wait()
	wg.Add(1)

	go func() {
		doWork()
		wg.Done()
	}()

	wg.Wait()
}

func TestChainAltAlt(t *testing.T) {
	wg := New()

	wg.Add(1)
	wg.Done()

	wg.Add(1)
	wg.Done()

	wg.Add(1)
	wg.Done()
}

func TestChainLoop(t *testing.T) {
	wg := New()

	for i := 0; i <= 1000; i++ {
		wg.Add(1)
		wg.Done()
	}
}

func TestAddMulti(t *testing.T) {
	wg := New()

	wg.Add(2)

	go func() {
		doWork()
		wg.Done()
	}()

	go func() {
		doWork()
		wg.Done()
	}()

	wg.Wait()
}

func TestEmptyChannelUnblocksOnWait(t *testing.T) {
	wg := New()

	wg.Wait()
}

func BenchmarkWaitgroup(b *testing.B) {
	wg := New()

	for i := 0; i < b.N; i++ {
		wg.Add(1)

		go func() {
			wg.Done()
		}()

		wg.Wait()
	}
}

func doWork() {
	time.Sleep(time.Millisecond * 100)
}
