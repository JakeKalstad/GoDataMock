# GoDataMock

Go library to mock data for testing so you don't have to

### Usage
    Create an array of the type you want to mock with a single instance of the type desired
    Interfaces can be left nil for a nil return or given an empty pointer to the concrete type desired
    Define how many instances you want back
    
### Example
```
type SubTestData struct {
	Key  int64
	Meta string
}

type Grouper interface {
}

type TestGroup struct {
	GroupKey  int64
	GroupName string
}

type TestData struct {
	Name  string
	Age   int
	Male  bool
	Info  SubTestData
	Group Grouper
}

func TestGen(t *testing.T) {
	gen := &Generator{}
	data := []TestData{TestData{Group: &TestGroup{}}}
	data = gen.Get(data, 10).([]TestData)
	fmt.Printf("%v+", data)
}

```
    
    
    
### Version
0.1

### Installation
    go get "github.com/JakeKalstad/GoDataMock"

### Todos

 - Write Tests
 - Rethink Github Save
 - Add Code Comments
 - Add Night Mode

License
----

MIT


**Free Software, Hell Yeah!**
