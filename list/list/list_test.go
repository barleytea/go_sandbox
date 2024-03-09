package list

import "testing"

func TestList_NewList(t *testing.T) {
	l := NewList[int]()
	if (l.head != nil) {
		t.Error("head should be nil")
	}
	if (l.tail != nil) {
		t.Error("tail should be nil")
	}
}

func TestList_Add(t *testing.T) {
	l := NewList[int]()
	l.Add(1)
	if (l.head.value != 1) {
		t.Error("head value should be 1")
	}
	if (l.tail.value != 1) {
		t.Error("tail value should be 1")
	}

	headAddress := &l.head

	l.Add(2)
	if (l.head.value != 1) {
		t.Error("head value should be 1")
	}
	if (l.tail.value != 2) {
		t.Error("tail value should be 2")
	}
	if (l.head.next.value != 2) {
		t.Error("head.next value should be 2")
	}
	if (l.tail.prev.value != 1) {
		t.Error("tail.prev value should be 1")
	}

	if (&l.head != headAddress) {
		t.Error("head address should not change")
	}
}

func TestGet(t *testing.T) {
	l := NewList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)

	if (l.Get(0) != 1) {
		t.Error("Get(0) should be 1")
	}
	if (l.Get(1) != 2) {
		t.Error("Get(1) should be 2")
	}
	if (l.Get(2) != 3) {
		t.Error("Get(2) should be 3")
	}
	if (l.Get(3) != 4) {
		t.Error("Get(3) should be 4")
	}
	if (l.Get(4) != 5) {
		t.Error("Get(4) should be 5")
	}
}

func TestFilter(t *testing.T) {
	l := NewList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)

	newList := l.Filter(func(x int) bool {
		return x % 2 == 0
	})

	if (newList.Get(0) != 2) {
		t.Error("newList.Get(0) should be 2")
	}
	if (newList.Get(1) != 4) {
		t.Error("newList.Get(1) should be 4")
	}
}

func TestMap(t *testing.T) {
	l := NewList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)

	newList := Map[int, int](l, func(x int) int {
		return x * 2
	})

	if (newList.Get(0) != 2) {
		t.Error("newList.Get(0) should be 2")
	}
	if (newList.Get(1) != 4) {
		t.Error("newList.Get(1) should be 4")
	}
	if (newList.Get(2) != 6) {
		t.Error("newList.Get(2) should be 6")
	}
	if (newList.Get(3) != 8) {
		t.Error("newList.Get(3) should be 8")
	}
	if (newList.Get(4) != 10) {
		t.Error("newList.Get(4) should be 10")
	}
}

func TestReduce(t *testing.T) {
	l := NewList[int]()
	l.Add(0)
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)

	result := l.Reduce(func(x int, y int) int {
		return x + y
	})

	if (result != 15) {
		t.Errorf("result should be 15. Got %d", result)
	}
}

func TestRemove(t *testing.T) {
	l := NewList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)

	l.Remove(2)

	if (l.Get(0) != 1) {
		t.Error("l.Get(0) should be 1")
	}
	if (l.Get(1) != 2) {
		t.Error("l.Get(1) should be 2")
	}
	if (l.Get(2) != 4) {
		t.Error("l.Get(2) should be 4")
	}
	if (l.Get(3) != 5) {
		t.Error("l.Get(3) should be 5")
	}
}