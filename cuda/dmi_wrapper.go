package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var adddmi_code cu.Function

type adddmi_args struct {
	arg_Hx unsafe.Pointer
	arg_Hy unsafe.Pointer
	arg_Hz unsafe.Pointer
	arg_mx unsafe.Pointer
	arg_my unsafe.Pointer
	arg_mz unsafe.Pointer
	arg_Dx float32
	arg_Dy float32
	arg_Dz float32
	arg_N0 int
	arg_N1 int
	arg_N2 int
	argptr [12]unsafe.Pointer
}

// Wrapper for adddmi CUDA kernel, asynchronous.
func k_adddmi_async(Hx unsafe.Pointer, Hy unsafe.Pointer, Hz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, Dx float32, Dy float32, Dz float32, N0 int, N1 int, N2 int, cfg *config, str cu.Stream) {
	if adddmi_code == 0 {
		adddmi_code = fatbinLoad(adddmi_map, "adddmi")
	}

	var a adddmi_args

	a.arg_Hx = Hx
	a.argptr[0] = unsafe.Pointer(&a.arg_Hx)
	a.arg_Hy = Hy
	a.argptr[1] = unsafe.Pointer(&a.arg_Hy)
	a.arg_Hz = Hz
	a.argptr[2] = unsafe.Pointer(&a.arg_Hz)
	a.arg_mx = mx
	a.argptr[3] = unsafe.Pointer(&a.arg_mx)
	a.arg_my = my
	a.argptr[4] = unsafe.Pointer(&a.arg_my)
	a.arg_mz = mz
	a.argptr[5] = unsafe.Pointer(&a.arg_mz)
	a.arg_Dx = Dx
	a.argptr[6] = unsafe.Pointer(&a.arg_Dx)
	a.arg_Dy = Dy
	a.argptr[7] = unsafe.Pointer(&a.arg_Dy)
	a.arg_Dz = Dz
	a.argptr[8] = unsafe.Pointer(&a.arg_Dz)
	a.arg_N0 = N0
	a.argptr[9] = unsafe.Pointer(&a.arg_N0)
	a.arg_N1 = N1
	a.argptr[10] = unsafe.Pointer(&a.arg_N1)
	a.arg_N2 = N2
	a.argptr[11] = unsafe.Pointer(&a.arg_N2)

	args := a.argptr[:]
	cu.LaunchKernel(adddmi_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for adddmi CUDA kernel, synchronized.
func k_adddmi(Hx unsafe.Pointer, Hy unsafe.Pointer, Hz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, Dx float32, Dy float32, Dz float32, N0 int, N1 int, N2 int, cfg *config) {
	str := stream()
	k_adddmi_async(Hx, Hy, Hz, mx, my, mz, Dx, Dy, Dz, N0, N1, N2, cfg, str)
	syncAndRecycle(str)
}

var adddmi_map = map[int]string{0: "",
	20: adddmi_ptx_20,
	30: adddmi_ptx_30,
	35: adddmi_ptx_35}

const (
	adddmi_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry adddmi(
	.param .u64 adddmi_param_0,
	.param .u64 adddmi_param_1,
	.param .u64 adddmi_param_2,
	.param .u64 adddmi_param_3,
	.param .u64 adddmi_param_4,
	.param .u64 adddmi_param_5,
	.param .f32 adddmi_param_6,
	.param .f32 adddmi_param_7,
	.param .f32 adddmi_param_8,
	.param .u32 adddmi_param_9,
	.param .u32 adddmi_param_10,
	.param .u32 adddmi_param_11
)
{
	.reg .pred 	%p<10>;
	.reg .s32 	%r<111>;
	.reg .f32 	%f<43>;
	.reg .s64 	%rd<38>;


	ld.param.u64 	%rd4, [adddmi_param_0];
	ld.param.u64 	%rd5, [adddmi_param_1];
	ld.param.u64 	%rd6, [adddmi_param_2];
	ld.param.u64 	%rd7, [adddmi_param_3];
	ld.param.u64 	%rd8, [adddmi_param_4];
	ld.param.u64 	%rd9, [adddmi_param_5];
	ld.param.f32 	%f16, [adddmi_param_6];
	ld.param.f32 	%f17, [adddmi_param_7];
	ld.param.f32 	%f18, [adddmi_param_8];
	ld.param.u32 	%r10, [adddmi_param_9];
	ld.param.u32 	%r11, [adddmi_param_10];
	ld.param.u32 	%r12, [adddmi_param_11];
	.loc 2 13 1
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r13, %ntid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r3, %r13, %r1, %r2;
	.loc 2 14 1
	mov.u32 	%r14, %ntid.y;
	mov.u32 	%r15, %ctaid.y;
	mov.u32 	%r16, %tid.y;
	mad.lo.s32 	%r4, %r14, %r15, %r16;
	.loc 2 16 1
	setp.lt.s32 	%p1, %r4, %r12;
	setp.lt.s32 	%p2, %r3, %r11;
	and.pred  	%p3, %p2, %p1;
	.loc 2 20 1
	setp.gt.s32 	%p4, %r10, 0;
	.loc 2 16 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_9;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 20 1
	mad.lo.s32 	%r109, %r12, %r3, %r4;
	mov.u32 	%r110, 0;
	cvta.to.global.u64 	%rd10, %rd4;
	cvta.to.global.u64 	%rd12, %rd5;
	cvta.to.global.u64 	%rd13, %rd6;

BB0_2:
	.loc 2 23 1
	mul.wide.s32 	%rd11, %r109, 4;
	add.s64 	%rd1, %rd10, %rd11;
	ld.global.f32 	%f41, [%rd1];
	add.s64 	%rd2, %rd12, %rd11;
	ld.global.f32 	%f42, [%rd2];
	add.s64 	%rd3, %rd13, %rd11;
	ld.global.f32 	%f40, [%rd3];
	.loc 2 25 1
	setp.eq.f32 	%p6, %f16, 0f00000000;
	@%p6 bra 	BB0_4;

	.loc 2 26 1
	add.s32 	%r27, %r110, 1;
	mov.u32 	%r28, 0;
	.loc 3 238 5
	max.s32 	%r29, %r27, %r28;
	.loc 2 26 1
	add.s32 	%r30, %r10, -1;
	.loc 3 210 5
	min.s32 	%r31, %r29, %r30;
	.loc 2 26 1
	add.s32 	%r32, %r11, -1;
	.loc 3 238 5
	max.s32 	%r33, %r3, %r28;
	.loc 3 210 5
	min.s32 	%r34, %r33, %r32;
	.loc 2 26 1
	mad.lo.s32 	%r35, %r31, %r11, %r34;
	add.s32 	%r36, %r12, -1;
	.loc 3 238 5
	max.s32 	%r37, %r4, %r28;
	.loc 3 210 5
	min.s32 	%r38, %r37, %r36;
	.loc 2 26 1
	mad.lo.s32 	%r39, %r35, %r12, %r38;
	cvta.to.global.u64 	%rd14, %rd7;
	.loc 2 26 1
	mul.wide.s32 	%rd15, %r39, 4;
	add.s64 	%rd16, %rd14, %rd15;
	add.s32 	%r40, %r110, -1;
	.loc 3 238 5
	max.s32 	%r41, %r40, %r28;
	.loc 3 210 5
	min.s32 	%r42, %r41, %r30;
	.loc 2 26 1
	mad.lo.s32 	%r43, %r42, %r11, %r34;
	mad.lo.s32 	%r44, %r43, %r12, %r38;
	mul.wide.s32 	%rd17, %r44, 4;
	add.s64 	%rd18, %rd14, %rd17;
	ld.global.f32 	%f19, [%rd18];
	ld.global.f32 	%f20, [%rd16];
	sub.f32 	%f21, %f20, %f19;
	fma.rn.f32 	%f40, %f21, %f16, %f40;
	cvta.to.global.u64 	%rd19, %rd9;
	.loc 2 27 1
	add.s64 	%rd20, %rd19, %rd15;
	add.s64 	%rd21, %rd19, %rd17;
	ld.global.f32 	%f22, [%rd21];
	ld.global.f32 	%f23, [%rd20];
	sub.f32 	%f24, %f23, %f22;
	mul.f32 	%f25, %f24, %f16;
	sub.f32 	%f41, %f41, %f25;

BB0_4:
	.loc 2 30 1
	setp.eq.f32 	%p7, %f17, 0f00000000;
	@%p7 bra 	BB0_6;

	mov.u32 	%r49, 0;
	.loc 3 238 5
	max.s32 	%r50, %r110, %r49;
	.loc 2 26 1
	add.s32 	%r51, %r10, -1;
	.loc 3 210 5
	min.s32 	%r52, %r50, %r51;
	.loc 2 31 1
	add.s32 	%r57, %r3, 1;
	.loc 3 238 5
	max.s32 	%r58, %r57, %r49;
	.loc 2 26 1
	add.s32 	%r59, %r11, -1;
	.loc 3 210 5
	min.s32 	%r60, %r58, %r59;
	.loc 2 31 1
	mad.lo.s32 	%r61, %r52, %r11, %r60;
	.loc 3 238 5
	max.s32 	%r66, %r4, %r49;
	.loc 2 26 1
	add.s32 	%r67, %r12, -1;
	.loc 3 210 5
	min.s32 	%r68, %r66, %r67;
	.loc 2 31 1
	mad.lo.s32 	%r69, %r61, %r12, %r68;
	cvta.to.global.u64 	%rd22, %rd8;
	.loc 2 31 1
	mul.wide.s32 	%rd23, %r69, 4;
	add.s64 	%rd24, %rd22, %rd23;
	add.s32 	%r70, %r3, -1;
	.loc 3 238 5
	max.s32 	%r71, %r70, %r49;
	.loc 3 210 5
	min.s32 	%r72, %r71, %r59;
	.loc 2 31 1
	mad.lo.s32 	%r73, %r52, %r11, %r72;
	add.s32 	%r74, %r4, -1;
	.loc 3 238 5
	max.s32 	%r75, %r74, %r49;
	.loc 3 210 5
	min.s32 	%r76, %r75, %r67;
	.loc 2 31 1
	mad.lo.s32 	%r77, %r73, %r12, %r76;
	mul.wide.s32 	%rd25, %r77, 4;
	add.s64 	%rd26, %rd22, %rd25;
	ld.global.f32 	%f26, [%rd26];
	ld.global.f32 	%f27, [%rd24];
	sub.f32 	%f28, %f27, %f26;
	fma.rn.f32 	%f40, %f28, %f17, %f40;
	cvta.to.global.u64 	%rd27, %rd9;
	.loc 2 32 1
	add.s64 	%rd28, %rd27, %rd23;
	add.s64 	%rd29, %rd27, %rd25;
	ld.global.f32 	%f29, [%rd29];
	ld.global.f32 	%f30, [%rd28];
	sub.f32 	%f31, %f30, %f29;
	mul.f32 	%f32, %f31, %f16;
	sub.f32 	%f42, %f42, %f32;

BB0_6:
	.loc 2 35 1
	setp.eq.f32 	%p8, %f18, 0f00000000;
	@%p8 bra 	BB0_8;

	mov.u32 	%r82, 0;
	.loc 3 238 5
	max.s32 	%r83, %r110, %r82;
	.loc 2 26 1
	add.s32 	%r84, %r10, -1;
	.loc 3 210 5
	min.s32 	%r85, %r83, %r84;
	.loc 2 26 1
	add.s32 	%r86, %r11, -1;
	.loc 3 238 5
	max.s32 	%r87, %r3, %r82;
	.loc 3 210 5
	min.s32 	%r88, %r87, %r86;
	.loc 2 36 1
	mad.lo.s32 	%r89, %r85, %r11, %r88;
	add.s32 	%r94, %r4, 1;
	.loc 3 238 5
	max.s32 	%r95, %r94, %r82;
	.loc 2 26 1
	add.s32 	%r96, %r12, -1;
	.loc 3 210 5
	min.s32 	%r97, %r95, %r96;
	.loc 2 36 1
	mad.lo.s32 	%r98, %r89, %r12, %r97;
	cvta.to.global.u64 	%rd30, %rd8;
	.loc 2 36 1
	mul.wide.s32 	%rd31, %r98, 4;
	add.s64 	%rd32, %rd30, %rd31;
	.loc 3 238 5
	max.s32 	%r99, %r4, %r82;
	.loc 3 210 5
	min.s32 	%r100, %r99, %r96;
	.loc 2 36 1
	mad.lo.s32 	%r101, %r89, %r12, %r100;
	mul.wide.s32 	%rd33, %r101, 4;
	add.s64 	%rd34, %rd30, %rd33;
	ld.global.f32 	%f33, [%rd34];
	ld.global.f32 	%f34, [%rd32];
	sub.f32 	%f35, %f34, %f33;
	fma.rn.f32 	%f41, %f35, %f16, %f41;
	cvta.to.global.u64 	%rd35, %rd7;
	.loc 2 37 1
	add.s64 	%rd36, %rd35, %rd31;
	add.s64 	%rd37, %rd35, %rd33;
	ld.global.f32 	%f36, [%rd37];
	ld.global.f32 	%f37, [%rd36];
	sub.f32 	%f38, %f37, %f36;
	mul.f32 	%f39, %f38, %f16;
	sub.f32 	%f42, %f42, %f39;

BB0_8:
	.loc 2 41 1
	st.global.f32 	[%rd1], %f41;
	.loc 2 42 1
	st.global.f32 	[%rd2], %f42;
	.loc 2 43 1
	st.global.f32 	[%rd3], %f40;
	.loc 2 20 1
	mad.lo.s32 	%r109, %r12, %r11, %r109;
	.loc 2 20 18
	add.s32 	%r110, %r110, 1;
	.loc 2 20 1
	setp.lt.s32 	%p9, %r110, %r10;
	@%p9 bra 	BB0_2;

BB0_9:
	.loc 2 45 2
	ret;
}


`
	adddmi_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry adddmi(
	.param .u64 adddmi_param_0,
	.param .u64 adddmi_param_1,
	.param .u64 adddmi_param_2,
	.param .u64 adddmi_param_3,
	.param .u64 adddmi_param_4,
	.param .u64 adddmi_param_5,
	.param .f32 adddmi_param_6,
	.param .f32 adddmi_param_7,
	.param .f32 adddmi_param_8,
	.param .u32 adddmi_param_9,
	.param .u32 adddmi_param_10,
	.param .u32 adddmi_param_11
)
{
	.reg .pred 	%p<10>;
	.reg .s32 	%r<100>;
	.reg .f32 	%f<43>;
	.reg .s64 	%rd<38>;


	ld.param.u64 	%rd4, [adddmi_param_0];
	ld.param.u64 	%rd5, [adddmi_param_1];
	ld.param.u64 	%rd6, [adddmi_param_2];
	ld.param.u64 	%rd7, [adddmi_param_3];
	ld.param.u64 	%rd8, [adddmi_param_4];
	ld.param.u64 	%rd9, [adddmi_param_5];
	ld.param.f32 	%f16, [adddmi_param_6];
	ld.param.f32 	%f17, [adddmi_param_7];
	ld.param.f32 	%f18, [adddmi_param_8];
	ld.param.u32 	%r11, [adddmi_param_9];
	ld.param.u32 	%r12, [adddmi_param_10];
	ld.param.u32 	%r13, [adddmi_param_11];
	.loc 2 13 1
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r14, %ntid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r3, %r14, %r1, %r2;
	.loc 2 14 1
	mov.u32 	%r15, %ntid.y;
	mov.u32 	%r16, %ctaid.y;
	mov.u32 	%r17, %tid.y;
	mad.lo.s32 	%r4, %r15, %r16, %r17;
	.loc 2 16 1
	setp.lt.s32 	%p1, %r4, %r13;
	setp.lt.s32 	%p2, %r3, %r12;
	and.pred  	%p3, %p2, %p1;
	.loc 2 20 1
	setp.gt.s32 	%p4, %r11, 0;
	.loc 2 16 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_9;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 20 1
	mad.lo.s32 	%r98, %r13, %r3, %r4;
	mul.lo.s32 	%r6, %r13, %r12;
	mov.u32 	%r99, 0;
	cvta.to.global.u64 	%rd10, %rd4;
	cvta.to.global.u64 	%rd12, %rd5;
	cvta.to.global.u64 	%rd13, %rd6;

BB0_2:
	.loc 2 23 1
	mul.wide.s32 	%rd11, %r98, 4;
	add.s64 	%rd1, %rd10, %rd11;
	ld.global.f32 	%f41, [%rd1];
	add.s64 	%rd2, %rd12, %rd11;
	ld.global.f32 	%f42, [%rd2];
	add.s64 	%rd3, %rd13, %rd11;
	ld.global.f32 	%f40, [%rd3];
	.loc 2 25 1
	setp.eq.f32 	%p6, %f16, 0f00000000;
	@%p6 bra 	BB0_4;

	.loc 2 26 1
	add.s32 	%r28, %r99, 1;
	mov.u32 	%r29, 0;
	.loc 3 238 5
	max.s32 	%r30, %r28, %r29;
	.loc 2 26 1
	add.s32 	%r31, %r11, -1;
	.loc 3 210 5
	min.s32 	%r32, %r30, %r31;
	.loc 2 26 1
	add.s32 	%r33, %r12, -1;
	.loc 3 238 5
	max.s32 	%r34, %r3, %r29;
	.loc 3 210 5
	min.s32 	%r35, %r34, %r33;
	.loc 2 26 1
	mad.lo.s32 	%r36, %r32, %r12, %r35;
	add.s32 	%r37, %r13, -1;
	.loc 3 238 5
	max.s32 	%r38, %r4, %r29;
	.loc 3 210 5
	min.s32 	%r39, %r38, %r37;
	.loc 2 26 1
	mad.lo.s32 	%r40, %r36, %r13, %r39;
	cvta.to.global.u64 	%rd14, %rd7;
	.loc 2 26 1
	mul.wide.s32 	%rd15, %r40, 4;
	add.s64 	%rd16, %rd14, %rd15;
	add.s32 	%r41, %r99, -1;
	.loc 3 238 5
	max.s32 	%r42, %r41, %r29;
	.loc 3 210 5
	min.s32 	%r43, %r42, %r31;
	.loc 2 26 1
	mad.lo.s32 	%r44, %r43, %r12, %r35;
	mad.lo.s32 	%r45, %r44, %r13, %r39;
	mul.wide.s32 	%rd17, %r45, 4;
	add.s64 	%rd18, %rd14, %rd17;
	ld.global.f32 	%f19, [%rd18];
	ld.global.f32 	%f20, [%rd16];
	sub.f32 	%f21, %f20, %f19;
	fma.rn.f32 	%f40, %f21, %f16, %f40;
	cvta.to.global.u64 	%rd19, %rd9;
	.loc 2 27 1
	add.s64 	%rd20, %rd19, %rd15;
	add.s64 	%rd21, %rd19, %rd17;
	ld.global.f32 	%f22, [%rd21];
	ld.global.f32 	%f23, [%rd20];
	sub.f32 	%f24, %f23, %f22;
	mul.f32 	%f25, %f24, %f16;
	sub.f32 	%f41, %f41, %f25;

BB0_4:
	.loc 2 30 1
	setp.eq.f32 	%p7, %f17, 0f00000000;
	@%p7 bra 	BB0_6;

	mov.u32 	%r50, 0;
	.loc 3 238 5
	max.s32 	%r51, %r99, %r50;
	.loc 2 26 1
	add.s32 	%r52, %r11, -1;
	.loc 3 210 5
	min.s32 	%r53, %r51, %r52;
	.loc 2 31 1
	add.s32 	%r54, %r3, 1;
	.loc 3 238 5
	max.s32 	%r55, %r54, %r50;
	.loc 2 26 1
	add.s32 	%r56, %r12, -1;
	.loc 3 210 5
	min.s32 	%r57, %r55, %r56;
	.loc 2 31 1
	mad.lo.s32 	%r58, %r53, %r12, %r57;
	.loc 2 26 1
	add.s32 	%r59, %r13, -1;
	.loc 3 238 5
	max.s32 	%r60, %r4, %r50;
	.loc 3 210 5
	min.s32 	%r61, %r60, %r59;
	.loc 2 31 1
	mad.lo.s32 	%r62, %r58, %r13, %r61;
	cvta.to.global.u64 	%rd22, %rd8;
	.loc 2 31 1
	mul.wide.s32 	%rd23, %r62, 4;
	add.s64 	%rd24, %rd22, %rd23;
	add.s32 	%r63, %r3, -1;
	.loc 3 238 5
	max.s32 	%r64, %r63, %r50;
	.loc 3 210 5
	min.s32 	%r65, %r64, %r56;
	.loc 2 31 1
	mad.lo.s32 	%r66, %r53, %r12, %r65;
	add.s32 	%r67, %r4, -1;
	.loc 3 238 5
	max.s32 	%r68, %r67, %r50;
	.loc 3 210 5
	min.s32 	%r69, %r68, %r59;
	.loc 2 31 1
	mad.lo.s32 	%r70, %r66, %r13, %r69;
	mul.wide.s32 	%rd25, %r70, 4;
	add.s64 	%rd26, %rd22, %rd25;
	ld.global.f32 	%f26, [%rd26];
	ld.global.f32 	%f27, [%rd24];
	sub.f32 	%f28, %f27, %f26;
	fma.rn.f32 	%f40, %f28, %f17, %f40;
	cvta.to.global.u64 	%rd27, %rd9;
	.loc 2 32 1
	add.s64 	%rd28, %rd27, %rd23;
	add.s64 	%rd29, %rd27, %rd25;
	ld.global.f32 	%f29, [%rd29];
	ld.global.f32 	%f30, [%rd28];
	sub.f32 	%f31, %f30, %f29;
	mul.f32 	%f32, %f31, %f16;
	sub.f32 	%f42, %f42, %f32;

BB0_6:
	.loc 2 35 1
	setp.eq.f32 	%p8, %f18, 0f00000000;
	@%p8 bra 	BB0_8;

	mov.u32 	%r75, 0;
	.loc 3 238 5
	max.s32 	%r76, %r99, %r75;
	.loc 2 26 1
	add.s32 	%r77, %r11, -1;
	.loc 3 210 5
	min.s32 	%r78, %r76, %r77;
	.loc 2 26 1
	add.s32 	%r79, %r12, -1;
	.loc 3 238 5
	max.s32 	%r80, %r3, %r75;
	.loc 3 210 5
	min.s32 	%r81, %r80, %r79;
	.loc 2 36 1
	mad.lo.s32 	%r82, %r78, %r12, %r81;
	add.s32 	%r83, %r4, 1;
	.loc 3 238 5
	max.s32 	%r84, %r83, %r75;
	.loc 2 26 1
	add.s32 	%r85, %r13, -1;
	.loc 3 210 5
	min.s32 	%r86, %r84, %r85;
	.loc 2 36 1
	mad.lo.s32 	%r87, %r82, %r13, %r86;
	cvta.to.global.u64 	%rd30, %rd8;
	.loc 2 36 1
	mul.wide.s32 	%rd31, %r87, 4;
	add.s64 	%rd32, %rd30, %rd31;
	.loc 3 238 5
	max.s32 	%r88, %r4, %r75;
	.loc 3 210 5
	min.s32 	%r89, %r88, %r85;
	.loc 2 36 1
	mad.lo.s32 	%r90, %r82, %r13, %r89;
	mul.wide.s32 	%rd33, %r90, 4;
	add.s64 	%rd34, %rd30, %rd33;
	ld.global.f32 	%f33, [%rd34];
	ld.global.f32 	%f34, [%rd32];
	sub.f32 	%f35, %f34, %f33;
	fma.rn.f32 	%f41, %f35, %f16, %f41;
	cvta.to.global.u64 	%rd35, %rd7;
	.loc 2 37 1
	add.s64 	%rd36, %rd35, %rd31;
	add.s64 	%rd37, %rd35, %rd33;
	ld.global.f32 	%f36, [%rd37];
	ld.global.f32 	%f37, [%rd36];
	sub.f32 	%f38, %f37, %f36;
	mul.f32 	%f39, %f38, %f16;
	sub.f32 	%f42, %f42, %f39;

BB0_8:
	.loc 2 41 1
	st.global.f32 	[%rd1], %f41;
	.loc 2 42 1
	st.global.f32 	[%rd2], %f42;
	.loc 2 43 1
	st.global.f32 	[%rd3], %f40;
	.loc 2 20 1
	add.s32 	%r98, %r98, %r6;
	.loc 2 20 18
	add.s32 	%r99, %r99, 1;
	.loc 2 20 1
	setp.lt.s32 	%p9, %r99, %r11;
	@%p9 bra 	BB0_2;

BB0_9:
	.loc 2 45 2
	ret;
}


`
	adddmi_ptx_35 = `
.version 3.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry adddmi(
	.param .u64 adddmi_param_0,
	.param .u64 adddmi_param_1,
	.param .u64 adddmi_param_2,
	.param .u64 adddmi_param_3,
	.param .u64 adddmi_param_4,
	.param .u64 adddmi_param_5,
	.param .f32 adddmi_param_6,
	.param .f32 adddmi_param_7,
	.param .f32 adddmi_param_8,
	.param .u32 adddmi_param_9,
	.param .u32 adddmi_param_10,
	.param .u32 adddmi_param_11
)
{
	.reg .pred 	%p<10>;
	.reg .s32 	%r<88>;
	.reg .f32 	%f<43>;
	.reg .s64 	%rd<38>;


	ld.param.u64 	%rd4, [adddmi_param_0];
	ld.param.u64 	%rd5, [adddmi_param_1];
	ld.param.u64 	%rd6, [adddmi_param_2];
	ld.param.u64 	%rd7, [adddmi_param_3];
	ld.param.u64 	%rd8, [adddmi_param_4];
	ld.param.u64 	%rd9, [adddmi_param_5];
	ld.param.f32 	%f16, [adddmi_param_6];
	ld.param.f32 	%f17, [adddmi_param_7];
	ld.param.f32 	%f18, [adddmi_param_8];
	ld.param.u32 	%r11, [adddmi_param_9];
	ld.param.u32 	%r12, [adddmi_param_10];
	ld.param.u32 	%r13, [adddmi_param_11];
	.loc 3 13 1
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r14, %ntid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r3, %r14, %r1, %r2;
	.loc 3 14 1
	mov.u32 	%r15, %ntid.y;
	mov.u32 	%r16, %ctaid.y;
	mov.u32 	%r17, %tid.y;
	mad.lo.s32 	%r4, %r15, %r16, %r17;
	.loc 3 16 1
	setp.lt.s32 	%p1, %r4, %r13;
	setp.lt.s32 	%p2, %r3, %r12;
	and.pred  	%p3, %p2, %p1;
	.loc 3 20 1
	setp.gt.s32 	%p4, %r11, 0;
	.loc 3 16 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB2_9;
	bra.uni 	BB2_1;

BB2_1:
	.loc 3 20 1
	mad.lo.s32 	%r86, %r13, %r3, %r4;
	mul.lo.s32 	%r6, %r13, %r12;
	mov.u32 	%r87, 0;
	cvta.to.global.u64 	%rd10, %rd4;
	cvta.to.global.u64 	%rd12, %rd5;
	cvta.to.global.u64 	%rd13, %rd6;

BB2_2:
	.loc 3 23 1
	mul.wide.s32 	%rd11, %r86, 4;
	add.s64 	%rd1, %rd10, %rd11;
	ld.global.f32 	%f41, [%rd1];
	add.s64 	%rd2, %rd12, %rd11;
	ld.global.f32 	%f42, [%rd2];
	add.s64 	%rd3, %rd13, %rd11;
	ld.global.f32 	%f40, [%rd3];
	.loc 3 25 1
	setp.eq.f32 	%p6, %f16, 0f00000000;
	@%p6 bra 	BB2_4;

	.loc 3 26 1
	add.s32 	%r28, %r87, 1;
	mov.u32 	%r29, 0;
	.loc 4 238 5
	max.s32 	%r30, %r28, %r29;
	.loc 3 26 1
	add.s32 	%r31, %r11, -1;
	.loc 4 210 5
	min.s32 	%r32, %r30, %r31;
	.loc 3 26 1
	add.s32 	%r33, %r12, -1;
	.loc 4 238 5
	max.s32 	%r34, %r3, %r29;
	.loc 4 210 5
	min.s32 	%r35, %r34, %r33;
	.loc 3 26 1
	mad.lo.s32 	%r36, %r32, %r12, %r35;
	add.s32 	%r37, %r13, -1;
	.loc 4 238 5
	max.s32 	%r38, %r4, %r29;
	.loc 4 210 5
	min.s32 	%r39, %r38, %r37;
	.loc 3 26 1
	mad.lo.s32 	%r40, %r36, %r13, %r39;
	cvta.to.global.u64 	%rd14, %rd7;
	.loc 3 26 1
	mul.wide.s32 	%rd15, %r40, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.nc.f32 	%f19, [%rd16];
	add.s32 	%r41, %r87, -1;
	.loc 4 238 5
	max.s32 	%r42, %r41, %r29;
	.loc 4 210 5
	min.s32 	%r43, %r42, %r31;
	.loc 3 26 1
	mad.lo.s32 	%r44, %r43, %r12, %r35;
	mad.lo.s32 	%r45, %r44, %r13, %r39;
	mul.wide.s32 	%rd17, %r45, 4;
	add.s64 	%rd18, %rd14, %rd17;
	ld.global.nc.f32 	%f20, [%rd18];
	sub.f32 	%f21, %f19, %f20;
	fma.rn.f32 	%f40, %f21, %f16, %f40;
	cvta.to.global.u64 	%rd19, %rd9;
	.loc 3 27 1
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.nc.f32 	%f22, [%rd20];
	add.s64 	%rd21, %rd19, %rd17;
	ld.global.nc.f32 	%f23, [%rd21];
	sub.f32 	%f24, %f22, %f23;
	mul.f32 	%f25, %f24, %f16;
	sub.f32 	%f41, %f41, %f25;

BB2_4:
	.loc 3 30 1
	setp.eq.f32 	%p7, %f17, 0f00000000;
	@%p7 bra 	BB2_6;

	mov.u32 	%r46, 0;
	.loc 4 238 5
	max.s32 	%r47, %r87, %r46;
	.loc 3 26 1
	add.s32 	%r48, %r11, -1;
	.loc 4 210 5
	min.s32 	%r49, %r47, %r48;
	.loc 3 31 1
	add.s32 	%r50, %r3, 1;
	.loc 4 238 5
	max.s32 	%r51, %r50, %r46;
	.loc 3 26 1
	add.s32 	%r52, %r12, -1;
	.loc 4 210 5
	min.s32 	%r53, %r51, %r52;
	.loc 3 31 1
	mad.lo.s32 	%r54, %r49, %r12, %r53;
	.loc 3 26 1
	add.s32 	%r55, %r13, -1;
	.loc 4 238 5
	max.s32 	%r56, %r4, %r46;
	.loc 4 210 5
	min.s32 	%r57, %r56, %r55;
	.loc 3 31 1
	mad.lo.s32 	%r58, %r54, %r13, %r57;
	cvta.to.global.u64 	%rd22, %rd8;
	.loc 3 31 1
	mul.wide.s32 	%rd23, %r58, 4;
	add.s64 	%rd24, %rd22, %rd23;
	ld.global.nc.f32 	%f26, [%rd24];
	add.s32 	%r59, %r3, -1;
	.loc 4 238 5
	max.s32 	%r60, %r59, %r46;
	.loc 4 210 5
	min.s32 	%r61, %r60, %r52;
	.loc 3 31 1
	mad.lo.s32 	%r62, %r49, %r12, %r61;
	add.s32 	%r63, %r4, -1;
	.loc 4 238 5
	max.s32 	%r64, %r63, %r46;
	.loc 4 210 5
	min.s32 	%r65, %r64, %r55;
	.loc 3 31 1
	mad.lo.s32 	%r66, %r62, %r13, %r65;
	mul.wide.s32 	%rd25, %r66, 4;
	add.s64 	%rd26, %rd22, %rd25;
	ld.global.nc.f32 	%f27, [%rd26];
	sub.f32 	%f28, %f26, %f27;
	fma.rn.f32 	%f40, %f28, %f17, %f40;
	cvta.to.global.u64 	%rd27, %rd9;
	.loc 3 32 1
	add.s64 	%rd28, %rd27, %rd23;
	ld.global.nc.f32 	%f29, [%rd28];
	add.s64 	%rd29, %rd27, %rd25;
	ld.global.nc.f32 	%f30, [%rd29];
	sub.f32 	%f31, %f29, %f30;
	mul.f32 	%f32, %f31, %f16;
	sub.f32 	%f42, %f42, %f32;

BB2_6:
	.loc 3 35 1
	setp.eq.f32 	%p8, %f18, 0f00000000;
	@%p8 bra 	BB2_8;

	mov.u32 	%r67, 0;
	.loc 4 238 5
	max.s32 	%r68, %r87, %r67;
	.loc 3 26 1
	add.s32 	%r69, %r11, -1;
	.loc 4 210 5
	min.s32 	%r70, %r68, %r69;
	.loc 3 26 1
	add.s32 	%r71, %r12, -1;
	.loc 4 238 5
	max.s32 	%r72, %r3, %r67;
	.loc 4 210 5
	min.s32 	%r73, %r72, %r71;
	.loc 3 36 1
	mad.lo.s32 	%r74, %r70, %r12, %r73;
	add.s32 	%r75, %r4, 1;
	.loc 4 238 5
	max.s32 	%r76, %r75, %r67;
	.loc 3 26 1
	add.s32 	%r77, %r13, -1;
	.loc 4 210 5
	min.s32 	%r78, %r76, %r77;
	.loc 3 36 1
	mad.lo.s32 	%r79, %r74, %r13, %r78;
	cvta.to.global.u64 	%rd30, %rd8;
	.loc 3 36 1
	mul.wide.s32 	%rd31, %r79, 4;
	add.s64 	%rd32, %rd30, %rd31;
	ld.global.nc.f32 	%f33, [%rd32];
	.loc 4 238 5
	max.s32 	%r80, %r4, %r67;
	.loc 4 210 5
	min.s32 	%r81, %r80, %r77;
	.loc 3 36 1
	mad.lo.s32 	%r82, %r74, %r13, %r81;
	mul.wide.s32 	%rd33, %r82, 4;
	add.s64 	%rd34, %rd30, %rd33;
	ld.global.nc.f32 	%f34, [%rd34];
	sub.f32 	%f35, %f33, %f34;
	fma.rn.f32 	%f41, %f35, %f16, %f41;
	cvta.to.global.u64 	%rd35, %rd7;
	.loc 3 37 1
	add.s64 	%rd36, %rd35, %rd31;
	ld.global.nc.f32 	%f36, [%rd36];
	add.s64 	%rd37, %rd35, %rd33;
	ld.global.nc.f32 	%f37, [%rd37];
	sub.f32 	%f38, %f36, %f37;
	mul.f32 	%f39, %f38, %f16;
	sub.f32 	%f42, %f42, %f39;

BB2_8:
	.loc 3 41 1
	st.global.f32 	[%rd1], %f41;
	.loc 3 42 1
	st.global.f32 	[%rd2], %f42;
	.loc 3 43 1
	st.global.f32 	[%rd3], %f40;
	.loc 3 20 1
	add.s32 	%r86, %r86, %r6;
	.loc 3 20 18
	add.s32 	%r87, %r87, 1;
	.loc 3 20 1
	setp.lt.s32 	%p9, %r87, %r11;
	@%p9 bra 	BB2_2;

BB2_9:
	.loc 3 45 2
	ret;
}


`
)
