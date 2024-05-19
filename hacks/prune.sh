#!/bin/bash

find . \
    \( \
    -name bun.lockb \
    -o -name .svelte-kit \
    -o -name 'vite.config.ts.*' \
    -o -type d -name node_modules \
    -o -name d -name cache \
    -o -type d -name build \
    -o -type d -name dist \
    \) \
    -exec rm -rf {} +

rm -rf ui/src/paraglide

echo "✅ Autogenerated files and directories were removed."
