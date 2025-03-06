# ec2WithLambda

just messing around with lambda's custom runtime 
using go binaries

i think these thin binaries are sexy. instead of having
to specify a language version all the time

## build

you might need to `chmod +x ./build.sh`

i have simple `zsh` script that builds a binary named
`bootstrap` and then compresses this in a zip using `7z`

### tricks
- build all bins at once with: 
```zsh
for d in cmd/*; do ./build.sh "${d#cmd/}"; done
```