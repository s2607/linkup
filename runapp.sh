#!/bin/sh
support = "someone@liberty.edu"
while true
do
	./linkup
	echo "The process has exited, please contact the current maintainers at $support"
	sleep 3;
done
