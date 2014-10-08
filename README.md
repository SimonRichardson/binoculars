binoculars
==========

Monaidc Lenses for data

Lenses are composable, immutable getters and setters. Composable in
that they allow updating of nested data structures. Immutable in that
the setters return copies of the whole data structure.

## Examples

### Nested updating

```go
type Person struct {
    Name     string
    Location Location
}
type Location struct {
    Number   int
    Street   string
    Postcode int
}

person := Person{
    Name: "Joe Smith",
    Location: Location{
        Number: 1006,
        Street: "Pearl St",
        Postcode: 80302,
    }
}

locationLens = binoculars.ObjectLens('Location'),
numberLens = binoculars.ObjectLens('Number'),
store = locationLens.AndThen(numberLens).Run(person);

console.log(store.Get());
// 1006

console.log(store.Set(1007));
// { Name: 'Joe Smith',
//   Location: { Number: 1007, Street: 'Pearl St', Postcode: 80302, },
// }
```
