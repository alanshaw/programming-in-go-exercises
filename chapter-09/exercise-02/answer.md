You'd use a embedded anonymous field to represent an "is-a" relationship, for situations when a type has a "super" type that is shared between multiple other types. For example a `Car` is-a `Vehicle`, and a `Moped` is also a `Vehicle`.

The embedded anonymous field allows the type to inherit the methods of the super type allowing them to be called directly on instances of the type as if they belonged to it.

For example:

```go
type Vehicle struct {
    Name string
}

func (v *Vehicle) Move() {
    fmt.Println("Moving", v.Name)
}

type Car struct {
    Vehicle
}

func main () {
    car := new(Car)
    car.Name = "Brum"
    car.Move() // or car.Vehicle.Move()
}
```