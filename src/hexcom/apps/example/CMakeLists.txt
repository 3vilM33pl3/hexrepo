set(BUILD_TESTING OFF)

IF (WIN32)
    add_definitions(/D_WIN32_WINNT=0x0601)
ENDIF()

find_package(cli CONFIG REQUIRED)

add_executable(hexworld_client hexworld_clientcli.cpp)
target_include_directories(hexworld_client PUBLIC ../../include)
target_link_libraries(hexworld_client hex_client_lib absl::base absl::flags absl::flags_parse cli::cli)
