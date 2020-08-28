# Copyright (c) 2014 Stefan.Eilemann@epfl.ch

# taken from https://github.com/pipedown/noisecpp/blob/master/CMake/Findrocksdb.cmake

# - Try to find the Rocksdb library
# Once done this will define
#
#  ROCKSDB_ROOT - Set this variable to the root installation
#
# Read-Only variables:
#  ROCKSDB_FOUND - system has the Rocksdb library
#  ROCKSDB_INCLUDE_DIR - the Rocksdb include directory
#  ROCKSDB_LIBRARIES - The libraries needed to use Rocksdb
#  ROCKSDB_VERSION - This is set to $major.$minor.$patch (eg. 0.9.8)

include(FindPackageHandleStandardArgs)

if(rocksdb_FIND_REQUIRED)
  set(_ROCKSDB_output_type FATAL_ERROR)
  set(_ROCKSDB_output 1)
else()
  set(_ROCKSDB_output_type STATUS)
  if(rocksdb_FIND_QUIETLY)
    set(_ROCKSDB_output)
  else()
    set(_ROCKSDB_output 1)
  endif()
endif()

find_path(_ROCKSDB_INCLUDE_DIR rocksdb/db.h
  HINTS $ENV{ROCKSDB_ROOT} ${ROCKSDB_ROOT}
  PATH_SUFFIXES include
  PATHS /usr /usr/local /opt /opt/local)

if(EXISTS "${_ROCKSDB_INCLUDE_DIR}/rocksdb/version.h")
  set(_ROCKSDB_Version_file "${_ROCKSDB_INCLUDE_DIR}/rocksdb/version.h")
elseif(EXISTS "${_ROCKSDB_INCLUDE_DIR}/rocksdb/db.h")
  set(_ROCKSDB_Version_file "${_ROCKSDB_INCLUDE_DIR}/rocksdb/db.h")
endif()
if(_ROCKSDB_INCLUDE_DIR AND _ROCKSDB_Version_file)
  file(READ ${_ROCKSDB_Version_file} _ROCKSDB_header_contents)
  string(REGEX REPLACE ".*kMajorVersion = ([0-9]+).*kMinorVersion = ([0-9]+).*"
    "\\1.\\2" _ROCKSDB_VERSION "${_ROCKSDB_header_contents}")
  string(REGEX REPLACE ".*ROCKSDB_MAJOR ([0-9]+).*ROCKSDB_MINOR ([0-9]+).*"
    "\\1.\\2" _ROCKSDB_VERSION "${_ROCKSDB_header_contents}")
  set(ROCKSDB_VERSION ${_ROCKSDB_VERSION} CACHE INTERNAL
    "The version of rocksdb which was detected")
else()
  set(_ROCKSDB_EPIC_FAIL TRUE)
  if(_ROCKSDB_output)
    message(${_ROCKSDB_output_type}
      "Can't find rocksdb header file rocksdb/db.h.")
  endif()
endif()

# Version checking
if(ROCKSDB_FIND_VERSION AND ROCKSDB_VERSION)
  if(ROCKSDB_FIND_VERSION_EXACT)
    if(NOT ROCKSDB_VERSION VERSION_EQUAL ${ROCKSDB_FIND_VERSION})
      set(_ROCKSDB_version_not_exact TRUE)
    endif()
  else()
    # version is too low
    if(NOT ROCKSDB_VERSION VERSION_EQUAL ${ROCKSDB_FIND_VERSION} AND
        NOT ROCKSDB_VERSION VERSION_GREATER ${ROCKSDB_FIND_VERSION})
      set(_ROCKSDB_version_not_high_enough TRUE)
    endif()
  endif()
endif()

find_library(ROCKSDB_LIBRARY rocksdb
  HINTS ${CMAKE_SOURCE_DIR}/../../.. $ENV{ROCKSDB_ROOT} ${ROCKSDB_ROOT}
  PATH_SUFFIXES lib lib64
  PATHS /usr /usr/local /opt /opt/local)

# Inform the users with an error message based on what version they
# have vs. what version was required.
if(NOT ROCKSDB_VERSION)
  set(_ROCKSDB_EPIC_FAIL TRUE)
  if(_ROCKSDB_output)
    message(${_ROCKSDB_output_type}
      "Version not found in ${_ROCKSDB_Version_file}.")
  endif()
elseif(_ROCKSDB_version_not_high_enough)
  set(_ROCKSDB_EPIC_FAIL TRUE)
  if(_ROCKSDB_output)
    message(${_ROCKSDB_output_type}
      "Version ${ROCKSDB_FIND_VERSION} or higher of rocksdb is required. "
      "Version ${ROCKSDB_VERSION} was found in ${_ROCKSDB_Version_file}.")
  endif()
elseif(_ROCKSDB_version_not_exact)
  set(_ROCKSDB_EPIC_FAIL TRUE)
  if(_ROCKSDB_output)
    message(${_ROCKSDB_output_type}
      "Version ${ROCKSDB_FIND_VERSION} of rocksdb is required exactly. "
      "Version ${ROCKSDB_VERSION} was found.")
  endif()
else()
  if(ROCKSDB_FIND_REQUIRED)
    if(ROCKSDB_LIBRARY MATCHES "ROCKSDB_LIBRARY-NOTFOUND")
      message(FATAL_ERROR "Missing the rocksdb library.\n"
        "Consider using CMAKE_PREFIX_PATH or the ROCKSDB_ROOT environment variable. "
        "See the ${CMAKE_CURRENT_LIST_FILE} for more details.")
    endif()
  endif()
  find_package_handle_standard_args(RocksDB DEFAULT_MSG
                                    ROCKSDB_LIBRARY _ROCKSDB_INCLUDE_DIR)
endif()

if(_ROCKSDB_EPIC_FAIL)
  # Zero out everything, we didn't meet version requirements
  set(ROCKSDB_FOUND FALSE)
  set(ROCKSDB_LIBRARY)
  set(_ROCKSDB_INCLUDE_DIR)
  set(ROCKSDB_INCLUDE_DIRS)
  set(ROCKSDB_LIBRARIES)
else()
  set(ROCKSDB_INCLUDE_DIRS ${_ROCKSDB_INCLUDE_DIR})
  set(ROCKSDB_LIBRARIES ${ROCKSDB_LIBRARY})

  add_library(RocksDB::RocksDB UNKNOWN IMPORTED)
  set_target_properties(RocksDB::RocksDB PROPERTIES
    IMPORTED_LOCATION "${ROCKSDB_LIBRARIES}"
    INTERFACE_INCLUDE_DIRECTORIES "${ROCKSDB_INCLUDE_DIRS}"
    IMPORTED_LINK_INTERFACE_LIBRARIES "dl"
  )
  if(_ROCKSDB_output)
    message(STATUS
      "Found rocksdb ${ROCKSDB_VERSION} in ${ROCKSDB_INCLUDE_DIRS};${ROCKSDB_LIBRARIES}")
  endif()
endif()
