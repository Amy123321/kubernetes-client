/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs defs_openbsd.go

package route

const (
	sysAF_UNSPEC = 0x0
	sysAF_INET   = 0x2
	sysAF_ROUTE  = 0x11
	sysAF_LINK   = 0x12
	sysAF_INET6  = 0x18

	sysSOCK_RAW = 0x3

	sysNET_RT_DUMP    = 0x1
	sysNET_RT_FLAGS   = 0x2
	sysNET_RT_IFLIST  = 0x3
	sysNET_RT_STATS   = 0x4
	sysNET_RT_TABLE   = 0x5
	sysNET_RT_IFNAMES = 0x6
	sysNET_RT_MAXID   = 0x7
)

const (
	sysCTL_MAXNAME = 0xc

	sysCTL_UNSPEC  = 0x0
	sysCTL_KERN    = 0x1
	sysCTL_VM      = 0x2
	sysCTL_FS      = 0x3
	sysCTL_NET     = 0x4
	sysCTL_DEBUG   = 0x5
	sysCTL_HW      = 0x6
	sysCTL_MACHDEP = 0x7
	sysCTL_DDB     = 0x9
	sysCTL_VFS     = 0xa
	sysCTL_MAXID   = 0xb
)

const (
	sysRTM_VERSION = 0x5

	sysRTM_ADD        = 0x1
	sysRTM_DELETE     = 0x2
	sysRTM_CHANGE     = 0x3
	sysRTM_GET        = 0x4
	sysRTM_LOSING     = 0x5
	sysRTM_REDIRECT   = 0x6
	sysRTM_MISS       = 0x7
	sysRTM_LOCK       = 0x8
	sysRTM_RESOLVE    = 0xb
	sysRTM_NEWADDR    = 0xc
	sysRTM_DELADDR    = 0xd
	sysRTM_IFINFO     = 0xe
	sysRTM_IFANNOUNCE = 0xf
	sysRTM_DESYNC     = 0x10
	sysRTM_INVALIDATE = 0x11
	sysRTM_BFD        = 0x12
	sysRTM_PROPOSAL   = 0x13

	sysRTA_DST     = 0x1
	sysRTA_GATEWAY = 0x2
	sysRTA_NETMASK = 0x4
	sysRTA_GENMASK = 0x8
	sysRTA_IFP     = 0x10
	sysRTA_IFA     = 0x20
	sysRTA_AUTHOR  = 0x40
	sysRTA_BRD     = 0x80
	sysRTA_SRC     = 0x100
	sysRTA_SRCMASK = 0x200
	sysRTA_LABEL   = 0x400
	sysRTA_BFD     = 0x800
	sysRTA_DNS     = 0x1000
	sysRTA_STATIC  = 0x2000
	sysRTA_SEARCH  = 0x4000

	sysRTAX_DST     = 0x0
	sysRTAX_GATEWAY = 0x1
	sysRTAX_NETMASK = 0x2
	sysRTAX_GENMASK = 0x3
	sysRTAX_IFP     = 0x4
	sysRTAX_IFA     = 0x5
	sysRTAX_AUTHOR  = 0x6
	sysRTAX_BRD     = 0x7
	sysRTAX_SRC     = 0x8
	sysRTAX_SRCMASK = 0x9
	sysRTAX_LABEL   = 0xa
	sysRTAX_BFD     = 0xb
	sysRTAX_DNS     = 0xc
	sysRTAX_STATIC  = 0xd
	sysRTAX_SEARCH  = 0xe
	sysRTAX_MAX     = 0xf
)

const (
	sizeofRtMsghdr = 0x60

	sizeofSockaddrStorage = 0x100
	sizeofSockaddrInet    = 0x10
	sizeofSockaddrInet6   = 0x1c
)