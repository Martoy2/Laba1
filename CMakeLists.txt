cmake_minimum_required(VERSION 4.0)
project(2C_Laba1)

set(CMAKE_CXX_STANDARD 20)

add_executable(dbms main.cpp
        structures/container.h
        structures/serialization.h
        structures/serialization.cpp
        structures/array.h
        structures/array.cpp
        structures/node.h
        structures/forward_list.h
        structures/forward_list.cpp
        structures/doubly_list.h
        structures/doubly_list.cpp
        structures/stack.h
        structures/stack.cpp
        structures/queue.h
        structures/queue.cpp
        structures/avl_tree.h
        structures/avl_tree.cpp
        engine/storage.h
        engine/storage.cpp
        engine/query.h
        engine/query.cpp
        engine/executor.h
        engine/executor.cpp)

set_target_properties(dbms PROPERTIES RUNTIME_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR})
