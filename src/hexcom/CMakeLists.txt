cmake_minimum_required(VERSION 3.16)
project(HexCom)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_EXTENSIONS ON)

IF(UNIX AND NOT APPLE)
    find_package(LLVM 9.0 REQUIRED CONFIG)
    message(STATUS "Found LLVM ${LLVM_PACKAGE_VERSION}")
    message(STATUS "Using LLVMConfig.cmake in: ${LLVM_DIR}")
    include_directories(${LLVM_INCLUDE_DIRS})
    add_definitions(${LLVM_DEFINITIONS})
    set(CMAKE_POSITION_INDEPENDENT_CODE ON)
ENDIF()

IF (WIN32)
    add_definitions(/D_WIN32_WINNT=0x0601)
    set(OPENSSL_NO_ASM ON)
ENDIF()

set(CMAKE_SUPPRESS_DEVELOPER_WARNINGS 1 CACHE INTERNAL "No dev warnings")

IF(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
    set(VCPKG_TARGET_ARCHITECTURE x86_64 arm64)

    SET(UE_ROOT "/Users/Shared/Epic Games/UE_5.0/")
    set(gRPC_ZLIB_PROVIDER package)
    set(ZLIB_INCLUDE_DIR "${UE_ROOT}/Engine/Source/ThirdParty/zlib/v1.2.8/include/Mac/")
    set(ZLIB_LIBRARY_DEBUG "${UE_ROOT}/Engine/Source/ThirdParty/zlib/v1.2.8/lib/Mac/libz.a")
    set(ZLIB_LIBRARY_RELEASE "${UE_ROOT}/Engine/Source/ThirdParty/zlib/v1.2.8/lib/Mac/libz.a")

    set(gRPC_SSL_PROVIDER=OpenSSL)

    set(gRPC_SSL_PROVIDER package)
    set(LIB_EAY_LIBRARY_DEBUG "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/lib/Mac/libcrypto.a")
    set(LIB_EAY_LIBRARY_RELEASE "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/lib/Mace/libcrypto.a")
    set(LIB_EAY_DEBUG "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/lib/Mac/libcrypto.a")
    set(LIB_EAY_RELEASE "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/lib/Mac/libcrypto.a")

    set(OPENSSL_INCLUDE_DIR "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/include/Mac")
    set(SSL_EAY_DEBUG "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/lib/Mac/libssl.a")
    set(SSL_EAY_LIBRARY_DEBUG "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/lib/Mac/libssl.a")
    set(SSL_EAY_LIBRARY_RELEASE "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/lib/Mac/libssl.a")
    set(SSL_EAY_RELEASE "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/lib/Mac/libssl.a")
ELSEIF(WIN32)
    set(UE_ROOT "D:/Epic Games/Games/UE_5.0")
    set(gRPC_ZLIB_PROVIDER package)
    set(ZLIB_INCLUDE_DIR "${UE_ROOT}/Engine/Source/ThirdParty/zlib/v1.2.8/include/Win64/VS2015")
    set(ZLIB_LIBRARY_DEBUG "${UE_ROOT}/Engine/Source/ThirdParty/zlib/v1.2.8/lib/Win64/VS2015/Release/zlibstatic.lib")
    set(ZLIB_LIBRARY_RELEASE "${UE_ROOT}/Engine/Source/ThirdParty/zlib/v1.2.8/lib/Win64/VS2015/Release/zlibstatic.lib")
    set(gRPC_SSL_PROVIDER package)
    set(LIB_EAY_LIBRARY_DEBUG "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/Lib/Win64/VS2015/Debug/libcrypto.lib")
    set(LIB_EAY_LIBRARY_RELEASE "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/Lib/Win64/VS2015/Release/libcrypto.lib")
    set(LIB_EAY_DEBUG "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/Lib/Win64/VS2015/Debug/libcrypto.lib")
    set(LIB_EAY_RELEASE "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/Lib/Win64/VS2015/Release/libcrypto.lib")

    set(OPENSSL_INCLUDE_DIR "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/include/Win64/VS2015")
    set(SSL_EAY_DEBUG "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/Lib/Win64/VS2015/Debug/libssl.lib")
    set(SSL_EAY_LIBRARY_DEBUG "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/Lib/Win64/VS2015/Debug/libssl.lib")
    set(SSL_EAY_LIBRARY_RELEASE "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/Lib/Win64/VS2015/Release/libssl.lib")
    set(SSL_EAY_RELEASE "${UE_ROOT}/Engine/Source/ThirdParty/OpenSSL/1.1.1k/Lib/Win64/VS2015/Release/libssl.lib")
ENDIF()

set(gRPC_BUILD_GRPC_CPP_PLUGIN ON)
set(gRPC_BUILD_GRPC_CSHARP_PLUGIN OFF)
set(gRPC_BUILD_GRPC_NODE_PLUGIN OFF)
set(gRPC_BUILD_GRPC_OBJECTIVE_C_PLUGIN OFF)
set(gRPC_BUILD_GRPC_PHP_PLUGIN OFF)
set(gRPC_BUILD_GRPC_PYTHON_PLUGIN OFF)
set(gRPC_BUILD_GRPC_RUBY_PLUGIN  OFF)


set(gRPC)

set(gRPC_USE_PROTO_LITE OFF)

include(FetchContent)
FetchContent_Declare(
        gRPC
        GIT_REPOSITORY https://github.com/grpc/grpc
        GIT_TAG        v1.41.0
)
set(FETCHCONTENT_QUIET OFF)
FetchContent_MakeAvailable(gRPC)

add_subdirectory(lib)
add_subdirectory(apps/example)

set(HEXCOM_INSTALL_DIR ${CMAKE_CURRENT_BINARY_DIR}/../hexcom-install/)
set(HEXCOM_INSTALL_DIR_HEADERS ${CMAKE_CURRENT_BINARY_DIR}/../hexcom-install/ThirdParty/HexWorldLibrary/include/hexworld)
set(HEXCOM_INSTALL_DIR_WIN64 ${CMAKE_CURRENT_BINARY_DIR}/../hexcom-install/ThirdParty/HexWorldLibrary/Win64/Release)

file(MAKE_DIRECTORY ${HEXCOM_INSTALL_DIR_HEADERS})
file(MAKE_DIRECTORY ${HEXCOM_INSTALL_DIR_WIN64})

set(HEXCOM_HEADER_FILES
        ${CMAKE_SOURCE_DIR}/include/hexworld/hex_client.h
        ${CMAKE_SOURCE_DIR}/include/hexworld/hex_com_state.h
        ${CMAKE_SOURCE_DIR}/include/hexworld/hex_lib.h
        )

if(WIN32)
    set(HEXCOM_LIBS_WIN64
            ${CMAKE_CURRENT_BINARY_DIR}/lib/hex_client_lib.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/address_sorting.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/cares/cares/lib/cares.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/gpr.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/grpc.lib 
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/grpc++.lib 
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/protobuf/libprotobufd.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/upb.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/abseil-cpp/absl/base/absl_base.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/abseil-cpp/absl/numeric/absl_int128.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/abseil-cpp/absl/base/absl_raw_logging_internal.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/abseil-cpp/absl/base/absl_throw_delegate.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/abseil-cpp/absl/time/absl_time.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/abseil-cpp/absl/time/absl_time_zone.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/abseil-cpp/absl/strings/absl_str_format_internal.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/abseil-cpp/absl/strings/absl_strings_internal.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/abseil-cpp/absl/strings/absl_strings.lib
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/abseil-cpp/absl/types/absl_bad_optional_access.lib
            )
#    set(UNREAL_PROJECT_DIR F:/Unreal/HexWorld/)
    set(UNREAL_PLATFORM Win64)
elseif(UNIX)
    set(HEXCOM_LIBS
            ${CMAKE_CURRENT_BINARY_DIR}/lib/libhex_client_lib.a
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/protobuf/libprotobuf.a
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/libgrpc++.a
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/libgrpc.a
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/libgpr.a
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/libaddress_sorting.a
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/boringssl-with-bazel/libssl.a
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/boringssl-with-bazel/libcrypto.a
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/libupb.a
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/zlib/libz.a
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/cares/cares/lib/libcares.a
            ${CMAKE_CURRENT_BINARY_DIR}/_deps/grpc-build/third_party/abseil-cpp/absl/base/*.a)
    set(UNREAL_PROJECT_DIR /home/olivier/Projects/UnrealProjects/HexWorld)
    set(UNREAL_PLATFORM Linux)
endif()

add_custom_target(copy_headers)
add_custom_target(copy_libs_win64)
add_custom_target(create_zip COMMAND ${CMAKE_COMMAND} -E tar "cvf" "archive.zip" --format=zip -- ${HEXCOM_INSTALL_DIR})

add_dependencies(copy_libs_win64 hex_client_lib)
add_dependencies(copy_headers copy_libs_win64)
add_dependencies(create_zip copy_headers)

foreach( FILE_I ${HEXCOM_LIBS_WIN64})
    add_custom_command(
            TARGET copy_libs_win64
            POST_BUILD
            COMMAND ${CMAKE_COMMAND}
            ARGS -E copy ${FILE_I} ${HEXCOM_INSTALL_DIR_WIN64}
    )
endforeach( FILE_I )

foreach( FILE_I ${HEXCOM_HEADER_FILES})
    add_custom_command(
            TARGET copy_headers
            POST_BUILD
            COMMAND ${CMAKE_COMMAND}
            ARGS -E copy ${FILE_I} ${HEXCOM_INSTALL_DIR_HEADERS}
    )
endforeach( FILE_I )
