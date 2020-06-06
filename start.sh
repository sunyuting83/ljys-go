#!/bin/bash
echo "start..."
basepath=$(cd `dirname $0`; pwd)
export PATH=$basepath:/sbin:/bin:/usr/bin:/usr/local/bin
pidpath="$basepath/movie.pid"
if [ -f $pidpath ]; then
  pid=`cat $pidpath`
  kill -HUP $pid
  sleep 1
  cd $basepath
  nohup ./movie -p 5002 > run.log 2 >run.log 2>&1 &
  echo $! > $pidpath
  echo "The process $! is running..."
else
  cd $basepath
  nohup ./movie -p 5002 > run.log 2 >run.log 2>&1 &
  echo $! > $pidpath
  echo "The process $! is running..."
fi