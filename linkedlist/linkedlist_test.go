package linkedlist

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestLinkedListCreation(t *testing.T) {
	expected := []int{33, 123, 325, 544}

	list := NewLinkedList(33, 123, 325, 544)

	cnt := 0
	iterator := list.GetIterator()
	for item := iterator.Next(); item != nil; item = iterator.Next() {
		if expected[cnt] != item.Data.(int) {
			t.Fatalf("Expected %d to be equal to %d", item.Data.(int), expected[cnt])
		}
		cnt++
	}
}
func TestEmptyLinkedListCreation(t *testing.T) {
	list := NewLinkedList()

	iterator := list.GetIterator()
	for item := iterator.Next(); item != nil; item = iterator.Next() {
		t.Fatal("Should not iterate because list was empty")
	}
}
func TestLinkedListInsert(t *testing.T) {
	expected := []int{33, 123, 325, 544}

	list := &LinkedList{}
	list.Insert(33, 123, 325, 544)

	cnt := 0
	iterator := list.GetIterator()
	for item := iterator.Next(); item != nil; item = iterator.Next() {
		if expected[cnt] != item.Data.(int) {
			t.Fatalf("Expected %d to be equal to %d", item.Data.(int), expected[cnt])
		}
		cnt++
	}
}
func TestLinkedListRemoveMiddle(t *testing.T) {
	expected := []int{33, 325, 544}

	list := &LinkedList{}
	list.Insert(33)
	list.Insert(123)
	list.Insert(325)
	list.Insert(544)
	middle := list.head.next
	list.Remove(middle)

	cnt := 0
	iterator := list.GetIterator()
	for item := iterator.Next(); item != nil; item = iterator.Next() {
		val := item.Data.(int)
		if expected[cnt] != val {
			t.Fatalf("Expected value to be %d was %d", expected[cnt], val)
		}
		cnt++
	}
}
func TestLinkedListRemoveHead(t *testing.T) {
	expected := []int{123, 325, 544}

	list := &LinkedList{}
	list.Insert(33)
	list.Insert(123)
	list.Insert(325)
	list.Insert(544)

	list.Remove(list.Head())
	cnt := 0
	iterator := list.GetIterator()
	for item := iterator.Next(); item != nil; item = iterator.Next() {
		val := item.Data.(int)
		if expected[cnt] != val {
			t.Fatalf("Expected value to be %d was %d", expected[cnt], val)
		}
		cnt++
	}
}
func TestLinkedListRemoveTail(t *testing.T) {
	expected := []int{33, 123, 325}

	list := &LinkedList{}
	list.Insert(33)
	list.Insert(123)
	list.Insert(325)
	list.Insert(544)

	list.Remove(list.Tail())
	cnt := 0
	iterator := list.GetIterator()
	for item := iterator.Next(); item != nil; item = iterator.Next() {
		val := item.Data.(int)
		if expected[cnt] != val {
			t.Fatalf("Expected value to be %d was %d", expected[cnt], val)
		}
		cnt++
	}
}
func TestMarshalJSON(t *testing.T) {
	expected := []byte("[33,123,325,555]")
	data := []int{33, 123, 325, 555}

	list := &LinkedList{}

	for _, item := range data {
		list.Insert(item)
	}

	marshaledData, err := json.Marshal(list)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expected, marshaledData) {
		t.Fatalf("Expected %s to be equal to %s", string(marshaledData), string(expected))
	}
}
func TestMarshalJSONComplex(t *testing.T) {
	expected := []byte("[{\"asd\":\"asd123\"}]")
	list := &LinkedList{}

	type tempTestStruct struct {
		Asd string `json:"asd"`
	}
	list.Insert(&tempTestStruct{Asd: "asd123"})
	marshaledData, err := json.Marshal(list)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expected, marshaledData) {
		t.Fatalf("Expected %s to be equal to %s", string(marshaledData), string(expected))
	}
}
func TestUnmarshalJSON(t *testing.T) {
	jsonData := []byte("[33,65,76,132]")
	expected := &LinkedList{}
	expected.Insert(33)
	expected.Insert(65)
	expected.Insert(76)
	expected.Insert(132)

	unmarshalResult := &LinkedList{}

	if err := json.Unmarshal(jsonData, unmarshalResult); err != nil {
		t.Fatal(err)
	}
	iterr := unmarshalResult.GetIterator()
	expectedIterr := expected.GetIterator()

	for resultItem := iterr.Next(); resultItem != nil; resultItem = iterr.Next() {
		result := resultItem.Data.(float64)
		test := expectedIterr.Next().Data.(int)
		if test != int(result) {
			t.Fatalf("Expected %f to be equal to %d", result, test)
		}
	}
}
func TestUnmarshalJSONComplex(t *testing.T) {
	jsonData := []byte("[{\"asd\":\"abcdefghijklm\"}]")
	expected := &LinkedList{}
	expected.Insert(map[string]string{"asd": "abcdefghijklm"})

	unmarshalResult := &LinkedList{}

	if err := json.Unmarshal(jsonData, unmarshalResult); err != nil {
		t.Fatal(err)
	}
	iterr := unmarshalResult.GetIterator()
	expectedIterr := expected.GetIterator()

	for resultItem := iterr.Next(); resultItem != nil; resultItem = iterr.Next() {
		result := resultItem.Data.(map[string]interface{})
		test := expectedIterr.Next().Data.(map[string]string)
		if test["asd"] != result["asd"].(string) {
			t.Fatalf("Expected %s to be equal to %s", result, test)
		}
	}
}
