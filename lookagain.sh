#!/bin/bash
find  -type f -name "*.sh"| cut -d '.' -f2 | tr -d '/'
