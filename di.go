package di

import "reflect"

type Storage struct {
	values map[string]interface{}
	TypeStorage
}

type TypeStorage struct {
	types map[reflect.Type]reflect.Value
	parent *TypeStorage
}

func MakeStorage() Storage {
	return Storage{
		TypeStorage: MakeTypeStorage(),
		values:      map[string]interface{}{},
	}
}

func (s *Storage) Invoke(k string) {
	Invoke(s.Get(k), s.TypeStorage)
}

func(s *Storage) Set(k string, v interface{}) {
	s.values[k] = v
}

func(s *Storage) Get(k string) interface{} {
	return s.values[k]
}

func MakeTypeStorage() TypeStorage {
	return TypeStorage{
		types: make(map[reflect.Type]reflect.Value),
	}
}

func(t *TypeStorage) Provide(v interface{}) {
	t.types[reflect.TypeOf(v)] = reflect.ValueOf(v)
}

func(t *TypeStorage) ProvideType(v interface{}, pt interface{}) {
	rpt := reflect.TypeOf(pt)
	if rpt.Kind() == reflect.Ptr {
		rpt = rpt.Elem()
	}
	if rpt.Kind() != reflect.Interface {
		panic("Error: ProvideType must provide pointer in second param")
	}
	t.types[rpt] = reflect.ValueOf(v)
}

func(t *TypeStorage) Get(k reflect.Type) reflect.Value {
	v := t.types[k]
	if v.IsValid() {
		return v
	}
	if t.parent != nil {
		t.parent.Get(k)
	}
	return reflect.ValueOf(nil)
}

func Invoke(f interface{},s TypeStorage) {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		panic("Error: Invoke only accept func type param")
	}
	params := make([]reflect.Value, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		argType := t.In(i)
		storageType := s.Get(argType)
		if storageType.IsValid() {
			params[i] = storageType
		} else {
			panic("Error: param type must be register firstly")
		}
	}
	fv := reflect.ValueOf(f)
	fv.Call(params)
}


