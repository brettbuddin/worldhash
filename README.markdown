Simple implementation of spatial hashing.

```go
type person struct {
    name string
    x int
    y int
    radius int
}

func (p *person) X() int {
    return p.x
}

func (p *person) Y() int {
    return p.y
}

func (p *person) Radius() int {
    return p.radius
}

func main() {
    world  := worldhash.NewWorld(1000, 1000, 20)
    jerry  := &person{"Jerry", 100, 50, 5}
    newman := &person{"Newman", 90, 45, 5}
    kramer := &person{"Kramer", 800, 800, 5}

    world.Add(jerry)
    world.Add(newman)
    world.Add(kramer)

    world.Nearby(jerry)
    // => []person{newman}

    world.Nearby(kramer)
    // => []person{}

    world.Remove(jerry)
    world.Remove(newman)
    world.Remove(kramer)
}

```
