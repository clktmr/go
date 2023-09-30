// Code generated by mksizeclasses_mcu64.go; DO NOT EDIT.
//go:generate go run mksizeclasses_mcu64.go

//go:build noos && (riscv64 || mips64)

package runtime

// class  bytes/obj  bytes/span  objects  tail waste  max waste  min align
//     1          8        2048      256           0     87.50%          8
//     2         16        2048      128           0     43.75%         16
//     3         32        2048       64           0     46.88%         32
//     4         48        2048       42          32     32.32%         16
//     5         64        2048       32           0     23.44%         64
//     6         80        2048       25          48     20.65%         16
//     7         96        2048       21          32     16.94%         32
//     8        112        2048       18          32     14.75%         16
//     9        128        2048       16           0     11.72%        128
//    10        144        2048       14          32     11.82%         16
//    11        160        2048       12         128     15.04%         32
//    12        176        2048       11         112     13.53%         16
//    13        192        2048       10         128     13.57%         64
//    14        224        2048        9          32     15.19%         32
//    15        256        2048        8           0     12.11%        256
//    16        288        2048        7          32     12.16%         32
//    17        320        2048        6         128     15.33%         64
//    18        352        4096       11         224     13.79%         32
//    19        384        2048        5         128     13.82%        128
//    20        416        4096        9         352     15.41%         32
//    21        512        2048        4           0     18.55%        512
//    22        576        4096        7          64     12.33%         64
//    23        640        2048        3         128     15.48%        128
//    24        768        6144        8           0     16.54%        256
//    25        768        4096        5         256      6.13%        256
//    26        832        6144        7         320     12.39%         64
//    27       1024        2048        2           0     18.65%       1024
//    28       1152        6144        5         384     16.59%        128
//    29       1280        4096        3         256     15.55%        256
//    30       1536        6144        4           0     16.60%        512
//    31       1664       10240        6         256      9.94%        128
//    32       2048        2048        1           0     18.70%       2048
//    33       2560       10240        4           0     19.96%        512
//    34       2688        8192        3         128      6.21%        128
//    35       3072        6144        2           0     12.47%       1024
//    36       3328       10240        3         256      9.97%        256
//    37       4096        4096        1           0     18.73%       2048

// alignment  bits  min obj size
//         8     3             8
//        16     4            16
//        32     5           192
//        64     6           512
//       128     7          1024
//       256     8          3072
//      2048    11          4096

const (
	_MaxSmallSize   = 4096
	smallSizeDiv    = 8
	smallSizeMax    = 256
	largeSizeDiv    = 128
	_NumSizeClasses = 38
	_PageShift      = 11
)

var class_to_size = [_NumSizeClasses]uint16{0, 8, 16, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 224, 256, 288, 320, 352, 384, 416, 512, 576, 640, 768, 768, 832, 1024, 1152, 1280, 1536, 1664, 2048, 2560, 2688, 3072, 3328, 4096}
var class_to_allocnpages = [_NumSizeClasses]uint8{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 2, 3, 1, 3, 2, 3, 5, 1, 5, 4, 3, 5, 2}
var class_to_divmagic = [_NumSizeClasses]uint32{0, ^uint32(0)/8 + 1, ^uint32(0)/16 + 1, ^uint32(0)/32 + 1, ^uint32(0)/48 + 1, ^uint32(0)/64 + 1, ^uint32(0)/80 + 1, ^uint32(0)/96 + 1, ^uint32(0)/112 + 1, ^uint32(0)/128 + 1, ^uint32(0)/144 + 1, ^uint32(0)/160 + 1, ^uint32(0)/176 + 1, ^uint32(0)/192 + 1, ^uint32(0)/224 + 1, ^uint32(0)/256 + 1, ^uint32(0)/288 + 1, ^uint32(0)/320 + 1, ^uint32(0)/352 + 1, ^uint32(0)/384 + 1, ^uint32(0)/416 + 1, ^uint32(0)/512 + 1, ^uint32(0)/576 + 1, ^uint32(0)/640 + 1, ^uint32(0)/768 + 1, ^uint32(0)/768 + 1, ^uint32(0)/832 + 1, ^uint32(0)/1024 + 1, ^uint32(0)/1152 + 1, ^uint32(0)/1280 + 1, ^uint32(0)/1536 + 1, ^uint32(0)/1664 + 1, ^uint32(0)/2048 + 1, ^uint32(0)/2560 + 1, ^uint32(0)/2688 + 1, ^uint32(0)/3072 + 1, ^uint32(0)/3328 + 1, ^uint32(0)/4096 + 1}
var size_to_class8 = [smallSizeMax/smallSizeDiv + 1]uint8{0, 1, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 14, 14, 15, 15, 15, 15}
var size_to_class128 = [(_MaxSmallSize-smallSizeMax)/largeSizeDiv + 1]uint8{15, 19, 21, 23, 24, 27, 27, 28, 29, 30, 30, 31, 32, 32, 32, 33, 33, 33, 33, 34, 35, 35, 35, 36, 36, 37, 37, 37, 37, 37, 37}
