package main

//import "fmt"

type emap struct{
    sharedMap map[string]int
    stop chan bool
    readers chan int
    writers chan string
    //Several channels
}

//Constructor?
func NewChannelMap() *emap {

    //Initialize map and channels
    em := emap{}
    em.sharedMap = make(map[string]int)
    em.stop = make(chan bool)
    em.readers = make(chan int, ASK_BUFFER_SIZE)
    em.writers = make(chan string, ADD_BUFFER_SIZE)
    return &em
}

//Implement interface functions
func (em *emap) Listen() {
    for {
        select {
            case <-em.writers:
                //------threads are writing here --------
                str := <-em.writers
                if count, exists := em.sharedMap[str]; exists {
                    //word exists in map
                    em.sharedMap[str] = count + 1
                } else{
                    //word doesn't exist in map
                    em.sharedMap[str] = 1
                }

            case <-em.readers: //threads want to write


            case <-em.stop: //stop
                return
        }
    }
}

func (em *emap) Stop() {
    em.stop <- true
}
func (em *emap) Reduce(functor ReduceFunc, accum_str string, accum_int int) (string, int) {
    return " ", 0
}
func (em *emap) AddWord(word string) {
    em.writers <- word
}
func (em *emap) GetCount(word string) int {

    return 0
}
