project(ctests)
cmake_minimum_required(VERSION 2.8)

macro(add_msvc_precompiled_header PrecompiledHeader PrecompiledSource SourcesVar)
  get_filename_component(PrecompiledBasename ${PrecompiledHeader} NAME_WE)
  set(PrecompiledBinary "$(IntDir)/${PrecompiledBasename}.pch")
  set(Sources ${${SourcesVar}})

  set_source_files_properties(${PrecompiledSource}
                              PROPERTIES COMPILE_FLAGS "/Yc\"${PrecompiledHeader}\" /Fp\"${PrecompiledBinary}\""
                                         OBJECT_OUTPUTS "${PrecompiledBinary}")
  set_source_files_properties(${Sources}
                              PROPERTIES COMPILE_FLAGS "/Yu\"${PrecompiledHeader}\" /FI\"${PrecompiledHeader}\" /Fp\"${PrecompiledBinary}\""
                                         OBJECT_DEPENDS "${PrecompiledBinary}")  

  list(APPEND ${SourcesVar} ${PrecompiledSource})
  list(APPEND ${SourcesVar} ${PrecompiledHeader})

  set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} /FI\"${PrecompiledHeader}\"")
endmacro(add_msvc_precompiled_header)

macro(setup_source_groups SourceFileList ProjectRootDir)
  get_filename_component(ProjectRootDir "${ProjectRootDir}" ABSOLUTE)
  string(REPLACE "\\" "/" ProjectRootDir "${ProjectRootDir}")

  foreach(source_file_path ${SourceFileList})
    get_filename_component(source_file_path "${source_file_path}" ABSOLUTE)
    string(REPLACE "\\" "/" source_file_path "${source_file_path}")

    string(REGEX REPLACE "^${ProjectRootDir}/" "" relative_path "${source_file_path}")

    get_filename_component(dirname "${relative_path}" DIRECTORY)

    string(REPLACE "/" "\\" group_name "${dirname}")
    source_group("${group_name}" FILES "${source_file_path}")
  endforeach()
endmacro()

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
