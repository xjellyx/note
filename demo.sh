#!/bin/bash

demoFunc(){
  echo "hello,world"
  read a
  read b
  return $(($a+$b))
}

demoFunc
echo "输出函数返回值：$?"