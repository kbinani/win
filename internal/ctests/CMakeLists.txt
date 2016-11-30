project(ctests)
cmake_minimum_required(VERSION 2.8)

set(CMAKE_MODULE_PATH ${CMAKE_CURRENT_SOURCE_DIR}/cmakemodules/add_msvc_precompiled_header
                      ${CMAKE_CURRENT_SOURCE_DIR}/cmakemodules/setup_source_groups)
include(setup_source_groups)
include(add_msvc_precompiled_header)

if (CMAKE_CL_64)
  file(GLOB_RECURSE ctests_src "amd64/*.cpp")
else()
  file(GLOB_RECURSE ctests_src "386/*.cpp")
endif()

add_msvc_precompiled_header("stdafx.h" "stdafx.cpp" ctests_src)
set(CMAKE_CXX_FLAGS "${CMAKE_C_FLAGS} /MP /EHsc")

add_executable(ctests ${ctests_src} googletest/googletest/src/gtest_main.cc googletest/googletest/src/gtest-all.cc)
include_directories(. googletest/googletest/include googletest/googletest)
setup_source_groups("${ctests_src}" "${CMAKE_CURRENT_SOURCE_DIR}")
