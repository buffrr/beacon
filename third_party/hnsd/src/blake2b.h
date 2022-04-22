/**
 * Parts of this software are based on BLAKE2:
 * https://github.com/BLAKE2/BLAKE2
 *
 * BLAKE2 reference source code package - reference C implementations
 *
 * Copyright 2012, Samuel Neves <sneves@dei.uc.pt>.  You may use this under
 * the terms of the CC0, the OpenSSL Licence, or the Apache Public License
 * 2.0, at your option.  The terms of these licenses can be found at:
 *
 * - CC0 1.0 Universal : http://creativecommons.org/publicdomain/zero/1.0
 * - OpenSSL license   : https://www.openssl.org/source/license.html
 * - Apache 2.0        : http://www.apache.org/licenses/LICENSE-2.0
 *
 * More information about the BLAKE2 hash function can be found at
 * https://blake2.net.
 */

#ifndef _HSK_BLAKE2_H
#define _HSK_BLAKE2_H

#include <stddef.h>
#include <stdint.h>

#if defined(_MSC_VER)
#define HSK_BLAKE2_PACKED(x) __pragma(pack(push, 1)) x __pragma(pack(pop))
#else
#define HSK_BLAKE2_PACKED(x) x __attribute__((packed))
#endif

enum hsk_blake2b_constant {
  HSK_BLAKE2B_BLOCKBYTES = 128,
  HSK_BLAKE2B_OUTBYTES = 64,
  HSK_BLAKE2B_KEYBYTES = 64,
  HSK_BLAKE2B_SALTBYTES = 16,
  HSK_BLAKE2B_PERSONALBYTES = 16
};

typedef struct hsk_blake2b_ctx__ {
  uint64_t h[8];
  uint64_t t[2];
  uint64_t f[2];
  uint8_t buf[HSK_BLAKE2B_BLOCKBYTES];
  size_t buflen;
  size_t outlen;
  uint8_t last_node;
} hsk_blake2b_ctx;

HSK_BLAKE2_PACKED(struct hsk_blake2b_param__ {
  uint8_t digest_length;
  uint8_t key_length;
  uint8_t fanout;
  uint8_t depth;
  uint32_t leaf_length;
  uint32_t node_offset;
  uint32_t xof_length;
  uint8_t node_depth;
  uint8_t inner_length;
  uint8_t reserved[14];
  uint8_t salt[HSK_BLAKE2B_SALTBYTES];
  uint8_t personal[HSK_BLAKE2B_PERSONALBYTES];
});

typedef struct hsk_blake2b_param__ hsk_blake2b_param;

enum {
  HSK_BLAKE2_DUMMY_1 = 1 / (sizeof(hsk_blake2b_param) == HSK_BLAKE2B_OUTBYTES)
};

int hsk_blake2b_init(hsk_blake2b_ctx *ctx, size_t outlen);

int hsk_blake2b_init_key(
  hsk_blake2b_ctx *ctx,
  size_t outlen,
  const void *key,
  size_t keylen
);

int hsk_blake2b_init_param(hsk_blake2b_ctx *ctx, const hsk_blake2b_param *P);

int hsk_blake2b_update(hsk_blake2b_ctx *ctx, const void *in, size_t inlen);

int hsk_blake2b_final(hsk_blake2b_ctx *ctx, void *out, size_t outlen);

int hsk_blake2b(
  void *out,
  size_t outlen,
  const void *in,
  size_t inlen,
  const void *key,
  size_t keylen
);

#endif
