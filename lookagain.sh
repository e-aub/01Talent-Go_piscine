#!/bin/bash
find  -type f -exec ls -r {} + -name "*.sh" | cut -d '.' -f2 | tr -d '/'