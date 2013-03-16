package worldhash

import (
    "math"
)

type Object interface {
    X() int
    Y() int
    Radius() int
}

type point struct {
    x int
    y int
}

type World struct {
    Width int
    Height int
    Subdivide int
    Objects map[int][]Object
}

// Creates a new World
func NewWorld(width int, height int, subdivide int) *World {
    return &World{
        Width: width,
        Height: height,
        Subdivide: subdivide,
        Objects: make(map[int][]Object),
    }
}

// Adds an Object to the objects map
func (w *World) Register(o Object) {
    ids := w.HashIds(o)
    for _, id := range ids {
        w.Objects[id] = append(w.Objects[id], o)
    }
}

// Removes references to Object from the Objects map
func (w *World) Remove(o Object) {
    ids := w.HashIds(o)

    for _, id := range ids {
        for j, other := range w.Objects[id] {
            if o == other {
                w.Objects[id] = append(w.Objects[id][:j],w.Objects[id][j+1:]...)
            }
        }
    }
}

// Returns a list of nearby Objects
func (w *World) Nearby(o Object) []Object {
    objects := []Object{}
    ids     := w.HashIds(o)

    _append := func(slice []Object, o Object) []Object {
        for _, other := range slice {
            if other == o {
                return slice
            }
        }

        return append(slice, o)
    }

    for _, id := range ids {
        for _, object := range w.Objects[id] {
            if object != o {
                objects = _append(objects, object)
            }
        }
    }

    return objects
}

// Returns the hash table IDs that an Object resides in
func (w *World) HashIds(o Object) []int {
    ids   := []int{}
    min   := point{o.X() - o.Radius(), o.Y() - o.Radius()}
    max   := point{o.X() + o.Radius(), o.Y() + o.Radius()}
    width := w.Width / w.Subdivide

    _append := func(slice []int, i int) []int {
        for _, other := range slice {
            if other == i {
                return slice
            }
        }

        return append(slice, i)
    }

    add := func(p point) {
        id := int(math.Floor(float64(p.x / w.Subdivide))) + 
              int(math.Floor(float64(p.y / w.Subdivide))) * width

        ids = _append(ids, id)
    }

    // make a list of all hash IDs that
    // are hit by the four corners of the
    // Object's bounding box
    add(point{min.x, max.y}) // top left
    add(point{max.x, max.y}) // top right
    add(point{max.x, min.y}) // bottom right
    add(min)                 // bottom left

    return ids
}
