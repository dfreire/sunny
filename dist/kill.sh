#!/bin/bash
kill $(ps -aux | grep "sunny" | grep -v "grep" | awk '{ print $2 }')