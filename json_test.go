package json

import (
	"testing"
	"reflect"
	"encoding/json"
)

type Group struct {
	Num         []int        `json:"num"`
	Right       bool         `json:"-,"`
	ID          int          `json:"id,omitempty"`
	Name        string       `json:"-"`
	Colors      []string     `json:",omitempty"`
	Dictionary  map[int]int  `json:"dictionary,omitempty"`
}

var group = Group{
	Num:	    []int{55, 8, -12},
	Name:       "Reds",
	Right:	    true,
	Colors:     []string{"Crimson", "<15", "Ruby"},
	Dictionary: map[int]int{16: 1, 75:100},
}

func TestJsonMarshal(t *testing.T) {
	bytes, err1 := JsonMarshal(group)
	expected_json, err2 := json.Marshal(group)	

	if (err1 != err2) {
		t.Errorf("failed to transfer from data to json")
	} else if string(bytes) != string(expected_json) {
		t.Errorf("failed to transfer from data to json")
	}
}

func TestGetKey(t *testing.T) {
	flags := make([]int, 6)
	keys := make([]string, 6)
	expected_flags := []int{1, 1, 0, 0, 0, 0}
	expected_keys := []string{"num", "-", "id", "", "Colors", "dictionary"}

	for i := 0; i < reflect.TypeOf(group).NumField(); i++ {
		t_field := reflect.TypeOf(group).Field(i)
		tagDetails := t_field.Tag.Get("json")		
	 	flags[i], keys[i] = getKey(t_field, tagDetails)

	 	if (flags[i] != expected_flags[i] || keys[i] != expected_keys[i]) {
	 		t.Errorf("failed to get keys")
	 	}
	}
}

func TestStringToJson(t *testing.T) {
	str := stringToJson("A<BB>C->&")
	expected := "A\\u003cBB\\u003eC-\\u003e\\u0026"

	if str != expected {
		t.Errorf("failed to transfer from string to json")
	}
}

func ExampleJsonMarshal() {
	type ColorGroup struct {
		Num         []int        `json:"num"`
		Right       bool         `json:"-,"`
		ID          int          `json:"id,omitempty"`
		Name        string       `json:"-"`
		Colors      []string     `json:",omitempty"`
		Dictionary  map[int]int  `json:"dictionary,omitempty"`
	}

	group := ColorGroup{
		Num:	    []int{55, 8, -12},
		Name:       "Reds",
		Right:	    true,
		Colors:     []string{"Crimson", "<15", "Ruby"},
		Dictionary: map[int]int{16: 1, 75:100},
	}

	b, err := Json.JsonMarshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	// Output: {"num":[55,8,-12],"-":true,"Colors":["Crimson","\u003c15","Ruby"],"dictionary":{"16":1,"75":100}}
}