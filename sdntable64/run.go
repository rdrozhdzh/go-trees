package sdntable64

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/willf/bloom"
)

type run struct {
	k []int64
	v []uint64
	i []uint32
}

func makeRun(k []int64, v uint64) run {
	keys := make([]int64, len(k))
	copy(keys, k)

	return run{
		k: keys,
		v: []uint64{v},
	}
}

func makeRunForBlock(n int) run {
	m := blocks[n-1]
	return run{
		k: make([]int64, m*n),
		v: make([]uint64, m),
	}
}

func (r run) insert(k []int64, v uint64) (run, bool) {
	if len(r.v) > 0 {
		if len(r.k) < len(r.v) {
			panic(fmt.Errorf("corrupted run (k: %d, v: %d) on inplace insert %d (%x)", len(r.k), len(r.v), len(k), v))
		}

		if len(r.k)/len(r.v) != len(k) {
			panic(fmt.Errorf("invalid key on inplace insert %d (%x) for run (k: %d, v: %d, n: %d)",
				len(k), v, len(r.k), len(r.v), len(r.k)/len(r.v)),
			)
		}

		left := 0
		right := len(r.v)
		for {
			m := (left + right) / 2
			b := m * len(k)
			d := compare(k, r.k[b:b+len(k)])
			if d == 0 {
				if r.v[m] == v {
					return r, false
				}

				tmp := make([]uint64, len(r.v))
				copy(tmp, r.v)
				tmp[m] = v
				r.v = tmp

				if len(r.i) > 0 {
					tmp := make([]uint32, len(r.i))
					copy(tmp, r.i)
					tmp[m] = math.MaxUint32
					r.i = tmp
				}

				return r, true
			}

			if d < 0 {
				right = m
				if left == right {
					break
				}
			} else {
				if left == m {
					left = right
					break
				}

				left = m
			}
		}

		r.k = append(r.k, k...)
		r.v = append(r.v, v)
		if len(r.i) > 0 {
			r.i = append(r.i, math.MaxUint32)
		}

		if right < len(r.v)-1 {
			i := len(k) * right
			copy(r.k[i+len(k):], r.k[i:])
			copy(r.k[i:], k)

			copy(r.v[right+1:], r.v[right:])
			r.v[right] = v

			if len(r.i) > 0 {
				copy(r.i[right+1:], r.i[right:])
				r.i[right] = math.MaxUint32
			}
		}

		return r, true
	}

	return makeRun(k, v), true
}

func (r run) inplaceInsert(k []int64, v uint64) run {
	if len(r.v) > 0 {
		if len(r.k) < len(r.v) {
			panic(fmt.Errorf("corrupted run (k: %d, v: %d) on inplace insert %d (%x)", len(r.k), len(r.v), len(k), v))
		}

		if len(r.k)/len(r.v) != len(k) {
			panic(fmt.Errorf("invalid key on inplace insert %d (%x) for run (k: %d, v: %d, n: %d)",
				len(k), v, len(r.k), len(r.v), len(r.k)/len(r.v)),
			)
		}

		left := 0
		right := len(r.v)
		for {
			m := (left + right) / 2
			b := m * len(k)
			d := compare(k, r.k[b:b+len(k)])
			if d == 0 {
				r.v[m] = v
				if len(r.i) > 0 {
					r.i[m] = math.MaxUint32
				}

				return r
			}

			if d < 0 {
				right = m
				if left == right {
					break
				}
			} else {
				if left == m {
					left = right
					break
				}

				left = m
			}
		}

		r.k = append(r.k, k...)
		r.v = append(r.v, v)
		if len(r.i) > 0 {
			r.i = append(r.i, math.MaxUint32)
		}

		if right < len(r.v)-1 {
			i := len(k) * right
			copy(r.k[i+len(k):], r.k[i:])
			copy(r.k[i:], k)

			copy(r.v[right+1:], r.v[right:])
			r.v[right] = v

			if len(r.i) > 0 {
				copy(r.i[right+1:], r.i[right:])
				r.i[right] = math.MaxUint32
			}
		}

		return r
	}

	return makeRun(k, v)
}

func (r run) clone() run {
	if len(r.k) > 0 {
		a := make([]int64, len(r.k))
		copy(a, r.k)
		r.k = a
	}

	if len(r.v) > 0 {
		a := make([]uint64, len(r.v))
		copy(a, r.v)
		r.v = a
	}

	if len(r.i) > 0 {
		a := make([]uint32, len(r.i))
		copy(a, r.i)
		r.i = a
	}

	return r
}

func (r run) append(k []int64, v uint64) run {
	if len(r.v) > 0 {
		if len(r.k) < len(r.v) {
			panic(fmt.Errorf("corrupted run (k: %d, v: %d) on append %d (%x)", len(r.k), len(r.v), len(k), v))
		}

		if len(r.k)/len(r.v) != len(k) {
			panic(fmt.Errorf("invalid key on append %d (%x) for run (k: %d, v: %d, n: %d)",
				len(k), v, len(r.k), len(r.v), len(r.k)/len(r.v)),
			)
		}

		r.k = append(r.k, k...)
		r.v = append(r.v, v)
		if len(r.i) > 0 {
			r.i = append(r.i, math.MaxUint32)
		}
		return r
	}

	return makeRun(k, v)
}

func (r run) normalize(log normalizeLoggers) run {
	n := len(r.v)
	if n > 0 {
		if len(r.k) < n {
			panic(fmt.Errorf("corrupted run (k: %d, v: %d) on normalize", len(r.k), len(r.v)))
		}

		m := len(r.k) / n

		if n > 1 {
			if log.before != nil {
				log.before(m, n)
			}

			ir := r.makeIndexedRun()

			if log.sort != nil {
				log.sort(m, n)
			}
			ir.sort()

			r = ir.makeRun()

			if log.after != nil {
				log.after(m, len(r.v))
			}
		}
	}

	return r
}

func (r run) get(k []int64) (uint64, uint32, uint32) {
	n := len(r.v)
	if n > 0 {
		if len(r.k) < n {
			panic(fmt.Errorf("corrupted run (k: %d, v: %d) on get %d", len(r.k), n, len(k)))
		}

		m := len(r.k) / n
		if m != len(k) {
			panic(fmt.Errorf("invalid key on get %d for run (k: %d, v: %d, n: %d)", len(k), len(r.k), n, m))
		}

		left := 0
		right := len(r.v)
		for {
			m := (left + right) / 2
			b := m * len(k)
			d := compare(k, r.k[b:b+len(k)])
			if d == 0 {
				return r.v[m], 0, 0
			}

			if left == m {
				if d < 0 {
					right = m
				}
				break
			}

			if d < 0 {
				right = m
			} else {
				left = m
			}
		}

		if len(r.i) > 0 {
			start, end := r.getDiskRange(right)
			return 0, start, end
		}
	}

	return 0, 0, 0
}

func (r run) getDiskRange(right int) (uint32, uint32) {
	start := uint32(0)
	end := uint32(math.MaxUint32)

	left := right
	if left > 0 {
		left--
		for left >= 0 && (left >= len(r.i) || r.i[left] >= math.MaxUint32) {
			left--
		}

		if left >= 0 {
			start = r.i[left]
		}
	}

	for right < len(r.i) && r.i[right] >= math.MaxUint32 {
		right++
	}

	if right < len(r.i) {
		end = r.i[right]
	}

	return start, end
}

func (r run) size() int {
	return (len(r.k) + len(r.v)) * 8
}

func (r run) size90() int {
	if len(r.v) > 0 {
		if len(r.k) < len(r.v) {
			panic(fmt.Errorf("corrupted run (k: %d, v: %d) on geting size", len(r.k), len(r.v)))
		}

		n := len(r.v)
		m := len(r.k) / n

		return n * 90 / 100 * 8 * (m + 1)
	}

	return 0
}

func (r run) truncate(n int) run {
	if n >= 0 && len(r.v) > n {
		if len(r.k) < len(r.v) {
			panic(fmt.Errorf("corrupted run (k: %d, v: %d) on truncate", len(r.k), len(r.v)))
		}

		m := len(r.k) / len(r.v)

		r.k = r.k[:m*n]
		r.v = r.v[:n]
	}

	return r
}

func (r run) read(f io.Reader) error {
	if err := binary.Read(f, binary.LittleEndian, r.k); err != nil {
		return err
	}

	if err := binary.Read(f, binary.LittleEndian, r.v); err != nil {
		return err
	}

	return nil
}

func (r run) write(f io.Writer) error {
	if err := binary.Write(f, binary.LittleEndian, r.k); err != nil {
		return err
	}

	if err := binary.Write(f, binary.LittleEndian, r.v); err != nil {
		return err
	}

	return nil
}

func (r run) merge(in *rRun, out *wRun) (*bloom.BloomFilter, run, error) {
	n := len(r.v)
	m := len(r.k) / n
	idx := make([]uint32, n)

	est := uint(n)
	if len(r.i) > 0 {
		for i := len(r.i) - 1; i >= 0; i-- {
			if r.i[i] < math.MaxUint32 {
				est += uint(r.i[i])
				break
			}
		}
	}

	fBuf := make([]byte, 8*m)
	filter := bloom.NewWithEstimates(est, 0.03)

	mK := r.k
	i := 0

	sK, sV, err := in.next()
	if err != nil {
		return nil, r, err
	}

	var (
		c uint32
		z int
	)

	for len(mK) > 0 || len(sK) > 0 {
		var d int64 = 1
		if len(mK) > 0 && len(sK) > 0 {
			d = compare(mK[:m], sK)
		} else if len(mK) > 0 {
			d = -1
		}

		if d <= 0 {
			if r.v[i] != 0 {
				if err := out.put(mK[:m], r.v[i]); err != nil {
					return nil, r, err
				}

				for i, n := range mK[:m] {
					binary.LittleEndian.PutUint64(fBuf[8*i:], uint64(n))
				}
				filter.Add(fBuf)

				idx[i-z] = c

				c++
				if c >= math.MaxUint32 {
					return nil, r, fmt.Errorf("run %d overflow", m)
				}
			} else {
				z++
			}

			i++
			mK = mK[m:]
		} else {
			if err := out.put(sK, sV); err != nil {
				return nil, r, err
			}

			for i, n := range sK {
				binary.LittleEndian.PutUint64(fBuf[8*i:], uint64(n))
			}
			filter.Add(fBuf)

			c++
			if c >= math.MaxUint32 {
				return nil, r, fmt.Errorf("run %d overflow", m)
			}
		}

		if d >= 0 {
			sK, sV, err = in.next()
			if err != nil {
				return nil, r, err
			}
		}
	}

	if err := out.flush(); err != nil {
		return nil, r, err
	}

	if z > 0 {
		keys := make([]int64, (n-z)*m)
		values := make([]uint64, n-z)

		a := 0
		j := 0
		for i, v := range r.v {
			if v != 0 {
				copy(keys[a:a+m], r.k[i*m:])
				values[j] = v

				a += m
				j++
			}
		}

		r.k = keys
		r.v = values
	}

	r.i = idx[:n-z]
	return filter, r, nil
}

func (r run) drop90() run {
	n := len(r.v)
	if n > 0 {
		if len(r.k) < n {
			panic(fmt.Errorf("corrupted run (k: %d, v: %d) on dropping 90%% entries", len(r.k), n))
		}

		m := len(r.k) / n

		n /= 10
		keys := make([]int64, n*m)
		values := make([]uint64, n)
		var idx []uint32
		if len(r.i) > 0 {
			idx = make([]uint32, n)
		}

		j := 0
		for i := range values {
			a := j * m
			copy(keys[i*m:], r.k[a:a+m])
			values[i] = r.v[j]
			if idx != nil {
				idx[i] = r.i[j]
			}

			j += 10
		}

		r = run{
			k: keys,
			v: values,
			i: idx,
		}
	}

	return r
}

var blocks = []int{
	6400,
	4266,
	3200,
	2560,
	2133,
	1828,
	1600,
	1422,
	1280,
	1163,
	1066,
	984,
	914,
	853,
	800,
	752,
	711,
	673,
	640,
	609,
	581,
	556,
	533,
	512,
	492,
	474,
	457,
	441,
	426,
	412,
	400,
	387,
	376,
	365,
	355,
	345,
	336,
	328,
	320,
	312,
	304,
	297,
	290,
	284,
	278,
	272,
	266,
	261,
	256,
	250,
	246,
	241,
	237,
	232,
	228,
	224,
	220,
	216,
	213,
	209,
	206,
	203,
	200,
	196,
	193,
	191,
	188,
	185,
	182,
	180,
	177,
	175,
	172,
	170,
	168,
	166,
	164,
	162,
	160,
	158,
	156,
	154,
	152,
	150,
	148,
	147,
	145,
	143,
	142,
	140,
	139,
	137,
	136,
	134,
	133,
	131,
	130,
	129,
	128,
	126,
	125,
	124,
	123,
	121,
	120,
	119,
	118,
	117,
	116,
	115,
	114,
	113,
	112,
	111,
	110,
	109,
	108,
	107,
	106,
	105,
	104,
	104,
	103,
	102,
	101,
	100,
	100,
}
