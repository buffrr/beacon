#ifndef _HSK_SIPHASH_H
#define _HSK_SIPHASH_H

#include <stdint.h>
#include <stdlib.h>

uint64_t
hsk_siphash(const uint8_t *data, size_t len, const uint8_t *key);

uint32_t
hsk_siphash32(uint32_t num, const uint8_t *key);

uint64_t
hsk_siphash64(uint64_t num, const uint8_t *key);

uint32_t
hsk_siphash32k256(uint32_t num, const uint8_t *key);

uint64_t
hsk_siphash64k256(uint64_t num, const uint8_t *key);

#endif
