#pragma once

#ifdef __WIN32
typedef unsigned long long xlong;
typedef long long xlong;
#else
typedef unsigned long xlong;
typedef long xxlong;
#endif
