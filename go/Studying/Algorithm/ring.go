package Algorithm

// 循环链表
type Ring struct {
	next, prev *Ring
	Value      interface{}
}

// 初始化空的循环链表，前驱和后驱都指向自己，因为是循环的
func (r *Ring) InitRing() *Ring {
	r.next = r
	r.prev = r
	return r
}
func NewRing(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	p.prev = p
	return nil
}

// 获取下一个节点
func (r *Ring) NextRing() *Ring {
	if r.next == nil {
		return r.InitRing()
	}
	return r.next
}

// 获取上一个节点
func (r *Ring) PrevRing() *Ring {
	if r.next == nil {
		return r.InitRing()
	}
	return r.prev
}

// 因为链表是循环的，当 n 为负数，表示从前面往前遍历，否则往后面遍历
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.InitRing()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}

	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// 往节点A，链接一个节点，并且返回之前节点A的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.NextRing()
	if s != nil {
		p := s.PrevRing()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}
