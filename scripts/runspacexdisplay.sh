#!/bin/bash

if [ "$TERM" == "linux" ] ; then
	while :
	do
		./../cmd/spacexdisplay/spacexdisplay
		sleep 1
	done
fi
