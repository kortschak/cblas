// Do not manually edit this file. It was created by the genBlas.pl script from cblas.h.

// Copyright ©2012 The bíogo.blas Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cblas implements the blas interfaces.
package cblas

/*
#cgo CFLAGS: -g -O2 -fPIC -m64 -pthread
#cgo LDFLAGS: -L/usr/lib/ -lblas
#include "cblas.h"
*/
import "C"

import (
	"github.com/kortschak/blas"
	"unsafe"
)

// Type check assertions:
var (
	_ blas.Float32    = Blas{}
	_ blas.Float64    = Blas{}
	_ blas.Complex64  = Blas{}
	_ blas.Complex128 = Blas{}
)

type Blas struct{}

// Special cases...

func (Blas) Srotg(a float32, b float32) (c float32, s float32, r float32, z float32) {
	C.cblas_srotg((*C.float)(&a), (*C.float)(&b), (*C.float)(&c), (*C.float)(&s))
	return c, s, a, b
}
func (Blas) Srotmg(d1 float32, d2 float32, b1 float32, b2 float32) (p *blas.SrotmParams, rd1 float32, rd2 float32, rb1 float32) {
	p = &blas.SrotmParams{}
	C.cblas_srotmg((*C.float)(&d1), (*C.float)(&d2), (*C.float)(&b1), C.float(b2), (*C.float)(unsafe.Pointer(p)))
	return p, d1, d2, b1
}
func (Blas) Srotm(n int, x []float32, incX int, y []float32, incY int, p *blas.SrotmParams) {
	C.cblas_srotm(C.int(n), (*C.float)(&x[0]), C.int(incX), (*C.float)(&y[0]), C.int(incY), (*C.float)(unsafe.Pointer(p)))
}
func (Blas) Drotg(a float64, b float64) (c float64, s float64, r float64, z float64) {
	C.cblas_drotg((*C.double)(&a), (*C.double)(&b), (*C.double)(&c), (*C.double)(&s))
	return c, s, a, b
}
func (Blas) Drotmg(d1 float64, d2 float64, b1 float64, b2 float64) (p *blas.DrotmParams, rd1 float64, rd2 float64, rb1 float64) {
	p = &blas.DrotmParams{}
	C.cblas_drotmg((*C.double)(&d1), (*C.double)(&d2), (*C.double)(&b1), C.double(b2), (*C.double)(unsafe.Pointer(p)))
	return p, d1, d2, b1
}
func (Blas) Drotm(n int, x []float64, incX int, y []float64, incY int, p *blas.DrotmParams) {
	C.cblas_drotm(C.int(n), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY), (*C.double)(unsafe.Pointer(p)))
}

func (Blas) Sdsdot(n int, alpha float32, x []float32, incX int, y []float32, incY int) float32 {
	return float32(C.cblas_sdsdot(C.int(n), C.float(alpha), (*C.float)(&x[0]), C.int(incX), (*C.float)(&y[0]), C.int(incY)))
}
func (Blas) Dsdot(n int, x []float32, incX int, y []float32, incY int) float64 {
	return float64(C.cblas_dsdot(C.int(n), (*C.float)(&x[0]), C.int(incX), (*C.float)(&y[0]), C.int(incY)))
}
func (Blas) Sdot(n int, x []float32, incX int, y []float32, incY int) float32 {
	return float32(C.cblas_sdot(C.int(n), (*C.float)(&x[0]), C.int(incX), (*C.float)(&y[0]), C.int(incY)))
}
func (Blas) Ddot(n int, x []float64, incX int, y []float64, incY int) float64 {
	return float64(C.cblas_ddot(C.int(n), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY)))
}
func (Blas) Cdotu(n int, x []complex64, incX int, y []complex64, incY int, dotu []complex64) {
	C.cblas_cdotu_sub(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&dotu[0]))
}
func (Blas) Cdotc(n int, x []complex64, incX int, y []complex64, incY int, dotc []complex64) {
	C.cblas_cdotc_sub(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&dotc[0]))
}
func (Blas) Zdotu(n int, x []complex128, incX int, y []complex128, incY int, dotu []complex128) {
	C.cblas_zdotu_sub(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&dotu[0]))
}
func (Blas) Zdotc(n int, x []complex128, incX int, y []complex128, incY int, dotc []complex128) {
	C.cblas_zdotc_sub(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&dotc[0]))
}
func (Blas) Snrm2(n int, x []float32, incX int) float32 {
	return float32(C.cblas_snrm2(C.int(n), (*C.float)(&x[0]), C.int(incX)))
}
func (Blas) Sasum(n int, x []float32, incX int) float32 {
	return float32(C.cblas_sasum(C.int(n), (*C.float)(&x[0]), C.int(incX)))
}
func (Blas) Dnrm2(n int, x []float64, incX int) float64 {
	return float64(C.cblas_dnrm2(C.int(n), (*C.double)(&x[0]), C.int(incX)))
}
func (Blas) Dasum(n int, x []float64, incX int) float64 {
	return float64(C.cblas_dasum(C.int(n), (*C.double)(&x[0]), C.int(incX)))
}
func (Blas) Scnrm2(n int, x []complex64, incX int) float32 {
	return float32(C.cblas_scnrm2(C.int(n), unsafe.Pointer(&x[0]), C.int(incX)))
}
func (Blas) Scasum(n int, x []complex64, incX int) float32 {
	return float32(C.cblas_scasum(C.int(n), unsafe.Pointer(&x[0]), C.int(incX)))
}
func (Blas) Dznrm2(n int, x []complex128, incX int) float64 {
	return float64(C.cblas_dznrm2(C.int(n), unsafe.Pointer(&x[0]), C.int(incX)))
}
func (Blas) Dzasum(n int, x []complex128, incX int) float64 {
	return float64(C.cblas_dzasum(C.int(n), unsafe.Pointer(&x[0]), C.int(incX)))
}
func (Blas) Isamax(n int, x []float32, incX int) int {
	return int(C.cblas_isamax(C.int(n), (*C.float)(&x[0]), C.int(incX)))
}
func (Blas) Idamax(n int, x []float64, incX int) int {
	return int(C.cblas_idamax(C.int(n), (*C.double)(&x[0]), C.int(incX)))
}
func (Blas) Icamax(n int, x []complex64, incX int) int {
	return int(C.cblas_icamax(C.int(n), unsafe.Pointer(&x[0]), C.int(incX)))
}
func (Blas) Izamax(n int, x []complex128, incX int) int {
	return int(C.cblas_izamax(C.int(n), unsafe.Pointer(&x[0]), C.int(incX)))
}
func (Blas) Sswap(n int, x []float32, incX int, y []float32, incY int) {
	C.cblas_sswap(C.int(n), (*C.float)(&x[0]), C.int(incX), (*C.float)(&y[0]), C.int(incY))
}
func (Blas) Scopy(n int, x []float32, incX int, y []float32, incY int) {
	C.cblas_scopy(C.int(n), (*C.float)(&x[0]), C.int(incX), (*C.float)(&y[0]), C.int(incY))
}
func (Blas) Saxpy(n int, alpha float32, x []float32, incX int, y []float32, incY int) {
	C.cblas_saxpy(C.int(n), C.float(alpha), (*C.float)(&x[0]), C.int(incX), (*C.float)(&y[0]), C.int(incY))
}
func (Blas) Dswap(n int, x []float64, incX int, y []float64, incY int) {
	C.cblas_dswap(C.int(n), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY))
}
func (Blas) Dcopy(n int, x []float64, incX int, y []float64, incY int) {
	C.cblas_dcopy(C.int(n), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY))
}
func (Blas) Daxpy(n int, alpha float64, x []float64, incX int, y []float64, incY int) {
	C.cblas_daxpy(C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY))
}
func (Blas) Cswap(n int, x []complex64, incX int, y []complex64, incY int) {
	C.cblas_cswap(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Ccopy(n int, x []complex64, incX int, y []complex64, incY int) {
	C.cblas_ccopy(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Caxpy(n int, alpha complex64, x []complex64, incX int, y []complex64, incY int) {
	C.cblas_caxpy(C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Zswap(n int, x []complex128, incX int, y []complex128, incY int) {
	C.cblas_zswap(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Zcopy(n int, x []complex128, incX int, y []complex128, incY int) {
	C.cblas_zcopy(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Zaxpy(n int, alpha complex128, x []complex128, incX int, y []complex128, incY int) {
	C.cblas_zaxpy(C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Srot(n int, x []float32, incX int, y []float32, incY int, c float32, s float32) {
	C.cblas_srot(C.int(n), (*C.float)(&x[0]), C.int(incX), (*C.float)(&y[0]), C.int(incY), C.float(c), C.float(s))
}
func (Blas) Drot(n int, x []float64, incX int, y []float64, incY int, c float64, s float64) {
	C.cblas_drot(C.int(n), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY), C.double(c), C.double(s))
}
func (Blas) Sscal(n int, alpha float32, x []float32, incX int) {
	C.cblas_sscal(C.int(n), C.float(alpha), (*C.float)(&x[0]), C.int(incX))
}
func (Blas) Dscal(n int, alpha float64, x []float64, incX int) {
	C.cblas_dscal(C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX))
}
func (Blas) Cscal(n int, alpha complex64, x []complex64, incX int) {
	C.cblas_cscal(C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Zscal(n int, alpha complex128, x []complex128, incX int) {
	C.cblas_zscal(C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Csscal(n int, alpha float32, x []complex64, incX int) {
	C.cblas_csscal(C.int(n), C.float(alpha), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Zdscal(n int, alpha float64, x []complex128, incX int) {
	C.cblas_zdscal(C.int(n), C.double(alpha), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Sgemv(o blas.Order, tA blas.Transpose, m int, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	C.cblas_sgemv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.int(m), C.int(n), C.float(alpha), (*C.float)(&a[0]), C.int(lda), (*C.float)(&x[0]), C.int(incX), C.float(beta), (*C.float)(&y[0]), C.int(incY))
}
func (Blas) Sgbmv(o blas.Order, tA blas.Transpose, m int, n int, kL int, kU int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	C.cblas_sgbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.int(m), C.int(n), C.int(kL), C.int(kU), C.float(alpha), (*C.float)(&a[0]), C.int(lda), (*C.float)(&x[0]), C.int(incX), C.float(beta), (*C.float)(&y[0]), C.int(incY))
}
func (Blas) Strmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float32, lda int, x []float32, incX int) {
	C.cblas_strmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), (*C.float)(&a[0]), C.int(lda), (*C.float)(&x[0]), C.int(incX))
}
func (Blas) Stbmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []float32, lda int, x []float32, incX int) {
	C.cblas_stbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), C.int(k), (*C.float)(&a[0]), C.int(lda), (*C.float)(&x[0]), C.int(incX))
}
func (Blas) Stpmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float32, x []float32, incX int) {
	C.cblas_stpmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), (*C.float)(&ap[0]), (*C.float)(&x[0]), C.int(incX))
}
func (Blas) Strsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float32, lda int, x []float32, incX int) {
	C.cblas_strsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), (*C.float)(&a[0]), C.int(lda), (*C.float)(&x[0]), C.int(incX))
}
func (Blas) Stbsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []float32, lda int, x []float32, incX int) {
	C.cblas_stbsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), C.int(k), (*C.float)(&a[0]), C.int(lda), (*C.float)(&x[0]), C.int(incX))
}
func (Blas) Stpsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float32, x []float32, incX int) {
	C.cblas_stpsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), (*C.float)(&ap[0]), (*C.float)(&x[0]), C.int(incX))
}
func (Blas) Dgemv(o blas.Order, tA blas.Transpose, m int, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dgemv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.int(m), C.int(n), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX), C.double(beta), (*C.double)(&y[0]), C.int(incY))
}
func (Blas) Dgbmv(o blas.Order, tA blas.Transpose, m int, n int, kL int, kU int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dgbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.int(m), C.int(n), C.int(kL), C.int(kU), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX), C.double(beta), (*C.double)(&y[0]), C.int(incY))
}
func (Blas) Dtrmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtrmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX))
}
func (Blas) Dtbmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), C.int(k), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX))
}
func (Blas) Dtpmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float64, x []float64, incX int) {
	C.cblas_dtpmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), (*C.double)(&ap[0]), (*C.double)(&x[0]), C.int(incX))
}
func (Blas) Dtrsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtrsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX))
}
func (Blas) Dtbsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtbsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), C.int(k), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX))
}
func (Blas) Dtpsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float64, x []float64, incX int) {
	C.cblas_dtpsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), (*C.double)(&ap[0]), (*C.double)(&x[0]), C.int(incX))
}
func (Blas) Cgemv(o blas.Order, tA blas.Transpose, m int, n int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	C.cblas_cgemv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Cgbmv(o blas.Order, tA blas.Transpose, m int, n int, kL int, kU int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	C.cblas_cgbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.int(m), C.int(n), C.int(kL), C.int(kU), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Ctrmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []complex64, lda int, x []complex64, incX int) {
	C.cblas_ctrmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Ctbmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []complex64, lda int, x []complex64, incX int) {
	C.cblas_ctbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), C.int(k), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Ctpmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []complex64, x []complex64, incX int) {
	C.cblas_ctpmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), unsafe.Pointer(&ap[0]), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Ctrsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []complex64, lda int, x []complex64, incX int) {
	C.cblas_ctrsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Ctbsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []complex64, lda int, x []complex64, incX int) {
	C.cblas_ctbsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), C.int(k), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Ctpsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []complex64, x []complex64, incX int) {
	C.cblas_ctpsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), unsafe.Pointer(&ap[0]), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Zgemv(o blas.Order, tA blas.Transpose, m int, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	C.cblas_zgemv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Zgbmv(o blas.Order, tA blas.Transpose, m int, n int, kL int, kU int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	C.cblas_zgbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.int(m), C.int(n), C.int(kL), C.int(kU), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Ztrmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []complex128, lda int, x []complex128, incX int) {
	C.cblas_ztrmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Ztbmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []complex128, lda int, x []complex128, incX int) {
	C.cblas_ztbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), C.int(k), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Ztpmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []complex128, x []complex128, incX int) {
	C.cblas_ztpmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), unsafe.Pointer(&ap[0]), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Ztrsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []complex128, lda int, x []complex128, incX int) {
	C.cblas_ztrsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Ztbsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []complex128, lda int, x []complex128, incX int) {
	C.cblas_ztbsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), C.int(k), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Ztpsv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []complex128, x []complex128, incX int) {
	C.cblas_ztpsv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(n), unsafe.Pointer(&ap[0]), unsafe.Pointer(&x[0]), C.int(incX))
}
func (Blas) Ssymv(o blas.Order, ul blas.Uplo, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	C.cblas_ssymv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.float(alpha), (*C.float)(&a[0]), C.int(lda), (*C.float)(&x[0]), C.int(incX), C.float(beta), (*C.float)(&y[0]), C.int(incY))
}
func (Blas) Ssbmv(o blas.Order, ul blas.Uplo, n int, k int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	C.cblas_ssbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.int(k), C.float(alpha), (*C.float)(&a[0]), C.int(lda), (*C.float)(&x[0]), C.int(incX), C.float(beta), (*C.float)(&y[0]), C.int(incY))
}
func (Blas) Sspmv(o blas.Order, ul blas.Uplo, n int, alpha float32, ap []float32, x []float32, incX int, beta float32, y []float32, incY int) {
	C.cblas_sspmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.float(alpha), (*C.float)(&ap[0]), (*C.float)(&x[0]), C.int(incX), C.float(beta), (*C.float)(&y[0]), C.int(incY))
}
func (Blas) Sger(o blas.Order, m int, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int) {
	C.cblas_sger(C.enum_CBLAS_ORDER(o), C.int(m), C.int(n), C.float(alpha), (*C.float)(&x[0]), C.int(incX), (*C.float)(&y[0]), C.int(incY), (*C.float)(&a[0]), C.int(lda))
}
func (Blas) Ssyr(o blas.Order, ul blas.Uplo, n int, alpha float32, x []float32, incX int, a []float32, lda int) {
	C.cblas_ssyr(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.float(alpha), (*C.float)(&x[0]), C.int(incX), (*C.float)(&a[0]), C.int(lda))
}
func (Blas) Sspr(o blas.Order, ul blas.Uplo, n int, alpha float32, x []float32, incX int, ap []float32) {
	C.cblas_sspr(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.float(alpha), (*C.float)(&x[0]), C.int(incX), (*C.float)(&ap[0]))
}
func (Blas) Ssyr2(o blas.Order, ul blas.Uplo, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int) {
	C.cblas_ssyr2(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.float(alpha), (*C.float)(&x[0]), C.int(incX), (*C.float)(&y[0]), C.int(incY), (*C.float)(&a[0]), C.int(lda))
}
func (Blas) Sspr2(o blas.Order, ul blas.Uplo, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32) {
	C.cblas_sspr2(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.float(alpha), (*C.float)(&x[0]), C.int(incX), (*C.float)(&y[0]), C.int(incY), (*C.float)(&a[0]))
}
func (Blas) Dsymv(o blas.Order, ul blas.Uplo, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dsymv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX), C.double(beta), (*C.double)(&y[0]), C.int(incY))
}
func (Blas) Dsbmv(o blas.Order, ul blas.Uplo, n int, k int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dsbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.int(k), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&x[0]), C.int(incX), C.double(beta), (*C.double)(&y[0]), C.int(incY))
}
func (Blas) Dspmv(o blas.Order, ul blas.Uplo, n int, alpha float64, ap []float64, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dspmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), (*C.double)(&ap[0]), (*C.double)(&x[0]), C.int(incX), C.double(beta), (*C.double)(&y[0]), C.int(incY))
}
func (Blas) Dger(o blas.Order, m int, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	C.cblas_dger(C.enum_CBLAS_ORDER(o), C.int(m), C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY), (*C.double)(&a[0]), C.int(lda))
}
func (Blas) Dsyr(o blas.Order, ul blas.Uplo, n int, alpha float64, x []float64, incX int, a []float64, lda int) {
	C.cblas_dsyr(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX), (*C.double)(&a[0]), C.int(lda))
}
func (Blas) Dspr(o blas.Order, ul blas.Uplo, n int, alpha float64, x []float64, incX int, ap []float64) {
	C.cblas_dspr(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX), (*C.double)(&ap[0]))
}
func (Blas) Dsyr2(o blas.Order, ul blas.Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	C.cblas_dsyr2(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY), (*C.double)(&a[0]), C.int(lda))
}
func (Blas) Dspr2(o blas.Order, ul blas.Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64) {
	C.cblas_dspr2(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), (*C.double)(&x[0]), C.int(incX), (*C.double)(&y[0]), C.int(incY), (*C.double)(&a[0]))
}
func (Blas) Chemv(o blas.Order, ul blas.Uplo, n int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	C.cblas_chemv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Chbmv(o blas.Order, ul blas.Uplo, n int, k int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	C.cblas_chbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.int(k), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Chpmv(o blas.Order, ul blas.Uplo, n int, alpha complex64, ap []complex64, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	C.cblas_chpmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&ap[0]), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Cgeru(o blas.Order, m int, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	C.cblas_cgeru(C.enum_CBLAS_ORDER(o), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&a[0]), C.int(lda))
}
func (Blas) Cgerc(o blas.Order, m int, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	C.cblas_cgerc(C.enum_CBLAS_ORDER(o), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&a[0]), C.int(lda))
}
func (Blas) Cher(o blas.Order, ul blas.Uplo, n int, alpha float32, x []complex64, incX int, a []complex64, lda int) {
	C.cblas_cher(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.float(alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&a[0]), C.int(lda))
}
func (Blas) Chpr(o blas.Order, ul blas.Uplo, n int, alpha float32, x []complex64, incX int, a []complex64) {
	C.cblas_chpr(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.float(alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&a[0]))
}
func (Blas) Cher2(o blas.Order, ul blas.Uplo, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	C.cblas_cher2(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&a[0]), C.int(lda))
}
func (Blas) Chpr2(o blas.Order, ul blas.Uplo, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, ap []complex64) {
	C.cblas_chpr2(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&ap[0]))
}
func (Blas) Zhemv(o blas.Order, ul blas.Uplo, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	C.cblas_zhemv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Zhbmv(o blas.Order, ul blas.Uplo, n int, k int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	C.cblas_zhbmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.int(k), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Zhpmv(o blas.Order, ul blas.Uplo, n int, alpha complex128, ap []complex128, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	C.cblas_zhpmv(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&ap[0]), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.int(incY))
}
func (Blas) Zgeru(o blas.Order, m int, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	C.cblas_zgeru(C.enum_CBLAS_ORDER(o), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&a[0]), C.int(lda))
}
func (Blas) Zgerc(o blas.Order, m int, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	C.cblas_zgerc(C.enum_CBLAS_ORDER(o), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&a[0]), C.int(lda))
}
func (Blas) Zher(o blas.Order, ul blas.Uplo, n int, alpha float64, x []complex128, incX int, a []complex128, lda int) {
	C.cblas_zher(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&a[0]), C.int(lda))
}
func (Blas) Zhpr(o blas.Order, ul blas.Uplo, n int, alpha float64, x []complex128, incX int, a []complex128) {
	C.cblas_zhpr(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.double(alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&a[0]))
}
func (Blas) Zher2(o blas.Order, ul blas.Uplo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	C.cblas_zher2(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&a[0]), C.int(lda))
}
func (Blas) Zhpr2(o blas.Order, ul blas.Uplo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, ap []complex128) {
	C.cblas_zhpr2(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&ap[0]))
}
func (Blas) Sgemm(o blas.Order, tA blas.Transpose, tB blas.Transpose, m int, n int, k int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int) {
	C.cblas_sgemm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_TRANSPOSE(tB), C.int(m), C.int(n), C.int(k), C.float(alpha), (*C.float)(&a[0]), C.int(lda), (*C.float)(&b[0]), C.int(ldb), C.float(beta), (*C.float)(&c[0]), C.int(ldc))
}
func (Blas) Ssymm(o blas.Order, s blas.Side, ul blas.Uplo, m int, n int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int) {
	C.cblas_ssymm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.int(m), C.int(n), C.float(alpha), (*C.float)(&a[0]), C.int(lda), (*C.float)(&b[0]), C.int(ldb), C.float(beta), (*C.float)(&c[0]), C.int(ldc))
}
func (Blas) Ssyrk(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha float32, a []float32, lda int, beta float32, c []float32, ldc int) {
	C.cblas_ssyrk(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), C.float(alpha), (*C.float)(&a[0]), C.int(lda), C.float(beta), (*C.float)(&c[0]), C.int(ldc))
}
func (Blas) Ssyr2k(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int) {
	C.cblas_ssyr2k(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), C.float(alpha), (*C.float)(&a[0]), C.int(lda), (*C.float)(&b[0]), C.int(ldb), C.float(beta), (*C.float)(&c[0]), C.int(ldc))
}
func (Blas) Strmm(o blas.Order, s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha float32, a []float32, lda int, b []float32, ldb int) {
	C.cblas_strmm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(m), C.int(n), C.float(alpha), (*C.float)(&a[0]), C.int(lda), (*C.float)(&b[0]), C.int(ldb))
}
func (Blas) Strsm(o blas.Order, s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha float32, a []float32, lda int, b []float32, ldb int) {
	C.cblas_strsm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(m), C.int(n), C.float(alpha), (*C.float)(&a[0]), C.int(lda), (*C.float)(&b[0]), C.int(ldb))
}
func (Blas) Dgemm(o blas.Order, tA blas.Transpose, tB blas.Transpose, m int, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	C.cblas_dgemm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_TRANSPOSE(tB), C.int(m), C.int(n), C.int(k), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&b[0]), C.int(ldb), C.double(beta), (*C.double)(&c[0]), C.int(ldc))
}
func (Blas) Dsymm(o blas.Order, s blas.Side, ul blas.Uplo, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	C.cblas_dsymm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.int(m), C.int(n), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&b[0]), C.int(ldb), C.double(beta), (*C.double)(&c[0]), C.int(ldc))
}
func (Blas) Dsyrk(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha float64, a []float64, lda int, beta float64, c []float64, ldc int) {
	C.cblas_dsyrk(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), C.double(alpha), (*C.double)(&a[0]), C.int(lda), C.double(beta), (*C.double)(&c[0]), C.int(ldc))
}
func (Blas) Dsyr2k(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	C.cblas_dsyr2k(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&b[0]), C.int(ldb), C.double(beta), (*C.double)(&c[0]), C.int(ldc))
}
func (Blas) Dtrmm(o blas.Order, s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int) {
	C.cblas_dtrmm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(m), C.int(n), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&b[0]), C.int(ldb))
}
func (Blas) Dtrsm(o blas.Order, s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int) {
	C.cblas_dtrsm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(m), C.int(n), C.double(alpha), (*C.double)(&a[0]), C.int(lda), (*C.double)(&b[0]), C.int(ldb))
}
func (Blas) Cgemm(o blas.Order, tA blas.Transpose, tB blas.Transpose, m int, n int, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int) {
	C.cblas_cgemm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_TRANSPOSE(tB), C.int(m), C.int(n), C.int(k), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb), unsafe.Pointer(&beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
func (Blas) Csymm(o blas.Order, s blas.Side, ul blas.Uplo, m int, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int) {
	C.cblas_csymm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb), unsafe.Pointer(&beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
func (Blas) Csyrk(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha complex64, a []complex64, lda int, beta complex64, c []complex64, ldc int) {
	C.cblas_csyrk(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
func (Blas) Csyr2k(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int) {
	C.cblas_csyr2k(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb), unsafe.Pointer(&beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
func (Blas) Ctrmm(o blas.Order, s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int) {
	C.cblas_ctrmm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb))
}
func (Blas) Ctrsm(o blas.Order, s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int) {
	C.cblas_ctrsm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb))
}
func (Blas) Zgemm(o blas.Order, tA blas.Transpose, tB blas.Transpose, m int, n int, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) {
	C.cblas_zgemm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_TRANSPOSE(tB), C.int(m), C.int(n), C.int(k), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb), unsafe.Pointer(&beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
func (Blas) Zsymm(o blas.Order, s blas.Side, ul blas.Uplo, m int, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) {
	C.cblas_zsymm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb), unsafe.Pointer(&beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
func (Blas) Zsyrk(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha complex128, a []complex128, lda int, beta complex128, c []complex128, ldc int) {
	C.cblas_zsyrk(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
func (Blas) Zsyr2k(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) {
	C.cblas_zsyr2k(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb), unsafe.Pointer(&beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
func (Blas) Ztrmm(o blas.Order, s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int) {
	C.cblas_ztrmm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb))
}
func (Blas) Ztrsm(o blas.Order, s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int) {
	C.cblas_ztrsm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(tA), C.enum_CBLAS_DIAG(d), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb))
}
func (Blas) Chemm(o blas.Order, s blas.Side, ul blas.Uplo, m int, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int) {
	C.cblas_chemm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb), unsafe.Pointer(&beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
func (Blas) Cherk(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha float32, a []complex64, lda int, beta float32, c []complex64, ldc int) {
	C.cblas_cherk(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), C.float(alpha), unsafe.Pointer(&a[0]), C.int(lda), C.float(beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
func (Blas) Cher2k(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta float32, c []complex64, ldc int) {
	C.cblas_cher2k(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb), C.float(beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
func (Blas) Zhemm(o blas.Order, s blas.Side, ul blas.Uplo, m int, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) {
	C.cblas_zhemm(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_SIDE(s), C.enum_CBLAS_UPLO(ul), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb), unsafe.Pointer(&beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
func (Blas) Zherk(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha float64, a []complex128, lda int, beta float64, c []complex128, ldc int) {
	C.cblas_zherk(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), C.double(alpha), unsafe.Pointer(&a[0]), C.int(lda), C.double(beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
func (Blas) Zher2k(o blas.Order, ul blas.Uplo, t blas.Transpose, n int, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta float64, c []complex128, ldc int) {
	C.cblas_zher2k(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(k), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb), C.double(beta), unsafe.Pointer(&c[0]), C.int(ldc))
}
