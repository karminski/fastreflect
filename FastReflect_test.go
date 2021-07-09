package fastreflect

import (
	"reflect"
	"testing"
)



type Rabbit struct {
	Id        int  
	Name      string
	Weight    float64
	IsRabbit  bool
	Location  Address
}

type Address struct {
	Name  string
}



func TestStructFieldByName(t *testing.T) {
	Id 		 := 12001
	Name     := "token 01\n"
	Weight   := 12.43
	IsRabbit := false
	Location := Address{"b content.\t"}
	rabbit   := Rabbit{Id, Name, Weight, IsRabbit, Location}
	rabbitI  := reflect.ValueOf(rabbit).Interface()
	// check resolve result
	resolvedId        := StructFieldByName(rabbitI, "Id")
	resolvedName      := StructFieldByName(rabbitI, "Name")
	resolvedWeight    := StructFieldByName(rabbitI, "Weight")
	resolvedIsRabbit  := StructFieldByName(rabbitI, "IsRabbit")
	resolvedLocation  := StructFieldByName(rabbitI, "Location")

	if resolvedId.(int) != Id {
		t.Errorf("got %v, want %d, nil", resolvedId, Id)
	}
	if resolvedName.(string) != Name {
		t.Errorf("got %v, want %s, nil", resolvedName, Name)
	}
	if resolvedWeight.(float64) != Weight {
		t.Errorf("got %v, want %f, nil", resolvedWeight, Weight)
	}
	if resolvedIsRabbit.(bool) != IsRabbit {
		t.Errorf("got %v, want %v, nil", resolvedIsRabbit, IsRabbit)
	}
	if resolvedLocation.(Address).Name != Location.Name {
		t.Errorf("got %v, want %v, nil", resolvedLocation.(Address).Name, Location.Name)
	}
}


func TestResolveSliceAllElements(t *testing.T) {
	SliceExample   := []int{9,8,7,6,5,4,3,2,1}
	SliceExampleI  := reflect.ValueOf(SliceExample).Interface()
	// check resolve result
	elem := SliceAllElements(SliceExampleI)
	i := 9
	for _, v := range elem {
		if v.(int) != i {
			t.Errorf("got %v, want %d, nil", v.(int), i)
		}
		i --
	}
}



func BenchmarkStructFieldByName(b *testing.B) {
	// prepare data
	Id 		 := 12001
	Name     := "token 01\n"
	Weight   := 12.43
	IsRabbit := false
	Location := Address{"b content.\t"}
	rabbit   := Rabbit{Id, Name, Weight, IsRabbit, Location}
	rabbitI  := reflect.ValueOf(rabbit).Interface()

	// let's rock
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		StructFieldByName(rabbitI, "Id")
		StructFieldByName(rabbitI, "Name")
		StructFieldByName(rabbitI, "Weight")
		StructFieldByName(rabbitI, "IsRabbit")
		StructFieldByName(rabbitI, "Location")
	}
}

func BenchmarkInternalFieldByName(b *testing.B) {
	// prepare data
	Id 		 := 12001
	Name     := "token 01\n"
	Weight   := 12.43
	IsRabbit := false
	Location := Address{"b content.\t"}
	rabbit   := Rabbit{Id, Name, Weight, IsRabbit, Location}
	rabbitI  := reflect.ValueOf(rabbit).Interface()

	// let's rock
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r := reflect.ValueOf(rabbitI)
    	reflect.Indirect(r).FieldByName("ID")
    	reflect.Indirect(r).FieldByName("Name")
    	reflect.Indirect(r).FieldByName("Weight")
    	reflect.Indirect(r).FieldByName("IsRabbit")
    	reflect.Indirect(r).FieldByName("Location")
	}
}


func BenchmarkSliceAllElements(b *testing.B) {
	SliceExample   := []int{9,8,7,6,5,4,3,2,1}
	SliceExampleI  := reflect.ValueOf(SliceExample).Interface()
	elem := make([]interface{}, 9)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		elem = SliceAllElements(SliceExampleI)
		if len(elem) != 9 {
			b.Errorf("BenchmarkSliceAllElements: failed")
		}
	}
}


func BenchmarkInternalIndex(b *testing.B) {
	SliceExample   := []int{9,8,7,6,5,4,3,2,1}
	SliceExampleI  := reflect.ValueOf(SliceExample).Interface()
	elem := make([]interface{}, 9)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s := reflect.ValueOf(SliceExampleI)
		slen := s.Len()
        for i := 0; i < slen; i++ {
            elem[i] = s.Index(i)
        }
		if len(elem) != 9 {
			b.Errorf("BenchmarkInternalIndex: failed")
		}
	}

}