# Test bolt db as a NoSQL

Check the ability of the NoSQL in bolt db
You can change `schema` of the database while running

## Related link

- BoltDB GitHub: https://github.com/boltdb/bolt
- BBoltDB GitHub: https://github.com/etcd-io/bbolt
- BoltHold GitHub: https://github.com/timshannon/bolthold
- Blog (in Korean):

## 1. bolt1.go

- Generate one instance and save it to db
- Read and print all the records

User structure of this moment

```
type User struct {
	ID int `boltholdKey:"ID"`
	// Group string
	// Email string
	Name string
	Age  int
}
```

## 2. bolt2.go

- Add `Email` field
- Generate one instance and save it to db
- Read and print all the records

User structure of this moment

```
type User struct {
	ID int `boltholdKey:"ID"`
	// Group string
	Email string
	Name  string
	Age   int
}
```

## 3. bolt3.go

- Delete `Name` field
- Generate one instance and save it to db
- Read and print all the records

User structure of this moment

```
type User struct {
	ID int `boltholdKey:"ID"`
	// Group string
	Email string
	// Name  string
	Age int
}
```

## 4. bolt4.go

- Let's restore `Name` field
- Generate one instance and save it to db
- Read and print all the records

User structure of this moment

```
type User struct {
	ID int `boltholdKey:"ID"`
	// Group string
	Email string
	Name  string
	Age   int
}
```

## Run the codes sequentially

```

$go run bolt1.go
([]main.User) (len=1 cap=1) {
 (main.User) {
  ID: (int) 1,
  Name: (string) (len=8) "John Doe",
  Age: (int) 21
 }
}

$go run bolt2.go
([]main.User) (len=2 cap=2) {
 (main.User) {
  ID: (int) 1,
  Email: (string) "",
  Name: (string) (len=8) "John Doe",
  Age: (int) 21
 },
 (main.User) {
  ID: (int) 2,
  Email: (string) (len=13) "aaa@gmail.com",
  Name: (string) (len=8) "John Doe",
  Age: (int) 21
 }
}

$go run bolt3.go
([]main.User) (len=3 cap=4) {
 (main.User) {
  ID: (int) 1,
  Email: (string) "",
  Age: (int) 21
 },
 (main.User) {
  ID: (int) 2,
  Email: (string) (len=13) "aaa@gmail.com",
  Age: (int) 21
 },
 (main.User) {
  ID: (int) 3,
  Email: (string) (len=13) "bbb@gmail.com",
  Age: (int) 33
 }
}

$go run bolt4.go
([]main.User) (len=3 cap=4) {
 (main.User) {
  ID: (int) 1,
  Email: (string) "",
  Name: (string) (len=8) "John Doe",
  Age: (int) 21
 },
 (main.User) {
  ID: (int) 2,
  Email: (string) (len=13) "aaa@gmail.com",
  Name: (string) (len=8) "John Doe",
  Age: (int) 21
 },
 (main.User) {
  ID: (int) 3,
  Email: (string) (len=13) "bbb@gmail.com",
  Name: (string) "",
  Age: (int) 33
 }
}
```
