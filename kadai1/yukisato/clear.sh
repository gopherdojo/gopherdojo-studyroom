#!/bin/bash

find ./files -type f -name "*.jpg" | xargs -I {} rm {}
