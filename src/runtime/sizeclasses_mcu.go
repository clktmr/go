// Code generated by mksizeclasses_mcu.go; DO NOT EDIT.
//go:generate go run mksizeclasses_mcu.go

//go:build noos && thumb

package runtime

// class  bytes/obj  bytes/span  objects  tail waste  max waste  min align
//     1          8         256       32           0     87.50%          8
//     2         16         256       16           0     43.75%         16
//     3         32         256        8           0     46.88%         32
//     4         48         256        5          16     35.55%         16
//     5         64         256        4           0     23.44%         64
//     6         80         256        3          16     23.83%         16
//     7         96         512        5          32     20.90%         32
//     8        128         256        2           0     24.22%        128
//     9        144         768        5          48     16.02%         16
//    10        160         512        3          32     15.04%         32
//    11        192         768        4           0     16.15%         64
//    12        208        1280        6          32      9.53%         16
//    13        256         256        1           0     18.36%        256
//    14        288        1280        4         128     19.69%         32
//    15        320        1024        3          64     15.33%         64
//    16        384         768        2           0     16.41%        128
//    17        416        1280        3          32      9.77%         32
//    18        512         512        1           0     18.55%        256

// alignment  bits  min obj size
//         8     3             8
//        16     4            16
//        32     5           256
//       256     8           512

const (
	_MaxSmallSize   = 512
	smallSizeDiv    = 8
	smallSizeMax    = 256
	largeSizeDiv    = 128
	_NumSizeClasses = 19
	_PageShift      = 8
)

var class_to_size = [_NumSizeClasses]uint16{0, 8, 16, 32, 48, 64, 80, 96, 128, 144, 160, 192, 208, 256, 288, 320, 384, 416, 512}
var class_to_allocnpages = [_NumSizeClasses]uint8{0, 1, 1, 1, 1, 1, 1, 2, 1, 3, 2, 3, 5, 1, 5, 4, 3, 5, 2}
var class_to_divmagic = [_NumSizeClasses]uint32{0, ^uint32(0)/8 + 1, ^uint32(0)/16 + 1, ^uint32(0)/32 + 1, ^uint32(0)/48 + 1, ^uint32(0)/64 + 1, ^uint32(0)/80 + 1, ^uint32(0)/96 + 1, ^uint32(0)/128 + 1, ^uint32(0)/144 + 1, ^uint32(0)/160 + 1, ^uint32(0)/192 + 1, ^uint32(0)/208 + 1, ^uint32(0)/256 + 1, ^uint32(0)/288 + 1, ^uint32(0)/320 + 1, ^uint32(0)/384 + 1, ^uint32(0)/416 + 1, ^uint32(0)/512 + 1}
var size_to_class8 = [smallSizeMax/smallSizeDiv + 1]uint8{0, 1, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 8, 8, 9, 9, 10, 10, 11, 11, 11, 11, 12, 12, 13, 13, 13, 13, 13, 13}
var size_to_class128 = [(_MaxSmallSize-smallSizeMax)/largeSizeDiv + 1]uint8{13, 16, 18}
