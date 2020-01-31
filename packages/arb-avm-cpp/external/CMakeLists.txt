include(ExternalProject)
ExternalProject_Add(secp256k1_ext
  GIT_REPOSITORY "https://github.com/bitcoin-core/secp256k1.git"
  GIT_TAG "master"
  SOURCE_DIR "${CMAKE_CURRENT_SOURCE_DIR}/secp256k1"
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

ExternalProject_Add(libff_ext
  GIT_REPOSITORY "https://github.com/scipr-lab/libff"
  GIT_TAG "master"
  SOURCE_DIR "${CMAKE_CURRENT_SOURCE_DIR}/libff"
  CMAKE_ARGS -DOPENSSL_ROOT_DIR=${OPENSSL_ROOT_DIR}
             -DOPENSSL_INCLUDE_DIR=${OPENSSL_ROOT_DIR}/include
             -DWITH_PROCPS=OFF
             -DCMAKE_INSTALL_PREFIX=<INSTALL_DIR>
)

add_library(libff IMPORTED STATIC GLOBAL)
add_dependencies (libff libff_ext)
ExternalProject_Get_Property(libff_ext install_dir)

file(MAKE_DIRECTORY ${install_dir}/include)

set_target_properties(libff PROPERTIES
    "IMPORTED_LOCATION" "${install_dir}/lib/libff.a"
    "INTERFACE_INCLUDE_DIRECTORIES" "${install_dir}/include"

)
