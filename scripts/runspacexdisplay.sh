#!/bin/bash

if [ "$TERM" == "linux" ] ; then
	while :
	do
		./spacexdisplay
		sleep 1
	done
fi
