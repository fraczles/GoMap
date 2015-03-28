package main

import "fmt"

type emap struct{
    sharedMap map[string]int
    stop chan int
    readerIn chan string
    readerOut chan int
    writers chan string
    reduceRequest chan int
    reduceResponse chan int
    doneReducing chan int

    //Several channels
}

//Constructor?
func NewChannelMap() *emap {

    //Initialize map and channels
    em := emap{}
    em.sharedMap = make(map[string]int)
    em.stop = make(chan int)
    em.readerIn = make(chan string, ASK_BUFFER_SIZE)
    em.readerOut = make(chan int, ASK_BUFFER_SIZE)
    em.writers = make(chan string, ADD_BUFFER_SIZE)
    em.reduceRequest = make(chan int,1)
    em.reduceResponse = make(chan int,1)
    em.doneReducing = make(chan int,1)
    return &em
}

//Implement interface functions
func (em *emap) Listen() {
    for {
        select {

            case <-em.reduceRequest:
                em.reduceRequest <- 0
                <- em.doneReducing

            case str := <-em.writers:
                //------threads are writing here --------
                if count, exists := em.sharedMap[str]; exists {
                    //word exists in map
                    em.sharedMap[str] = count + 1
                } else{
                    //word doesn't exist in map
                    em.sharedMap[str] = 1
                }
            case str := <-em.readerIn: //threads want to read


                if count, exists := em.sharedMap[str]; exists {
                    //word exists in map
                    em.readerOut <- count

                } else{

                    //word doesn't exist in map
                    em.readerOut <- 1
                }

            case <-em.stop: //stoP
                return

        }
    }
}

func (em *emap) Stop() {
    em.stop <- 1
}
func (em *emap) Reduce(functor ReduceFunc, accum_str string, accum_int int) (string, int) {
    fmt.Println("============FROM REDUCE  1============")
    em.reduceRequest <- 1
    fmt.Println("============FROM REDUCE  1.25============")
    fmt.Println(<-em.reduceRequest)
    fmt.Println("============FROM REDUCE  1.5============")
    <-em.reduceResponse
    fmt.Println("============FROM REDUCE  2============")
    for k,v := range em.sharedMap{
        if accum_int == 0{
            accum_str, accum_int = k, v
        }
        accum_str, accum_int = functor(accum_str,accum_int,k,v)
    }
    em.doneReducing <- 1
    return accum_str, accum_int

}
func (em *emap) AddWord(word string) {
    em.writers <- word
}
func (em *emap) GetCount(word string) int {
    em.readerIn <- word
    count := <-em.readerOut
    return count
}

func min_word(key1 string, val1 int, key2 string, val2 int) (string,int) {
    if val1 > val2 {
        return key2,val2
    }
    return key1,val1
}
