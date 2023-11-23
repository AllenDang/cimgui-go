#pragma once

typedef void (*VoidCallback)();

#ifdef __WIN32
typedef unsigned long long xulong;
typedef long long xlong;
#else
typedef unsigned long xulong;
typedef long xlong;
#endif
