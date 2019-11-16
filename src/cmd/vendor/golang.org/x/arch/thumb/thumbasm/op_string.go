// Code generated by "stringer -type Op"; DO NOT EDIT.

package thumbasm

import "strconv"

const _Op_name = "ADDADD_SADCADC_SANDAND_SASRASR_SBBCCBCSBEQBFCBFIBGEBGTBHIBICBIC_SBKPTBLBLEBLSBLTBLXBMIBNEBPLBVCBVSBXCBNZCBZCLREXCLZCMNCMPCPSIDCPSIEDMBDSBEOREOR_SISBITITEITEEITEEEITEETITETITETEITETTITTITTEITTEEITTETITTTITTTEITTTTLDMIALDMDBLDRLDRBLDRDLDREXLDREXBLDREXHLDRHLDRSBLDRSHLSLLSL_SLSRLSR_SMLAMLSMOVMOV_SMOVTMOVWMRSMSRMVNMVN_SMULNOPORNORN_SORRORR_SPLDPLIPOPPUSHRBITREVREV16REVSHRORROR_SRSBRSB_SSBCSBC_SSBFXSDIVSELSEVSMLABBSMLABTSMLADSMLADXSMLALSMLATBSMLATTSMLAWBSMLAWTSMLSDSMLSDXSMMLASMMLSSMULLSTMIASTMDBSTRSTRBSTRDSTREXSTREXBSTREXHSTRHSUBSUB_SSVCSXTBSXTHTBBTBHTSTTEQUBFXUDFUDIVUMLALUMULLUXTBUXTHVABS_F32VABS_F64VADD_F32VADD_F64VCMP_F32VCMP_F64VCMPE_F32VCMPE_F64VCVT_F32_FXS16VCVT_F32_FXS32VCVT_F32_FXU16VCVT_F32_FXU32VCVT_F64_FXS16VCVT_F64_FXS32VCVT_F64_FXU16VCVT_F64_FXU32VCVT_F32_U32VCVT_F32_S32VCVT_F64_U32VCVT_F64_S32VCVT_F64_F32VCVT_F32_F64VCVT_FXS16_F32VCVT_FXS16_F64VCVT_FXS32_F32VCVT_FXS32_F64VCVT_FXU16_F32VCVT_FXU16_F64VCVT_FXU32_F32VCVT_FXU32_F64VCVTB_F32_F16VCVTB_F16_F32VCVTT_F32_F16VCVTT_F16_F32VCVTR_U32_F32VCVTR_U32_F64VCVTR_S32_F32VCVTR_S32_F64VCVT_U32_F32VCVT_U32_F64VCVT_S32_F32VCVT_S32_F64VDIV_F32VDIV_F64VLDRVMLA_F32VMLA_F64VMLS_F32VMLS_F64VMOVVMOV_32VMOV_F32VMOV_F64VMRSVMSRVMUL_F32VMUL_F64VNEG_F32VNEG_F64VNMLS_F32VNMLS_F64VNMLA_F32VNMLA_F64VNMUL_F32VNMUL_F64VSQRT_F32VSQRT_F64VSTRVSUB_F32VSUB_F64WFEWFIYIELD"

var _Op_index = [...]uint16{0, 3, 8, 11, 16, 19, 24, 27, 32, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 65, 69, 71, 74, 77, 80, 83, 86, 89, 92, 95, 98, 100, 104, 107, 112, 115, 118, 121, 126, 131, 134, 137, 140, 145, 148, 150, 153, 157, 162, 167, 171, 176, 181, 184, 188, 193, 198, 202, 207, 212, 217, 222, 225, 229, 233, 238, 244, 250, 254, 259, 264, 267, 272, 275, 280, 283, 286, 289, 294, 298, 302, 305, 308, 311, 316, 319, 322, 325, 330, 333, 338, 341, 344, 347, 351, 355, 358, 363, 368, 371, 376, 379, 384, 387, 392, 396, 400, 403, 406, 412, 418, 423, 429, 434, 440, 446, 452, 458, 463, 469, 474, 479, 484, 489, 494, 497, 501, 505, 510, 516, 522, 526, 529, 534, 537, 541, 545, 548, 551, 554, 557, 561, 564, 568, 573, 578, 582, 586, 594, 602, 610, 618, 626, 634, 643, 652, 666, 680, 694, 708, 722, 736, 750, 764, 776, 788, 800, 812, 824, 836, 850, 864, 878, 892, 906, 920, 934, 948, 961, 974, 987, 1000, 1013, 1026, 1039, 1052, 1064, 1076, 1088, 1100, 1108, 1116, 1120, 1128, 1136, 1144, 1152, 1156, 1163, 1171, 1179, 1183, 1187, 1195, 1203, 1211, 1219, 1228, 1237, 1246, 1255, 1264, 1273, 1282, 1291, 1295, 1303, 1311, 1314, 1317, 1322}

func (i Op) String() string {
	i -= 1
	if i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
