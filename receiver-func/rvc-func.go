
//https://medium.com/@adityaa803/wth-is-a-go-receiver-function-84da20653ca2#:~:text=Unlike%20other%20languages%2C%20in%20go,which%20can%20receive%20a%20person.

package main
import "fmt"

type Space struct{
    occupied bool
}

type ParkingLot struct{
    spaces []Space
}

//normal function
func occupySpace(lot *ParkingLot, spaceNum int){
    lot.spaces[spaceNum - 1].occupied = true
}

//receiver func
func (lot *ParkingLot) occupySpace (spaceNum int) {
    lot.spaces[spaceNum - 1].occupied = true
}

func (lot *ParkingLot) vacateSpace (spaceNum int) {
    lot.spaces[spaceNum - 1].occupied = false
}

func main(){
    lot := ParkingLot{spaces: make([] Space, 4)}
    fmt.Println(lot)
    // receiver func
    lot.occupySpace(1)
    // normal func
    occupySpace(&lot, 2)
    fmt.Println("After occupyiong space",lot)

    lot.vacateSpace(2)
    fmt.Println("After vacating space",lot)

}