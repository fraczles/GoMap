package main


type emap struct{
    sharedMap map[string]int
    stop chan bool
    readers chan string
    //Several channels
}

//Constructor?
func NewChannelMap() *emap {
    c := emap{}
    c.sharedMap = make(map[string]int)
    c.stop <- false
    return &c
}

//Implement interface functions
func (c *emap) Listen() {
    fmt.Println("Yo")
    // for {
    //     select {
    //     case <-c.stop: //stuff
    //         return
    //     }
    // }
}

func (c *emap) Stop() {
    c.stop <- true
}
func (c *emap) Reduce(functor ReduceFunc, accum_str string, accum_int int) (string, int) {
    return "", 0
}
func (c *emap) AddWord(word string) {
    // Word exists in map
    if num, exists := c.sharedMap[word]; exists {
        num = num + 1
    } else {
        c.sharedMap[word] = 1
    }
    
    

}
func (c *emap) GetCount(word string) int {
    return 0
}
