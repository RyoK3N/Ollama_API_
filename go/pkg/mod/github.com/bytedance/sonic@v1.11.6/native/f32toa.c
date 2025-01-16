/* Copyright 2020 Alexander Bolz
 * 
 * Boost Software License - Version 1.0 - August 17th, 2003
 * 
 * Permission is hereby granted, free of charge, to any person or organization
 * obtaining a copy of the software and accompanying documentation covered by
 * this license (the "Software") to use, reproduce, display, distribute,
 * execute, and transmit the Software, and to prepare derivative works of the
 * Software, and to permit third-parties to whom the Software is furnished to
 * do so, all subject to the following:
 * 
 * The copyright notices in the Software and this entire statement, including
 * the above license grant, this restriction and the following disclaimer,
 * must be included in all copies of the Software, in whole or in part, and
 * all derivative works of the Software, unless such copies or derivative
 * works are solely in the form of machine-executable object code generated by
 * a source language processor.
 *
 * Unless required by applicable law or agreed to in writing, this software
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.
 * 
 * This file may have been modified by ByteDance authors. All ByteDance
 * Modifications are Copyright 2022 ByteDance Authors.
 */

#include "native.h"
#include "tab.h"
#include "test/xassert.h"

#define F32_BITS         32
#define F32_EXP_BITS     8
#define F32_SIG_BITS     23
#define F32_EXP_MASK     0x7F800000u // middle 8 bits
#define F32_SIG_MASK     0x007FFFFFu // lower 23 bits
#define F32_EXP_BIAS     127
#define F32_INF_NAN_EXP  0xFF
#define F32_HIDDEN_BIT   0x00800000u

typedef struct {
    uint32_t sig;
    int32_t exp;
} f32_dec;

typedef __uint128_t uint128_t;

static always_inline unsigned ctz10_u32(const uint32_t v) {
    xassert(0 <= v && v < 1000000000u);
    if (v >= 100000) {
        if (v <   1000000) return 6;
        if (v <  10000000) return 7;
        if (v < 100000000) return 8;
        else               return 9;
    } else {
        if (v <        10) return 1;
        if (v <       100) return 2;
        if (v <      1000) return 3;
        if (v <     10000) return 4;
        else               return 5;
    }
}

static always_inline char* format_significand_f32(uint32_t sig, char *out, int cnt) {
    char *r = out + cnt;
    int ctz = 0;

    /* at most 9 digits here */
    if  (sig >= 10000) {
        uint32_t c = sig - 10000 * (sig / 10000);
        sig /= 10000;
        if (c != 0) {
            uint32_t c0 = (c % 100) << 1;
            uint32_t c1 = (c / 100) << 1;
            copy_two_digs(r - 2, Digits + c0);
            copy_two_digs(r - 4, Digits + c1);
        } else {
            ctz = 4;
        }
        r -= 4;
    }

    while (sig >= 100) {
        uint32_t c = (sig % 100) << 1;
        sig /= 100;
        copy_two_digs(r - 2, Digits + c);
        r -= 2;
    }

    if (sig >= 10) {
        uint32_t c = sig << 1;
        copy_two_digs(out, Digits + c);
    } else {
        *out = (char) ('0' + sig);
    }

    return out + cnt - ctz;
}

static always_inline char* format_integer_u32(uint32_t sig, char *out, unsigned cnt) {
    char *r = out + cnt;

    /* at most 9 digits here */
    if  (sig >= 10000) {
        uint32_t c = sig - 10000 * (sig / 10000);
        sig /= 10000;
        uint32_t c0 = (c % 100) << 1;
        uint32_t c1 = (c / 100) << 1;
        copy_two_digs(r - 2, Digits + c0);
        copy_two_digs(r - 4, Digits + c1);
        r -= 4;
    }

    while (sig >= 100) {
        uint32_t c = (sig % 100) << 1;
        sig /= 100;
        copy_two_digs(r - 2, Digits + c);
        r -= 2;
    }

    if (sig >= 10) {
        uint32_t c = sig << 1;
        copy_two_digs(out, Digits + c);
    } else {
        *out = (char) ('0' + sig);
    }

    return out + cnt;
}

static always_inline char* format_exponent_f32(f32_dec v, char *out, int cnt) {
    char* p = out + 1;
    char* end = format_significand_f32(v.sig, p, cnt);
    while (*(end - 1) == '0') end--;

    /* Print decimal point if needed */
    *out = *p;
    if (end - p > 1) {
        *p = '.';
    } else {
        end--;
    }

    /* Print the exponent */
    *end++ = 'e';
    int32_t exp = v.exp + (int32_t) cnt - 1;
    if (exp < 0) {
        *end++ = '-';
        exp = -exp;
    } else {
        *end++ = '+';
    }

    if (exp >= 100) {
        int32_t c = exp % 10;
        copy_two_digs(end, Digits + 2 * (exp / 10));
        end[2] = (char) ('0' + c);
        end += 3;
    } else if (exp >= 10) {
        copy_two_digs(end, Digits + 2 * exp);
        end += 2;
    } else {
        *end++ = (char) ('0' + exp);
    }
    return end;
}

static always_inline char* format_decimal_f32(f32_dec v, char* out, int cnt) {
    char* p = out;
    char* end;
    int point = cnt + v.exp;

    /* print leading zeros if fp < 1 */
    if (point <= 0) {
        *p++ = '0', *p++ = '.';
        for (int i = 0; i < -point; i++) {
            *p++ = '0';
        }
    }

    /* add the remaining digits */
    end = format_significand_f32(v.sig, p, cnt);
    while (*(end - 1) == '0') end--;
    if (point <= 0) {
        return end;
    }

    /* insert point or add trailing zeros */
    int digs = end - p, frac = digs - point;
    if (digs > point) {
        for (int i = 0; i < frac; i++) {
            *(end - i) = *(end - i - 1);
        }
        p[point] = '.';
        end++;
    } else {
        for (int i = 0; i < point - digs; i++) {
            *end++ = '0';
        }
    }
    return end;
}

static always_inline char* write_dec_f32(f32_dec dec, char* p) {
    int cnt = ctz10_u32(dec.sig);
    int dot = cnt + dec.exp;
    int sci_exp = dot - 1;
    bool exp_fmt = sci_exp < -6 || sci_exp > 20;
    bool has_dot = dot < cnt;

    if (exp_fmt) {
        return format_exponent_f32(dec, p, cnt);
    }
    if (has_dot) {
        return format_decimal_f32(dec, p, cnt);
    }

    char* end = p + dot;
    p = format_integer_u32(dec.sig, p, cnt);
    while (p < end) *p++ = '0';
    return end;
}

static always_inline uint32_t f32toraw(float fp) {
    union {
        uint32_t u32;
        float   f32;
    } uval;
    uval.f32 = fp;
    return uval.u32;
}

static always_inline uint64_t pow10_ceil_sig_f32(int32_t k)
{
    // There are unique beta and r such that 10^k = beta 2^r and
    // 2^63 <= beta < 2^64, namely r = floor(log_2 10^k) - 63 and
    // beta = 2^-r 10^k.
    // Let g = ceil(beta), so (g-1) 2^r < 10^k <= g 2^r, with the latter
    // value being a pretty good overestimate for 10^k.

    // NB: Since for all the required exponents k, we have g < 2^64,
    //     all constants can be stored in 128-bit integers.
    // reference from: 
    //  https://github.com/abolz/Drachennest/blob/master/src/schubfach_32.cc#L144

#define KMAX 45
#define KMIN -31
    static const uint64_t g[KMAX - KMIN + 1] = {
        0x81CEB32C4B43FCF5, // -31
        0xA2425FF75E14FC32, // -30
        0xCAD2F7F5359A3B3F, // -29
        0xFD87B5F28300CA0E, // -28
        0x9E74D1B791E07E49, // -27
        0xC612062576589DDB, // -26
        0xF79687AED3EEC552, // -25
        0x9ABE14CD44753B53, // -24
        0xC16D9A0095928A28, // -23
        0xF1C90080BAF72CB2, // -22
        0x971DA05074DA7BEF, // -21
        0xBCE5086492111AEB, // -20
        0xEC1E4A7DB69561A6, // -19
        0x9392EE8E921D5D08, // -18
        0xB877AA3236A4B44A, // -17
        0xE69594BEC44DE15C, // -16
        0x901D7CF73AB0ACDA, // -15
        0xB424DC35095CD810, // -14
        0xE12E13424BB40E14, // -13
        0x8CBCCC096F5088CC, // -12
        0xAFEBFF0BCB24AAFF, // -11
        0xDBE6FECEBDEDD5BF, // -10
        0x89705F4136B4A598, //  -9
        0xABCC77118461CEFD, //  -8
        0xD6BF94D5E57A42BD, //  -7
        0x8637BD05AF6C69B6, //  -6
        0xA7C5AC471B478424, //  -5
        0xD1B71758E219652C, //  -4
        0x83126E978D4FDF3C, //  -3
        0xA3D70A3D70A3D70B, //  -2
        0xCCCCCCCCCCCCCCCD, //  -1
        0x8000000000000000, //   0
        0xA000000000000000, //   1
        0xC800000000000000, //   2
        0xFA00000000000000, //   3
        0x9C40000000000000, //   4
        0xC350000000000000, //   5
        0xF424000000000000, //   6
        0x9896800000000000, //   7
        0xBEBC200000000000, //   8
        0xEE6B280000000000, //   9
        0x9502F90000000000, //  10
        0xBA43B74000000000, //  11
        0xE8D4A51000000000, //  12
        0x9184E72A00000000, //  13
        0xB5E620F480000000, //  14
        0xE35FA931A0000000, //  15
        0x8E1BC9BF04000000, //  16
        0xB1A2BC2EC5000000, //  17
        0xDE0B6B3A76400000, //  18
        0x8AC7230489E80000, //  19
        0xAD78EBC5AC620000, //  20
        0xD8D726B7177A8000, //  21
        0x878678326EAC9000, //  22
        0xA968163F0A57B400, //  23
        0xD3C21BCECCEDA100, //  24
        0x84595161401484A0, //  25
        0xA56FA5B99019A5C8, //  26
        0xCECB8F27F4200F3A, //  27
        0x813F3978F8940985, //  28
        0xA18F07D736B90BE6, //  29
        0xC9F2C9CD04674EDF, //  30
        0xFC6F7C4045812297, //  31
        0x9DC5ADA82B70B59E, //  32
        0xC5371912364CE306, //  33
        0xF684DF56C3E01BC7, //  34
        0x9A130B963A6C115D, //  35
        0xC097CE7BC90715B4, //  36
        0xF0BDC21ABB48DB21, //  37
        0x96769950B50D88F5, //  38
        0xBC143FA4E250EB32, //  39
        0xEB194F8E1AE525FE, //  40
        0x92EFD1B8D0CF37BF, //  41
        0xB7ABC627050305AE, //  42
        0xE596B7B0C643C71A, //  43
        0x8F7E32CE7BEA5C70, //  44
        0xB35DBF821AE4F38C, //  45
    };

    xassert(k >= KMIN && k <= KMAX);
    return g[k - KMIN];
#undef KMIN
#undef KMAX
}

static always_inline uint32_t round_odd_f32(uint64_t g, uint32_t cp) {
    const uint128_t p = ((uint128_t)g) * cp;
    const uint32_t y1 = (uint64_t)(p >> 64);
    const uint32_t y0 = ((uint64_t)(p)) >> 32;
    return y1 | (y0 > 1);
}

/**
 Rendering float point number into decimal.
 The function used Schubfach algorithm, reference:
 The Schubfach way to render doubles, Raffaello Giulietti, 2022-03-20.
 https://drive.google.com/file/d/1gp5xv4CAa78SVgCeWfGqqI4FfYYYuNFb
 https://mail.openjdk.java.net/pipermail/core-libs-dev/2021-November/083536.html
 https://github.com/openjdk/jdk/pull/3402 (Java implementation)
 https://github.com/abolz/Drachennest (C++ implementation)
 */
static always_inline f32_dec f32todec(uint32_t rsig, int32_t rexp, uint32_t c, int32_t q) {
    uint32_t cbl, cb, cbr, vbl, vb, vbr, lower, upper, s;
    int32_t k, h;
    bool even, irregular, w_inside, u_inside;
    f32_dec dec;

    even = !(c & 1);
    irregular = rsig == 0 && rexp > 1;

    cbl = 4 * c - 2 + irregular;
    cb = 4 * c;
    cbr = 4 * c + 2;

    k = (q * 1262611 - (irregular ? 524031 : 0)) >> 22;
    h = q + ((-k) * 1741647 >> 19) + 1;
    uint64_t pow10 = pow10_ceil_sig_f32(-k);
    vbl = round_odd_f32(pow10, cbl << h);
    vb = round_odd_f32(pow10, cb << h);
    vbr = round_odd_f32(pow10, cbr << h);


    lower = vbl + !even;
    upper = vbr - !even;

    s = vb / 4;
    if (s >= 10) {
        uint64_t sp = s / 10;
        bool up_inside = lower <= (40 * sp);
        bool wp_inside = (40 * sp + 40) <= upper;
        if (up_inside != wp_inside) {
            dec.sig = sp + wp_inside;
            dec.exp = k + 1;
            return dec;
        }
    }

    u_inside = lower <= (4 * s);
    w_inside = (4 * s + 4) <= upper;
    if (u_inside != w_inside) {
        dec.sig = s + w_inside;
        dec.exp = k;
        return dec;
    }

    uint64_t mid = 4 * s + 2;
    bool round_up = vb > mid || (vb == mid && (s & 1) != 0);
    dec.sig = s + round_up;
    dec.exp = k;
    return dec;
}

 int f32toa(char *out, float fp) {
    char* p = out;
    uint32_t raw = f32toraw(fp);
    bool neg;
    uint32_t rsig, c;
    int32_t rexp, q;

    neg = ((raw >> (F32_BITS - 1)) != 0);
    rsig = raw & F32_SIG_MASK;
    rexp = (int32_t)((raw & F32_EXP_MASK) >> F32_SIG_BITS);

    /* check infinity and nan */
    if (unlikely(rexp == F32_INF_NAN_EXP)) {
        return 0;
    }

    /* check negative numbers */
    *p = '-';
    p += neg;

    /* simple case of 0.0 */
    if ((raw << 1) ==  0) {
        *p++ = '0';
        return p - out;
    }

    if (likely(rexp != 0)) {
        /* double is normal */
        c = rsig | F32_HIDDEN_BIT;
        q = rexp - F32_EXP_BIAS - F32_SIG_BITS;

        /* fast path for integer */
        if (q <= 0 && q >= -F32_SIG_BITS && is_div_pow2(c, -q)) {
            uint32_t u = c >> -q;
            p = format_integer_u32(u, p, ctz10_u32(u));
            return p - out;
        }
    } else {
        c = rsig;
        q = 1 - F32_EXP_BIAS - F32_SIG_BITS;
    }

    f32_dec dec = f32todec(rsig, rexp, c, q);
    p = write_dec_f32(dec, p);
    return p - out;
}

#undef F32_BITS       
#undef F32_EXP_BITS   
#undef F32_SIG_BITS   
#undef F32_EXP_MASK   
#undef F32_SIG_MASK   
#undef F32_EXP_BIAS   
#undef F32_INF_NAN_EXP
#undef F32_HIDDEN_BIT 
