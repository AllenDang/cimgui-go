# Automatically generated by scripts/boost/generate-ports.ps1

vcpkg_from_github(
    OUT_SOURCE_PATH SOURCE_PATH
    REPO boostorg/yap
    REF boost-${VERSION}
    SHA512 5ed8930d8764eab2edcefd9627315dc812ff80eda8e2982f0c2fcb89ff828e13f4998fa852e57578116863da09b0cf8ca83afea9aa18b6ee2c4852277b14afce
    HEAD_REF master
)

include(${CURRENT_INSTALLED_DIR}/share/boost-vcpkg-helpers/boost-modular-headers.cmake)
boost_modular_headers(SOURCE_PATH ${SOURCE_PATH})