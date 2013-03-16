Simple implementation of spatial hashing.

```go
type person struct {
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
    world   := &World{1000, 1000, 20}
    person1 := &person{100, 50, 5}
    person2 := &person{90, 45, 5}

    world.Add(person1)
    world.Add(person2)

    world.Nearby(person1)
    // => []person{person2}

    world.Remove(person1)
    world.Remove(person2)
}

```
