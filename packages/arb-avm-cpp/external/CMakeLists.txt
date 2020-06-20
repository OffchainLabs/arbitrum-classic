include(ExternalProject)

find_package(Git)
if(GIT_FOUND)
# Update submodules as needed
    message(STATUS "Submodule update")
    execute_process(COMMAND ${GIT_EXECUTABLE} submodule update --init --recursive
                    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
                    RESULT_VARIABLE GIT_SUBMOD_RESULT)
    if(NOT GIT_SUBMOD_RESULT EQUAL "0")
        message(FATAL_ERROR "git submodule update --init failed with ${GIT_SUBMOD_RESULT}, please checkout submodules")
    endif()
endif()

if(NOT EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/libff/CMakeLists.txt")
    message(FATAL_ERROR "The submodules were not downloaded! GIT_SUBMODULE was turned off or failed. Please update submodules and try again.")
endif()

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
SET(OPENSSL_ROOT_DIR ${OPENSSL_ROOT_DIR})
SET(OPENSSL_INCLUDE_DIR ${OPENSSL_ROOT_DIR}/include)
SET(WITH_PROCPS OFF CACHE BOOL "" FORCE)
add_subdirectory(libff)
