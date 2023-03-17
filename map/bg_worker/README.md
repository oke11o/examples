# Map with and without races

Dont do like this
```go
func (s *serviceWithRace) get() string {
	v := s.d[rand.Intn(s.n)]
	return *v
}

func (s *serviceWithRace) update() {
	for _, v := range s.d {
		newVal := strconv.Itoa(rand.Int())
		*v = newVal
	}
}
```

Just run for compare
```shell
make run_with_race
make run_without_race
```