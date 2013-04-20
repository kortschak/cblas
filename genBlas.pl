#!/usr/bin/env perl
# Copyright ©2012 The bíogo.blas Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

use strict;
use warnings;

my $cblasHeader = "cblas.h";
my $LIB = "/usr/lib/";

my $excludeComplex = 0;
my $excludeAtlas = 1;


open(my $cblas, "<", $cblasHeader) or die;
open(my $goblas, ">", "blas.go") or die;

my %done = ("cblas_errprn" => 1,
	        "cblas_srotg"  => 1,
	        "cblas_srotmg" => 1,
	        "cblas_srotm"  => 1,
	        "cblas_drotg"  => 1,
	        "cblas_drotmg" => 1,
	        "cblas_drotm"  => 1,
	        "cblas_crotg"  => 1,
	        "cblas_zrotg"  => 1,
	        );

my $atlas = "";
if ($excludeAtlas) {
	$done{'cblas_csrot'} = 1;
	$done{'cblas_zdrot'} = 1;
} else {
	$atlas = " -latlas";
}
printf $goblas <<EOH;
// Do not manually edit this file. It was created by the genBlas.pl script from ${cblasHeader}.

// Copyright ©2012 The bíogo.blas Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cblas implements the blas interfaces.
package cblas

/*
#cgo CFLAGS: -g -O2 -fPIC -m64 -pthread
#cgo LDFLAGS: -L${LIB} -lblas${atlas}
#include "${cblasHeader}"
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
EOH

printf $goblas <<EOH unless $excludeAtlas;
func (Blas) Crotg(a complex64, b complex64) (c complex64, s complex64, r complex64, z complex64) {
	C.cblas_srotg(unsafe.Pointer(&a), unsafe.Pointer(&b), unsafe.Pointer(&c), unsafe.Pointer(&s))
	return c, s, a, b
}
func (Blas) Zrotg(a complex128, b complex128) (c complex128, s complex128, r complex128, z complex128) {
	C.cblas_drotg(unsafe.Pointer(&a), unsafe.Pointer(&b), unsafe.Pointer(&c), unsafe.Pointer(&s))
	return c, s, a, b
}
EOH

print $goblas "\n";

$/ = undef;
my $header = <$cblas>;

# horrible munging of text...
$header =~ s/#[^\n\r]*//g;                 # delete cpp lines
$header =~ s/\n +([^\n\r]*)/\n$1/g;        # remove starting space
$header =~ s/(?:\n ?\n)+/\n/g;             # delete empty lines
$header =~ s! ((['"]) (?: \\. | .)*? \2) | # skip quoted strings
             /\* .*? \*/ |                 # delete C comments
             // [^\n\r]*                   # delete C++ comments just in case
             ! $1 || ' '                   # change comments to a single space
             !xseg;    	                   # ignore white space, treat as single line
                                           # evaluate result, repeat globally
$header =~ s/([^;])\n/$1/g;                # join prototypes into single lines
$header =~ s/, +/,/g;
$header =~ s/ +/ /g;
$header =~ s/ +}/}/g;
$header =~ s/\n+//;

$/ = "\n";
my @lines = split ";\n", $header;

our %retConv = (
	"int" => "int ",
	"float" => "float32 ",
	"double" => "float64 ",
	"CBLAS_INDEX" => "int ",
	"void" => ""
);

foreach my $line (@lines) {
	process($line);
}

close($goblas);
`go fmt .`;

sub process {
	my $line = shift;
	chomp $line;
	if (not $line =~ m/^enum/) {
		processProto($line);
	}
}

sub processProto {
	my $proto = shift;
	my ($func, $paramList) = split /[()]/, $proto;
	(my $ret, $func) = split ' ', $func;
	if ($done{$func} or $excludeComplex && $func =~ m/_[isd]?[zc]/ or $excludeAtlas && $func =~ m/^catlas_/) {
		return
	}
	$done{$func} = 1;
	my $GoRet = $retConv{$ret};
	my $complexType = $func;
	$complexType =~ s/.*_[isd]?([zc]).*/$1/;
	print $goblas "func (Blas) ".Gofunc($func)."(".processParamToGo($func, $paramList, $complexType).") ".$GoRet."{\n";
	print $goblas "\t";
	if ($ret ne 'void') {
		chop($GoRet);
		print $goblas "return ".$GoRet."(";
	}
	print $goblas "C.$func(".processParamToC($func, $paramList).")";
	if ($ret ne 'void') {
		print $goblas ")";
	}
	print $goblas "\n}\n";
}

sub Gofunc {
	my $fnName = shift;
	$fnName =~ s/_sub//;
	my ($pack, $func, $tail) = split '_', $fnName;
	if ($pack eq 'cblas') {
		$pack = "";
	} else {
		$pack = substr $pack, 1;
	}

	return ucfirst $pack . ucfirst $func . ucfirst $tail if $tail;
	return ucfirst $pack . ucfirst $func;
}

sub processParamToGo {
	my $func = shift;
	my $paramList = shift;
	my $complexType = shift;
	my @processed;
	my @params = split ',', $paramList;
	foreach my $param (@params) {
		my @parts = split /[ *]/, $param;
		my $var = lcfirst $parts[scalar @parts - 1];
		$param =~ m/^(?:const )?int/ && do {
			push @processed, $var." int"; next;
		};
		$param =~ m/^(?:const )?void/ && do {
			my $type;
			if ($var eq "alpha" || $var eq "beta") {
				$type = " ";
			} else {
				$type = " []";
			}
			if ($complexType eq 'c') {
				push @processed, $var.$type."complex64"; next;
			} elsif ($complexType eq 'z') {
				push @processed, $var.$type."complex128"; next;
			} else {
				die "unexpected complex type for '$func' - '$complexType'";
			}
		};
		$param =~ m/^(?:const )?char \*/ && do {
			push @processed, $var." *byte"; next;
		};
		$param =~ m/^(?:const )?float \*/ && do {
			push @processed, $var." []float32"; next;
		};
		$param =~ m/^(?:const )?double \*/ && do {
			push @processed, $var." []float64"; next;
		};
		$param =~ m/^(?:const )?float/ && do {
			push @processed, $var." float32"; next;
		};
		$param =~ m/^(?:const )?double/ && do {
			push @processed, $var." float64"; next;
		};
		$param =~ m/^const enum/ && do {
			$var eq "order" && do {
				$var = "o";
				push @processed, $var." blas.Order"; next;
			};
			$var =~ /trans/ && do {
				$var =~ s/trans([AB]?)/t$1/;
				push @processed, $var." blas.Transpose"; next;
			};
			$var eq "uplo" && do {
				$var = "ul";
				push @processed, $var." blas.Uplo"; next;
			};
			$var eq "diag" && do {
				$var = "d";
				push @processed, $var." blas.Diag"; next;
			};
			$var eq "side" && do {
				$var = "s";
				push @processed, $var." blas.Side"; next;
			};
		};
	}
	die "missed Go parameters from '$func', '$paramList'" if scalar @processed != scalar @params;
	return join ", ", @processed;
}

sub processParamToC {
	my $func = shift;
	my $paramList = shift;
	my @processed;
	my @params = split ',', $paramList;
	foreach my $param (@params) {
		my @parts = split /[ *]/, $param;
		my $var = lcfirst $parts[scalar @parts - 1];
		$param =~ m/^(?:const )?int \*[a-zA-Z]/ && do {
			push @processed, "(*C.int)(&".$var.")"; next;
		};
		$param =~ m/^(?:const )?void \*[a-zA-Z]/ && do {
			my $type;
			if ($var eq "alpha" || $var eq "beta") {
				$type = "";
			} else {
				$type = "[0]";
			}
			push @processed, "unsafe.Pointer(&".$var.$type.")"; next;
		};
		$param =~ m/^(?:const )?char \*[a-zA-Z]/ && do {
			push @processed, "(*C.char)(&".$var.")"; next;
		};
		$param =~ m/^(?:const )?float \*[a-zA-Z]/ && do {
			push @processed, "(*C.float)(&".$var."[0])"; next;
		};
		$param =~ m/^(?:const )?double \*[a-zA-Z]/ && do {
			push @processed, "(*C.double)(&".$var."[0])"; next;
		};
		$param =~ m/^(?:const )?int [a-zA-Z]/ && do {
			push @processed, "C.int(".$var.")"; next;
		};
		$param =~ m/^(?:const )float [a-zA-Z]/ && do {
			push @processed, "C.float(".$var.")"; next;
		};
		$param =~ m/^(?:const )double [a-zA-Z]/ && do {
			push @processed, "C.double(".$var.")"; next;
		};
		$param =~ m/^const enum [a-zA-Z]/ && do {
			$var eq "order" && do {
				$var = "o";
				push @processed, "C.enum_$parts[scalar @parts - 2](".$var.")"; next;
			};
			$var =~ /trans/ && do {
				$var =~ s/trans([AB]?)/t$1/;
				push @processed, "C.enum_$parts[scalar @parts - 2](".$var.")"; next;
			};
			$var eq "uplo" && do {
				$var = "ul";
				push @processed, "C.enum_$parts[scalar @parts - 2](".$var.")"; next;
			};
			$var eq "diag" && do {
				$var = "d";
				push @processed, "C.enum_$parts[scalar @parts - 2](".$var.")"; next;
			};
			$var eq "side" && do {
				$var = "s";
				push @processed, "C.enum_$parts[scalar @parts - 2](".$var.")"; next;
			};
		};
	}
	die "missed C parameters from '$func', '$paramList'" if scalar @processed != scalar @params;
	return join ", ", @processed;
}
