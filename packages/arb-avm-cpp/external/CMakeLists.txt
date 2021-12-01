include(ExternalProject)

ExternalProject_Add(secp256k1_ext
  SOURCE_DIR "${CMAKE_CURRENT_SOURCE_DIR}/secp256k1"
  INSTALL_DIR "${CMAKE_BINARY_DIR}"
  CONFIGURE_COMMAND ./autogen.sh COMMAND <SOURCE_DIR>/configure --enable-module-recovery --prefix <INSTALL_DIR>
  BUILD_COMMAND make
  UPDATE_COMMAND ""
  BUILD_IN_SOURCE 1
)

add_library(secp256k1 IMPORTED STATIC GLOBAL)
add_dependencies (secp256k1 secp256k1_ext)
ExternalProject_Get_Property(secp256k1_ext install_dir)

file(MAKE_DIRECTORY ${install_dir}/include)

set_target_properties(secp256k1 PROPERTIES
    "IMPORTED_LOCATION" "${install_dir}/lib/libsecp256k1.a"
    "INTERFACE_INCLUDE_DIRECTORIES" "${install_dir}/include"

)

SET(WITH_PROCPS OFF CACHE BOOL "" FORCE)
add_subdirectory(libff)

add_subdirectory(Catch2)

set(JSON_BuildTests OFF CACHE INTERNAL "")
add_subdirectory(json)

add_library(picosha2 INTERFACE)
set_property(TARGET picosha2 PROPERTY INTERFACE_INCLUDE_DIRECTORIES ${CMAKE_CURRENT_SOURCE_DIR})
