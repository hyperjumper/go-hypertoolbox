package interval

import (
	"errors"
	"fmt"
)

func NewIntRange(a, b int) *IntRange {
	if a > b {
		return &IntRange{
			From: b,
			To:   a,
		}
	}
	return &IntRange{
		From: a,
		To:   b,
	}
}

type IntRange struct {
	From int
	To int
}

func (r *IntRange) Size() int {
	if r.From > r.To {
		return r.From - r.To + 1
	}
	return r.To - r.From + 1
}

func (r *IntRange) String() string {
	return fmt.Sprintf("[%d:%d]", r.From, r.To)
}

func (r *IntRange) IsInRange(value int) bool {
	return value >= r.From && value <= r.To
}

func (r *IntRange) IsTouching(that *IntRange) bool {
	return r.To + 1 == that.From || r.From - 1 == that.To
}

func (r *IntRange) IsSlicing(that *IntRange) bool {
	return (r.IsInRange(that.From) && !r.IsInRange(that.To)) || (!r.IsInRange(that.From) && r.IsInRange(that.To))
}

func (r *IntRange) IsInside(that *IntRange) bool {
	return that.IsInRange(r.To) && that.IsInRange(r.From)
}

func (r *IntRange) IsCovering(that *IntRange) bool {
	return r.IsInRange(that.To) && r.IsInRange(that.From)
}

func (r *IntRange) CanMerge(that *IntRange) bool {
	return r.IsCovering(that) || r.IsInside(that) || r.IsSlicing(that) || r.IsTouching(that)
}

func (r *IntRange) Merge(that *IntRange) error {
	if r.CanMerge(that) {
		nfrom := r.From
		nto := r.To
		if r.From > that.From {
			nfrom = that.From
		}
		if r.To < that.To {
			nto = that.To
		}
		r.From = nfrom
		r.To = nto
		return  nil
	}
	return errors.New("can not merge separated range")
}

func (r *IntRange) MergeTo(that *IntRange) error {
	if r.CanMerge(that) {
		nfrom := r.From
		nto := r.To
		if r.From > that.From {
			nfrom = that.From
		}
		if r.To < that.To {
			nto = that.To
		}
		that.From = nfrom
		that.To = nto
		return  nil
	}
	return errors.New("can not merge separated range")
}

func MergeIntRange(one, two *IntRange) (*IntRange, error) {
	if one.CanMerge(two) {
		nfrom := one.From
		nto := one.To
		if one.From > two.From {
			nfrom = two.From
		}
		if one.To < two.To {
			nto = two.To
		}
		return &IntRange{
			From: nfrom,
			To:   nto,
		}, nil
	}
	return nil, errors.New("can not merge separated range")
}


