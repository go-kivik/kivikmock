package kivikmock

import "reflect"

func meets(a, e expectation) bool {
	if reflect.TypeOf(a).Elem().Name() != reflect.TypeOf(e).Elem().Name() {
		return false
	}
	if !dbMeetsExpectation(a.DB(), e.DB()) {
		return false
	}
	return a.met(e)
}

func dbMeetsExpectation(a, e *MockDB) bool {
	if e == nil {
		return true
	}
	return e.name == a.name && e.id == a.id
}
