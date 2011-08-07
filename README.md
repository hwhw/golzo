# A simple Go interface for liblzo2

This is a minimalistic, quick 'n dirty wrapper for liblzo2 - I didn't find another and was in a hurry and am by no means proud of my creation.

Currently, only `lzo_init()` and `lzo1x_999_compress()` are interfaced. But you could easily add other `_compress`/`_uncompress` functions given the example.

The software is free software, licensed under a GPLv2 license. See the *COPYING* file for the terms.
