#!/bin/bash

bee pack -be GOOS=darwin -o bin/darwin --exr='bin/*'
bee pack -be GOOS=linux -o bin/linux --exr='bin/*'