// 1. we collect the data inside the struct because "GO" does not have a Class and access the struct using the methods we provide (Receivers).

// 2. The Arr need 3 things which is "Capacity" to track the total allocated memory slots and "Size" to track how many slots are currently used and "Static Array" to holds the actual data

type DynamicArray struct {
	capacity int
	size int
	arr []int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		capacity: capacity,
		size: 0,
		arr: make([]int, capacity),
	}
}

func (da *DynamicArray) Get(i int) int {
	if i < 0 || i >= da.size {
		return 0
	}
	return da.arr[i]
}

func (da *DynamicArray) Set(i int, n int) {
	if i < 0 || i >= da.size {
		return
	}
	da.arr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if da.capacity == da.size {
		da.resize()
	}
	da.arr[da.size] = n
	da.size++

}

func (da *DynamicArray) Popback() int {
	if da.size == 0 {
		return 0
	}
	da.size--
	return da.arr[da.size]
}

func (da *DynamicArray) resize() {
	newArr := make([]int, da.capacity * 2)
	for i := 0; i < da.size; i++ {
		newArr[i] = da.arr[i]
	}
	da.arr = newArr
	da.capacity *= 2
}

func (da *DynamicArray) GetSize() int {
	return da.size
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
