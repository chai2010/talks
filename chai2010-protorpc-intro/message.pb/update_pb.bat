:: Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

:: gen cxx code
protoc.exe --cxx_out=. arith.proto
protoc.exe --cxx_out=. echo.proto

:: gen go code
protoc.exe --go_out=. arith.proto
protoc.exe --go_out=. echo.proto
