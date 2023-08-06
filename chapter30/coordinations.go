package main

import (
	"chapter30/utils"
	"context"
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

func coordination1_1(iterationNumber int, counter *int) {
	for i := 0; i < iterationNumber; i++ {
		*counter++
	}
}

func Coordination1() {
	utils.Printfln("Coordination1()")

	counter := 0
	coordination1_1(5000, &counter)
	utils.Printfln("Counter: %v", counter)
}

func coordination2_2(iterationNumber int, counter *int) {
	for i := 0; i < iterationNumber; i++ {
		*counter++
	}
}

func Coordination2() {
	utils.Printfln("Coordination2()")

	counter := 0
	go coordination2_2(5000, &counter)
	utils.Printfln("Counter: %v", counter)
}

func coordination3_3(iterationNumber int, counter *int, waitGroup *sync.WaitGroup) {
	for i := 0; i < iterationNumber; i++ {
		*counter++
	}

	waitGroup.Done()
}

func Coordination3() {
	utils.Printfln("Coordination3()")

	waitGroup := sync.WaitGroup{}

	counter := 0
	go coordination3_3(5000, &counter, &waitGroup)
	waitGroup.Add(1)
	waitGroup.Wait()
	utils.Printfln("Counter: %v", counter)
}

func coordination4_4(iterationNumber int, counter *int, waitGroup *sync.WaitGroup) {
	time.Sleep(time.Second)

	for i := 0; i < iterationNumber; i++ {
		*counter++
	}

	waitGroup.Done()
}

func Coordination4() {
	utils.Printfln("Coordination4()")

	waitGroup := sync.WaitGroup{}
	counter := 0
	threadNum := 3
	waitGroup.Add(threadNum)

	for i := 0; i < threadNum; i++ {
		go coordination4_4(5000, &counter, &waitGroup)
	}

	waitGroup.Wait()
	utils.Printfln("Reading common data without synchronisation mechanism")
	utils.Printfln("Counter: %v", counter)
}

func coordination5_5(iterationNumber int, counter *int, waitGroup *sync.WaitGroup, mutex *sync.Mutex) {
	time.Sleep(time.Second)

	for i := 0; i < iterationNumber; i++ {
		mutex.Lock()
		*counter++
		mutex.Unlock()
	}

	waitGroup.Done()
}

func Coordination5() {
	utils.Printfln("Coordination5()")

	waitGroup := sync.WaitGroup{}
	mutex := sync.Mutex{}
	counter := 0
	threadNum := 3
	waitGroup.Add(threadNum)

	for i := 0; i < threadNum; i++ {
		go coordination5_5(5000, &counter, &waitGroup, &mutex)
	}

	waitGroup.Wait()
	utils.Printfln("Reading common data with synchronisation mechanism")
	utils.Printfln("Counter: %v", counter)
}

func coordination6_6(iterationsNumber int, max int, squaresCache *map[int]int, waitGroup *sync.WaitGroup, rwMutex *sync.RWMutex) {
	for i := 0; i < iterationsNumber; i++ {
		number := rand.Intn(max)
		rwMutex.RLock()
		cachedSquare, cachedSquareExists := (*squaresCache)[number]
		rwMutex.RUnlock()

		if cachedSquareExists {
			utils.Printfln("Cached value for square of (%v) is %v", number, cachedSquare)
		} else {
			rwMutex.Lock()
			if cachedSquare, cachedSquareExists = (*squaresCache)[number]; !cachedSquareExists {
				(*squaresCache)[number] = int(math.Pow(float64(number), 2))
				utils.Printfln("Calculated value for square of (%v) is %v", number, (*squaresCache)[number])
			} else {
				utils.Printfln("Cached value for square of (%v) is %v", number, cachedSquare)
			}
			rwMutex.Unlock()
		}
	}

	waitGroup.Done()
}

func Coordination6() {
	utils.Printfln("Coordination6()")

	rand.Seed(time.Now().UnixNano())
	squares := make(map[int]int, 10)
	waitGroup := sync.WaitGroup{}
	rwMutex := sync.RWMutex{}
	threadNum := 3
	waitGroup.Add(threadNum)
	for i := 0; i < threadNum; i++ {
		go coordination6_6(5, 10, &squares, &waitGroup, &rwMutex)
	}
	waitGroup.Wait()
}

func coordination7_7(max int, squaresCache *map[int]int, waitGroup *sync.WaitGroup, rwMutex *sync.RWMutex, mutexCondition *sync.Cond) {
	utils.Printfln("Generating squares...")

	rwMutex.Lock()
	for i := 0; i < max; i++ {
		(*squaresCache)[i] = int(math.Pow(float64(i), 2))
	}
	rwMutex.Unlock()

	utils.Printfln("Done with squares generation. Broadcasting it... ")
	mutexCondition.Broadcast()
	waitGroup.Done()
}

func coordination7_7_7(id int, iterationsNumber int, max int, squaresCache *map[int]int, waitGroup *sync.WaitGroup, mutexCondition *sync.Cond) {
	mutexCondition.L.Lock()
	if len(*squaresCache) == 0 {
		mutexCondition.Wait()
	}

	for i := 0; i < iterationsNumber; i++ {
		number := rand.Intn(max)
		cachedSquare, _ := (*squaresCache)[number]
		utils.Printfln("#%v Cached value for square of (%v) is %v", id, number, cachedSquare)
		time.Sleep(time.Millisecond * 100)
	}

	mutexCondition.L.Unlock()
	waitGroup.Done()
}

func Coordination7() {
	utils.Printfln("Coordination7()")

	rand.Seed(time.Now().UnixNano())
	squares := make(map[int]int, 10)
	waitGroup := sync.WaitGroup{}
	rwMutex := sync.RWMutex{}
	mutexCondition := sync.NewCond(rwMutex.RLocker())

	readThreadNum := 2
	writeThreadNum := 1

	waitGroup.Add(readThreadNum)
	for i := 0; i < readThreadNum; i++ {
		go coordination7_7_7(i, 5, 10, &squares, &waitGroup, mutexCondition)
	}

	waitGroup.Add(writeThreadNum)
	for i := 0; i < writeThreadNum; i++ {
		go coordination7_7(10, &squares, &waitGroup, &rwMutex, mutexCondition)
	}

	waitGroup.Wait()
}

func coordination8_8(max int, squaresCache *map[int]int, waitGroup *sync.WaitGroup, rwMutex *sync.RWMutex, mutexCondition *sync.Cond) {
	utils.Printfln("Generating squares...")

	rwMutex.Lock()
	for i := 0; i < max; i++ {
		(*squaresCache)[i] = int(math.Pow(float64(i), 2))
	}
	rwMutex.Unlock()

	utils.Printfln("Done with squares generation. Broadcasting it... ")
	mutexCondition.Broadcast()
	waitGroup.Done()
}

func coordination8_8_8(id int, iterationsNumber int, max int, squaresCache *map[int]int, waitGroup *sync.WaitGroup, mutexCondition *sync.Cond) {
	mutexCondition.L.Lock()
	for len(*squaresCache) == 0 {
		mutexCondition.Wait()
	}

	for i := 0; i < iterationsNumber; i++ {
		number := rand.Intn(max)
		cachedSquare, _ := (*squaresCache)[number]
		utils.Printfln("#%v Cached value for square of (%v) is %v", id, number, cachedSquare)
		time.Sleep(time.Millisecond * 100)
	}

	mutexCondition.L.Unlock()
	waitGroup.Done()
}

func Coordination8() {
	utils.Printfln("Coordination8()")

	rand.Seed(time.Now().UnixNano())
	squares := make(map[int]int, 10)
	waitGroup := sync.WaitGroup{}
	rwMutex := sync.RWMutex{}
	mutexCondition := sync.NewCond(&rwMutex)

	readThreadNum := 2
	writeThreadNum := 1

	waitGroup.Add(readThreadNum)
	for i := 0; i < readThreadNum; i++ {
		go coordination8_8_8(i, 5, 10, &squares, &waitGroup, mutexCondition)
	}

	waitGroup.Add(writeThreadNum)
	for i := 0; i < writeThreadNum; i++ {
		go coordination8_8(10, &squares, &waitGroup, &rwMutex, mutexCondition)
	}

	waitGroup.Wait()
}

func coordination9_9(max int, squaresCache *map[int]int) {
	utils.Printfln("Generating squares...")

	for i := 0; i < max; i++ {
		(*squaresCache)[i] = int(math.Pow(float64(i), 2))
	}

	utils.Printfln("Done with squares generation. ")
}

func coordination9_9_9(id int, iterationsNumber int, max int, squaresCache *map[int]int, waitGroup *sync.WaitGroup, once *sync.Once) {
	once.Do(func() {
		coordination9_9(max, squaresCache)
	})

	for i := 0; i < iterationsNumber; i++ {
		number := rand.Intn(max)
		cachedSquare, _ := (*squaresCache)[number]
		utils.Printfln("#%v Cached value for square of (%v) is %v", id, number, cachedSquare)
		time.Sleep(time.Millisecond * 100)
	}

	waitGroup.Done()
}

func Coordination9() {
	utils.Printfln("Coordination9()")

	rand.Seed(time.Now().UnixNano())
	squares := make(map[int]int, 10)
	waitGroup := sync.WaitGroup{}
	once := sync.Once{}

	readThreadNum := 2
	waitGroup.Add(readThreadNum)
	for i := 0; i < readThreadNum; i++ {
		go coordination9_9_9(i, 5, 10, &squares, &waitGroup, &once)
	}
	waitGroup.Wait()
}

func coordination10_10(waitGroup *sync.WaitGroup, count int) {
	total := 0

	for i := 0; i < count; i++ {
		utils.Printfln("Processing request: %v", total)
		total++
		time.Sleep(time.Millisecond * 250)
	}

	utils.Printfln("Request processed: %v", total)
	waitGroup.Done()
}

func Coordination10() {
	utils.Printfln("Coordination10()")

	waitGroup := sync.WaitGroup{}
	requestCount := 1
	waitGroup.Add(requestCount)

	utils.Printfln("Request dispatched...")
	for i := 0; i < requestCount; i++ {
		go coordination10_10(&waitGroup, 10)
	}

	waitGroup.Wait()
}

func coordination11_11(ctx context.Context, waitGroup *sync.WaitGroup, count int) {
	total := 0

	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			utils.Printfln("Stopping processing - request cancelled.")
			goto end
		default:
			utils.Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
	}

	utils.Printfln("Request processed: %v", total)
end:
	waitGroup.Done()
}

func Coordination11() {
	utils.Printfln("Coordination11()")

	waitGroup := sync.WaitGroup{}
	requestCount := 1
	waitGroup.Add(requestCount)

	ctx, cancel := context.WithCancel(context.Background())

	utils.Printfln("Request dispatched...")
	go coordination11_11(ctx, &waitGroup, 10)

	time.Sleep(time.Second)
	utils.Printfln("Cancelling request")
	cancel()

	waitGroup.Wait()
}

func coordination12_12(ctx context.Context, waitGroup *sync.WaitGroup, count int) {
	total := 0

	deadline, ok := ctx.Deadline()

	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			utils.Printfln("Stopping processing - request cancelled.")
			goto end
		default:
			utils.Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}

		if ok && time.Now().Compare(deadline) > 0 {
			utils.Printfln("Stopping processing - deadline exceeded.")
			goto end
		}
	}

	utils.Printfln("Request processed: %v", total)
end:
	waitGroup.Done()
}

func Coordination12() {
	utils.Printfln("Coordination12()")

	waitGroup := sync.WaitGroup{}
	requestCount := 1
	waitGroup.Add(requestCount)

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second))

	utils.Printfln("Request dispatched...")
	go coordination12_12(ctx, &waitGroup, 10)

	waitGroup.Wait()
}

func coordination13_13(ctx context.Context, waitGroup *sync.WaitGroup, count int) {
	total := 0

	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.Canceled) {
				utils.Printfln("Stopping processing - request cancelled.")
			} else {
				utils.Printfln("Stopping processing - deadline exceeded.")
			}

			goto end
		default:
			utils.Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
	}

	utils.Printfln("Request processed: %v", total)
end:
	waitGroup.Done()
}

func Coordination13() {
	utils.Printfln("Coordination13()")

	waitGroup := sync.WaitGroup{}
	requestCount := 1
	waitGroup.Add(requestCount)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	utils.Printfln("Request dispatched...")
	go coordination13_13(ctx, &waitGroup, 10)

	waitGroup.Wait()
}

func coordination14_14(ctx context.Context, waitGroup *sync.WaitGroup) {
	total := 0

	count := ctx.Value("count").(int)
	sleepDuration := ctx.Value("sleepDuration").(time.Duration)

	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.Canceled) {
				utils.Printfln("Stopping processing - request cancelled.")
			} else {
				utils.Printfln("Stopping processing - deadline exceeded.")
			}

			goto end
		default:
			utils.Printfln("Processing request: %v", total)
			total++
			time.Sleep(sleepDuration)
		}
	}

	utils.Printfln("Request processed: %v", total)
end:
	waitGroup.Done()
}

func Coordination14() {
	utils.Printfln("Coordination14()")

	waitGroup := sync.WaitGroup{}
	requestCount := 1
	waitGroup.Add(requestCount)

	ctx := context.WithValue(context.Background(), "count", 7)
	ctx = context.WithValue(ctx, "sleepDuration", time.Millisecond*150)

	utils.Printfln("Request dispatched...")
	go coordination14_14(ctx, &waitGroup)

	waitGroup.Wait()
}
