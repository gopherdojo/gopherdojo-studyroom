#!/bin/bash

find ./files -type f -name "*.png" | xargs -I {} rm {}
