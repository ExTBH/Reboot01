package pusher

import "fmt"

func pa(a, b *[]int) {
	*a = append(*a, (*b)[0])
	reverse(*a)
	*b = (*b)[1:]
	fmt.Println("pa")
}

func pb(a, b *[]int) {
	*b = append(*b, (*a)[0])
	reverse(*b)
	*a = (*a)[1:]
	fmt.Println("pb")
}

func sa(a []int) {
	a[0], a[1] = a[1], a[0]
	fmt.Println("sa")
}

func sb(b []int) {
	b[0], b[1] = b[1], b[0]
	fmt.Println("sb")
}

func ss(a, b []int) {
	sa(a)
	sb(b)
	fmt.Println("ss")
}

func ra(a []int) {
	for i := 0; i < len(a)-1; i++ {
		a[i], a[i+1] = a[i+1], a[i]
	}
	fmt.Println("ra")
}

func rb(b []int) {
	for i := 0; i < len(b)-1; i++ {
		b[i], b[i+1] = b[i+1], b[i]
	}
	fmt.Println("rb")
}

func rr(a, b []int) {
	ra(a)
	rb(b)
	fmt.Println("rr")
}

func rra(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		a[i], a[i-1] = a[i-1], a[i]
	}
	fmt.Println("rra")
}

func rrb(b []int) {
	for i := len(b) - 1; i > 0; i-- {
		b[i], b[i-1] = b[i-1], b[i]
	}
	fmt.Println("rrb")
}

func rrr(a, b []int) {
	rra(a)
	rrb(b)
	fmt.Println("rrr")
}
