package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"sync"
	"unsafe"
)

// CUDA handle for settemperature kernel
var settemperature_code cu.Function

// Stores the arguments for settemperature kernel invocation
type settemperature_args_t struct {
	arg_B            unsafe.Pointer
	arg_noise        unsafe.Pointer
	arg_kB2_VgammaDt float32
	arg_tempRedLUT   unsafe.Pointer
	arg_regions      unsafe.Pointer
	arg_N            int
	argptr           [6]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for settemperature kernel invocation
var settemperature_args settemperature_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	settemperature_args.argptr[0] = unsafe.Pointer(&settemperature_args.arg_B)
	settemperature_args.argptr[1] = unsafe.Pointer(&settemperature_args.arg_noise)
	settemperature_args.argptr[2] = unsafe.Pointer(&settemperature_args.arg_kB2_VgammaDt)
	settemperature_args.argptr[3] = unsafe.Pointer(&settemperature_args.arg_tempRedLUT)
	settemperature_args.argptr[4] = unsafe.Pointer(&settemperature_args.arg_regions)
	settemperature_args.argptr[5] = unsafe.Pointer(&settemperature_args.arg_N)
}

// Wrapper for settemperature CUDA kernel, asynchronous.
func k_settemperature_async(B unsafe.Pointer, noise unsafe.Pointer, kB2_VgammaDt float32, tempRedLUT unsafe.Pointer, regions unsafe.Pointer, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	settemperature_args.Lock()
	defer settemperature_args.Unlock()

	if settemperature_code == 0 {
		settemperature_code = fatbinLoad(settemperature_map, "settemperature")
	}

	settemperature_args.arg_B = B
	settemperature_args.arg_noise = noise
	settemperature_args.arg_kB2_VgammaDt = kB2_VgammaDt
	settemperature_args.arg_tempRedLUT = tempRedLUT
	settemperature_args.arg_regions = regions
	settemperature_args.arg_N = N

	args := settemperature_args.argptr[:]
	cu.LaunchKernel(settemperature_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

// maps compute capability on PTX code for settemperature kernel.
var settemperature_map = map[int]string{0: "",
	20: settemperature_ptx_20,
	30: settemperature_ptx_30,
	35: settemperature_ptx_35}

// settemperature PTX code for various compute capabilities.
const (
	settemperature_ptx_20 = `
.version 4.1
.target sm_20
.address_size 64


.visible .entry settemperature(
	.param .u64 settemperature_param_0,
	.param .u64 settemperature_param_1,
	.param .f32 settemperature_param_2,
	.param .u64 settemperature_param_3,
	.param .u64 settemperature_param_4,
	.param .u32 settemperature_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<7>;
	.reg .s64 	%rd<17>;


	ld.param.u64 	%rd1, [settemperature_param_0];
	ld.param.u64 	%rd2, [settemperature_param_1];
	ld.param.f32 	%f1, [settemperature_param_2];
	ld.param.u64 	%rd3, [settemperature_param_3];
	ld.param.u64 	%rd4, [settemperature_param_4];
	ld.param.u32 	%r2, [settemperature_param_5];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd5, %rd1;
	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd3;
	cvta.to.global.u64 	%rd8, %rd4;
	cvt.s64.s32	%rd9, %r1;
	add.s64 	%rd10, %rd8, %rd9;
	ld.global.u8 	%rd11, [%rd10];
	shl.b64 	%rd12, %rd11, 2;
	add.s64 	%rd13, %rd7, %rd12;
	mul.wide.s32 	%rd14, %r1, 4;
	add.s64 	%rd15, %rd6, %rd14;
	ld.global.f32 	%f2, [%rd13];
	mul.f32 	%f3, %f2, %f1;
	sqrt.rn.f32 	%f4, %f3;
	ld.global.f32 	%f5, [%rd15];
	mul.f32 	%f6, %f5, %f4;
	add.s64 	%rd16, %rd5, %rd14;
	st.global.f32 	[%rd16], %f6;

BB0_2:
	ret;
}


`
	settemperature_ptx_30 = `
.version 4.1
.target sm_30
.address_size 64


.visible .entry settemperature(
	.param .u64 settemperature_param_0,
	.param .u64 settemperature_param_1,
	.param .f32 settemperature_param_2,
	.param .u64 settemperature_param_3,
	.param .u64 settemperature_param_4,
	.param .u32 settemperature_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<7>;
	.reg .s64 	%rd<17>;


	ld.param.u64 	%rd1, [settemperature_param_0];
	ld.param.u64 	%rd2, [settemperature_param_1];
	ld.param.f32 	%f1, [settemperature_param_2];
	ld.param.u64 	%rd3, [settemperature_param_3];
	ld.param.u64 	%rd4, [settemperature_param_4];
	ld.param.u32 	%r2, [settemperature_param_5];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd5, %rd1;
	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd3;
	cvta.to.global.u64 	%rd8, %rd4;
	cvt.s64.s32	%rd9, %r1;
	add.s64 	%rd10, %rd8, %rd9;
	ld.global.u8 	%rd11, [%rd10];
	shl.b64 	%rd12, %rd11, 2;
	add.s64 	%rd13, %rd7, %rd12;
	mul.wide.s32 	%rd14, %r1, 4;
	add.s64 	%rd15, %rd6, %rd14;
	ld.global.f32 	%f2, [%rd13];
	mul.f32 	%f3, %f2, %f1;
	sqrt.rn.f32 	%f4, %f3;
	ld.global.f32 	%f5, [%rd15];
	mul.f32 	%f6, %f5, %f4;
	add.s64 	%rd16, %rd5, %rd14;
	st.global.f32 	[%rd16], %f6;

BB0_2:
	ret;
}


`
	settemperature_ptx_35 = `
.version 4.1
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
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.visible .entry settemperature(
	.param .u64 settemperature_param_0,
	.param .u64 settemperature_param_1,
	.param .f32 settemperature_param_2,
	.param .u64 settemperature_param_3,
	.param .u64 settemperature_param_4,
	.param .u32 settemperature_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<7>;
	.reg .s64 	%rd<18>;


	ld.param.u64 	%rd1, [settemperature_param_0];
	ld.param.u64 	%rd2, [settemperature_param_1];
	ld.param.f32 	%f1, [settemperature_param_2];
	ld.param.u64 	%rd3, [settemperature_param_3];
	ld.param.u64 	%rd4, [settemperature_param_4];
	ld.param.u32 	%r2, [settemperature_param_5];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB5_2;

	cvta.to.global.u64 	%rd5, %rd1;
	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd3;
	cvta.to.global.u64 	%rd8, %rd4;
	cvt.s64.s32	%rd9, %r1;
	add.s64 	%rd10, %rd8, %rd9;
	ld.global.nc.u8 	%rs1, [%rd10];
	cvt.u64.u16	%rd11, %rs1;
	and.b64  	%rd12, %rd11, 255;
	shl.b64 	%rd13, %rd12, 2;
	add.s64 	%rd14, %rd7, %rd13;
	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd6, %rd15;
	ld.global.nc.f32 	%f2, [%rd14];
	mul.f32 	%f3, %f2, %f1;
	sqrt.rn.f32 	%f4, %f3;
	ld.global.nc.f32 	%f5, [%rd16];
	mul.f32 	%f6, %f5, %f4;
	add.s64 	%rd17, %rd5, %rd15;
	st.global.f32 	[%rd17], %f6;

BB5_2:
	ret;
}


`
)
