#
# Copyright (C) 2018 by George Cave - gcave@stablecoder.ca
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may not
# use this file except in compliance with the License. You may obtain a copy of
# the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations under
# the License.

# USAGE: To enable any code coverage instrumentation/targets, the single CMake
# option of `CODE_COVERAGE` needs to be set to 'ON', either by GUI, ccmake, or
# on the command line.
#
# From this point, there are two primary methods for adding instrumentation to
# targets: 1 - A blanket instrumentation by calling `add_code_coverage()`, where
# all targets in that directory and all subdirectories are automatically
# instrumented. 2 - Per-target instrumentation by calling
# `target_code_coverage(<TARGET_NAME>)`, where the target is given and thus only
# that target is instrumented. This applies to both libraries and executables.
#
# To add coverage targets, such as calling `make ccov` to generate the actual
# coverage information for perusal or consumption, call
# `target_code_coverage(<TARGET_NAME>)` on an *executable* target.
#
# Example 1: All targets instrumented
#
# In this case, the coverage information reported will will be that of the
# `theLib` library target and `theExe` executable.
#
# 1a: Via global command
#
# ~~~
# add_code_coverage() # Adds instrumentation to all targets
#
# add_library(theLib lib.cpp)
#
# add_executable(theExe main.cpp)
# target_link_libraries(theExe PRIVATE theLib)
# target_code_coverage(theExe) # As an executable target, adds the 'ccov-theExe' target (instrumentation already added via global anyways) for generating code coverage reports.
# ~~~
#
# 1b: Via target commands
#
# ~~~
# add_library(theLib lib.cpp)
# target_code_coverage(theLib) # As a library target, adds coverage instrumentation but no targets.
#
# add_executable(theExe main.cpp)
# target_link_libraries(theExe PRIVATE theLib)
# target_code_coverage(theExe) # As an executable target, adds the 'ccov-theExe' target and instrumentation for generating code coverage reports.
# ~~~
#
# Example 2: Target instrumented, but with regex pattern of files to be excluded
# from report
#
# ~~~
# add_executable(theExe main.cpp non_covered.cpp)
# target_code_coverage(theExe EXCLUDE non_covered.cpp test/*) # As an executable target, the reports will exclude the non-covered.cpp file, and any files in a test/ folder.
# ~~~
#
# Example 3: Target added to the 'ccov' and 'ccov-all' targets
#
# ~~~
# add_code_coverage_all_targets(EXCLUDE test/*) # Adds the 'ccov-all' target set and sets it to exclude all files in test/ folders.
#
# add_executable(theExe main.cpp non_covered.cpp)
# target_code_coverage(theExe AUTO ALL EXCLUDE non_covered.cpp test/*) # As an executable target, adds to the 'ccov' and ccov-all' targets, and the reports will exclude the non-covered.cpp file, and any files in a test/ folder.
# ~~~

# Options
option(
  CODE_COVERAGE
  "Builds targets with code coverage instrumentation. (Requires GCC or Clang)"
  OFF)

# Programs
find_program(LLVM_COV_PATH llvm-cov)
find_program(LLVM_PROFDATA_PATH llvm-profdata)
find_program(LCOV_PATH lcov)
find_program(GENHTML_PATH genhtml)

# Variables
set(CMAKE_COVERAGE_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/ccov)

# Common initialization/checks
if(CODE_COVERAGE AND NOT CODE_COVERAGE_ADDED)
  set(CODE_COVERAGE_ADDED ON)

  # Common Targets
  add_custom_target(ccov-preprocessing
                    COMMAND ${CMAKE_COMMAND}
                            -E
                            make_directory ${CMAKE_COVERAGE_OUTPUT_DIRECTORY}
                    DEPENDS ccov-clean)

  if("${CMAKE_C_COMPILER_ID}" MATCHES "(Apple)?[Cc]lang"
     OR "${CMAKE_CXX_COMPILER_ID}" MATCHES "(Apple)?[Cc]lang")
    # Messages
    message(STATUS "Building with llvm Code Coverage Tools")

    if(NOT LLVM_COV_PATH)
      message(FATAL_ERROR "llvm-cov not found! Aborting.")
    else()
      # Version number checking for 'EXCLUDE' compatability
      execute_process(COMMAND ${LLVM_COV_PATH} --version
                      OUTPUT_VARIABLE LLVM_COV_VERSION_CALL_OUTPUT)
      string(REGEX MATCH
                   "[0-9]+\\.[0-9]+\\.[0-9]+"
                   LLVM_COV_VERSION
                   ${LLVM_COV_VERSION_CALL_OUTPUT})

      if(LLVM_COV_VERSION VERSION_LESS "7.0.0")
        message(
          WARNING
            "target_code_coverage()/add_code_coverage_all_targets() 'EXCLUDE' option only available on llvm-cov >= 7.0.0"
          )
      endif()
    endif()

    # Targets
    add_custom_target(ccov-clean
                      COMMAND rm
                              -f
                              ${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/binaries.list
                      COMMAND rm
                              -f
                              ${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/profraw.list)

  elseif(CMAKE_COMPILER_IS_GNUCXX)
    # Messages
    message(STATUS "Building with lcov Code Coverage Tools")

    if(CMAKE_BUILD_TYPE)
      string(TOUPPER ${CMAKE_BUILD_TYPE} upper_build_type)
      if(NOT ${upper_build_type} STREQUAL "DEBUG")
        message(
          WARNING
            "Code coverage results with an optimized (non-Debug) build may be misleading"
          )
      endif()
    else()
      message(
        WARNING
          "Code coverage results with an optimized (non-Debug) build may be misleading"
        )
    endif()
    if(NOT LCOV_PATH)
      message(FATAL_ERROR "lcov not found! Aborting...")
    endif()
    if(NOT GENHTML_PATH)
      message(FATAL_ERROR "genhtml not found! Aborting...")
    endif()

    # Targets
    add_custom_target(ccov-clean
                      COMMAND ${LCOV_PATH}
                              --directory ${CMAKE_BINARY_DIR}
                              --zerocounters)

  else()
    message(FATAL_ERROR "Code coverage requires Clang or GCC. Aborting.")
  endif()
endif()

# Adds code coverage instrumentation to a library, or instrumentation/targets
# for an executable target.
# ~~~
# EXECUTABLE ADDED TARGETS:
# GCOV/LCOV:
# ccov : Generates HTML code coverage report for every target added with 'AUTO' parameter.
# ccov-${TARGET_NAME} : Generates HTML code coverage report for the associated named target.
# ccov-all : Generates HTML code coverage report, merging every target added with 'ALL' parameter into a single detailed report.
#
# LLVM-COV:
# ccov : Generates HTML code coverage report for every target added with 'AUTO' parameter.
# ccov-report : Generates HTML code coverage report for every target added with 'AUTO' parameter.
# ccov-${TARGET_NAME} : Generates HTML code coverage report.
# ccov-report-${TARGET_NAME} : Prints to command line summary per-file coverage information.
# ccov-show-${TARGET_NAME} : Prints to command line detailed per-line coverage information.
# ccov-all : Generates HTML code coverage report, merging every target added with 'ALL' parameter into a single detailed report.
# ccov-all-report : Prints summary per-file coverage information for every target added with ALL' parameter to the command line.
#
# Required:
# TARGET_NAME - Name of the target to generate code coverage for.
# Optional:
# AUTO - Adds the target to the 'ccov' target so that it can be run in a batch with others easily. Effective on executable targets.
# ALL - Adds the target to the 'ccov-all' and 'ccov-all-report' targets, which merge several executable targets coverage data to a single report. Effective on executable targets.
# EXCLUDE <REGEX_PATTERNS> - Excludes files of the patterns provided from coverage. **These do not copy to the 'all' targets.**
# ~~~
function(target_code_coverage TARGET_NAME)
  # Argument parsing
  set(options AUTO ALL)
  set(multi_value_keywords EXCLUDE)
  cmake_parse_arguments(target_code_coverage
                        "${options}"
                        ""
                        "${multi_value_keywords}"
                        ${ARGN})

  if(CODE_COVERAGE)
    # Instrumentation
    if("${CMAKE_C_COMPILER_ID}" MATCHES "(Apple)?[Cc]lang"
       OR "${CMAKE_CXX_COMPILER_ID}" MATCHES "(Apple)?[Cc]lang")
      target_compile_options(
        ${TARGET_NAME}
        PRIVATE -fprofile-instr-generate -fcoverage-mapping)
      set_property(TARGET ${TARGET_NAME}
                   APPEND_STRING
                   PROPERTY LINK_FLAGS "-fprofile-instr-generate ")
      set_property(TARGET ${TARGET_NAME}
                   APPEND_STRING
                   PROPERTY LINK_FLAGS "-fcoverage-mapping ")
    elseif(CMAKE_COMPILER_IS_GNUCXX)
      target_compile_options(${TARGET_NAME}
                             PRIVATE -fprofile-arcs -ftest-coverage)
      target_link_libraries(${TARGET_NAME} PRIVATE gcov)
    endif()

    # Targets
    get_target_property(target_type ${TARGET_NAME} TYPE)
    if(target_type STREQUAL "EXECUTABLE")
      if("${CMAKE_C_COMPILER_ID}" MATCHES "(Apple)?[Cc]lang"
         OR "${CMAKE_CXX_COMPILER_ID}" MATCHES "(Apple)?[Cc]lang")
        add_custom_target(
          ccov-run-${TARGET_NAME}
          COMMAND LLVM_PROFILE_FILE=${TARGET_NAME}.profraw
                  $<TARGET_FILE:${TARGET_NAME}>
          COMMAND echo
                  "-object=$<TARGET_FILE:${TARGET_NAME}>"
                  >>
                  ${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/binaries.list
          COMMAND echo
                  "${CMAKE_CURRENT_BINARY_DIR}/${TARGET_NAME}.profraw "
                  >>
                  ${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/profraw.list
          DEPENDS ccov-preprocessing ${TARGET_NAME})

        add_custom_target(ccov-processing-${TARGET_NAME}
                          COMMAND ${LLVM_PROFDATA_PATH} merge
                                  -sparse
                                  ${TARGET_NAME}.profraw
                                  -o
                                  ${TARGET_NAME}.profdata
                          DEPENDS ccov-run-${TARGET_NAME})

        if(LLVM_COV_VERSION VERSION_GREATER_EQUAL "7.0.0")
          foreach(EXCLUDE_ITEM ${target_code_coverage_EXCLUDE})
            set(EXCLUDE_REGEX ${EXCLUDE_REGEX}
                              -ignore-filename-regex='${EXCLUDE_ITEM}')
          endforeach()
        endif()

        add_custom_target(ccov-show-${TARGET_NAME}
                          COMMAND ${LLVM_COV_PATH} show $<TARGET_FILE:${TARGET_NAME}>
                                  -instr-profile=${TARGET_NAME}.profdata
                                  -show-line-counts-or-regions
                                  ${EXCLUDE_REGEX}
                          DEPENDS ccov-processing-${TARGET_NAME})

        add_custom_target(ccov-report-${TARGET_NAME}
                          COMMAND ${LLVM_COV_PATH} report $<TARGET_FILE:${TARGET_NAME}>
                                  -instr-profile=${TARGET_NAME}.profdata
                                  ${EXCLUDE_REGEX}
                          DEPENDS ccov-processing-${TARGET_NAME})

        add_custom_target(
          ccov-${TARGET_NAME}
          COMMAND ${LLVM_COV_PATH} show $<TARGET_FILE:${TARGET_NAME}>
                  -instr-profile=${TARGET_NAME}.profdata
                  -show-line-counts-or-regions
                  -output-dir=${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/${TARGET_NAME}
                  -format="html"
                  ${EXCLUDE_REGEX}
          DEPENDS ccov-processing-${TARGET_NAME})

      elseif(CMAKE_COMPILER_IS_GNUCXX)
        set(COVERAGE_INFO
            "${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/${TARGET_NAME}.info")

        add_custom_target(ccov-run-${TARGET_NAME}
                          COMMAND $<TARGET_FILE:${TARGET_NAME}>
                          DEPENDS ccov-preprocessing ${TARGET_NAME})

        foreach(EXCLUDE_ITEM ${target_code_coverage_EXCLUDE})
          set(EXCLUDE_REGEX
              ${EXCLUDE_REGEX}
              --remove
              ${COVERAGE_INFO}
              '${EXCLUDE_ITEM}')
        endforeach()

        if(EXCLUDE_REGEX)
          set(EXCLUDE_COMMAND
              ${LCOV_PATH}
              ${EXCLUDE_REGEX}
              --output-file
              ${COVERAGE_INFO})
        else()
          set(EXCLUDE_COMMAND ;)
        endif()

        add_custom_target(
          ccov-${TARGET_NAME}
          COMMAND ${CMAKE_COMMAND} -E remove ${COVERAGE_INFO}
          COMMAND ${LCOV_PATH} --directory ${CMAKE_BINARY_DIR} --zerocounters
          COMMAND $<TARGET_FILE:${TARGET_NAME}>
          COMMAND ${LCOV_PATH}
                  --directory ${CMAKE_BINARY_DIR}
                  --base-directory ${CMAKE_SOURCE_DIR}
                  --capture
                  --no-external
                  --output-file ${COVERAGE_INFO}
          COMMAND ${EXCLUDE_COMMAND}
          COMMAND ${GENHTML_PATH}
                  -o
                  ${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/${TARGET_NAME}
                  ${COVERAGE_INFO}
          DEPENDS ccov-preprocessing ${TARGET_NAME})
      endif()

      add_custom_command(
        TARGET ccov-${TARGET_NAME} POST_BUILD
        COMMAND ;
        COMMENT
          "Open ${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/${TARGET_NAME}/index.html in your browser to view the coverage report."
        )

      # AUTO
      if(target_code_coverage_AUTO)
        if(NOT TARGET ccov)
          add_custom_target(ccov)
        endif()
        add_dependencies(ccov ccov-${TARGET_NAME})

        if(NOT CMAKE_COMPILER_IS_GNUCXX)
          if(NOT TARGET ccov-report)
            add_custom_target(ccov-report)
          endif()
          add_dependencies(ccov-report ccov-report-${TARGET_NAME})
        endif()
      endif()

      # ALL
      if(target_code_coverage_ALL)
        if(NOT TARGET ccov-all-processing)
          message(
            FATAL_ERROR
              "Calling target_code_coverage with 'ALL' must be after a call to 'add_code_coverage_all_targets'."
            )
        endif()

        add_dependencies(ccov-all-processing ccov-run-${TARGET_NAME})
      endif()
    endif()
  endif()
endfunction()

# Adds code coverage instrumentation to all targets in the current directory and
# any subdirectories. To add coverage instrumentation to only specific targets,
# use `target_code_coverage`.
function(add_code_coverage)
  if("${CMAKE_C_COMPILER_ID}" MATCHES "(Apple)?[Cc]lang"
     OR "${CMAKE_CXX_COMPILER_ID}" MATCHES "(Apple)?[Cc]lang")
    add_compile_options(-fprofile-instr-generate -fcoverage-mapping)
    set(CMAKE_EXE_LINKER_FLAGS
        "${CMAKE_EXE_LINKER_FLAGS} -fprofile-instr-generate -fcoverage-mapping")
  elseif(CMAKE_COMPILER_IS_GNUCXX)
    add_compile_options(-fprofile-arcs -ftest-coverage)
    link_libraries(gcov)
  endif()
endfunction()

# Adds the 'ccov-all' type targets that calls all targets added via
# `target_code_coverage` with the `ALL` parameter, but merges all the coverage
# data from them into a single large report  instead of the numerous smaller
# reports.
# ~~~
# Optional:
# EXCLUDE <REGEX_PATTERNS> - Excludes files of the regex patterns provided from coverage.
# ~~~
function(add_code_coverage_all_targets)
  # Argument parsing
  set(multi_value_keywords EXCLUDE)
  cmake_parse_arguments(add_code_coverage_all_targets
                        ""
                        ""
                        "${multi_value_keywords}"
                        ${ARGN})

  if(CODE_COVERAGE)
    if("${CMAKE_C_COMPILER_ID}" MATCHES "(Apple)?[Cc]lang"
       OR "${CMAKE_CXX_COMPILER_ID}" MATCHES "(Apple)?[Cc]lang")
      # Targets
      add_custom_target(
        ccov-all-processing
        COMMAND ${LLVM_PROFDATA_PATH} merge
                -o
                ${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/all-merged.profdata
                -sparse
                `cat ${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/profraw.list`)

      if(LLVM_COV_VERSION VERSION_GREATER_EQUAL "7.0.0")
        foreach(EXCLUDE_ITEM ${add_code_coverage_all_targets_EXCLUDE})
          set(EXCLUDE_REGEX ${EXCLUDE_REGEX}
                            -ignore-filename-regex=${EXCLUDE_ITEM})
        endforeach()
      endif()

      add_custom_target(
        ccov-all-report
        COMMAND
          ${LLVM_COV_PATH}
          report
          `cat
          ${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/binaries.list`
          -instr-profile=${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/all-merged.profdata
          ${EXCLUDE_REGEX}
        DEPENDS ccov-all-processing)

      add_custom_target(
        ccov-all
        COMMAND
          ${LLVM_COV_PATH}
          show
          `cat
          ${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/binaries.list`
          -instr-profile=${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/all-merged.profdata
          -show-line-counts-or-regions
          -output-dir=${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/all-merged
          -format="html"
          ${EXCLUDE_REGEX}
        DEPENDS ccov-all-processing)

    elseif(CMAKE_COMPILER_IS_GNUCXX)
      # Targets
      set(COVERAGE_INFO "${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/all-merged.info")

      add_custom_target(ccov-all-processing COMMAND ;)

      foreach(EXCLUDE_ITEM ${add_code_coverage_all_targets_EXCLUDE})
        set(EXCLUDE_REGEX
            ${EXCLUDE_REGEX}
            --remove
            ${COVERAGE_INFO}
            '${EXCLUDE_ITEM}')
      endforeach()

      if(EXCLUDE_REGEX)
        set(EXCLUDE_COMMAND
            ${LCOV_PATH}
            ${EXCLUDE_REGEX}
            --output-file
            ${COVERAGE_INFO})
      else()
        set(EXCLUDE_COMMAND ;)
      endif()

      add_custom_target(ccov-all
                        COMMAND ${LCOV_PATH}
                                --directory ${CMAKE_BINARY_DIR}
                                --capture
                                --output-file ${COVERAGE_INFO}
                        COMMAND ${EXCLUDE_COMMAND}
                        COMMAND ${GENHTML_PATH}
                                -o
                                ${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/all-merged
                                ${COVERAGE_INFO}
                        COMMAND ${CMAKE_COMMAND} -E remove ${COVERAGE_INFO}
                        DEPENDS ccov-all-processing)

    endif()

    add_custom_command(
      TARGET ccov-all POST_BUILD
      COMMAND ;
      COMMENT
        "Open ${CMAKE_COVERAGE_OUTPUT_DIRECTORY}/all-merged/index.html in your browser to view the coverage report."
      )
  endif()
endfunction()