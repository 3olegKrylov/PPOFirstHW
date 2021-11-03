package test

import (
	"PPO/internal/models"
	"testing"
)

func TestCheckCap(t *testing.T) {
	lruc, err := models.CreatLRUCache(3)
	if err != nil {
		t.Error(err)
	}

	if lruc.Capacity != 3 {
		t.Error("capacity is ", lruc.Capacity, "but should to be 3")
	}
}

func TestAddElements(t *testing.T) {
	lruc, err := models.CreatLRUCache(3)
	if err != nil {
		t.Error(err)
	}
	if lruc.Get("1") != nil {
		t.Error("expectation is nil result is ")
	}

}

func TestSetLru(t *testing.T) {
	lruc, err := models.CreatLRUCache(3)
	if err != nil {
		t.Error(err)
	}

	key := "key"

	for i := 0; i < 100; i++ {
		lruc.Set(key, i)

		if lruc.Get(key) != i {
			panic("error: TestSetLru")
		}
	}
}

func TestCapacity(t *testing.T) {
	_, err := models.CreatLRUCache(-4)
	if err == nil {
		t.Error("error: TestCapacity, have tocapacity >0")
	}
}

func TestTypeLRUC(t *testing.T) {
	testTable := []struct {
		key     string
		value   interface{}
		exepted interface{}
	}{
		{"string", "string", "string"},
		{"integer", 1, "int"},
		{"float", 1.4, "float"},
	}

	lruc, err := models.CreatLRUCache(3)
	if err != nil {
		t.Error(err)
	}

	for _, testCase := range testTable {
		lruc.Set(testCase.key, testCase.value)
	}

	for _, testCase := range testTable {
		if lruc.Get(testCase.key) != testCase.value {
			panic("excepted value is not true TestTypeLRUC")
		}

		switch _ := testCase.value.(type) {
		case string:
			if testCase.exepted != "string" {
				panic("TestTypeLRUC expected type string")
			}
			break
		case int:
			if testCase.exepted != "int" {
				panic("TestTypeLRUC expected type int")
			}
			break
		case float64:
			if testCase.exepted != "float64" {
				panic("TestTypeLRUC expected type float64")
			}
			break
		case []int:
			if testCase.exepted != "[]int" {
				panic("TestTypeLRUC expected type []int")
			}
			break
		case interface{}:
			panic("TestTypeLRUC expected type interface")
		default:

		}

	}

}
