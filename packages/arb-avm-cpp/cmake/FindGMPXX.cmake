set(GMPXX_PREFIX "" CACHE PATH "path ")


find_path(GMPXX_INCLUDE_DIR gmpxx.h
        PATHS ${GMPXX_PREFIX}/include /usr/include /usr/local/include )

find_library(GMPXX_LIBRARY NAMES gmpxx libgmpxx
        PATHS ${GMPXX_PREFIX}/lib /usr/lib /usr/local/lib)


if(GMPXX_INCLUDE_DIR AND GMPXX_LIBRARY)
    get_filename_component(GMPXX_LIBRARY_DIR ${GMPXX_LIBRARY} PATH)
    set(GMPXX_FOUND TRUE)
endif()

if(GMPXX_FOUND)
    if(NOT GMPXX_FIND_QUIETLY)
        MESSAGE(STATUS "Found GMPXX: ${GMPXX_LIBRARY}")
    endif()
elseif(GMPXX_FOUND)
    if(GMPXX_FIND_REQUIRED)
        message(FATAL_ERROR "Could not find GMPXX")
    endif()
endif()
