package set

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestHashSet(t *testing.T) {
	s := NewHashSet()
	s.Add("abc")
	s.Add(123)
	if !reflect.DeepEqual([]interface{}{"abc", 123}, s.Elements()) {
		t.Fatalf("Wrong set\nexp: %#v\n\ngot: %#v\n", []interface{}{"abc", 123}, s.Elements())
	}

	s2 := NewHashSet()
	s2.Add("abc")
	s2.Add(123)

	if !s.Same(s2) {
		t.Fatalf("Fail equal test\nexp: %v\n\ngot %v\n", s.Elements(), s2.Elements())
	}

	s3 := NewHashSet()
	if !s.IsSuperset(s3) {
		t.Fatalf("Fail superset test\nexp: %v\n\ngot %v\n", s.Elements(), s3.Elements())
	}

	if !s.Contains("abc") {
		t.Fatalf("Fail to test Contains\nexp: %v\n\ngot: %v\n", true, s.Contains("abc"))
	}

	if s.Len() != 2 {
		t.Fatalf("Wrong len\nexp: %v\n\ngot %v\n", 2, s.Len())
	}

	s.Remove(123)

	if s.Contains(123) {
		t.Fatalf("Fail to test Contains\nexp: %v\n\ngot: %v\n", false, s.Contains(123))
	}

	if s.Len() != 1 {
		t.Fatalf("Wrong len\nexp: %v\n\ngot %v\n", 1, s.Len())
	}

	s.Clear()

	if !reflect.DeepEqual([]interface{}{}, s.Elements()) {
		_, file, line, _ := runtime.Caller(0)
		t.Fatalf("Wrong set\nexp: %#v\n\ngot: %#v\n", filepath.Base(file), line, []interface{}{}, s.Elements())
	}
}

func BenchmarkAdd(b *testing.B) {
	s := NewHashSet()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
	b.ReportAllocs()
}

func BenchmarkElements(b *testing.B) {
	s := NewHashSet()
	for i := 0; i < b.N; i++ {
		s.Add(i)
		s.Elements()
	}
	b.ReportAllocs()
}

func BenchmarkSame(b *testing.B) {
	s1, s2 := NewHashSet(), NewHashSet()
	for i := 0; i < b.N; i++ {
		s1.Add(i)
		s2.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s1.Same(s2)
	}
	b.ReportAllocs()
}

func BenchmarkIsSuperset(b *testing.B) {
	s1, s2 := NewHashSet(), NewHashSet()
	for i := 0; i < b.N; i++ {
		s1.Add(i)
		s2.Add(i / 2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s1.IsSuperset(s2)
	}
	b.ReportAllocs()
}
