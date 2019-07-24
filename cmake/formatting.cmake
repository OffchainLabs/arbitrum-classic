#
# Copyright (C) 2019 by George Cave - gcave@stablecoder.ca
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

#
# clang-format
#
find_program(CLANG_FORMAT_EXE "clang-format")
if(CLANG_FORMAT_EXE)
  message(STATUS "clang-format found: ${CLANG_FORMAT_EXE}")
else()
  message(STATUS "clang-format not found!")
endif()

# Generates a 'format' target using a custom name, files, and include
# directories all being parameters.
#
# Do note that in order for sources to be inherited properly, the source paths
# must be reachable from where the macro is called, or otherwise require a full
# path for proper inheritance.
#
# ~~~
# Required:
# TARGET_NAME - The name of the target to create.
#
# Optional: ARGN - The list of targets OR files to format. Relative and absolute
# paths are accepted.
# ~~~
function(clang_format TARGET_NAME)
  if(CLANG_FORMAT_EXE)
    # Check through the ARGN's, determine existant files
    foreach(item
            IN
            LISTS
            ARGN)
      if(TARGET ${item})
        # If the item is a target, then we'll attempt to grab the associated
        # source files from it.
        get_target_property(_TARGET_TYPE ${item} TYPE)
        if(NOT _TARGET_TYPE STREQUAL "INTERFACE_LIBRARY")
          get_property(_TEMP TARGET ${item} PROPERTY SOURCES)
          foreach(iter
                  IN
                  LISTS
                  _TEMP)
            if(EXISTS ${iter})
              set(FORMAT_FILES ${FORMAT_FILES} ${iter})
            endif()
          endforeach()
        endif()
      elseif(EXISTS ${item})
        # Check if it's a full file path
        set(FORMAT_FILES ${FORMAT_FILES} ${item})
      elseif(EXISTS ${CMAKE_CURRENT_SOURCE_DIR}/${item})
        # Check if it's based on the current source dir
        set(FORMAT_FILES ${FORMAT_FILES} ${CMAKE_CURRENT_SOURCE_DIR}/${item})
      endif()
    endforeach()

    # Make the target
    if(FORMAT_FILES)
      if(TARGET ${TARGET_NAME})
        message(
          ERROR
          "Cannot create clang-format target '${TARGET_NAME}', already exists.")
      else()
        add_custom_target(${TARGET_NAME}
                          COMMAND ${CLANG_FORMAT_EXE}
                                  -i -style=file
                                  ${FORMAT_FILES})

        if(NOT TARGET format)
          add_custom_target(format)
        endif()

        add_dependencies(format ${TARGET_NAME})
      endif()
    endif()

  endif()
endfunction()

#
# cmake-format
#
find_program(CMAKE_FORMAT_EXE "cmake-format")
if(CMAKE_FORMAT_EXE)
  message(STATUS "cmake-format found: ${CMAKE_FORMAT_EXE}")
else()
  message(STATUS "cmake-format not found!")
endif()

# When called, this function will call 'cmake-format' program on all listed
# files (if both the program and the files exist and are found)
# ~~~
# Required:
# TARGET_NAME - The name of the target to create.
#
# Optional:
# ARGN - Any  arguments passed in will be considered as 'files' to perform the
# formatting on. Any items that are not files will be ignored. Both relative and
# absolute paths are accepted.
# ~~~
function(cmake_format TARGET_NAME)
  if(CMAKE_FORMAT_EXE)
    # Determine files that exist
    foreach(iter
            IN
            LISTS
            ARGN)
      if(EXISTS ${iter})
        set(FORMAT_FILES ${FORMAT_FILES} ${iter})
      elseif(EXISTS ${CMAKE_CURRENT_SOURCE_DIR}/${iter})
        set(FORMAT_FILES ${FORMAT_FILES} ${CMAKE_CURRENT_SOURCE_DIR}/${iter})
      endif()
    endforeach()

    # Generate target
    if(FORMAT_FILES)
      if(TARGET ${TARGET_NAME})
        message(
          ERROR
          "Cannot create cmake-format target '${TARGET_NAME}', already exists.")
      else()
        add_custom_target(${TARGET_NAME}
                          COMMAND ${CMAKE_FORMAT_EXE} -i ${FORMAT_FILES})

        if(NOT TARGET cmake-format)
          add_custom_target(cmake-format)
        endif()
        add_dependencies(cmake-format ${TARGET_NAME})
      endif()
    endif()
  endif()
endfunction()