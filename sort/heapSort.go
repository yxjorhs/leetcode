package sort

type heap struct {
	data []int
}

func (h *heap) in(n int) {
	h.data = append([]int{n}, h.data...)
	h.adjustUp(len(h.data) - 1)
}

func (h *heap) out() (int, string) {
	if len(h.data) == 0 {
		return 0, "err"
	}

	v := h.data[0]

	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.adjustDown(0)

	return v, "ok"
}

// 向上调节
func (h *heap) adjustUp(i int) {
	if i == 0 {
		return
	}
	aim := (i - 1) / 2

	if h.data[aim] <= h.data[i] {
		return
	}

	temp := h.data[aim]
	h.data[aim] = h.data[i]
	h.data[i] = temp

	h.adjustUp(aim)
}

// 向下调节
func (h *heap) adjustDown(i int) {
	aim := i*2 + 1

	if len(h.data) <= aim {
		return
	}

	if len(h.data) > aim+1 && h.data[aim+1] < h.data[aim] {
		aim++
	}

	if h.data[aim] >= h.data[i] {
		return
	}

	temp := h.data[i]
	h.data[i] = h.data[aim]
	h.data[aim] = temp

	h.adjustDown(aim)
}

func headSort(nums []int) []int {
	hp := new(heap)

	for i := 0; i < len(nums); i++ {
		hp.in(nums[i])
	}

	ret := []int{}

	for {
		v, msg := hp.out()

		if msg == "err" {
			break
		}

		ret = append(ret, v)
	}

	return ret
}
