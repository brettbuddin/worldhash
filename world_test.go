package worldhash

import (
    "testing"
)

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

func buildWorld() (*World, []*person) {
    world  := New(10, 10, 2)

    people := []*person{
        &person{"Jerry", 3, 2, 1}, 
        &person{"Newman", 9, 9, 1},
        &person{"Kramer", 1, 1, 1},
    }

    for i := range people {
        world.Add(people[i])
    }

    return world, people
}

func markSeen(world *World, people []*person) []bool {
    seen := make([]bool, len(people))

    for _, slice := range world.Objects {
        for i := range slice {
            for j := range people {
                if slice[i] == people[j] {
                    seen[j] = true
                }
            }
        }
    }

    return seen
}

func TestAdding(t *testing.T) {
    world, people := buildWorld()
    seen := markSeen(world, people)

    for i, seenPerson := range seen {
        if !seenPerson {
            t.Errorf("%s added but not seen in objects map", people[i])
        }
    }
}

func TestRemoving(t *testing.T) {
    world, people := buildWorld()

    for i := range people {
        world.Remove(people[i])
    }

    seen := markSeen(world, people)

    for i, seenPerson := range seen {
        if seenPerson {
            t.Errorf("%s removed but was seen in objects map", people[i])
        }
    }
}

func TestNearby(t *testing.T) {
    world, people := buildWorld()

    nearby := world.Nearby(people[0])
    interested := people[2]

    found := false
    for i := range nearby {
        if interested == nearby[i] {
            found = true 
        }
    }

    if !found {
        t.Errorf("%s should be nearby %s, but isn't", people[2], people[0])
    }
}
