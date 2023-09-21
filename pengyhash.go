package pengyhash

import "encoding/binary"

func V64(p []byte) uint64 {
	return binary.LittleEndian.Uint64(p)
}

func pengyhash(buf []byte, size int, seed uint64) uint64 {
	p := buf
	s := [4]uint64{0}
	f := [4]uint64{0}

	s[0] = uint64(size)

	for size >= 32 {

		s[1] += V64(p[8:])
		s[0] += s[1] + V64(p[:])
		s[1] = s[0] ^ (s[1]<<14 | s[1]>>50)

		s[3] += V64(p[24:])
		s[2] += s[3] + V64(p[16:])
		s[3] = s[2] ^ (s[3]<<23 | s[3]>>41)

		s[3] += V64(p[24:])
		s[0] += s[3] + V64(p[:])
		s[3] = s[0] ^ (s[3]<<11 | s[3]>>53)

		s[1] += V64(p[8:])
		s[2] += s[1] + V64(p[16:])
		s[1] = s[2] ^ (s[1]<<40 | s[1]>>24)

		size -= 32
		p = p[32:]
	}

	for i := 0; i+8 <= size; i += 8 {
		f[i/8] = V64(p[i:])
	}

	for i := 0; i < size; i++ {
		f[i/8] |= uint64(p[i]) << (i % 8 * 8)
	}

	for i := 0; i < 6; i++ {

		s[1] += f[1] + seed
		s[0] += s[1] + f[0]
		s[1] = s[0] ^ (s[1]<<14 | s[1]>>50)

		s[3] += f[3]
		s[2] += s[3] + f[2]
		s[3] = s[2] ^ (s[3]<<23 | s[3]>>41)

		s[3] += f[3]
		s[0] += s[3] + f[0]
		s[3] = s[0] ^ (s[3]<<9 | s[3]>>55)

		s[1] += f[1]
		s[2] += s[1] + f[2]
		s[1] = s[2] ^ (s[1]<<40 | s[1]>>24)

	}

	return s[0] + s[1] + s[2] + s[3]
}
