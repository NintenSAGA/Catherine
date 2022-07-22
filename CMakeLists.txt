cmake_minimum_required(VERSION 3.17)
project(Push_Blocks C)

set(CMAKE_C_STANDARD 99)

add_executable(Push_Blocks main.c game.c game.h Global.h function.c function.h map_data.h Title.c Title.h)